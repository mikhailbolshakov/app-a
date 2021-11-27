package proto

import "fmt"

type AppAProto struct {}

func (a *AppAProto) Print() {
	fmt.Println("AppA")
}
