package interfaces

import (
	"strings"
)

type PKeyed struct {
	PrimaryKey0 string `json:"pk0"`
}

func (a PKeyed) Compare(b PKeyed) int {
	return strings.Compare(a.PrimaryKey0, b.PrimaryKey0)
}
