package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Single struct {
}

var singleInstance *Single

func getInstance() *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &Single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

func main() {
	for i := 0; i < 5; i++ {
		go getInstance()
	}

	fmt.Scanln()
}
