package cli

import (
	"fmt"
	"os"
	"login-srun/srun"
)

type CliFunc func(cmd string, params ...string)


func AccountHandle(cmd string, params ...string) {
	if len(params) == 0 {
		account, gErr := GetAccount()
		if gErr != nil {
			fmt.Println(gErr)
			os.Exit(1)
		}
		fmt.Println(account.String())
	} else if len(params) == 2{
		accessKey := params[0]
		secretKey := params[1]
		sErr := SetAccount(accessKey, secretKey)
		if sErr != nil {
			fmt.Println(sErr)
			os.Exit(1)
		}
	} else {
		CmdHelp(cmd)
	}
}

func LoginHandle(cmd string, params ...string) {
	if len(params) == 0 {
		account, gErr := GetAccount()
		if gErr != nil {
			fmt.Println(gErr)
			os.Exit(1)
		}
		srun.Login(account.Username, account.Password)
	} else {
		CmdHelp(cmd)
	}
}