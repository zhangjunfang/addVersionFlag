package codeVersion

import (
	"github.com/zhangjunfang/command/command"
)

func GetCodeVersion() (v string, e error) {
	cmd, e := command.ParseCommand("git rev-parse HEAD")
	if e != nil {
		return v, e
	} else {
		b, e := cmd.CombinedOutput()
		if e != nil {
			return v, e
		} else {
			return string(b), nil
		}
	}

}
