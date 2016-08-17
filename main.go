package main

import (
	"os"
)

var (
	buildTime = "2016-01-09 UTC"
	gitHash   = "master"
)

func init() {
	file, e := os.OpenFile("./.version", os.O_CREATE|os.O_RDWR, os.ModePerm)
	file.WriteString()
}
func main() {

}
