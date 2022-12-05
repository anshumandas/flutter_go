package schemata

import (
	"time"

	interfaces "github.com/flutter_go/framework/base"
	goInterfaces "github.com/flutter_go/framework/gointerfaces"
	enums "github.com/flutter_go/framework/schemata/enums"
)

//This is used to define the Schema of the Data
//do not use field name starting with _ here or any other schema. it is only in the interfaces package where we use in the json tag

type Schema struct {
	interfaces.Referred                             //Schema is a Data
	goInterfaces.Detail                             //Schema itself is a Detail
	Type                enums.SchemaType            `json:"type"`
	Fields              map[string]Field            `json:"fields,omitempty"`
	Namespace           string                      `json:"ns,omitempty"`
	Plural              string                      `json:"plural,omitempty"`
	Confidentiality     enums.ConfidentialityString `json:"confidentiality"`
	State               enums.SchemaState           `json:"state"`
	Retention           time.Duration               `json:"retention"`
	Workflow            string                      `json:"workflow,omitempty"`
	Within              interfaces.Reference        `json:"within,omitempty"` //Used incase this is a nested schema
	Tags                []string                    `json:"tags,omitempty"`   //can be used as choices for enums and sets
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
