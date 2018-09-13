# login-srun

[![Build Status](https://travis-ci.org/monigo/login-srun.svg?branch=master)](https://travis-ci.org/monigo/login-srun)
[![Gitter](https://img.shields.io/gitter/room/nwjs/nw.js.svg)](https://gitter.im/monigo-dev/project-login-srun)
![License](https://img.shields.io/packagist/l/doctrine/orm.svg)

北京理工大学校园网命令行登录工具~支持linux、maxOS、windows
基于go语言实现

## 下载最新编译好release

[latest release(包含linux、macos与win)](https://github.com/monigo/login-srun/releases/tag/v0.1.1)

## 更新日志
2018.9.1
- 实现登录与设置账号的功能

## 功能与原理

主要功能
- 保存账号
- 使用账号快速登录校园网

原理

工具会把账号信息保存为json存放到 `~/.srun/account.json` 下
执行的时候自动读取账号信息事实现登录


## 使用方法例
> 假设把编译好的可执行文件命名为srun， 加入系统path

```bash
$ srun
login success!
```

### 查看帮助

```bash
$ srun -h
Srun v0.1.1

Options:
	-d                  Show debug message
	-v                  Show version
	-h                  Show help

Commands:
	account             Get/Set Username and Password
	login               Login Srun

$ srun account -h
Usage: srun account [<Username> <Password>]
  Get/Set Username and Password

$ srun login -h
Usage: srun [login]
  Login Srun


```


### 查看账号

```bash
$ srun account
Username: <your-username>
Password: <your-password>
```

### 设置账号

```bash
$ srun account <your-username> <your-password>

```

## 登录校园网

```bash
$ srun login
login success!

$ srun
login success!

```


## develop

### 编译

执行build.sh 文件即可编译

```bash
$ chmod +x build.sh && ./build.sh
```

编译好的可执行文件在bin文件夹中




