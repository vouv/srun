package cli

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/monigo/srun-cmd"
	"github.com/monigo/srun-cmd/model"
	"github.com/monigo/srun-cmd/pkg/term"
	"github.com/monigo/srun-cmd/store"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

type Func func(cmd string, params ...string)

var cmdDocs = map[string][]string{
	"config": {"srun config", "Set Username and Password"},
	"login":  {"srun [login] [xyw|yd|lt]", "Login Srun"},
	"logout": {"srun logout", "Logout Srun"},
	"info":   {"srun info", "Get Srun Info"},
	"update": {"srun update", "Update srun"},
}

var CommandMap = map[string]Func{
	"config": SetAccount(),
	"login":  Login(),
	"logout": Logout(),
	"info":   GetInfo(),
	"update": Update(),

	"help": Help(),
}

type Client struct {
	LogLevel log.Level
	Cmd      string
	Params   []string

	debugMode   bool
	helpMode    bool
	versionMode bool
}

func New() *Client {
	return &Client{
		LogLevel: log.InfoLevel,
	}
}

func (s *Client) Handle() {
	s.parse()

	switch {
	case s.helpMode:
		s.Cmd = "help"
		Help()(s.Cmd)
		return
	case s.versionMode:
		Version()
		return
	case s.debugMode:
		s.LogLevel = log.DebugLevel
	}

	log.SetOutput(os.Stdout)
	log.SetLevel(s.LogLevel)
	log.SetFormatter(&log.TextFormatter{
		//DisableTimestamp: true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})

	if handle, ok := CommandMap[s.Cmd]; ok {
		handle(s.Cmd, s.Params...)
	} else {
		s.Cmd = "help"
		Help()(s.Cmd, s.Params...)
	}

}

func (s *Client) parse() {
	flag.BoolVar(&s.debugMode, "d", false, "debug mode")
	flag.BoolVar(&s.helpMode, "h", false, "show help")
	flag.BoolVar(&s.versionMode, "v", false, "show version")

	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		s.Cmd = args[0]
		s.Params = args[1:]
	} else {
		s.Cmd = "login"
	}
}

func SetAccount() Func {
	return func(cmd string, params ...string) {

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

	setDef:
		fmt.Print("设置默认登录模式( 校园网(默认): 1 | 移动: 2 | 联通: 3 )\n>")
		def := readInput(in)
		switch def {
		case "", "1":
			def = "校园网"
		case "2":
			def = "移动"
		case "3":
			def = "联通"
		default:
			goto setDef
		}

		// trim
		username = strings.TrimSpace(username)
		pwd = strings.TrimSpace(pwd)

		if err := store.SetAccount(username, pwd, def); err != nil {
			log.Error(err)
			os.Exit(1)
		}
		log.Info("账号密码已被保存")
	}
}

func readInput(in io.Reader) string {
	reader := bufio.NewReader(in)
	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	return string(line)
}

func Login() Func {
	return func(cmd string, params ...string) {
		account, gErr := store.LoadAccount()
		if gErr != nil {
			log.Error(gErr)
			os.Exit(1)
		}
		username := account.Username
		server := account.Server
		if len(params) == 0 {
			username = model.AddSuffix(username, server)
		} else {
			switch params[0] {
			case "xyw":
				server = model.ServerTypeSrun
			case "yd":
				server = model.ServerTypeCMCC
			case "lt":
				server = model.ServerTypeWCDMA
			default:
				CmdList()
				return
			}
		}
		log.Info("正在登录: ", server)

		username = model.AddSuffix(username, server)
		info, err := srun.Login(username, account.Password)
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
}

func Logout() Func {
	return func(cmd string, params ...string) {
		if len(params) == 0 {
			var err error
			account, err := store.LoadAccount()
			if err != nil {
				log.Error(err)
				os.Exit(1)
			}
			if err = srun.Logout(account.Username); err != nil {
				log.Error(err)
				os.Exit(1)
			}
			log.Infof("账号 %s 注销成功!", account.Username)
		} else {
			CmdList()
		}
	}
}

func GetInfo() Func {
	return func(cmd string, params ...string) {
		if len(params) == 0 {
			var err error
			account, err := store.LoadAccount()
			if err != nil {
				log.Error(err)
				os.Exit(1)
			}
			log.Info("当前校园网登录账号:", account.Username)
			if err = srun.Info(account); err != nil {
				log.Error(err)
				os.Exit(1)
			}
		} else {
			CmdList()
		}
	}
}
