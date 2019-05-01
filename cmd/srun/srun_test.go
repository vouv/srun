package srun

import (
	"fmt"
	"testing"
)

func Test_genInfo(t *testing.T) {

	ql := QLogin{
		Username: "awlsx@yidong",
		Password: "Lsx767400405",
		Acid:     1,
		Ip:       "10.63.63.95",
	}
	fmt.Println(genInfo(ql, "274414592c076f62250947f3af02a0254514cf1839db5d7074f9b4e4a90f3669"))
}

func TestLogin(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name      string
		args      args
		wantToken string
		wantIp    string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotToken, gotIp := Login(tt.args.username, tt.args.password)
			if gotToken != tt.wantToken {
				t.Errorf("Login() gotToken = %v, want %v", gotToken, tt.wantToken)
			}
			if gotIp != tt.wantIp {
				t.Errorf("Login() gotIp = %v, want %v", gotIp, tt.wantIp)
			}
		})
	}
}

func TestInfo(t *testing.T) {
	type args struct {
		username string
		token    string
		ip       string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.username, tt.args.token, tt.args.ip)
		})
	}
}

func TestLogout(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Logout(tt.args.username)
		})
	}
}
