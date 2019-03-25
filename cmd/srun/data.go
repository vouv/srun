package srun

// query challenge
type QChallenge struct {
	Username string `json:"username"`
	Ip       string `json:"ip"`
}

// response challenge
type RChallenge struct {
	Challenge string `json:"challenge"`
	ClientIp  string `json:"client_ip"`
}

// query login
type QLogin struct {
	Action   string `json:"action"`
	Username string `json:"username"`
	Password string `json:"password"`
	Acid     int    `json:"ac_id"`
	Ip       string `json:"ip"`
	Info     string `json:"info"`
	Chksum   string `json:"chksum"`
	N        int    `json:"n"`
	Type     int    `json:"type"`
}

// query info
type QInfo struct {
	Acid        int    `json:"ac_id"`
	Username    string `json:"username"`
	ClientIp    string `json:"client_ip"`
	AccessToken string `json:"access_token"`
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

// query logout
type QLogout struct {
	Action   string `json:"action"`
	Username string `json:"username"`
	Acid     int    `json:"ac_id"`
	Ip       string `json:"ip"`
}

func NewQChallenge(username string) QChallenge {
	return QChallenge{
		Username: username,
		Ip:       "",
	}
}

func NewQLogin(username, password string) QLogin {
	return QLogin{
		Action:   "login",
		Username: username,
		Password: password,
		Acid:     8,
		Ip:       "",
		Info:     "",
		Chksum:   "",
		N:        200,
		Type:     1,
	}
}

func NewQLogout(username, password string) QLogout {
	return QLogout{
		Username: username,
	}
}

//func NewLoginData(username, password string) map[string]interface{} {
//	return map[string]interface{}{
//		"action": "login",
//		"username": username,
//		"password": password,
//		"ac_id": 1,
//		"ip": "",
//		"info": "",
//		"chksum": "",
//		"n": 200,
//		"type": 1,
//	}
//}
