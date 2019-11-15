package srun

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/monigo/srun-cmd/hash"
	"github.com/monigo/srun-cmd/model"
	"github.com/monigo/srun-cmd/resp"
	"github.com/monigo/srun-cmd/utils"
	log "github.com/sirupsen/logrus"
	"net/url"
)

const (
	challengeUrl       = "http://10.0.0.55/cgi-bin/get_challenge"
	portalUrl          = "http://10.0.0.55/cgi-bin/srun_portal"
	succeedUrl         = "http://10.0.0.55/srun_portal_pc_succeed.php"
	succeedUrlYidong   = "http://10.0.0.55/srun_portal_pc_succeed_yys.php"
	succeedUrlLiantong = "srun_portal_pc_succeed_yys_cucc.php"
)

// api Login
// step 1: get acid
// step 2: get challenge
// step 3: do login
func Login(username, password string) (info model.QInfo, err error) {
	// 先获取acid
	// 并检查是否已经联网
	acid, err := utils.GetAcid()
	if err != nil {
		log.Debug(err)
		err = ErrConnected
		return
	}

	// 创建登录表单
	formLogin := model.Login(username, password, acid)

	//	get token
	qc := model.Challenge(username)

	rc := resp.Challenge{}
	if err = utils.GetJson(challengeUrl, qc, &rc); err != nil {
		log.Debug(err)
		err = ErrRequest
		return
	}

	info.AccessToken = rc.Challenge
	info.ClientIp = rc.ClientIp

	formLogin.Set("ip", info.ClientIp)
	formLogin.Set("info", hash.GenInfo(formLogin, info.AccessToken))
	formLogin.Set("password", hash.PwdHmd5("", info.AccessToken))
	formLogin.Set("chksum", hash.Checksum(formLogin, info.AccessToken))

	// response
	ra := resp.RAction{}

	if err = utils.GetJson(portalUrl, formLogin, &ra); err != nil {
		log.Debug(err)
		err = ErrRequest
		return
	}
	if ra.Res != "ok" {
		msg := ra.Res
		if msg == "" {
			msg = ra.ErrorMsg
		}
		log.Debug("登录失败: ", msg)
		log.Debug(ra)
		err = ErrFailed
		return
	}
	// 余量查询
	return
}

// api info
func Info(account model.Account) (err error) {
	qs := model.Info(
		1,
		account.Username,
		account.Ip,
		account.AccessToken,
	)
	// 余量查询
	log.Info("服务器: ", account.Server)
	switch account.Server {
	case model.ServerTypeCMCC:
		err = ParseHtml(succeedUrlYidong, qs)
	case model.ServerTypeWCDMA:
		err = ParseHtml(succeedUrlLiantong, qs)
	default:
		err = ParseHtml(succeedUrl, qs)
	}
	return
}

// api logout
func Logout(username string) (err error) {
	q := model.Logout(username)
	ra := resp.RAction{}
	if err = utils.GetJson(portalUrl, q, &ra); err != nil {
		log.Debug(err)
		err = ErrRequest
		return
	}
	if ra.Error != "ok" {
		log.Debug(ra)
		err = ErrRequest
	}
	return
}

// get the info page and parse the html code
func ParseHtml(url string, data url.Values) (err error) {
	resp, err := utils.DoRequest(url, data)
	if err != nil {
		log.Debug(err)
		err = ErrRequest
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Error(err)
		err = ErrRequest
		return
	}

	// find the items
	bytes := doc.Find("span#sum_bytes").Last().Text()
	times := doc.Find("span#sum_seconds").Text()
	balance := doc.Find("span#user_balance").Text()
	fmt.Println("已用流量:", bytes)
	fmt.Println("已用时长:", times)
	fmt.Println("账户余额:", balance)
	return
}
