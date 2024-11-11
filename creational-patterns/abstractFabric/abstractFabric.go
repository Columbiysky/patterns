package main

import (
	"fmt"
	"reflect"
	"strings"
)

type IObject interface {
	print()
}

type IUpper interface {
	print()
}

type IReversed interface {
	print()
}

type UpperObject struct{}

type ReversedObject struct{}

type FirstUpperObject struct{}

func (f *FirstUpperObject) print() {
	fmt.Println(strings.ToUpper(reflect.TypeOf(FirstUpperObject{}).Name()))
}

type FirstReversedObject struct{}

func (f *FirstReversedObject) print() {
	fmt.Println(reverseString(reflect.TypeOf(FirstReversedObject{}).Name()))
}

type SecondUpperObject struct{}

func (f *SecondUpperObject) print() {
	fmt.Println(strings.ToUpper(reflect.TypeOf(SecondUpperObject{}).Name()))
}

type SecondReversedObject struct{}

func (f *SecondReversedObject) print() {
	fmt.Println(reverseString(reflect.TypeOf(SecondReversedObject{}).Name()))
}

type IAbstractFactory interface {
	createUpper() IUpper
	createReversed() IReversed
}

type FirstFactory struct{}

func (f *FirstFactory) createUpper() IUpper {
	return &FirstUpperObject{}
}

func (f *FirstFactory) createReversed() IReversed {
	return &FirstReversedObject{}
}

type SecondFactory struct{}

func (f *SecondFactory) createUpper() IUpper {
	return &SecondUpperObject{}
}

func (f *SecondFactory) createReversed() IReversed {
	return &SecondReversedObject{}
}

func getFactory(t string) (IAbstractFactory, error) {
	if t == "first" {
		return &FirstFactory{}, nil
	}
	if t == "second" {
		return &SecondFactory{}, nil
	}
	return nil, fmt.Errorf("Wrong factory type passed")
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fFactory, _ := getFactory("first")
	sFactory, _ := getFactory("second")

	fuObj := fFactory.createUpper()
	frObj := fFactory.createReversed()

	suObj := sFactory.createUpper()
	srObj := sFactory.createReversed()

	fuObj.print()
	frObj.print()

	suObj.print()
	srObj.print()

	fmt.Println("Final")
}
