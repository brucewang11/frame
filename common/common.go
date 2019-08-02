package common

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
	"strings"
)

type httpCfg struct {
	Port string
	GinMode string
}
var HttpInfo httpCfg



func InitCommon(){
	if _, err := toml.DecodeFile(getCurrentDirectory()+"/configs/http.toml", &HttpInfo); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}
