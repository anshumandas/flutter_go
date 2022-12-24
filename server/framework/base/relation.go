package interfaces

import (
	"time"

	"github.com/flutter_go/framework/base/enums"
)

//This is used to define the Field of the Data
//do not use field name starting with _ here or any other schema. it is only in the interfaces package where we use in the json tag

type Relation struct {
	Reference
	DataState0  enums.DataState `json:"st0"`  //Delete just changes the state of the last version that gets added
	AddedBy0    string          `json:"cBy0"` //we keep this for easier access
	AddedOn0    time.Time       `json:"cOn0"`
	LastAction0 Audited         `json:"last0,omitempty"`
}
