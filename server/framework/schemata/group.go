package schemata

import (
	interfaces "github.com/flutter_go/framework/gointerfaces"
	enums "github.com/flutter_go/framework/schemata/enums"
)

//This is used to define the Field of the Data
//do not use field name starting with _ here or any other schema. it is only in the interfaces package where we use in the json tag

type Group struct {
	//whenever we see this type the UX should allow choosing one of the sub classes
	interfaces.Detail //represented as tag => name_key:value
}

type OneOfGroup struct {
	Group
	Fields []string //name of the fields that is linked as key
}

type AnyOfGroup struct {
	Group
	Fields []string //name of the fields that is linked as key
}

type VisibilityGroup struct {
	Group
	Visibility map[string]string //name of the fields that is linked as key. can be comma delimited. Condition as value
}

type ContentGroup struct {
	Group
	Content map[string]string //name of the fields that is linked as key. can be comma delimited. Condition as value
}

type IndexGroup struct {
	Group
	Fields    []string //name of the fields that is linked as key.
	IndexType enums.IndexedString
}
