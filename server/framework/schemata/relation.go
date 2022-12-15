package schemata

//This is used to define the Field of the Data
//do not use field name starting with _ here or any other schema. it is only in the interfaces package where we use in the json tag

type Relation struct {
	Field
	ReverseName        string `json:"rName,omitempty"`
	LinkedSchema       string `json:"schema,omitempty"`
	IsEmbedded         bool   `json:"embed,omitempty"`
	CounterName        string `json:"counter,omitempty"`
	ReverseCounterName string `json:"rCounter,omitempty"`
}
