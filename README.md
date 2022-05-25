# apidoc

This is a tool and golang toolkit that uses swagger2.0 definitions to generate API static documents (such as pdf) based
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

### From swagger file

```shell
apidoc --src <your-swagger-json> [--dest <your-output-dir>]
```

### From url

```shell
apidoc --src https://petstore.swagger.io/v2/swagger.json
```

## Toolkit Example

Please visit the [example](./example/main.go)
