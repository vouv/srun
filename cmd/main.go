package main

import (
	"github.com/monigo/srun-cmd/cli"
)

func main() {
	// set cpu count
	//runtime.GOMAXPROCS(runtime.NumCPU())

	cmd := cli.New()
	cmd.Handle()

	// has update
	// todo 修改更新逻辑, 减少更新频率
	//if ok, url := cli.HasUpdate(); ok {
	//	fmt.Print("更新: " + url)
	//	fmt.Println(" 当前版本: " + config.Version)
	//}

}
