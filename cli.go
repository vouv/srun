package main

import (
	"bufio"
	"fmt"
	"github.com/monigo/srun/core"
	"github.com/monigo/srun/model"
	"github.com/monigo/srun/pkg/term"
	"github.com/monigo/srun/store"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime"
	"strings"
)

type Func func(cmd string, params ...string)

var DefaultClient = &Client{}

type Client struct{}

var serverTypes = map[string]string{
	"xyw": model.ServerTypeOrigin,
	"yd":  model.ServerTypeCMCC,
	"lt":  model.ServerTypeWCDMA,
}

// 登录
func (s *Client) Login(cmd string, params ...string) {
	account, gErr := store.LoadAccount()
	if gErr != nil {
		log.Error(gErr)
		os.Exit(1)
	}
	username := account.Username
	server := account.Server

	if len(params) != 0 {
		if t, ok := serverTypes[params[0]]; ok {
			server = t
		} else {
			s.CmdList()
			return
		}
	}
	log.Info("正在登录: ", server)

	username = model.AddSuffix(username, server)
	info, err := core.Login(username, account.Password)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	log.Info("登录成功!")
	log.Info("在线IP: ", info.ClientIp)

	err = store.SetInfo(info.AccessToken, info.ClientIp)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func (s *Client) Logout(cmd string, params ...string) {
	if len(params) == 0 {
		var err error
		account, err := store.LoadAccount()
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
		if err = core.Logout(account.Username); err != nil {
			log.Error(err)
			os.Exit(1)
		}
		log.Info("注销成功!")
	} else {
		s.CmdList()
	}
}

func (s *Client) GetInfo(cmd string, params ...string) {
	if len(params) == 0 {
		var err error
		account, err := store.LoadAccount()
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
		log.Info("当前校园网登录账号:", account.Username)
		if err = core.Info(account); err != nil {
			log.Error(err)
			os.Exit(1)
		}
	} else {
		s.CmdList()
	}
}

func (Client) SetAccount(cmd string, params ...string) {

	in := os.Stdin
	fmt.Print("设置校园网账号:\n>")
	username := readInput(in)

	// 终端API
	fd, _ := term.GetFdInfo(in)
	oldState, err := term.SaveState(fd)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	fmt.Print("设置校园网密码:\n>")

	// read in stdin
	_ = term.DisableEcho(fd, oldState)
	pwd := readInput(in)
	_ = term.RestoreTerminal(fd, oldState)

	fmt.Println()

setServer:
	fmt.Print("设置默认登录模式( 校园网(默认): 1 | 移动: 2 | 联通: 3 )\n>")
	server := readInput(in)
	switch server {
	case "", "1":
		server = model.ServerTypeOrigin
	case "2":
		server = model.ServerTypeCMCC
	case "3":
		server = model.ServerTypeWCDMA
	default:
		goto setServer
	}

	// trim
	username = strings.TrimSpace(username)
	pwd = strings.TrimSpace(pwd)

	if err := store.SetAccount(username, pwd, server); err != nil {
		log.Error(err)
		os.Exit(1)
	}
	log.Info("账号密码已被保存")
}

func readInput(in io.Reader) string {
	reader := bufio.NewReader(in)
	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	return string(line)
}

func (Client) ShowVersion() {
	fmt.Println("System:")
	fmt.Printf("\tOS:%s ARCH:%s GOVERSION:%s\n", runtime.GOOS, runtime.GOARCH, runtime.Version())
	fmt.Println("About:")
	fmt.Printf("\tVersion:%s\n", Version)
	fmt.Println("\n\t</> with ❤ By monigo")
}

// srun help [COMMAND]
func (s *Client) CmdHelp(cmd string, params ...string) {
	if len(params) == 0 {
		fmt.Println(s.CmdList())
	} else {
		if c, ok := cmdDocs[params[0]]; ok {
			fmt.Println("Usage: ", c[0])
		} else {
			fmt.Println(s.CmdList())
		}
	}
}

func (Client) CmdList() string {
	sb := &strings.Builder{}
	sb.WriteString(fmt.Sprint("\r\nUsage:	srun [OPTIONS] COMMAND \r\n\r\n"))

	sb.WriteString("A efficient client for BIT campus network\r\n\r\n")

	sb.WriteString("Options:\r\n")
	for k, v := range optionDocs {
		sb.WriteString(fmt.Sprintf("  %-10s%-20s\r\n", k, v))
	}

	sb.WriteString("\r\nCommands:\r\n")
	for k, v := range cmdDocs {
		sb.WriteString(fmt.Sprintf("  %-10s%-20s\r\n", k, v[1]))
	}
	return sb.String()
}

func (Client) Update(cmd string, params ...string) {
	ok, v, d := HasUpdate()
	if !ok {
		log.Info("当前已是最新版本:", Version)
		return
	}
	log.Info("发现新版本: ", v, "当前版本: ", Version)
	log.Info("打开链接下载: ", d)
}
