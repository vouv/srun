package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var errParse = errors.New("error-parse")

// generate callback function string
func genCallback() string {
	return fmt.Sprintf("jsonp%d", int(time.Now().Unix()))
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
