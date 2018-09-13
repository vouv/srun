package srun

import (
	"encoding/base64"
	"fmt"
	"time"
	"net/http"
	"os"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

const (
	get_challenge_url = "http://10.0.0.55/cgi-bin/get_challenge"
	srun_portal_url = "http://10.0.0.55/cgi-bin/srun_portal"
	url = "http://10.0.0.55"
)

func data_info(get_data map[string]interface{}, token string) string  {
	x_encode_json := map[string]interface{} {
		"username": get_data["username"],
		"password": get_data["password"],
		"ip": get_data["ip"],
		"acid": get_data["ac_id"],
		"enc_ver": "srun_bx1",
	}
	x_encode_raw, err := json.Marshal(x_encode_json);
	if err != nil {
		logs.Error(err)
		return ""
	}
	xen := string(x_encode_raw)
	x_encode_res := X_encode(xen, token)
	dict_key := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="
	dict_val := "LVoJPiCN2R8G90yg+hmFHuacZ1OWMnrsSTXkYpUq/3dlbfKwv6xztjI7DeBE45QA="
	dict := map[string]string{}
	for idx, v := range dict_key {
		dict[string(v)] = dict_val[idx:idx+1]
	}
	b64_arr := []byte{}
	for _,c := range x_encode_res {
		b64_arr = append(b64_arr, byte(c))
	}
	b64_res := base64.StdEncoding.EncodeToString(b64_arr)
	target := ""
	for _, s := range b64_res {
		target += dict[string(s)]
	}
	return "{SRBX1}" + target
}

func getJson(url string, data map[string]interface{}) (res map[string]interface{}) {
	callback := fmt.Sprintf("jsonp%s", int(time.Now().Unix()))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	q := req.URL.Query()
	q.Add("callback", callback)
	for k,v := range data {
		if val,ok := v.(int); ok{
			q.Add(k, strconv.Itoa(val))
		}
		if val,ok := v.(string); ok{
			q.Add(k, val)
		}
	}
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logs.Error(err)
		return nil
	}
	defer resp.Body.Close()
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err)
		return nil
	}
	dt := string(raw)[len(callback)+1: len(raw)-1]
	if err := json.Unmarshal([]byte(dt), &res); err != nil {
		logs.Error(err)
		return nil
	}
	return res
}


func Login(username, password string)  {
	get_data := NewLoginData(username, password)
	//	get token
	req := map[string]interface{}{
		"username": username,
	}
	challenge_json := getJson(get_challenge_url, req)
	token := challenge_json["challenge"].(string)
	client_ip := challenge_json["client_ip"]
	get_data["ip"] = client_ip.(string)
	info := data_info(get_data, token)
	get_data["info"] = info
	get_data["password"] = Pwd_hmd5("", token)
	get_data["chksum"] = Checksum(get_data, token)
	res := getJson(srun_portal_url, get_data)
	if res["res"] == "ok" {
		fmt.Println("login success!")
	}
}
