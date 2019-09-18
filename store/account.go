package store

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

var RootPath string

type Account struct {
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	AccessToken string `json:"access_token"`
	Ip          string `json:"ip"`
	Server      string `json:"server"`
}

func (acc *Account) ToJson() (jsonStr string, err error) {
	jsonData, mErr := json.Marshal(acc)
	if mErr != nil {
		err = fmt.Errorf("marshal account data error, %s", mErr)
		return
	}
	jsonStr = string(jsonData)
	return
}

func (acc *Account) String() string {
	return fmt.Sprintln("用户名:", acc.Username)
}

func getAccountFilename() (fileSrc string, err error) {
	storageDir := filepath.Join(RootPath, ".srun")
	if _, sErr := os.Stat(storageDir); sErr != nil {
		if mErr := os.MkdirAll(storageDir, 0755); mErr != nil {
			err = fmt.Errorf("mkdir `%s` error, %s", storageDir, mErr)
			return
		}
	}
	fileSrc = filepath.Join(storageDir, "account.json")
	return
}

func init() {
	curUser, gErr := user.Current()
	if gErr != nil {
		log.Error("无法读取账户信息, 请重新设置账户信息")
		os.Exit(1)
	}

	RootPath = curUser.HomeDir
}

//写入账号信息到文件
func SetAccount(username, password, def string) (err error) {

	accountFilename, err := getAccountFilename()
	if err != nil {
		log.Error(err)
		return
	}
	file, openErr := os.OpenFile(accountFilename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if openErr != nil {
		err = fmt.Errorf("open account file error, %s", openErr)
		return
	}
	defer file.Close()

	//write to local dir
	var account Account
	account.Username = b64Encode(username)
	account.Password = b64Encode(password)
	account.Server = def

	jsonStr, mErr := account.ToJson()
	if mErr != nil {
		err = mErr
		return
	}
	_, wErr := file.WriteString(jsonStr)
	if wErr != nil {
		err = fmt.Errorf("write account info error, %s", wErr)
		return
	}

	return
}

func SetInfo(token, ip string) (err error) {
	//write to local dir
	account, err := LoadAccount()
	if err != nil {
		log.Error(err)
		return
	}
	accountFilename, err := getAccountFilename()
	if err != nil {
		log.Error(err)
		return
	}
	file, openErr := os.OpenFile(accountFilename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if openErr != nil {
		err = fmt.Errorf("open account file error, %s", openErr)
		return
	}
	defer file.Close()
	account.Username = b64Encode(account.Username)
	account.Password = b64Encode(account.Password)
	account.AccessToken = token
	account.Ip = ip

	jsonStr, mErr := account.ToJson()
	if mErr != nil {
		err = mErr
		return
	}
	_, wErr := file.WriteString(jsonStr)
	if wErr != nil {
		err = fmt.Errorf("write account info error, %s", wErr)
		return
	}

	return
}

func LoadAccount() (account Account, err error) {

	accountFilename, err := getAccountFilename()
	if err != nil {
		log.Error(err)
		return
	}

	file, openErr := os.Open(accountFilename)
	if openErr != nil {
		err = fmt.Errorf("Open account file error, %s, please use `account` to set Username and Password first. ", openErr)
		return
	}
	defer file.Close()

	accountBytes, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		err = fmt.Errorf("read account file error, %s. ", readErr)
		return
	}

	if umError := json.Unmarshal(accountBytes, &account); umError != nil {
		err = fmt.Errorf("parse account file error, %s. ", umError)
		return
	}
	account.Username = b64Decode(account.Username)
	account.Password = b64Decode(account.Password)

	log.Debug("load account from ", accountFilename)
	return
}

func b64Decode(b64 string) string {
	bs, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		log.Error(err)
	}
	return string(bs)
}

func b64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
