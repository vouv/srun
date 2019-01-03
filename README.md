# login-srun	[![Build Status](https://travis-ci.org/monigo/login-srun.svg?branch=master)](https://travis-ci.org/monigo/login-srun)	[![Gitter](https://img.shields.io/gitter/room/nwjs/nw.js.svg)](https://gitter.im/monigo-dev/project-login-srun)	![License](https://img.shields.io/packagist/l/doctrine/orm.svg)

北京理工大学校园网命令行登录工具
- 支持linux、maxOS、windows
- 基于Go语言实现


## Release

[latest release(包含linux、macos与win)](https://github.com/monigo/login-srun/releases/tag/v0.1.3)

## 更新日志

2019.1.3
- 实现无缓冲输入密码（在macOS上测试通过）
- 修复宿舍无法登录移动网的bug

2018.11.24
- 增加登出功能
- 增加查询流量和余额功能

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
> 如没有权限添加运行权限 chmod +x


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
$ srun account
```

PS: 如果在宿舍并连接Bit-web，要登录移动请把用户名添加尾缀`@yidong`，例如`monigo@yidong`

### 查看账号

```bash
$ srun account get
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

> 要求先安装好golang环境 go version > 1.10

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


## Thanks to

- [beego](https://github.com/astaxie/beego)
- [goquery](https://github.com/PuerkitoBio/goquery)




