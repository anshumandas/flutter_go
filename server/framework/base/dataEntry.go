package interfaces

//This is the interface for a single data entry which is immutable
//All Data will embed this struct and have their own immutable fields
//all immutables should be in version = 0
//when we get data it is retrieved from v0 to v specified or current
//each version only contains the edited field
//json TAG is used for DB unless there is a override with DB (or specifc postgre/msql/oracle) TAG

import (
	"encoding/json"
)

type DataEntry struct {
	Data
	Audited
}

func (d *DataEntry) String() string {
	res, _ := json.Marshal(d)
	return string(res)
}
