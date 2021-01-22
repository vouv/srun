package model

import (
	"fmt"
	"strings"

	"github.com/vouv/srun/utils"
)

type ChallengeResp struct {
	Challenge string `json:"challenge"`
	ClientIp  string `json:"client_ip"`
}

type ActionResp struct {
	Res      string      `json:"res"`
	Error    string      `json:"error"`
	Ecode    interface{} `json:"ecode"`
	ErrorMsg string      `json:"error_msg"`
	ClientIp string      `json:"client_ip"`
}

type InfoResp struct {
	ServerFlag    int64   `json:"ServerFlag"`
	AddTime       int64   `json:"add_time"`
	AllBytes      int64   `json:"all_bytes"`
	BytesIn       int64   `json:"bytes_in"`
	BytesOut      int64   `json:"bytes_out"`
	CheckoutDate  int64   `json:"checkout_date"`
	Domain        string  `json:"domain"`
	Error         string  `json:"error"`
	GroupID       string  `json:"group_id"`
	KeepaliveTime int64   `json:"keepalive_time"`
	OnlineIP      string  `json:"online_ip"`
	ProductsName  string  `json:"products_name"`
	RealName      string  `json:"real_name"`
	RemainBytes   int64   `json:"remain_bytes"`
	RemainSeconds int64   `json:"remain_seconds"`
	SumBytes      int64   `json:"sum_bytes"`
	SumSeconds    int64   `json:"sum_seconds"`
	UserBalance   float64 `json:"user_balance"`
	UserCharge    float64 `json:"user_charge"`
	UserMac       string  `json:"user_mac"`
	UserName      string  `json:"user_name"`
	WalletBalance float64 `json:"wallet_balance"`
}

func (r *InfoResp) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf(" 在线IP: %s\n", r.OnlineIP))
	sb.WriteString(fmt.Sprintf("上网账号: %s\n", r.UserName))
	sb.WriteString(fmt.Sprintf("电子钱包: ￥%.2f\n", r.WalletBalance))
	sb.WriteString(fmt.Sprintf("套餐余额: ￥%.2f\n", r.UserBalance))
	sb.WriteString(fmt.Sprintf("已用流量: %s\n", utils.FormatFlux(r.SumBytes)))
	sb.WriteString(fmt.Sprintf("在线时长: %s\n", utils.FormatTime(r.SumSeconds)))

	return sb.String()
}

type InfoResult struct {
	Acid        int    `json:"ac_id"`
	Username    string `json:"username"`
	ClientIp    string `json:"client_ip"`
	AccessToken string `json:"access_token"`
}
