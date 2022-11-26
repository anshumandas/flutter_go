package schemata

import "github.com/flutter_go/base/interfaces"

//This is used to define the Schema of the Data
//do not use field name starting with _ here or any other schema. it is only in the interfaces package where we use in the json tag

type Schema struct {
	interfaces.Referred        //Schema is a Data
	Detail                     //Schema itself is a Detail
	IsColumnar          bool   `json:"isColumnar"`
	TableName           string `json:"tableName"`
	HasEvents           bool   `json:"hasEvents"`
	HasTimeseries       bool   `json:"hasTimeseries"`
	IsRegionSegregated  bool   `json:"isRegionSegregated"`
	IsDateSegregated    bool   `json:"isDateSegregated"`
	Path                string `json:"path"`

	Fields     map[string]Field `json:"fields"`
	Key        string           `json:"key"`
	SchemaType SchemaType       `json:"SchemaType"`
	Version    string           `json:"_version"`

	Datatype     string                 `json:"datatype"`
	DbAlias      string                 `json:"dbAlias"`
	Alias        string                 `json:"alias"`
	SchemaState  SchemaState            `json:"schemaState"`
	StateVersion string                 `json:"stateVersion"`
	Required     string                 `json:"required"`
	Visibility   VisibilityString       `json:"visibility"`
	Editability  EditabilityString      `json:"editability"`
	Default      map[string]interface{} `json:"_default"`
	Indexed      IndexedString          `json:"indexed"`
	Workflow     string                 `json:"workflow"`
	// Permissions  map[string]access.Permission `json:"permissions"`
	Audit string `json:"_audit"`
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
