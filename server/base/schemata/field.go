package schemata

import "github.com/flutter_go/base/interfaces"

//This is used to define the Field of the Data
//do not use field name starting with _ here or any other schema. it is only in the interfaces package where we use in the json tag

type Field struct {
	Named
	interfaces.Embedded //Field is Embedded and not Referable
}

func NewField() *Field {
	s := Field{}
	return &s
}

func NewFieldByName(name string) *Field {
	s := Field{}
	s.Name = name
	return &s
}
