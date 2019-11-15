package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const demoUrl = "http://t.cn"

var errParse = errors.New("error-parse")
var errNotLogin = errors.New("not-login")
var errLogin = errors.New("已经联网")

var reg, _ = regexp.Compile(`index_[\d]\.html`)

// generate callback function string
func genCallback() string {
	return fmt.Sprintf("jsonp%d", int(time.Now().Unix()))
}

// get acid
func GetAcid() (acid int, err error) {
	acid = 8
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if strings.Contains(req.URL.String(), "10.0.0.5") {
				// get acid
				if reg.MatchString(req.URL.String()) {
					res := reg.FindString(req.URL.String())
					acids := strings.TrimRight(strings.TrimLeft(res, "index_"), ".html")
					acid, err = strconv.Atoi(acids)
					if err != nil {
						acid = 8
					}
					return errNotLogin
				}
			}
			return nil
		},
	}

	req, err := http.NewRequest(http.MethodGet, demoUrl, nil)
	if err != nil {
		log.Error(err)
		return acid, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return acid, nil
	} else {
		_ = resp.Body.Close()
		return acid, errLogin
	}
}

// make request with data
func DoRequest(url string, params url.Values) (*http.Response, error) {

	// add callback
	params.Add("callback", genCallback())
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Debug(err)
		return nil, err
	}
	//req.AddCookie(&http.Cookie{Cmd: "username", Value: params.Get("username"), HttpOnly: true})
	req.URL.RawQuery = params.Encode()
	client := http.DefaultClient
	client.Timeout = time.Second * 3

	resp, err := client.Do(req)
	if err != nil {
		log.Error("network error")
		log.Debug(err)
		return nil, err
	}
	return resp, nil
}

// request for login and get json response
func GetJson(url string, data url.Values, res interface{}) (err error) {
	resp, err := DoRequest(url, data)
	if err != nil {
		log.Error("network error")
		log.Debug(err)
		return
	}
	defer resp.Body.Close()
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("network error")
		log.Debug(err)
		return
	}
	rawStr := string(raw)

	// cut jsonp
	start := strings.Index(rawStr, "(")
	end := strings.LastIndex(rawStr, ")")
	if start == -1 && end == -1 {
		log.Error(rawStr)
		return errParse
	}
	dt := string(raw)[start+1 : end]

	if err = json.Unmarshal([]byte(dt), &res); err != nil {
		return
	}
	return nil
}
