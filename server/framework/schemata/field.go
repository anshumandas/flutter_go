package schemata

import (
	interfaces "github.com/flutter_go/framework/base"
	goInterfaces "github.com/flutter_go/framework/base/gointerfaces"
	enums "github.com/flutter_go/framework/schemata/enums"
)

//This is used to define the Field of the Data
//do not use field name starting with _ here or any other schema. it is only in the interfaces package where we use in the json tag

type Field struct {
	goInterfaces.Ordered
	interfaces.Embedded //Field is Embedded and not Referable
	Detail
	DefaultValue    interface{}                 //any value type
	Type            enums.DataType              `json:"type"`
	Confidentiality enums.ConfidentialityString `json:"confidentiality"`
	State           enums.SchemaState           `json:"state"`
}

func (a Field) Compare(b Field) int { //first by name and then by the order
	var r = a.Named.Compare(b.Named)
	if r != 0 { //if name same then no need to check order
		var o = a.Ordered.Compare(b.Ordered)
		if o != 0 {
			return 0
		}
	}
	return r
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
