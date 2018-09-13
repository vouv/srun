package srun

func NewLoginData(username, password string) map[string]interface{} {
	return map[string]interface{}{
		"action": "login",
		"username": username,
		"password": password,
		"ac_id": 1,
		"ip": "",
		"info": "",
		"chksum": "",
		"n": 200,
		"type": 1,
	}
}