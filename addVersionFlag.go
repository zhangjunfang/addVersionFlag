package addVersionFlag

import (
	"bytes"
	"os"
	"time"

	"github.com/zhangjunfang/addVersionFlag/codeVersion"
	"github.com/zhangjunfang/addVersionFlag/golang"
)

func VersionWithGit() {
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

func VersionWithCustomer(v string) {
	var buff bytes.Buffer
	file, _ := os.OpenFile("./.version", os.O_CREATE|os.O_RDWR, os.ModePerm)
	buff.WriteString(time.Now().Format("2006-1-2-15:04:05"))
	buff.WriteString("-")
	buff.WriteString(golang.GetRunTimeVersion())
	buff.WriteString("-")
	buff.WriteString(v)
	file.WriteString(buff.String())
	buff.Reset()
	file.Close()
}
func Version(s string) {
	var buff bytes.Buffer
	file, _ := os.OpenFile("./.version", os.O_CREATE|os.O_RDWR, os.ModePerm)
	buff.WriteString(s)
	file.WriteString(buff.String())
	buff.Reset()
	file.Close()
}
