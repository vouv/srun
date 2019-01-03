package cli

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io"
	"login-srun/cmd/srun"
	"login-srun/cmd/term"
	"os"
	"strings"
)

type CliFunc func(cmd string, params ...string)

type accountOptions struct {
	user          string
	password      string
}

func AccountH(cmd string, params ...string)  {
	if len(params) == 0 {
		setAccount()
	} else if params[0] == "get" {
		account, gErr := GetAccount()
		if gErr != nil {
			fmt.Println(gErr)
			os.Exit(1)
		}
		fmt.Println("当前校园网登录账号:", account.Username)

	}else {
		CmdHelp(cmd)
	}
}

func setAccount()  {
	in := os.Stdin
	fmt.Print("请输入校园网账号：")
	username := readInput(in)

	fd, _ := term.GetFdInfo(in)
	oldState, err := term.SaveState(fd)
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	fmt.Print("请输入校园网密码：🔑")
	term.DisableEcho(fd, oldState)
	pwd := readInput(in)

	// restore
	term.RestoreTerminal(fd, oldState)

	// trim
	username = strings.TrimSpace(username)
	pwd = strings.TrimSpace(pwd)

	sErr := SetAccount(username, pwd)
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
	if len(params) == 0 {
		account, gErr := GetAccount()
		if gErr != nil {
			fmt.Println(gErr)
			os.Exit(1)
		}
 		tk, ip := srun.Login(account.Username, account.Password)
 		SetInfo(tk, ip)
	} else {
		CmdHelp(cmd)
	}
}

func InfoH(cmd string, params ...string)  {
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

func LogoutH(cmd string, params ...string)  {
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