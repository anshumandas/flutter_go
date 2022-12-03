package schemata

//This is used to define the Field of the Data
//do not use field name starting with _ here or any other schema. it is only in the interfaces package where we use in the json tag

type Relation struct {
	Field
	LinkedType string
	Embedded   bool
	//field could be a Referred schema but in an embeded form
}