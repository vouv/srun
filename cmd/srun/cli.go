package main

import (
	"bufio"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/vouv/srun/core"
	"github.com/vouv/srun/pkg/term"
	"github.com/vouv/srun/store"
	"io"
	"os"
	"runtime"
	"strings"
)

var ErrReadAccount = errors.New("读取账号文件错误, 请执行`srun config`配置账号信息")

type Func func(cmd string, params ...string)

var DefaultClient = &Client{}

type Client struct{}

// 登录
func (s *Client) Login(cmd string, params ...string) {
	account, gErr := store.ReadAccount()
	if gErr != nil {
		log.Error(ErrReadAccount.Error())
		log.Debug(gErr)
		os.Exit(1)
	}

	log.Info("尝试登录...")

	//username = model.AddSuffix(username, server)
	info, err := core.Login(&account)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	log.Info("登录成功!")

	err = store.SetInfo(info.AccessToken, info.ClientIp)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	s.GetInfo(cmd, params...)
}

func (s *Client) Logout(cmd string, params ...string) {
	if len(params) == 0 {
		var err error
		account, err := store.ReadAccount()
		if err != nil {
			log.Error(ErrReadAccount.Error())
			log.Debug(err)
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
		res, err := core.Info()
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
		fmt.Println(res.String())
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

	// trim
	username = strings.TrimSpace(username)
	pwd = strings.TrimSpace(pwd)

	if err := store.SetAccount(username, pwd); err != nil {
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
	fmt.Printf("\tVersion: %s\n", Version)
	fmt.Println("\n\t</> with ❤ By vouv")
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
	sb.WriteString("Srun " + Version + "\r\n")
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
