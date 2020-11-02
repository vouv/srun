package model

import (
	"fmt"
	"net/url"
)

func Challenge(username string) url.Values {
	return url.Values{
		"username": {username},
		"ip":       {""},
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
