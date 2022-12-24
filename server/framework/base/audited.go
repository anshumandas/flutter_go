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

	"github.com/flutter_go/framework/base/enums"
	goInterfaces "github.com/flutter_go/framework/gointerfaces"
)

type Audited struct {
	goInterfaces.Versioned
	Action0   enums.ActionType `json:"ac0"`
	By0       string           `json:"aBy0"` //for v=0 its created & other version it is updated. this stored the id only and links to User table
	On0       time.Time        `json:"aOn0"`
	WorkItem0 string           `json:"w0,omitempty"` //v=0 is creation and other v are mod workflows. Only straight WF. For parallel use documents
}

func (d *Audited) String() string {
	res, _ := json.Marshal(d)
	return string(res)
}
