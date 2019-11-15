package store

import (
	"encoding/base64"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/user"
)

// 写入账号信息到文件
// 统一错误
func SetAccount(username, password, def string) (err error) {
	accountFilename, err := getAccountFilename()
	if err != nil {
		log.Debugf("打开账号文件错误, %s", err)
		err = ErrReadFile
		return
	}
	file, openErr := os.OpenFile(accountFilename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if openErr != nil {
		log.Debugf("打开账号文件错误, %s", openErr)
		err = ErrReadFile
		return
	}
	defer file.Close()

	//write to local dir
	var account Account
	account.Username = b64Encode(username)
	account.Password = b64Encode(password)
	account.Server = def

	jsonStr, mErr := account.toJSON()
	if mErr != nil {
		log.Debugf("序列化账号错误, %s", mErr)
		err = ErrParse
		return
	}
	_, wErr := file.WriteString(jsonStr)
	if wErr != nil {
		log.Debugf("写入账号文件错误, %s", wErr)
		err = ErrWriteFile
		return
	}

	return
}

// 保存token
func SetInfo(token, ip string) (err error) {
	//write to local dir
	account, err := LoadAccount()
	if err != nil {
		return
	}

	accountFilename, err := getAccountFilename()
	if err != nil {
		log.Debugf("打开账号文件错误, %s", err)
		err = ErrReadFile
		return
	}
	file, openErr := os.OpenFile(accountFilename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if openErr != nil {
		log.Debugf("打开账号文件错误, %s", openErr)
		err = ErrReadFile
		return
	}
	defer file.Close()
	account.Username = b64Encode(account.Username)
	account.Password = b64Encode(account.Password)
	account.AccessToken = token
	account.Ip = ip

	jsonStr, mErr := account.toJSON()
	if mErr != nil {
		log.Debugf("序列化账号错误, %s", mErr)
		err = ErrParse
		return
	}
	_, wErr := file.WriteString(jsonStr)
	if wErr != nil {
		log.Debugf("写入账号文件错误, %s", wErr)
		err = ErrWriteFile
		return
	}
	return
}

func LoadAccount() (account Account, err error) {
	accountFilename, err := getAccountFilename()
	if err != nil {
		log.Debug(err)
		err = ErrReadFile
		return
	}
	file, openErr := os.Open(accountFilename)
	if openErr != nil {
		log.Debugf("打开账号文件错误, %s,", openErr)
		err = ErrReadFile
		return
	}
	defer file.Close()

	accountBytes, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		log.Debugf("读取账号文件错误, %s", readErr)
		err = ErrReadFile
		return
	}

	if umError := json.Unmarshal(accountBytes, &account); umError != nil {
		log.Debugf("解析账号文件错误, %s", umError)
		err = ErrReadFile
		return
	}
	account.Username, err = b64Decode(account.Username)
	if err != nil {
		err = ErrReadFile
	}
	account.Password, err = b64Decode(account.Password)
	if err != nil {
		err = ErrReadFile
	}
	log.Debug("读取账号信息: ", accountFilename)
	return
}

func b64Decode(b64 string) (res string, err error) {
	bs, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return
	}
	res = string(bs)
	return
}

func b64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func init() {
	curUser, gErr := user.Current()
	if gErr != nil {
		log.Fatalln("无法读取账户信息, 请重新设置账户信息")
	}
	RootPath = curUser.HomeDir
}
