package interfaces

//This is the interface for a single data entry which is immutable
//All Data will embed this struct and have their own immutable fields
//all immutables should be in version = 0
//when we get data it is retrieved from v0 to v specified or current
//each version only contains the edited field
//json TAG is used for DB unless there is a override with DB (or specifc postgre/msql/oracle) TAG

import (
	"encoding/json"
	"time"

	enums "github.com/flutter_go/base/enums"
	goInterfaces "github.com/flutter_go/base/interfaces/go"
)

type Data struct {
	Event
	goInterfaces.PKeyed
	goInterfaces.Versioned
	AddedBy0   string          `json:"aBy0"` //for v=0 its created & other version it is updated. this stored the id only and links to User table
	AddedOn0   time.Time       `json:"aOn0"`
	DataState0 enums.DataState `json:"st0"`          //Delete just changes the state of the last version that gets added
	WorkItem0  string          `json:"w0,omitempty"` //v=0 is creation and other v are mod workflows. Only straight WF. For parallel use documents
}

func (d *Data) String() string {
	res, _ := json.Marshal(d)
	return string(res)
}
