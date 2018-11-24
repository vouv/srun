package cli

import (
	"fmt"
	"os"
	"runtime"
)

var version = "v0.1.1"

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
	"account": []string{"srun account [<Username> <Password>]", "Get/Set Username and Password"},
	"login": []string{"srun [login]", "Login Srun"},
	"logout": []string{"srun logout", "Logout Srun"},
	"info": []string{"srun info", "Get Srun Info"},
	}

func Version() {
	fmt.Printf("Srun version/%s (OS:%s ARCH:%s GOVERSION:%s)\n", version, runtime.GOOS, runtime.GOARCH, runtime.Version())
}

func Help(cmd string, params ...string) {
	if len(params) == 0 {
		fmt.Println(CmdList())
	} else {
		CmdHelps(params...)
	}
}

func CmdList() string {
	helpAll := fmt.Sprintf("Srun %s\r\n\r\n", version)
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

func UserAgent() string {
	return fmt.Sprintf("Srun Tool By Monigo /%s (%s; %s; %s)", version, runtime.GOOS, runtime.GOARCH, runtime.Version())
}
