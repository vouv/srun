package cli

import (
	"github.com/astaxie/beego/logs"
	"net/http"
	"srun-cmd/config"
	"strings"
)

const url = "https://github.com/monigo/srun-cmd/releases/latest"

var client = http.Transport{}

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
