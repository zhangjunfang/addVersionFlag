package main

import (
	"bytes"
	"os"
	"time"

	"github.com/zhangjunfang/addVersionFlag/codeVersion"
	"github.com/zhangjunfang/addVersionFlag/golang"
)

var (
	buildTime = "2016-01-09 UTC"
	gitHash   = "master"
)

func init() {
	var buff bytes.Buffer
	file, _ := os.OpenFile("./.version", os.O_CREATE|os.O_RDWR, os.ModePerm)
	buff.WriteString(time.Now().Format("2006-1-2-15:04:05"))
	buff.WriteString("-")
	buff.WriteString(golang.GetRunTimeVersion())
	buff.WriteString("-")
	v, _ := codeVersion.GetCodeVersion()
	buff.WriteString(v)
	file.WriteString(buff.String())
	//os.Setenv("codeVersion", buff.String())
	buff.Reset()
	file.Close()
}
func main() {

}
