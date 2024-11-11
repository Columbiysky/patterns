package main

import (
	"fmt"
)

type Object struct {
}

type IObject interface {
	foo()
}

func (c *Object) foo(com IObject) {
	fmt.Println("foo of abstract object")
	com.foo()
}

type SupportedObject struct {
}

func (m *SupportedObject) foo() {
	fmt.Println("foo of SupportedObject")
}

type UnsupportedObject struct{}

func (w *UnsupportedObject) foo() {
	fmt.Println("foo of UnsupportedObject")
}

type Adapter struct {
	unsupportedObj *UnsupportedObject
}

func (w *Adapter) foo() {
	fmt.Println("Adapter converts unsupported to supported")
	w.unsupportedObj.foo()
}

func main() {
	obj := &Object{}
	supportedObj := &SupportedObject{}

	obj.foo(supportedObj)

	unsupportedObj := &UnsupportedObject{}
	adapter := &Adapter{
		unsupportedObj: unsupportedObj,
	}

	obj.foo(adapter)
}
