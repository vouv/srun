package main

import (
	"context"
	"net"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const repo = "https://github.com/vouv/srun/releases/latest"
const updateTimeout = 3 * time.Second

var client = http.Transport{
	DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
		conn, err := net.DialTimeout(network, addr, updateTimeout)
		if err != nil {
			return nil, err
		}
		_ = conn.SetDeadline(time.Now().Add(updateTimeout))
		return conn, nil
	},
}

func HasUpdate() (ok bool, version string, dist string) {
	req, err := http.NewRequest("GET", repo, nil)

	if err != nil {
		log.Debug("请求错误", err)
		return
	}
	res, err := client.RoundTrip(req)
	if err != nil {
		log.Debug("请求错误", err)
		return
	}
	dist = res.Header.Get("Location")
	arr := strings.Split(dist, "/")
	version = arr[len(arr)-1]

	log.Debug("最新版本", version)

	ok = version != Version
	return

}
