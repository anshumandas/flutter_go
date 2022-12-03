package schemata

import (
	"time"

	interfaces "github.com/flutter_go/infinity/base"
	enums "github.com/flutter_go/infinity/schemata/enums"
)

//This is used to define the Schema of the Data
//do not use field name starting with _ here or any other schema. it is only in the interfaces package where we use in the json tag

type Schema struct {
	interfaces.Referred                             //Schema is a Data
	Detail                                          //Schema itself is a Detail
	Fields              map[string]Field            `json:"fields"`
	Namespace           string                      `json:"ns"`
	Type                enums.SchemaType            `json:"type"`
	Plural              string                      `json:"plural"`
	Confidentiality     enums.ConfidentialityString `json:"confidentiality"`
	State               enums.SchemaState           `json:"state"`
	Retention           time.Duration               `json:"retention"`
	Workflow            string                      `json:"workflow,omitempty"`
	Within              interfaces.Reference        `json:"within"` //Used incase this is a nested schema
	Tags                []string                    `json:"tags"`
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

func NewByDetail(detail Detail) *Schema {
	s := Schema{}
	s.Detail = detail
	return &s
}
