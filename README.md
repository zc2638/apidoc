# apidoc

![LICENSE](https://img.shields.io/github/license/zc2638/apidoc.svg?style=flat-square&color=blue)
[![GoDoc](https://pkg.go.dev/badge/github.com/zc2638/apidoc)](https://pkg.go.dev/github.com/zc2638/apidoc)
[![Go Report Card](https://goreportcard.com/badge/github.com/zc2638/apidoc?style=flat-square)](https://goreportcard.com/report/github.com/zc2638/apidoc)

English | [简体中文](./README_zh.md)

This is a tool and `Golang` package that uses `swagger2.0` definitions to generate API static documents (such as pdf) based
on template files.

## Preconditions

The `wkhtmltopdf` package must be installed.

- MacOS: `brew install Caskroom/cask/wkhtmltopdf`
- multiple operating systems can be found
  at [https://wkhtmltopdf.org/downloads.html](https://wkhtmltopdf.org/downloads.html).

## Installation

### Source Code

```shell
go install github.com/zc2638/apidoc/cmd/apidoc@latest
```

## Use

### From File

```shell
apidoc --src <your-swagger-json> [--dest <your-output-dir>]
```

### From URL

```shell
apidoc --src https://petstore.swagger.io/v2/swagger.json
```

## Toolkit Example

Please visit the [example](./example/main.go)
