package cli

import (
	"context"
	"github.com/astaxie/beego/logs"
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

func HasUpdate() (bool, string) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		logs.Debug("update request error", err)
		return false, ""
	}
	res, err := client.RoundTrip(req)

	if err != nil {
		logs.Debug("update request error", err)
		return false, ""
	}
	dist := res.Header.Get("Location")
	arr := strings.Split(dist, "/")
	version := arr[len(arr)-1]
	//fmt.Println(version)
	return version > config.Version, dist

}
