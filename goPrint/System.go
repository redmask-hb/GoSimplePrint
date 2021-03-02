package goPrint

import (
	"runtime"
)

var OS string

func init()  {
	initOS()
	initColor()
}

func initOS()  {
	sysStr:=runtime.GOOS
	if sysStr=="windows" {
		OS=sysStr
	}else {
		OS="linux"
	}
}
