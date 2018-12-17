package cli

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

var SrunRootPath string

type Account struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	AccessToken string `json:"access_token"`
	Ip string `json:"ip"`
}

func (acc *Account) ToJson() (jsonStr string, err error) {
	jsonData, mErr := json.Marshal(acc)
	if mErr != nil {
		err = fmt.Errorf("Marshal account data error, %s", mErr)
		return
	}
	jsonStr = string(jsonData)
	return
}

func (acc *Account) String() string {
	return fmt.Sprintln("用户名:", acc.Username)
}


func getAccountFilename() (accountFname string, err error) {
	storageDir := filepath.Join(SrunRootPath, ".srun")
	if _, sErr := os.Stat(storageDir); sErr != nil {
		if mErr := os.MkdirAll(storageDir, 0755); mErr != nil {
			err = fmt.Errorf("Mkdir `%s` error, %s", storageDir, mErr)
			return
		}
	}
	accountFname = filepath.Join(storageDir, "account.json")
	return
}

func init() {
	curUser, gErr := user.Current()
	if gErr != nil {
		logs.Error("无法读取账户信息, 请重新设置账户信息")
		logs.Debug(gErr)
		os.Exit(1)
	}

	SrunRootPath = curUser.HomeDir
}



//写入账号信息到文件
func SetAccount(username string, password string) (err error) {

	accountFname, err := getAccountFilename()
	if err != nil {
		logs.Error(err)
		return
	}
	accountFh, openErr := os.OpenFile(accountFname, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if openErr != nil {
		err = fmt.Errorf("Open account file error, %s", openErr)
		return
	}
	defer accountFh.Close()

	//write to local dir
	var account Account
	account.Username = Encode(username)
	account.Password = Encode(password)


	jsonStr, mErr := account.ToJson()
	if mErr != nil {
		err = mErr
		return
	}
	_, wErr := accountFh.WriteString(jsonStr)
	if wErr != nil {
		err = fmt.Errorf("Write account info error, %s", wErr)
		return
	}

	return
}

func SetInfo(token, ip string) (err error) {

	//write to local dir
	account, err := GetAccount()
	if err != nil {
		logs.Error(err)
		return
	}
	accountFname, err := getAccountFilename()
	if err != nil {
		logs.Error(err)
		return
	}
	accountFh, openErr := os.OpenFile(accountFname, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if openErr != nil {
		err = fmt.Errorf("Open account file error, %s", openErr)
		return
	}
	defer accountFh.Close()
	account.Username = Encode(account.Username)
	account.Password = Encode(account.Password)
	account.AccessToken = token
	account.Ip = ip

	jsonStr, mErr := account.ToJson()
	if mErr != nil {
		err = mErr
		return
	}
	_, wErr := accountFh.WriteString(jsonStr)
	if wErr != nil {
		err = fmt.Errorf("Write account info error, %s", wErr)
		return
	}

	return
}

func GetAccount() (account Account, err error) {

	accountFname, err := getAccountFilename()
	if err != nil {
		logs.Error(err)
		return
	}

	accountFh, openErr := os.Open(accountFname)
	if openErr != nil {
		err = fmt.Errorf("Open account file error, %s, please use `account` to set Username and Password first. ", openErr)
		return
	}
	defer accountFh.Close()

	accountBytes, readErr := ioutil.ReadAll(accountFh)
	if readErr != nil {
		err = fmt.Errorf("Read account file error, %s. ", readErr)
		return
	}

	if umError := json.Unmarshal(accountBytes, &account); umError != nil {
		err = fmt.Errorf("Parse account file error, %s. ", umError)
		return
	}
	account.Username = Decode(account.Username)
	account.Password = Decode(account.Password)

	logs.Debug("Load account from %s", accountFname)
	return
}


func Decode(b64 string) string {
	bs, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		logs.Error(err)
	}
	return string(bs)
}

func Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}