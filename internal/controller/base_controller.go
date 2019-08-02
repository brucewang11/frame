package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/alecthomas/log4go"
	"github.com/brucewang11/frame/internal/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"text/template"
)
const headerLangKey = "content-language"
const headerZhLang = "zh"
const headerEnLang = "en"
var AllMsg = make(map[string]map[int]string)





func init (){
	msgCh := make(map[int]string)
	msgEn := make(map[int]string)
	err := readFile("./static/lang/msg_ch.json",&msgCh)
	if err !=nil {
		log.Error("读取json错误")
		return
	}
	err = readFile("./static/lang/msg_en.json",&msgEn)
	if err !=nil {
		log.Error("读取json错误")
		return
	}
	AllMsg[headerZhLang] = msgEn
	AllMsg[headerEnLang] = msgCh

}

func  ResData(e *service.CodeModel,ctx *gin.Context) {
	lan := ctx.GetHeader(headerLangKey)
	if lan == ""{
		lan = headerZhLang
	}
	if e.Data == nil {
		e.Data = "{}"
	}
	msg := AllMsg[lan][e.Code]
	if len(e.TemplateData) != 0 {
		tmpl, err := template.New("tem").Parse(msg)
		if err!=nil {
			log.Error("模板参数转化错误")
			ctx.JSON(http.StatusOK, gin.H{"Code": e.Code, "Data": e.Data, "Msg":msg})
			return
		}
		tmplBuffer := new(bytes.Buffer)
		err = tmpl.Execute(tmplBuffer, e.TemplateData)
		if err!=nil {
			log.Error("模板参数转化错误")
			ctx.JSON(http.StatusOK, gin.H{"Code": e.Code, "Data": e.Data, "Msg":msg})
			return
		}
		msg = tmplBuffer.String()
	}
	if e.Code !=0 {
		log.Trace("clientIP:"+ctx.ClientIP()+" "+ctx.Request.RequestURI,e.Code,msg)
	}

	ctx.JSON(http.StatusOK, gin.H{"Code": e.Code, "Data": e.Data, "Msg":msg})

}

func readFile(filename string,msg *map[int]string) error{
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return  err
	}
	if err = json.Unmarshal(bytes, msg); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return  err
	}
	return  nil
}

