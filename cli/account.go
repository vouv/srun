package cli

import (
	"path/filepath"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"os/user"
)

var SrunRootPath string

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
		fmt.Println("Error: get current user error,", gErr)
		os.Exit(1)
	}

	SrunRootPath = curUser.HomeDir
}
type Account struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
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
	return fmt.Sprintf("Username: %s\nPassword: %s", acc.Username, acc.Password)
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
	account.Username = username
	account.Password = password

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
		err = fmt.Errorf("Open account file error, %s, please use `account` to set AccessKey and SecretKey first. ", openErr)
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

	logs.Debug("Load account from %s", accountFname)
	return
}
