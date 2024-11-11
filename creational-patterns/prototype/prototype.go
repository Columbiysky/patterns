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
	return &Object{prop: o.prop}
}

func main() {
	initObj := &Object{prop: "copy"}
	copyObj := initObj.clone()

	fmt.Println(initObj.prop)
	fmt.Println(copyObj.prop)
}
