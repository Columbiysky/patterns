package main

import "fmt"

const (
	FirstWrap  = "FirstWrap"
	SecondWrap = "SecondWrap"
)

var (
	colorizedWrappersFactorySingleInstance = &ColorizedWrappersFactory{
		dressMap: make(map[string]ColorizedWrapper),
	}
)

type ColorizedWrappersFactory struct {
	dressMap map[string]ColorizedWrapper
}

func (d *ColorizedWrappersFactory) getColorizedWrapperByType(wrapperType string) (ColorizedWrapper, error) {
	if d.dressMap[wrapperType] != nil {
		return d.dressMap[wrapperType], nil
	}

	if wrapperType == FirstWrap {
		d.dressMap[wrapperType] = newBlackColorizedWrapper()
		return d.dressMap[wrapperType], nil
	}
	if wrapperType == SecondWrap {
		d.dressMap[wrapperType] = newWhiteColorizedWrapper()
		return d.dressMap[wrapperType], nil
	}

	return nil, fmt.Errorf("Wrong wrapper type passed")
}

func getColorizedWrappersFactorySingleInstance() *ColorizedWrappersFactory {
	return colorizedWrappersFactorySingleInstance
}

type ColorizedWrapper interface {
	getColor() string
}

type BlackColorizedWrapper struct {
	color string
}

func (t *BlackColorizedWrapper) getColor() string {
	return t.color
}

func newBlackColorizedWrapper() *BlackColorizedWrapper {
	return &BlackColorizedWrapper{color: "black"}
}

type WhiteColorizedWrapper struct {
	color string
}

func (c *WhiteColorizedWrapper) getColor() string {
	return c.color
}

func newWhiteColorizedWrapper() *WhiteColorizedWrapper {
	return &WhiteColorizedWrapper{color: "white"}
}

type Object struct {
	colorizedWrapper ColorizedWrapper
	objType          string
	lat              int
	long             int
}

func newObject(playerType, colorizedWrapperType string) *Object {
	colorizedWrapper, _ := getColorizedWrappersFactorySingleInstance().getColorizedWrapperByType(colorizedWrapperType)
	return &Object{
		objType:          playerType,
		colorizedWrapper: colorizedWrapper,
	}
}

func (p *Object) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

type ObjectsField struct {
	terrorists        []*Object
	counterTerrorists []*Object
}

func newField() *ObjectsField {
	return &ObjectsField{
		terrorists:        make([]*Object, 1),
		counterTerrorists: make([]*Object, 1),
	}
}

func (c *ObjectsField) addObjOne(dressType string) {
	player := newObject("One", dressType)
	c.terrorists = append(c.terrorists, player)
	return
}

func (c *ObjectsField) addObjTwo(dressType string) {
	player := newObject("Two", dressType)
	c.counterTerrorists = append(c.counterTerrorists, player)
	return
}

func main() {
	field := newField()

	//Add ObjOne
	field.addObjOne(FirstWrap)
	field.addObjOne(FirstWrap)
	field.addObjOne(FirstWrap)

	//Add dObjTwo
	field.addObjTwo(SecondWrap)
	field.addObjTwo(SecondWrap)
	field.addObjTwo(SecondWrap)

	colorizedWrappersFactoryInstance := getColorizedWrappersFactorySingleInstance()

	for wrapperType, wrapper := range colorizedWrappersFactoryInstance.dressMap {
		fmt.Printf("WrapperColorType: %s\nWrapperColor: %s\n", wrapperType, wrapper.getColor())
	}
}
