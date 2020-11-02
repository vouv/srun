package model

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	AccessToken string `json:"access_token"`
	Acid        int    `json:"acid"`
}

func (a *Account) JSONString() (jsonStr string, err error) {
	jsonData, err := json.Marshal(a)
	if err != nil {
		return
	}
	jsonStr = string(jsonData)
	return
}

func (a *Account) JSONBytes() (jsonData []byte, err error) {
	return json.Marshal(a)
}

func (a *Account) String() string {
	return fmt.Sprintln("用户名:", a.Username)
}
