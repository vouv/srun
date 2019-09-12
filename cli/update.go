package cli

import (
	"context"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"srun/config"
	"strings"
	"time"
)

const url = "https://github.com/monigo/srun-cmd/releases/latest"
const timeOut = 3 * time.Second

var client = http.Transport{
	DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
		conn, err := net.DialTimeout(network, addr, timeOut)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(timeOut))
		return conn, nil
	},
}

func HasUpdate() (ok bool, version string, dist string) {
	version = config.Version
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Debug("update request error", err)
		return
	}
	res, err := client.RoundTrip(req)

	if err != nil {
		log.Debug("update request error", err)
		return
	}
	dist = res.Header.Get("Location")
	arr := strings.Split(dist, "/")
	version = arr[len(arr)-1]
	//fmt.Println(version)
	ok = version > config.Version
	return

}

func Update() Func {
	return func(cmd string, params ...string) {
		ok, v, d := HasUpdate()
		if !ok {
			log.Info("当前已是最新版本:", config.Version)
			return
		}
		log.Info("发现新版本: ", v)
		log.Info("打开链接下载: ", d)
	}
}
