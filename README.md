# login-srun 北京理工大学校园网命令行登录工具

[![Build Status](https://travis-ci.org/monigo/login-srun.svg?branch=master)](https://travis-ci.org/monigo/login-srun)
[![Gitter](https://img.shields.io/gitter/room/nwjs/nw.js.svg)](https://gitter.im/monigo-dev/project-login-srun)
![License](https://img.shields.io/packagist/l/doctrine/orm.svg)

北京理工大学校园网命令行登录工具,支持linux、maxOS、windows
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


## 使用方法
> 假设运行前把编译好的可执行文件命名为srun, 加入系统path


### 查看帮助

```bash
$ srun -h
Srun v0.1.1

Options:
	-v                  Show version
	-h                  Show help
	-d                  Show debug message

Commands:
	account             Get/Set Username and Password
	login               Login Srun
	info                Get Srun Info
	logout              Logout Srun

```

### 设置账号

```bash
$ srun account <your-username> <your-password>
```

### 查看账号

```bash
$ srun account
Username: <your-username>
Password(Encoded): <your-password>
```


### 登录校园网（要求先设置好账号密码）
```bash
$ srun
登录成功!
ip: 10.62.41.249
已用流量: 54,418.87M
已用时长: 366小时38分48秒
账户余额: ￥19.68
```

### 查看余额
```bash
$ srun info
已用流量: 54,418.87M
已用时长: 366小时38分48秒
账户余额: ￥19.68
```

### 登出校园网
```bash
$ srun logout
下线成功！
```


## Contribute

### 编译

先克隆项目

```bash
$ git clone https://github.com/monigo/login-srun
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




