# sa [show adapters]
[![CircleCI](https://circleci.com/gh/vpoletaev11/sa.svg?style=svg)](https://circleci.com/gh/vpoletaev11/sa)
[![Coverage Status](https://coveralls.io/repos/github/vpoletaev11/sa/badge.svg?branch=master)](https://coveralls.io/github/vpoletaev11/sa?branch=master)

**sa** - linux utility that shows system information about network adapters in your system.

# About project
**sa** uses "ip" utility as dependence.<br />
Output of **sa** include: *name, mac, mode*
```shell
1: lo     00:00:00:00:00:00 DEFAULT
2: enp1s0 00:8c:fa:6a:09:87 DEFAULT
3: wlp2s0 e0:b9:a5:1e:e0:88 DORMANT
```
# Project setup
## Clone project
```shell
git clone https://github.com/vpoletaev11/sa
```
## Install dependencies
This project uses modules, because of this to install dependencies you just need to run tests
```shell
$ go test ./...
```
## Build program
```shell
$ go build main.go
```
## Run program
```shell
$ ./main
```