go-config-loader
----

[![GoDoc][1]][2] [![License: MIT][3]][4] [![Release][5]][6] [![Build Status][7]][8] [![Codecov Coverage][11]][12] [![Go Report Card][13]][14] [![Code Climate][19]][20] [![BCH compliance][21]][22]

[1]: https://godoc.org/github.com/evalphobia/go-config-loader?status.svg
[2]: https://godoc.org/github.com/evalphobia/go-config-loader
[3]: https://img.shields.io/badge/License-MIT-blue.svg
[4]: LICENSE.md
[5]: https://img.shields.io/github/release/evalphobia/go-config-loader.svg
[6]: https://github.com/evalphobia/go-config-loader/releases/latest
[7]: https://travis-ci.org/evalphobia/go-config-loader.svg?branch=master
[8]: https://travis-ci.org/evalphobia/go-config-loader
[9]: https://coveralls.io/repos/evalphobia/go-config-loader/badge.svg?branch=master&service=github
[10]: https://coveralls.io/github/evalphobia/go-config-loader?branch=master
[11]: https://codecov.io/github/evalphobia/go-config-loader/coverage.svg?branch=master
[12]: https://codecov.io/github/evalphobia/go-config-loader?branch=master
[13]: https://goreportcard.com/badge/github.com/evalphobia/go-config-loader
[14]: https://goreportcard.com/report/github.com/evalphobia/go-config-loader
[15]: https://img.shields.io/github/downloads/evalphobia/go-config-loader/total.svg?maxAge=1800
[16]: https://github.com/evalphobia/go-config-loader/releases
[17]: https://img.shields.io/github/stars/evalphobia/go-config-loader.svg
[18]: https://github.com/evalphobia/go-config-loader/stargazers
[19]: https://codeclimate.com/github/evalphobia/go-config-loader/badges/gpa.svg
[20]: https://codeclimate.com/github/evalphobia/go-config-loader
[21]: https://bettercodehub.com/edge/badge/evalphobia/go-config-loader?branch=master
[22]: https://bettercodehub.com/

`go-config-loader` supports to load config files and convert to map values.


## Supported format

- environment variable
- toml
- json

## Quick Usage


```go
import (
	config "github.com/evalphobia/go-config-loader"
)

func main(){
	confType := "toml"
	basePath := "config"

	conf := config.NewConfig()
	// load toml files from `config/dev/`
	conf.LoadConfigs(basePath + config.Sep + "dev", confType)

	// load toml files from `config/`
	conf.LoadConfigs(basePath, confType)

	var name string
	// find username parameter
	name = conf.ValueString("username")
	name = conf.ValueStringDefault("username", "John") // return John if `username` is not set
	name = conf.ValueString("empty_value") // return "" if `username` is not set

	var userid int
	// find userid parameter
	id = conf.ValueInt("userid")
	id = conf.ValueIntDefault("userid", 100) // return 100 if `userid` is not set
	id = conf.ValueInt("empty_value") // return 0 if `userid` is not set

	var isProd bool
	// find production parameter
	isProd = conf.ValueBool("production") // return false if `production` is not set
}
```
