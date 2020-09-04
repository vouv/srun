package core

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"github.com/vouv/srun/hash"
	"github.com/vouv/srun/model"
	"github.com/vouv/srun/resp"
	"github.com/vouv/srun/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

const (
	demoUrl = "http://t.cn"

	challengeUrl = "http://10.0.0.55/cgi-bin/get_challenge"
	portalUrl    = "http://10.0.0.55/cgi-bin/srun_portal"

	succeedUrlOrigin = "http://10.0.0.55/srun_portal_pc_succeed.php"
)

// api Login
// step 1: get acid
// step 2: get challenge
// step 3: do login
func Login(account *model.Account) (result model.QInfo, err error) {
	log.Debug("Username: ", account.Username)
	// 先获取acid
	// 并检查是否已经联网
	acid, err := getAcid()
	if err != nil {
		log.Debug("get acid eror:", err)
		err = ErrConnected
		return
	}

	username := account.GenUsername()
	// 创建登录表单
	formLogin := model.Login(username, account.Password, acid)

	//	get token
	rc, err := getChallenge(username)
	if err != nil {
		log.Debug("get challenge error", err)
		err = ErrRequest
		return
	}

	token := rc.Challenge
	ip := rc.ClientIp

	formLogin.Set("ip", ip)
	formLogin.Set("info", hash.GenInfo(formLogin, token))
	formLogin.Set("password", hash.PwdHmd5("", token))
	formLogin.Set("chksum", hash.Checksum(formLogin, token))

	// response
	ra := resp.RAction{}

	if err = utils.GetJson(portalUrl, formLogin, &ra); err != nil {
		log.Debug("request error", err)
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

	result = model.QInfo{
		Acid:        acid,
		Username:    username,
		ClientIp:    rc.ClientIp,
		AccessToken: rc.Challenge,
	}
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
	err = parseHtml(succeedUrlOrigin, qs)
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

var reg, _ = regexp.Compile(`index_[\d]\.html`)

// get acid
func getAcid() (acid int, err error) {
	res, err := http.Get(demoUrl)
	if err != nil {
		return 1, ErrConnected
	}
	bs, _ := ioutil.ReadAll(res.Body)
	data := string(bs)
	if strings.Contains(data, "10.0.0.5") && reg.MatchString(data) {
		res := reg.FindString(data)
		acids := strings.TrimRight(strings.TrimLeft(res, "index_"), ".html")
		acid, _ = strconv.Atoi(acids)
		log.Debug("Acid:", acid)
		return acid, nil
	}
	return 1, ErrConnected
}

func getChallenge(username string) (res resp.Challenge, err error) {
	qc := model.Challenge(username)
	err = utils.GetJson(challengeUrl, qc, &res)
	return
}

// get the info page and parse the html code
func parseHtml(url string, data url.Values) (err error) {
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
