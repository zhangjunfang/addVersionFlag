package plugin

import (
	"bytes"
	"time"

	"github.com/coreos/etcd/clientv3"
	//"github.com/coreos/etcd/pkg/transport"
	"github.com/zhangjunfang/addVersionFlag/codeVersion"
	"github.com/zhangjunfang/addVersionFlag/golang"
	"golang.org/x/net/context"
)

func VersionEtcd(key string, ip []string) (err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   ip,
		DialTimeout: 15 * time.Second,
	})
	if err != nil {
		return err
	}
	defer cli.Close()
	var buff bytes.Buffer
	buff.WriteString(time.Now().Format("2006-1-2-15:04:05"))
	buff.WriteString("-")
	buff.WriteString(golang.GetRunTimeVersion())
	buff.WriteString("-")
	v, _ := codeVersion.GetCodeVersion()
	buff.WriteString(v)
	_, err = cli.Put(context.TODO(), key, buff.String())
	buff.Reset()
	if err != nil {
		return err
	}
	return nil
}
