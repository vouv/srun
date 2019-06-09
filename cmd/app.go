package main

import (
	"flag"
	"fmt"
	"github.com/astaxie/beego/logs"
	"runtime"
	"srun-cmd/cmd/cli"
	"srun-cmd/config"
)

var supportedCmds = map[string]cli.Func{
	"account": cli.AccountH,
	"login":   cli.LoginH,
	"logout":  cli.LogoutH,
	"help":    cli.Help,
	"info":    cli.InfoH,
}

func main() {
	//set cpu count
	runtime.GOMAXPROCS(runtime.NumCPU())

	//parse command
	logs.SetLevel(logs.LevelInformational)
	logs.SetLogger(logs.AdapterConsole)

	//global options
	var debugMode bool
	var helpMode bool
	var versionMode bool

	flag.BoolVar(&debugMode, "d", false, "debug mode")
	flag.BoolVar(&helpMode, "h", false, "show help")
	flag.BoolVar(&versionMode, "v", false, "show version")
	flag.Parse()

	if helpMode {
		cli.Help("help")
		return
	}

	if versionMode {
		cli.Version()
		return
	}

	//set log level
	if debugMode {
		logs.SetLevel(logs.LevelDebug)
	}

	//set cmd and params
	flag.Parse()
	args := flag.Args()
	if len(args) >= 1 {
		cmd := args[0]
		params := args[1:]

		if cliFunc, ok := supportedCmds[cmd]; ok {
			cliFunc(cmd, params...)
		} else {
			supportedCmds["login"]("login", args...)
		}
	} else {
		supportedCmds["login"]("login")
	}

	// has update
	if ok, url := cli.HasUpdate(); ok {
		fmt.Print("更新: " + url)
		fmt.Println(" 当前版本: " + config.Version)
	}

}
