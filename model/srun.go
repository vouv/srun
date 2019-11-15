package model

import (
	"fmt"
	"net/url"
)

// query challenge
type QChallenge struct {
	Username string `json:"username"`
	Ip       string `json:"ip"`
}

// query login
type QLogin struct {
	Action   string `json:"action"`
	Username string `json:"username"`
	Password string `json:"password"`
	Acid     int    `json:"ac_id"`
	Ip       string `json:"ip"`
	Info     string `json:"info"`
	Chksum   string `json:"chksum"`
	N        int    `json:"n"`
	Type     int    `json:"type"`
}

// query info
type QInfo struct {
	Acid        int    `json:"ac_id"`
	Username    string `json:"username"`
	ClientIp    string `json:"client_ip"`
	AccessToken string `json:"access_token"`
}

// query logout
type QLogout struct {
	Action   string `json:"action"`
	Username string `json:"username"`
	Acid     int    `json:"ac_id"`
	Ip       string `json:"ip"`
}

func Challenge(username string) url.Values {
	return url.Values{
		"username": {username},
		"ip":       {""},
	}
}

func Info(acid int, username, clientIp, accessToken string) url.Values {
	return url.Values{
		"ac_id":        {fmt.Sprint(acid)},
		"username":     {username},
		"client_ip":    {clientIp},
		"access_token": {accessToken},
	}
}

func Login(username, password string, acid int) url.Values {
	return url.Values{
		"action":   {"login"},
		"username": {username},
		"password": {password},
		"ac_id":    {fmt.Sprint(acid)},
		"ip":       {""},
		"info":     {},
		"chksum":   {},
		"n":        {"200"},
		"type":     {"1"},
	}
}

func Logout(username string) url.Values {
	return url.Values{
		"action":   {"logout"},
		"username": {username},
	}
}

//func NewLoginData(username, password string) map[string]interface{} {
//	return map[string]interface{}{
//		"action": "login",
//		"username": username,
//		"password": password,
//		"ac_id": 1,
//		"ip": "",
//		"info": "",
//		"chksum": "",
//		"n": 200,
//		"type": 1,
//	}
//}
