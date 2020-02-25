package main

import (
	"fmt"
	"reflect"
)

type Brand struct {
	name string
}

func (brand Brand) Show() {
	fmt.Println(brand.name)
}

type FakeBrand = Brand

type Vehicle struct {
	FakeBrand
	Brand
}

func main() {
	var a Vehicle
	a.FakeBrand.name = "yanle"

	a.FakeBrand.Show()

	ta := reflect.TypeOf(a)

	//fmt.Println(a.FakeBrand.name)

	for i := 0; i < ta.NumField(); i++ {
		f := ta.Field(i)
		fmt.Printf("FieldName: %v, FiledType: %v\n", f.Name, f.Type.Name())
	}
}
