package cli

import (
	"fmt"
	"runtime"
	"srun/config"
)

var optionDocs = map[string]string{
	"-d": "Show debug message",
	"-v": "Show version",
	"-h": "Show help",
}

func Version() {
	fmt.Println("System:")
	fmt.Printf("\tOS:%s ARCH:%s GOVERSION:%s\n", runtime.GOOS, runtime.GOARCH, runtime.Version())
	fmt.Println("About:")
	fmt.Printf("\tVersion:%s\n", config.Version)
	fmt.Println("\n\t</> with ❤ By monigo")
}

func Help() Func {
	return func(cmd string, params ...string) {
		if len(params) == 0 {
			fmt.Println(CmdList())
			Version()
		} else {
			if c, ok := cmdDocs[params[0]]; ok {
				fmt.Println("Usage: ", c[0])
			} else {
				fmt.Println(CmdList())
				Version()
			}
		}
	}
}

func CmdList() string {
	helpAll := fmt.Sprintf("Srun-cmd %s\r\n\r\n", config.Version)
	helpAll += "Options:\r\n"
	for k, v := range optionDocs {
		helpAll += fmt.Sprintf("\t%-20s%-20s\r\n", k, v)
	}
	helpAll += "\r\n"
	helpAll += "Commands:\r\n"
	for k, v := range cmdDocs {
		helpAll += fmt.Sprintf("\t%-20s%-20s\r\n", k, v[1])
	}
	return helpAll
}
