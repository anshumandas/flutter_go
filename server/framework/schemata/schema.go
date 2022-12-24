package schemata

import (
	"encoding/json"

	interfaces "github.com/flutter_go/framework/base"
	goInterfaces "github.com/flutter_go/framework/gointerfaces"
	enums "github.com/flutter_go/framework/schemata/enums"
)

//This is used to define the Schema of the Data
//do not use field name starting with _ here or any other schema. it is only in the interfaces package where we use in the json tag

type Schema struct {
	interfaces.Referred //Schema is a Data
	goInterfaces.Detail //Schema itself is a Detail
	goInterfaces.Tagged
	Type          enums.SchemaType  `json:"type"`
	Namespace     string            `json:"ns,omitempty"`
	Plural        string            `json:"plural,omitempty"`
	State         enums.SchemaState `json:"state"`
	UIPersistence enums.Persistence `json:"ui,omitempty"`

	FormSchema    `visibility_Type:"-EnumSchema,-SetSchema"` //Group visibility for Type. - means exclude. can be comma delimited
	ChoicesSchema `visibility_Type:"EnumSchema,SetSchema"`   //Group visibility for Type. - means exclude. can be comma delimited
}

func New() *Schema {
	s := Schema{}
	return &s
}

func NewByName(name string) *Schema {
	s := Schema{}
	s.Name = name
	return &s
}

func NewByDetail(detail goInterfaces.Detail) *Schema {
	s := Schema{}
	s.Detail = detail
	return &s
}

func (w Schema) MarshalJSON() ([]byte, error) {
	//TODO: This method must be overridden
	return json.Marshal(struct {
		Schema
		Type0 string `json:"t0"`
	}{
		//add all other fields
		Type0: w.Type0(),
	})
}

func (p Schema) Type0() string {
	return "Schema#" + p.Type.String()
}
