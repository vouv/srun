package cli

import (
	"fmt"
	"os"
	"runtime"
	"srun-cmd/config"
)

var optionDocs = map[string]string{
	"-d": "Show debug message",
	"-v": "Show version",
	"-h": "Show help",
}

var cmds = []string{
	"account",
	"login",
	"help",
	"info",
	"logout",
}
var cmdDocs = map[string][]string{
	"account": []string{"srun account [get]", "Set Username and Password"},
	"login":   []string{"srun [login] [xyw] [yd] [lt]", "Login Srun"},
	"logout":  []string{"srun logout", "Logout Srun"},
	"info":    []string{"srun info", "Get Srun Info"},
}

func Version() {
	fmt.Printf("Srun version/%s (OS:%s ARCH:%s GOVERSION:%s)\n", config.Version, runtime.GOOS, runtime.GOARCH, runtime.Version())
	fmt.Println("Srun Tool </> with ❤️ By Monigo")
}

func Help(cmd string, params ...string) {
	if len(params) == 0 {
		fmt.Println(CmdList())
		Version()
	} else {
		CmdHelps(params...)
	}
}

func CmdList() string {
	helpAll := fmt.Sprintf("Srun %s\r\n\r\n", config.Version)
	helpAll += "Options:\r\n"
	for k, v := range optionDocs {
		helpAll += fmt.Sprintf("\t%-20s%-20s\r\n", k, v)
	}
	helpAll += "\r\n"
	helpAll += "Commands:\r\n"
	for _, cmd := range cmds {
		if help, ok := cmdDocs[cmd]; ok {
			cmdDesc := help[1]
			helpAll += fmt.Sprintf("\t%-20s%-20s\r\n", cmd, cmdDesc)
		}
	}
	return helpAll
}

func CmdHelps(cmds ...string) {
	defer os.Exit(1)
	if len(cmds) == 0 {
		fmt.Println(CmdList())
	} else {
		for _, cmd := range cmds {
			CmdHelp(cmd)
		}
	}
}

func CmdHelp(cmd string) {
	docStr := fmt.Sprintf("Unknow cmd `%s`", cmd)
	if cmdDoc, ok := cmdDocs[cmd]; ok {
		docStr = fmt.Sprintf("Usage: %s\r\n  %s\r\n", cmdDoc[0], cmdDoc[1])
	}
	fmt.Println(docStr)
}
