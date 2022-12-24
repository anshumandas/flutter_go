package schemata

import (
	"time"

	interfaces "github.com/flutter_go/framework/base"
	enums "github.com/flutter_go/framework/schemata/enums"
)

//This is used to define the Schema of the Data
//do not use field name starting with _ here or any other schema. it is only in the interfaces package where we use in the json tag

type FormSchema struct {
	Fields          map[string]Field            `json:"fields,omitempty"`
	Relations       map[string]Relationship     `json:"relations,omitempty"`
	Groups          []Group                     `json:"group,omitempty"`
	Inherits        []interfaces.Reference      `json:"inherits,omitempty"`
	InheritedBy     []interfaces.Reference      `json:"inheritedBy,omitempty"`
	Confidentiality enums.ConfidentialityString `json:"confidentiality,omitempty"`
	Retention       time.Duration               `json:"retention,omitempty"` //nano to 290 year possible. we will use milli to 10 years. Can be an enum instead
	Workflow        string                      `json:"workflow,omitempty"`
	Within          interfaces.Reference        `json:"within,omitempty"` //Used incase this is a nested schema //this is also a relation

}
