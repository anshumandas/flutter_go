package schemata

import (
	"encoding/json"

	interfaces "github.com/flutter_go/framework/base"
	goInterfaces "github.com/flutter_go/framework/gointerfaces"
)

//This is used to define the Field of the Data
//do not use field name starting with _ here or any other schema. it is only in the interfaces package where we use in the json tag

type Choice struct {
	goInterfaces.Ordered
	interfaces.Embedded //Field is Embedded and not Referable
	goInterfaces.Detail
	goInterfaces.Tagged
	Value interface{} `json:"value,omitempty"` //any value type
}

func (a Choice) Compare(b Choice) int { //first by name and then by the order
	var r = a.Named.Compare(b.Named)
	if r != 0 { //if name same then no need to check order
		var o = a.Ordered.Compare(b.Ordered)
		if o != 0 {
			return 0
		}
	}
	return r
}

func NewChoice() *Choice {
	s := Choice{}
	return &s
}

func NewChoiceByName(name string) *Choice {
	s := Choice{}
	s.Name = name
	return &s
}

func (w Choice) MarshalJSON() ([]byte, error) {
	//TODO: This method must be overridden
	return json.Marshal(struct {
		Field
		Type0 string `json:"t0"`
	}{
		//add all other fields
		Type0: w.Type0(),
	})
}

func (p Choice) Type0() string {
	return "Choice"
}
