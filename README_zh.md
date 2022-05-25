# apidoc

![LICENSE](https://img.shields.io/github/license/zc2638/apidoc.svg?style=flat-square&color=blue)
[![GoDoc](https://pkg.go.dev/badge/github.com/zc2638/apidoc)](https://pkg.go.dev/github.com/zc2638/apidoc)
[![Go Report Card](https://goreportcard.com/badge/github.com/zc2638/apidoc?style=flat-square)](https://goreportcard.com/report/github.com/zc2638/apidoc)

[English](./README.md) | 简体中文

这是一个使用 `swagger2.0` 定义，基于模板文件生成API静态文档（如pdf）的命令行工具和 `Golang`工具包。

## 前提条件

`wkhtmltopdf` 必须被安装。

- MacOS: `brew install Caskroom/cask/wkhtmltopdf`
- 其它操作系统安装方式见 [https://wkhtmltopdf.org/downloads.html](https://wkhtmltopdf.org/downloads.html).

## 安装

### 源码

```shell
go install github.com/zc2638/apidoc/cmd/apidoc@latest
```

## 使用

### 文件

```shell
apidoc --src <your-swagger-json> [--dest <your-output-dir>]
```

### 根据URL

```shell
apidoc --src https://petstore.swagger.io/v2/swagger.json
```

## 工具包使用示例

请查看 [example](./example/main.go)
