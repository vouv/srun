package core

import (
	"bufio"
	"net"
	"net/http"
	"time"
)

func get(addr string) (*http.Response, error) {
	req, _ := http.NewRequest(http.MethodGet, addr, nil)
	return request(req)
}

func request(req *http.Request) (*http.Response, error) {
	conn, err := net.DialTimeout("tcp", req.URL.Hostname()+":http", time.Second)
	if err != nil {
		return nil, err
	}
	_ = req.Write(conn)
	return http.ReadResponse(bufio.NewReader(conn), req)
}
