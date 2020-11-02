package model

import (
	"fmt"
	"github.com/vouv/srun/utils"
	"net/url"
	"strings"
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
	ServerFlag    int64   `json:"ServerFlag"`
	AddTime       int64   `json:"add_time"`
	AllBytes      int64   `json:"all_bytes"`
	BytesIn       int64   `json:"bytes_in"`
	BytesOut      int64   `json:"bytes_out"`
	CheckoutDate  int64   `json:"checkout_date"`
	Domain        string  `json:"domain"`
	Error         string  `json:"error"`
	GroupId       string  `json:"group_id"`
	KeepaliveTime int64   `json:"keepalive_time"`
	OnlineIp      string  `json:"online_ip"`
	ProductsName  string  `json:"products_name"`
	RealName      string  `json:"real_name"`
	RemainBytes   int64   `json:"remain_bytes"`
	RemainSeconds int64   `json:"remain_seconds"`
	SumBytes      int64   `json:"sum_bytes"`
	SumSeconds    int64   `json:"sum_seconds"`
	UserBalance   float64 `json:"user_balance"`
	UserCharge    int     `json:"user_charge"`
	UserMac       string  `json:"user_mac"`
	UserName      string  `json:"user_name"`
	WalletBalance float64 `json:"wallet_balance"`
}

func (r *RInfo) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("用户名: %s\n", r.UserName))
	sb.WriteString(fmt.Sprintf("IP地址: %s\n", r.OnlineIp))
	sb.WriteString(fmt.Sprintf("已用流量: %s\n", utils.FormatFlux(r.SumBytes)))
	sb.WriteString(fmt.Sprintf("已用时长: %s\n", utils.FormatTime(r.SumSeconds)))
	sb.WriteString(fmt.Sprintf("账户余额: ￥%.2f\n", r.UserBalance))

	return sb.String()
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
