package interfaces

import "strings"

type Named struct {
	Name string `json:"name"`
}

func (a Named) Compare(b Named) int {
	return strings.Compare(b.Name, b.Name)
}
