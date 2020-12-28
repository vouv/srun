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
$ go get -u -v github.com/vouv/srun/cmd/srun
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





