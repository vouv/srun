package cli

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io"
	"os"
	"srun"
	"srun/term"
	"strings"
)

type Func func(cmd string, params ...string)

func AccountH(cmd string, params ...string) {
	if len(params) == 0 {
		setAccount()
	} else if params[0] == "get" {
		account, gErr := GetAccount()
		if gErr != nil {
			fmt.Println(gErr)
			os.Exit(1)
		}
		fmt.Println("当前校园网登录账号:", account.Username)

	} else {
		CmdHelp(cmd)
	}
}

func setAccount() {
	in := os.Stdin
	fmt.Print("设置校园网账号:\n>")
	username := readInput(in)

	fd, _ := term.GetFdInfo(in)
	oldState, err := term.SaveState(fd)
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	fmt.Print("设置校园网密码:\n>")

	// read in stdin
	term.DisableEcho(fd, oldState)
	pwd := readInput(in)
	term.RestoreTerminal(fd, oldState)

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

	sErr := SetAccount(username, pwd, def)

	if sErr != nil {
		fmt.Println(sErr)
		os.Exit(1)
	}
	fmt.Println("账号密码已被保存")
}

func readInput(in io.Reader) string {
	reader := bufio.NewReader(in)
	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	return string(line)
}

func LoginH(cmd string, params ...string) {
	account, gErr := GetAccount()
	if gErr != nil {
		fmt.Println(gErr)
		os.Exit(1)
	}
	username := account.Username
	if len(params) == 0 {
		switch account.Server {
		case "移动":
			fmt.Println("正在登录移动...")
			username = account.Username + "@yidong"
		case "联通":
			fmt.Println("正在登录联通...")
			username = account.Username + "@liantong"
		default:
			fmt.Println("正在登录校园网...")
		}
	} else {
		switch params[0] {
		case "yd":
			fmt.Println("正在登录移动...")
			username = account.Username + "@yidong"
		case "lt":
			fmt.Println("正在登录联通...")
			username = account.Username + "@liantong"
		case "xyw":
			fmt.Println("正在登录校园网...")
			username = account.Username
		default:
			CmdHelp(cmd)
			return
		}
	}
	tk, ip := srun.Login(username, account.Password)
	_ = SetInfo(tk, ip)
}

func InfoH(cmd string, params ...string) {
	if len(params) == 0 {
		account, gErr := GetAccount()
		if gErr != nil {
			fmt.Println(gErr)
			os.Exit(1)
		}
		srun.Info(account.Username, account.AccessToken, account.Ip)

	} else {
		CmdHelp(cmd)
	}
}

func LogoutH(cmd string, params ...string) {
	if len(params) == 0 {
		account, gErr := GetAccount()
		if gErr != nil {
			fmt.Println(gErr)
			os.Exit(1)
		}
		srun.Logout(account.Username)
	} else {
		CmdHelp(cmd)
	}
}
