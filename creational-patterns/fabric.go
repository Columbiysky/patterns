package main

import (
	"fmt"
	"reflect"
)

type IObject interface {
	print()
}

type FirstObject struct{}

type SecondObject struct{}

func (f *FirstObject) print() {
	fmt.Println(reflect.TypeOf(FirstObject{}).Name())
}

func (s *SecondObject) print() {
	fmt.Println(reflect.TypeOf(SecondObject{}).Name())
}

func getObj(name string) (IObject, error) {
	switch name {
	case reflect.TypeOf(FirstObject{}).Name():
		return &FirstObject{}, nil

	case reflect.TypeOf(SecondObject{}).Name():
		return &SecondObject{}, nil
	}

	return nil, fmt.Errorf("Wrong type passed")
}

func main() {
	fobj, _ := getObj("FirstObject")
	sobj, _ := getObj("SecondObject")

	fobj.print()
	sobj.print()

	fmt.Println("Final")
}
