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

func getJson(url string, data interface{}, res interface{}) (err error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}


	q := req.URL.Query()
	callback := fmt.Sprintf("jsonp%d", int(time.Now().Unix()))
	q.Add("callback", callback)

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
		return
	}
	defer resp.Body.Close()
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	dt := string(raw)[len(callback)+1: len(raw)-1]
	if err = json.Unmarshal([]byte(dt), &res); err != nil {
		return
	}
	return nil
}

func getHtml(url string, data interface{}) (res string, err error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	q := req.URL.Query()
	callback := fmt.Sprintf("jsonp%d", int(time.Now().Unix()))
	q.Add("callback", callback)

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

	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return string(raw), nil
}


func parseHtml(url string, data interface{}) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	q := req.URL.Query()
	callback := fmt.Sprintf("jsonp%d", int(time.Now().Unix()))
	q.Add("callback", callback)

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

	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logs.Error(err)
		return
	}

	// Find the review items
	bytes := doc.Find("span#sum_bytes").Last().Text()
	times := doc.Find("span#sum_seconds").Text()
	balance := doc.Find("span#user_balance").Text()
	fmt.Println("已用流量:", bytes)
	fmt.Println("已用时长:", times)
	fmt.Println("账户余额:", balance)
	return
}