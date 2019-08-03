package common

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type httpCfg struct {
	Port string
	GinMode string
}
var HttpInfo httpCfg



func InitCommon(){
	if _, err := toml.DecodeFile("./configs/http.toml", &HttpInfo); err != nil {
		fmt.Println(err)
		panic(err)
	}
}


