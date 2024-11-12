package main

import "fmt"

type Container interface {
	search(string)
}

type Dir struct {
	containers []Container
	name       string
}

func (f *Dir) search(keyword string) {
	fmt.Printf("Serching recursively for keyword '%s' in dir '%s'\n", keyword, f.name)
	for _, composite := range f.containers {
		composite.search(keyword)
	}
}

func (f *Dir) add(c Container) {
	f.containers = append(f.containers, c)
}

type Object struct {
	name string
}

func (f *Object) search(keyword string) {
	fmt.Printf("Searching for keyword '%s' in object '%s'\n", keyword, f.name)
}

func main() {
	objOne := &Object{name: "O1"}
	objTwo := &Object{name: "O2"}
	objThree := &Object{name: "O3"}

	dirOne := &Dir{
		name: "D1",
	}

	dirOne.add(objOne)

	dirTwo := &Dir{
		name: "D2",
	}
	dirTwo.add(objTwo)
	dirTwo.add(objThree)
	dirTwo.add(dirOne)

	dirTwo.search("test")
	fmt.Println()
}
