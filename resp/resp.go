package resp

// response challenge
type Challenge struct {
	Challenge string `json:"challenge"`
	ClientIp  string `json:"client_ip"`
}

// example
// map[res:login_error
// srun_ver:SRunCGIAuthIntfSvr V1.01 B20180306
// st:1.543044956e+09
// client_ip:10.62.44.153
// ecode:E2616
// error:login_error
// error_msg:E2616: Average users.
// online_ip:10.62.44.153]
type RAction struct {
	Res      string      `json:"res"`
	Error    string      `json:"error"`
	Ecode    interface{} `json:"ecode"`
	ErrorMsg string      `json:"error_msg"`
	ClientIp string      `json:"client_ip"`
}
