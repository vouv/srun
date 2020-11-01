package main

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vouv/srun/core"
	"github.com/vouv/srun/pkg/term"
	"github.com/vouv/srun/store"
	"io"
	"os"
	"runtime"
	"strings"
)

func LoginE(cmd *cobra.Command, args []string) error {
	account, err := store.ReadAccount()
	if err != nil {
		return err
	}
	log.Info("尝试登录...")

	//username = model.AddSuffix(username, server)
	info, err := core.Login(&account)
	if err != nil {
		return err
	}
	log.Info("登录成功!")

	err = store.SetInfo(info.AccessToken, info.Acid)
	if err != nil {
		return err
	}

	//GetInfo(cmd, params...)
	return nil
}

func LogoutE(cmd *cobra.Command, args []string) error {
	var err error
	account, err := store.ReadAccount()
	if err != nil {
		return err
	}
	if err = core.Logout(account.Username); err != nil {
		return err
	}
	log.Info("注销成功!")
	return nil
}

func InfoE(cmd *cobra.Command, args []string) error {
	res, err := core.Info()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	fmt.Println(res.String())
	return nil
}

func ConfigE(cmd *cobra.Command, args []string) error {

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
	return nil
}

func readInput(in io.Reader) string {
	reader := bufio.NewReader(in)
	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	return string(line)
}

func VersionString() string {
	return fmt.Sprintln("System:") +
		fmt.Sprintf("\tOS:%s ARCH:%s GO:%s\n", runtime.GOOS, runtime.GOARCH, runtime.Version()) +
		fmt.Sprintln("About:") +
		fmt.Sprintf("\tVersion: %s\n", Version) +
		fmt.Sprintln("\n\t</> with ❤ By vouv")
}

func Update(cmd string, params ...string) {
	ok, v, d := HasUpdate()
	if !ok {
		log.Info("当前已是最新版本:", Version)
		return
	}
	log.Info("发现新版本: ", v, "当前版本: ", Version)
	log.Info("打开链接下载: ", d)
}
