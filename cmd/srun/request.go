package srun

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// generate callback function string
func genCallback() string {
	return fmt.Sprintf("jsonp%d", int(time.Now().Unix()))
}

// make get request with data
// returns http.Response
func doGet(url string, data interface{}) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logs.Debug(err)
		return nil, err
	}

	q := req.URL.Query()
	q.Add("callback", genCallback())

	jd, _ := json.Marshal(data)
	md := map[string]interface{}{}
	json.Unmarshal(jd, &md)

	for k,v := range md {
		if val,ok := v.(float64); ok{
			q.Add(k, strconv.Itoa(int(val)))
		}
		if val,ok := v.(string); ok{
			q.Add(k, val)
		}
	}

	req.AddCookie(&http.Cookie{Name: "username",Value: q.Get("username"), HttpOnly: true})

	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logs.Error("network error")
		logs.Debug(err)
		return nil, err
	}
	return resp, nil
}

// request for login and get json response
func getJson(url string, data interface{}, res interface{}) (err error) {

	resp, err := doGet(url, data)
	if err != nil {
		logs.Error("network error")
		logs.Debug(err)
		return
	}
	defer resp.Body.Close()
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("network error")
		logs.Debug(err)
		return
	}
	dt := string(raw)[len(genCallback())+1: len(raw)-1]
	if err = json.Unmarshal([]byte(dt), &res); err != nil {
		return
	}
	return nil
}

// get the info page and parse the html code
func parseHtml(url string, data interface{}) {
	resp, err := doGet(url, data)
	if err != nil {
		logs.Error("network error")
		logs.Debug(err)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logs.Error(err)
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