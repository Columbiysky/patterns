package main

import (
	"fmt"
	"strings"
)

type Object struct {
	prop string
}

type IBuilder interface {
	setProp()
	getObject() Object
}

type BigBuilder struct {
	prop string
}

type ReversedBuilder struct {
	prop string
}

func newBigBuilder() *BigBuilder {
	return &BigBuilder{}
}

func (b *BigBuilder) setProp() {
	b.prop = strings.ToUpper("prop")
}

func (b *BigBuilder) getObject() Object {
	return Object{
		prop: b.prop,
	}
}

func newReversedBuilder() *ReversedBuilder {
	return &ReversedBuilder{}
}

func (r *ReversedBuilder) setProp() {
	r.prop = reverseString("prop")
}

func (r *ReversedBuilder) getObject() Object {
	return Object{
		prop: r.prop,
	}
}

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(builder IBuilder) {
	d.builder = builder
}

func (d *Director) buildObject() Object {
	d.builder.setProp()
	return d.builder.getObject()
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	bigBuilder := newBigBuilder()
	reversedBuilder := newReversedBuilder()

	director := newDirector(bigBuilder)
	bigObject := director.buildObject()

	director.setBuilder(reversedBuilder)
	reversedObject := director.buildObject()

	fmt.Println(bigObject.prop)
	fmt.Println(reversedObject.prop)
}
