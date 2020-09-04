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

type RInfo struct {
	Bytes   string `json:"bytes"`   // 已用流量
	Times   string `json:"times"`   // 已用时长
	Balance string `json:"balance"` // 余额
}

func (r *RInfo) String() string {
	return fmt.Sprintf("已用流量: %s\n已用时长: %s\n账户余额: %s\n", r.Bytes, r.Times, r.Balance)
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
