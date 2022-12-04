package interfaces

import (
	goInterfaces "github.com/flutter_go/framework/gointerfaces"
)

type Reference struct {
	goInterfaces.PKeyed
	goInterfaces.Versioned
	goInterfaces.Summaried
}

func (a Reference) Compare(b Reference) int {
	if a.PKeyed.Compare(b.PKeyed) == 0 {
		//May need version check
		return 0
	} else {
		return a.Summaried.Compare(b.Summaried)
	}
}
