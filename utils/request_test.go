package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func Test_doGet(t *testing.T) {

	req, err := http.NewRequest(http.MethodGet, "http://baidu.com", nil)
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.URL.String())
			if strings.Contains(req.URL.String(), "http://10.0.0.55") {
				return errors.New("check")
			}
			return nil
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp.Status)
	raw, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(raw))
}
