package interfaces

type IdTyped struct {
	Typed
	Id
}

func (a IdTyped) Compare(b IdTyped) int {
	if a.Typed.Compare(b.Typed) && a.Id0 == b.Id0 {
		return 0
	} else {
		return 1
	}
}
