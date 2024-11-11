package main

import "fmt"

type Object interface {
	Print()
	SetExecutor(Executor)
}

type FirstObject struct {
	printer Executor
}

func (m *FirstObject) Print() {
	fmt.Println("Print from FirstObject")
	m.printer.Execute()
}

func (m *FirstObject) SetExecutor(p Executor) {
	m.printer = p
}

type SecondObject struct {
	printer Executor
}

func (w *SecondObject) Print() {
	fmt.Println("Print from SecondObject")
	w.printer.Execute()
}

func (w *SecondObject) SetExecutor(p Executor) {
	w.printer = p
}

type Executor interface {
	Execute()
}

type FirstExecutor struct {
}

func (p *FirstExecutor) Execute() {
	fmt.Println("Print from FirstExecutor")
}

type SecondExecutor struct {
}

func (p *SecondExecutor) Execute() {
	fmt.Println("Print from SecondExecutor")
}

func main() {
	firstExecutor := &FirstExecutor{}
	secondExecutor := &SecondExecutor{}

	firstObject := &FirstObject{}
	firstObject.SetExecutor(firstExecutor)
	firstObject.Print()
	firstObject.SetExecutor(secondExecutor)
	firstObject.Print()
	fmt.Println()

	secondObject := &SecondObject{}
	secondObject.SetExecutor(firstExecutor)
	secondObject.Print()
	secondObject.SetExecutor(secondExecutor)
	secondObject.Print()
	fmt.Println()
}
