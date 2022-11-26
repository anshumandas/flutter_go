package interfaces

type Id struct {
	Id0 string `json:"id0"`
}

func (a Id) Compare(b Id) int {
	if a.Id0 == b.Id0 {
		return 0
	} else {
		return 1
	}
}
