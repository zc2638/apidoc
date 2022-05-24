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

## Example

Please visit the [example](./example/main.go)
