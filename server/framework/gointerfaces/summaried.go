package interfaces

import (
	"strings"
)

type Summaried struct {
	Summary0 []string `json:"s0,omitempty"` //used with opaque Ids, can have more than one fields. The display pattern should be in the UX system and not heres. We do not need the field names as that is part of schema def
}

func (a Summaried) Compare(b Summaried) int {
	var i = strings.Compare(strings.Join(a.Summary0, ","), strings.Join(b.Summary0, ","))
	if i == 0 {
		return 1 //even if summary same but id is different and hence cannot be considered same
	} else {
		return i
	}
}
