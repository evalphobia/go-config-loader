# go-config-loader

`go-config-loader` supports to load config files and convert to map values.


## Supported format

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
