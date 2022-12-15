package schemata

//This is used to define the Schema of the Data
//do not use field name starting with _ here or any other schema. it is only in the interfaces package where we use in the json tag

type ChoicesSchema struct {
	Choices map[string]Choice `json:"choices,omitempty"`
}
