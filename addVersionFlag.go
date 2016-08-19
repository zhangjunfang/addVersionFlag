package addVersionFlag

import (
	"bytes"
	"os"
	"time"

	"github.com/zhangjunfang/addVersionFlag/codeVersion"
	"github.com/zhangjunfang/addVersionFlag/golang"
	"github.com/zhangjunfang/addVersionFlag/plugin"
)

//使用etcd 存储版本号
func VersionWithEtcd(key string, ip []string) (err error) {
	return plugin.VersionEtcd(key, ip)
}

//使用本机管理  上线版本不多
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

//使用本机管理  上线版本不多
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

//使用本机管理  上线版本不多
func Version(s string) {
	var buff bytes.Buffer
	file, _ := os.OpenFile("./.version", os.O_CREATE|os.O_RDWR, os.ModePerm)
	buff.WriteString(s)
	file.WriteString(buff.String())
	buff.Reset()
	file.Close()
}
