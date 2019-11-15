package model

import (
	"encoding/json"
	"fmt"
)

const (
	ServerTypeCMCC  = "移动"
	ServerTypeWCDMA = "联通"
	ServerTypeSrun  = "校园网"

	suffixCMCC  = "@yidong"
	suffixWCDMA = "@liantong"
)

type Account struct {
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	AccessToken string `json:"access_token"`
	Ip          string `json:"ip"`
	Server      string `json:"server"`
}

func (acc *Account) JSONString() (jsonStr string, err error) {
	jsonData, err := json.Marshal(acc)
	if err != nil {
		return
	}
	jsonStr = string(jsonData)
	return
}

func (acc *Account) String() string {
	return fmt.Sprintln("用户名:", acc.Username)
}

func AddSuffix(name, server string) string {
	switch server {
	case ServerTypeCMCC:
		return name + suffixCMCC
	case ServerTypeWCDMA:
		return name + suffixWCDMA
	default:
		return name
	}
}
