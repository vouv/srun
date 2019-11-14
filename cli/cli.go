package cli

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/monigo/srun-cmd"
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
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("账号密码已被保存")
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
			fmt.Println(gErr)
			os.Exit(1)
		}
		username := account.Username
		if len(params) == 0 {
			switch account.Server {
			case "移动":
				log.Info("正在登录移动...")
				username = account.Username + "@yidong"
			case "联通":
				fmt.Println("正在登录联通...")
				username = account.Username + "@liantong"
			default:
				fmt.Println("正在登录校园网...")
			}
		} else {
			switch params[0] {
			case "xyw":
				log.Info("正在登录校园网...")
				username = account.Username
			case "yd":
				log.Info("正在登录移动...")
				username = account.Username + "@yidong"
			case "lt":
				log.Info("正在登录联通...")
				username = account.Username + "@liantong"
			default:
				CmdList()
				return
			}
		}
		tk, ip := srun.Login(username, account.Password)
		_ = store.SetInfo(tk, ip)
	}
}

func Logout() Func {
	return func(cmd string, params ...string) {
		if len(params) == 0 {
			account, gErr := store.LoadAccount()
			if gErr != nil {
				fmt.Println(gErr)
				os.Exit(1)
			}
			srun.Logout(account.Username)
		} else {
			CmdList()
		}
	}
}

func GetInfo() Func {
	return func(cmd string, params ...string) {
		if len(params) == 0 {
			account, gErr := store.LoadAccount()
			if gErr != nil {
				fmt.Println(gErr)
				os.Exit(1)
			}
			fmt.Println("当前校园网登录账号:", account.Username)
			srun.Info(account.Username, account.AccessToken, account.Ip)
		} else {
			CmdList()
		}
	}
}
