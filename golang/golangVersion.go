package golang

import (
	"runtime"
)

func GetRunTimeVersion() string {
	return runtime.Version()
}
