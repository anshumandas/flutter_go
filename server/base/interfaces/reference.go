package interfaces

import "strings"

type Reference struct {
	IdTyped
	PrimaryKey0 string   `json:"pk0"`          //combination of Type and Id
	Summary0    []string `json:"s0,omitempty"` //used with opaque Ids, can have more than one fields. The display pattern should be in the UX system and not heres. We do not need the field names as that is part of schema def
	Version0    int      `json:"v0,omitempty"` //if not defined then take latest
}

func (a Reference) Compare(b Reference) int {
	if a.PrimaryKey0 == b.PrimaryKey0 {
		//May need version check
		return 0
	} else {
		var i = strings.Compare(strings.Join(a.Summary0, ","), strings.Join(b.Summary0, ","))
		if i == 0 {
			return 1 //even if summary same but id is different and hence cannot be considered same
		} else {
			return i
		}
	}
}
