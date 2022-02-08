package main

import (
	"fmt"
)

type IRealObject interface {
	performAction()
}

type RealObject struct{}

// RealObject class implements IRealObject interface. The class has method
// performAction.
//RealObject class method performAction
func (realObject *RealObject) performAction() {
	fmt.Println("RealObject performAction()")
}

type VirtualProxy struct {
	realObject *RealObject
}

//VirtualProxy class method performAction
func (virtualProxy *VirtualProxy) performAction() {
	if virtualProxy.realObject == nil {
		virtualProxy.realObject = &RealObject{}
	}
	fmt.Println("Virtual Proxy performAction()")
	virtualProxy.realObject.performAction()
}

func main() {
	var object VirtualProxy = VirtualProxy{}
	object.performAction()
}
