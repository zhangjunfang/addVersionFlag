# addVersionFlag
golang开发程序增加版本号

## Install

```bash
go get github.com/zhangjunfang/addVersionFlag
```
## example
```go
addVersionFlag.VersionWithEtcd(key,ips) //use etcd
addVersionFlag.VersionWithGit()  //save File 
addVersionFlag.VersionWithCustomer(v) //the custom version number
addVersionFlag.Version(v) //completely customize the version number
``` 
