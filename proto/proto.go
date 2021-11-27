package proto

import (
	"github.com/mikhailbolshakov/app-b/proto"
)

type AppAProto struct {}

func (a *AppAProto) Print() {
	v := &proto.AppBProto{}
	v.Print()
}
