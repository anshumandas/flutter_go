package interfaces

//This is the interface for a Data that can be union of multiple entries in the DB
//goes from v=0 to last to build this struct
//this is mutable unlike DataEntry. However the field mutability is defined in the schema

import (
	"encoding/json"
	"time"

	enums "github.com/flutter_go/framework/base/enums"
	goInterfaces "github.com/flutter_go/framework/gointerfaces"
)

type Referred struct {
	goInterfaces.Versioned                   //holds the current version
	Data                                     //extends by embedding Entry. This holds the current consolidated version only
	DataState0             enums.DataState   `json:"st0"`  //Delete just changes the state of the last version that gets added
	CreatedBy0             string            `json:"cBy0"` //we keep this for easier access
	CreatedOn0             time.Time         `json:"cOn0"`
	LastAction0            Audited           `json:"last0,omitempty"`
	Old0                   []Reference       `json:"old0,omitempty"`
	History0               map[int]DataEntry `json:"history0,omitempty"` //we use map as query need not be from v1. version is the key. Only the changed fields per entry. The containing data should be the one before minVersion asked. If current version them this field is not present.
}

func (d *Referred) String() string {
	res, _ := json.Marshal(d)
	return string(res)
}
