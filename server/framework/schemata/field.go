package schemata

import (
	"encoding/json"

	interfaces "github.com/flutter_go/framework/base"
	goInterfaces "github.com/flutter_go/framework/gointerfaces"
	enums "github.com/flutter_go/framework/schemata/enums"
)

//This is used to define the Field of the Data
//do not use field name starting with _ here or any other schema. it is only in the interfaces package where we use in the json tag

type Field struct {
	goInterfaces.Ordered
	interfaces.Embedded //Field is Embedded and not Referable
	goInterfaces.Detail
	goInterfaces.Tagged
	Type            enums.DataType              `json:"type"`
	DefaultValue    interface{}                 `json:"default,omitempty"`         //any value type
	Confidentiality enums.ConfidentialityString `json:"confidentiality,omitempty"` //default is of parent schema
	Required        bool                        `json:"required,omitempty"`
	Mutability      enums.Editability           `json:"editability,omitempty"` //default is mutable
	MinValue        interface{}                 `json:"min,omitempty"`         //used for numbers and date
	MaxValue        interface{}                 `json:"max,omitempty"`         //used for numbers and date
	StepValue       interface{}                 `json:"step,omitempty"`        //used for numbers and date
	Options         string                      `json:"options,omitempty"`     //either links a schema or [values,]
	Excludes        string                      `json:"excludes,omitempty"`    //either links a schema or [values,]
	AllowOther      bool                        `json:"other,omitempty"`       //allow an Other option in the dropdown
	Pattern         string                      `json:"pattern,omitempty"`     //regex
	State           enums.SchemaState           `json:"state,omitempty"`
	Index           enums.Indexed               `json:"index,omitempty"`
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

func (w Field) MarshalJSON() ([]byte, error) {
	//TODO: This method must be overridden
	return json.Marshal(struct {
		Field
		Type0 string `json:"t0"`
	}{
		//add all other fields
		Type0: w.Type0(),
	})
}

func (p Field) Type0() string {
	return "Field"
}
