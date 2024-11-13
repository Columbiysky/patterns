package main

import (
	"fmt"
	"log"
)

type Facade struct {
	objOne *ObjectOne
	objTwo *ObjectTwo
}

func newFacade(objOneName string) *Facade {
	walletFacacde := &Facade{
		objOne: newObjectOne(objOneName),
		objTwo: newObjectTwo(),
	}
	fmt.Println("Created new facade")
	return walletFacacde
}

func (w *Facade) addValToObjTwo(objOneName string, amount int) error {
	fmt.Println("Add objTwo val if objOne is exising")
	err := w.objOne.isObjOneExisting(objOneName)
	if err != nil {
		return err
	}
	w.objTwo.increaseIntVal(amount)
	return nil
}

func (w *Facade) decreaseValFromObjTwo(accountID string, amount int) error {
	fmt.Println("Take away objTwo val if objOne is exising")
	err := w.objOne.isObjOneExisting(accountID)
	if err != nil {
		return err
	}

	err = w.objTwo.decreaseIntVal(amount)
	if err != nil {
		return err
	}
	return nil
}

type ObjectOne struct {
	name string
}

func newObjectOne(name string) *ObjectOne {
	return &ObjectOne{
		name: name,
	}
}

func (a *ObjectOne) isObjOneExisting(name string) error {
	if a.name != name {
		return fmt.Errorf("Obj one Name is incorrect")
	}
	fmt.Println("Obj one existing")
	return nil
}

type ObjectTwo struct {
	val int
}

func newObjectTwo() *ObjectTwo {
	return &ObjectTwo{
		val: 0,
	}
}

func (w *ObjectTwo) increaseIntVal(amount int) {
	w.val += amount
	fmt.Println("ObjectTwo val increased")
	return
}

func (w *ObjectTwo) decreaseIntVal(amount int) error {
	if w.val < amount {
		return fmt.Errorf("ObjectTwo val is less than amount")
	}
	fmt.Println("ObjectTwo val decreased")
	w.val = w.val - amount
	return nil
}

func main() {
	fmt.Println()
	facade := newFacade("FacadeOne")
	fmt.Println()

	err := facade.addValToObjTwo("FacadeOne", 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = facade.decreaseValFromObjTwo("FacadeOne", 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
