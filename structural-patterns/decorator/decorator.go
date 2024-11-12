package main

import "fmt"

type IObject interface {
	print()
}

type Base struct {
}

func (p *Base) print() {
	fmt.Println("Base")
}

type FirstWrapper struct {
	obj IObject
}

func (c *FirstWrapper) print() {
	c.obj.print()
	fmt.Println("FirstWrapper")
}

type AnotherWrapper struct {
	obj IObject
}

func (c *AnotherWrapper) print() {
	c.obj.print()
	fmt.Println("AnotherWrapper")
}

func main() {
	baseObject := &Base{}

	objectWithFirstWrapper := &FirstWrapper{
		obj: baseObject,
	}

	fmt.Println("Obj with only first wrapper:")
	objectWithFirstWrapper.print()

	//Add tomato topping
	objWithTwoWrappers := &AnotherWrapper{
		obj: objectWithFirstWrapper,
	}

	fmt.Println("Obj with two wrappers:")
	objWithTwoWrappers.print()
}
