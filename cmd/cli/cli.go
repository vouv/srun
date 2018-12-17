package cli

import (
	"fmt"
	"login-srun/cmd/srun"
	"os"
)

type CliFunc func(cmd string, params ...string)


func AccountH(cmd string, params ...string) {
	if len(params) == 0 {
		account, gErr := GetAccount()
		if gErr != nil {
			fmt.Println(gErr)
			os.Exit(1)
		}
		account.Password = Encode(account.Password)
		fmt.Println(account.String())
	} else if len(params) == 2{
		accessKey := params[0]
		secretKey := params[1]
		sErr := SetAccount(accessKey, secretKey)
		if sErr != nil {
			fmt.Println(sErr)
			os.Exit(1)
		}
		fmt.Println("账号密码已被保存")
	} else {
		CmdHelp(cmd)
	}
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