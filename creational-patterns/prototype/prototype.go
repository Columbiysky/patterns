package main

import (
	"fmt"
)

type IObject interface {
	clone()
}

type Object struct {
	prop string
}

func (o *Object) clone() *Object {
	return &Object{prop: o.prop + "_cloned"}
}

func main() {
	initObj := &Object{prop: "prop"}
	copyObj := initObj.clone()

	fmt.Println(initObj.prop)
	fmt.Println(copyObj.prop)
}
