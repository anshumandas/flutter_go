package interfaces

//This is the interface for a Data that can be union of multiple entries in the DB
//goes from v=0 to last to build this struct
//this is mutable unlike DataEntry. However the field mutability is defined in the schema

import (
	"encoding/json"
	"time"
)

type Referred struct {
	Entry               //extends by embedding DataEntry
	CreatedBy string    `json:"cBy0"` //ChangedBy of v=0
	CreatedOn time.Time `json:"cOn0"`
	DeletedBy string    `json:"dBy0,omitempty"` //ChangedBy of v=last
	DeletedOn time.Time `json:"dOn0,omitempty"`
}

func (d *Referred) String() string {
	res, _ := json.Marshal(d)
	return string(res)
}
