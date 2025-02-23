# Srun

[![Build Status](https://travis-ci.org/vouv/srun.svg?branch=master)](https://travis-ci.org/vouv/srun) [![Go Report Card](https://goreportcard.com/badge/github.com/vouv/srun)](https://goreportcard.com/report/github.com/vouv/srun) ![License](https://img.shields.io/packagist/l/doctrine/orm.svg) [![GoDoc](https://godoc.org/github.com/vouv/srun?status.svg)](https://godoc.org/github.com/vouv/srun/core)

> An efficient client for BIT campus network

北京理工大学校园网命令行登录工具
- 支持linux、maxOS、windows
- 基于Go语言实现

Related Projects

- macOS客户端: [SrunBar](https://github.com/vouv/SrunBar)

## Install

1. Homebrew(macOS only)

```bash
$ brew tap vouv/tap
$ brew install srun
$ srun config
```

2. Curl(for Linux amd64) [Release](https://github.com/vouv/srun/releases/latest)

```bash
# linux
$ curl -L -o srun https://github.com/vouv/srun/releases/latest/download/srun-linux
$ chmod +x ./srun
$ ./srun config
```

3. go get

如果已经[安装并配置GO环境](https://golang.google.cn/doc/install), 执行如下命令即可

```bash
$ go install github.com/vouv/srun/cmd/srun@latest
$ $GOPATH/bin/srun config
```


## Usage

### Show Help

```
$ srun -h
```

### Config

```
$ srun config
```

### Login

```
$ srun
$ srun login
```

### Info
```
$ srun info
```

## Update Log

2020.12.18

- 自动构建切换到Github Actions

2020.11.3

- 优化新版登录逻辑
- 优化命令行框架
- 删除无用代码，优化代码结构

2020.9.6

- 修复一些bug
- 移除不用的移动联通登录模式

2019.11.16

- 更新安装方式
- 优化项目api与项目结构

2019.9.10

- 修改优化登录逻辑
- 修复一些bug

2019.1.3
- 实现无缓冲输入密码（在macOS上测试通过）
- 修复宿舍无法登录移动网的bug

2018.11.24
- 增加登出功能
- 增加查询流量和余额功能

2018.9.1
- 实现登录与设置账号的功能


### Contribute

> 要求先安装好golang环境 go version > 1.10

先克隆项目

```
$ git clone https://github.com/vouv/srun && cd srun
```

macOS下编译

```bash
$ make
```
或
```bash
$ make darwin
```

Windows下编译
```bash
$ make windows
```

Linux下编译
```bash
$ make linux
```

编译好的可执行文件在bin文件夹中

### About

主要功能与原理

- 本地保存账号到`$HOME/.srun/account.json`（对安全性有疑问的请自行看代码）
- 使用账号快速登录校园网，环境支持的情况下也可以一键登录


### Thanks to

- [logrus](https://github.com/sirupsen/logrus)
- [cobra](https://github.com/spf13/cobra)

# Srun Release v1.1.5-patch

> 源代码：https://github.com/vouv/srun

原发布版本Srun Release v1.1.5（2021-1-2）已经无法正常连接校园网，报错如下：
```bash
./srunINFO[xxxx-xx-xx xx:xx:xx]尝试登录...
Error: {login error login error E2620 E2620: You are already online. xxx.xxx.xxx.xxx}
```

基于原版Srun修改后，解决了相应连接问题。发布Release v1.1.5-patch

## 代码修正

修改后的代码如下：
```python
#!/usr/bin/env python3
# -*-coding:utf-8-*-

import base64
import hashlib
import hmac
import json
import os
import time
import re
import requests
import argparse
import configparser

def char_code_at(stri, index):
    return 0 if index >= len(stri) else ord(stri[index])

def xencode(msg: str, key):
    '''
    xEncode
    '''
    def s(a: str, b: bool):
        c = len(a)
        v = []
        for i in range(0, c, 4):
            v.append(char_code_at(a, i) | (char_code_at(a, i+1) << 8) |
                     (char_code_at(a, i+2) << 16) | (char_code_at(a, i+3) << 24))
        if b:
            v.append(c)
        return v

    def l(a, b):
        d = len(a)
        c = (d-1) << 2
        if b:
            m = a[d-1]
            if (m < c-3) or (m > c):
                return None
            c = m
        for i in range(0, d):
            a[i] = ''.join([chr(a[i] & 0xff), chr((a[i] >> 8) & 0xff), chr(
                (a[i] >> 16) & 0xff), chr((a[i] >> 24) & 0xff)])
        if b:
            return (''.join(a))[0:c]
        else:
            return ''.join(a)

    if msg == "":
        return ""
    v = s(msg, True)
    k = s(key, False)
    # print(v)
    # print(k)
    n = len(v) - 1
    z = v[n]
    y = v[0]
    c = 0x86014019 | 0x183639A0
    m = 0
    e = 0
    p = 0
    q = 6 + 52 // (n + 1)
    d = 0
    while 0 < q:
        q -= 1
        d = d + c & (0x8CE0D9BF | 0x731F2640)
        e = d >> 2 & 3
        for p in range(0, n):
            y = v[p+1]
            m = z >> 5 ^ y << 2
            m += (y >> 3 ^ z << 4) ^ (d ^ y)
            m += k[(p & 3) ^ e] ^ z
            z = v[p] = v[p] + m & (0xEFB8D130 | 0x10472ECF)
        y = v[0]
        m = z >> 5 ^ y << 2
        m += (y >> 3 ^ z << 4) ^ (d ^ y)
        m += k[(n & 3) ^ e] ^ z
        z = v[n] = v[n] + m & (0xBB390742 | 0x44C6F8BD)
    # print(v)
    return l(v, False)

def get_json(url, data):
    '''Http GET, return json
    '''
    callback = "jsonp%s" % int(time.time()*1000)
    data["callback"] = callback

    response = requests.get(url, data)
    response_content = response.content.decode('utf-8')[len(callback)+1:-1]
    response_json = json.loads(response_content)

    return response_json

# 读取配置文件
def read_config():
    config = configparser.ConfigParser()
    config.read('srun_config.ini')

    # 如果没有配置文件或配置项不存在，返回空字典
    if not config.has_section('user') or not config.has_option('user', 'username'):
        return None

    return {
        'username': config.get('user', 'username'),
        'password': config.get('user', 'password')
    }

# 保存配置文件
def save_config(username, password):
    config = configparser.ConfigParser()
    if not config.has_section('user'):
        config.add_section('user')

    config.set('user', 'username', username)
    config.set('user', 'password', password)

    with open('srun_config.ini', 'w') as configfile:
        config.write(configfile)

# 显示帮助信息
def print_help():
    help_text = """
Usage:
  ./srun login    Log in to the campus network
  ./srun logout   Log out of the campus network
  ./srun config   Configure account and password
  ./srun -h       Show this help message
"""
    print(help_text)

def srun_login_logout(username, password, action):
    '''srun login and logout
    Args:
        username: username
        password: password
        action: 'login' or 'logout'
    Returns:
        a json object.
    '''
    def data_info(get_data, token):
        if get_data['action'] == 'login':
            x_encode_json = {
                "username": get_data['username'],
                "password": get_data['password'],
                "ip": get_data['ip'],
                "acid": get_data['ac_id'],
                "enc_ver": enc
            }
        else:
            x_encode_json = {
                "username": get_data['username'],
                "ip": get_data['ip'],
                "acid": get_data['ac_id'],
                "enc_ver": enc
            }

        x_encode_str = json.dumps(x_encode_json, separators=(',', ':'))
        x_encode_key = token
        x_encode_res = xencode(x_encode_str, x_encode_key)
        # print("x_encode('%s', '%s')" % (x_encode_str, x_encode_key))
        # print('x_encode_res(len: %s): %s' % (len(x_encode_res), x_encode_res))
        # print("x_encode_res unicode:", [ord(s) for s in x_encode_res])

        # base64_encode
        mapping = dict(zip("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=",
                           "LVoJPiCN2R8G90yg+hmFHuacZ1OWMnrsSTXkYpUq/3dlbfKwv6xztjI7DeBE45QA="))
        b64_res = base64.b64encode(
            bytes([ord(s) for s in x_encode_res])).decode()
        base64_encode_res = ''.join([mapping[b] for b in b64_res])
        # print('base64 encode res(len: %s): %s' % (len(base64_encode_res), base64_encode_res))

        return "{SRBX1}" + base64_encode_res

    def pwd_hmd5(password, token):
        hmac_key = token.encode('utf-8')
        hmac_msg = password.encode('utf-8')
        hmd5 = hmac.new(hmac_key, hmac_msg, digestmod='MD5').hexdigest()
        # print(hmd5)
        return '{MD5}' + hmd5

    def checksum(get_data, token):
        if get_data['action'] == 'login':
            str_list = ['', get_data['username'], get_data['password'][5:],
                        get_data['ac_id'], get_data['ip'], str(n), str(type_), get_data['info']]
        else:
            str_list = ['', get_data['username'], get_data['ac_id'],
                        get_data['ip'], str(n), str(type_), get_data['info']]
        chksum_str = token.join(str_list)
        chksum = hashlib.sha1(chksum_str.encode('utf-8')).hexdigest()
        return chksum

    if action not in ['login', 'logout']:
        print('action must be "login" or "logout".')
        return
    enc = "srun_bx1"
    n = 200
    type_ = 1
    get_challenge_url = "http://10.0.0.55/cgi-bin/get_challenge"
    srun_portal_url = "http://10.0.0.55/cgi-bin/srun_portal"
    url = 'http://detectportal.firefox.com/success.txt'
    #Check if Redirect, when not, set to default
    try:
        r = requests.get(url, timeout=0.1)
        ac_id=re.findall(r'index_(\d*).html',r.url)[0]
    except requests.exceptions.Timeout:
            ac_id=1
    if not ac_id:
        ac_id=1
    ac_id=str(ac_id)
    if action == 'login':
        get_data = {
            "action": action,
            "username": username,
            "password": password,
            "ac_id": ac_id,
            "ip": '',
            "info": '',
            "chksum": '',
            "n": n,
            "type": type_
        }
    else:
        get_data = {
            "action": action,
            "username": username,
            # "password": password, # logout,
            "ac_id": ac_id,
            "ip": '',
            "info": '',
            "chksum": '',
            "n": n,
            "type": type_
        }
    # get token
    challenge_json = get_json(
        get_challenge_url, {"username": get_data['username']})
    if challenge_json['res'] != "ok":
        print('Error getting challenge. %s failed.' % action)
        print('Server response:\n%s' % json.dumps(challenge_json, indent=4))
        return
    token = challenge_json['challenge']
    get_data['ip'] = challenge_json['client_ip']
    get_data['info'] = data_info(get_data, token)

    if action == 'login':
        get_data['password'] = pwd_hmd5('', token)
        # get_data['password'] = pwd_hmd5(get_data['password'], token) # srun's bug

    get_data['chksum'] = checksum(get_data, token)

    # print('get data: %s' % json.dumps(get_data, indent=4))
    res = get_json(srun_portal_url, get_data)
    # print("Server response: %s" % json.dumps(res, indent=4))

    if res['error'] == 'ok':
        print('%s success.' % action)
    else:
        print("%s failed.\n%s %s" % (action, res['error'], res['error_msg']))

    return res

# 主函数
def main():
    parser = argparse.ArgumentParser(description="Campus Network Management Tool", add_help=False)
    parser.add_argument("action", choices=["login", "logout", "config"], nargs='?', help="Action to perform")
    parser.add_argument("-h", "--help", action="store_true", help="Show this help message")
    args = parser.parse_args()
    
    if args.help:
        print_help()
        return

    if args.action == "config":
        # 配置账号密码
        print("Please enter your campus network username and password.")
        username = input("Username: ")
        password = input("Password: ")
        save_config(username, password)
        print("Configuration saved.")
    elif args.action == "login" or args.action == "logout":
        config = read_config()

        if not config:
            print("No configuration found. Please run './srun config' to set your username and password.")
            return

        username = config['username']
        password = config['password']
        

        if args.action == "login":
            print(f"Attempting to login as {username}...")
            srun_login_logout(username, password, action="login")
            print(f"Sucessfully login as {username}...")
        elif args.action == "logout":
            print(f"Attempting to logout as {username}...")
            srun_login_logout(username, None, action="logout")
            print(f"Sucessfully logout as {username}...")
    elif args.action is None:
        print_help()
    else:
        print_help()

if __name__ == "__main__":
    main()
```

配置文件的默认路径：./srun_config.ini，文件内容包括用户名和密码，格式如下：
```ini
[user]
username = your_username
password = your_password
```

## 重新编译

使用PyInstaller将这个Python脚本编译为Linux系统的可执行文件。首先，确保您已经安装PyInstaller：
```bash
pip install pyinstaller
```

进入你的Srun源码（srun_source_code.py）所在目录，运行以下命令来编译代码：
```bash
pyinstaller --onefile --distpath ./dist --name srun srun_source_code.py
# --onefile表示生成一个单独的可执行文件。
# --distpath ./dist指定输出目录为./dist。
# --name srun设定输出文件名为srun。
```
## 使用方法

编译完成后，在./dist目录下可找到./srun可执行文件，具体使用方法可通过./srun -h指令打印
```bash
Usage:
  ./srun login    Log in to the campus network
  ./srun logout   Log out of the campus network
  ./srun config   Configure account and password
  ./srun -h       Show this help messag
```

## 连接验证

```bash
# 若连接成功则可以可以ping通github.com （通过ipv4地址连接）
ping www.github.com 
# 对于www.baidu.com，由于该网站支持ipv6，而BIT-Web不需要登录也可使用ipv6，故即便没有连接上校园网，也能ping通www.baidu.com (将通过ipv6地址连接)
```


