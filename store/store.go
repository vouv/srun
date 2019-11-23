package store

import (
	"encoding/base64"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/vouv/srun/model"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

const accountFileName = "account.json"

var RootPath string

// 写入账号信息到文件
// 统一错误
func SetAccount(username, password, server string) (err error) {

	//write to local dir
	account, err := ReadAccount()
	if err != nil {
		err = ErrReadFile
	}
	account.Username = username
	account.Password = password
	account.Server = server

	return WriteAccount(account)
}

// 保存token
func SetInfo(token, ip string) (err error) {
	//write to local dir
	account, err := ReadAccount()
	if err != nil {
		return
	}

	account.AccessToken = token
	account.Ip = ip

	return WriteAccount(account)
}

// 从文件系统读取账号信息
func ReadAccount() (account model.Account, err error) {
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

	readed, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		log.Debugf("读取账号文件错误, %s", readErr)
		err = ErrReadFile
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(string(readed))
	if err != nil {
		log.Debugf("解析账号文件错误, %s", err)
		err = ErrReadFile
	}

	if umError := json.Unmarshal(decoded, &account); umError != nil {
		log.Debugf("解析账号文件错误, %s", umError)
		err = ErrReadFile
		return
	}
	log.Debug("读取账号信息: ", accountFilename)
	return
}

// 写入文件信息到文件系统
func WriteAccount(account model.Account) (err error) {
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
	jBytes, mErr := account.JSONBytes()
	if mErr != nil {
		log.Debugf("序列化账号错误, %s", mErr)
		err = ErrParse
		return
	}
	// b64 encode
	str := base64.StdEncoding.EncodeToString(jBytes)
	_, wErr := file.WriteString(str)
	if wErr != nil {
		log.Debugf("写入账号文件错误, %s", wErr)
		err = ErrWriteFile
		return
	}
	return
}

// 初始化账号信息
func InitAccount() error {
	return WriteAccount(model.Account{
		Username:    "",
		Password:    "",
		AccessToken: "",
		Ip:          "",
		Server:      model.ServerTypeOrigin,
	})
}

func getAccountFilename() (fileSrc string, err error) {
	storageDir := filepath.Join(RootPath, ".srun")
	if _, sErr := os.Stat(storageDir); sErr != nil {
		if mErr := os.MkdirAll(storageDir, 0755); mErr != nil {
			log.Debugf("mkdir `%s` error, %s", storageDir, mErr)
			return
		}
	}
	fileSrc = filepath.Join(storageDir, accountFileName)
	return
}

func init() {
	curUser, gErr := user.Current()
	if gErr != nil {
		log.Fatalln("无法读取账户信息, 请重新设置账户信息")
	} else {
		RootPath = curUser.HomeDir
	}
}
