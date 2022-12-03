package interfaces

type Ordered struct {
	Order0 int `json:"od0"` //used to rank
}

func (a Ordered) Compare(b Ordered) int {
	if a.Order0 > b.Order0 {
		return 1
	}
	if a.Order0 < b.Order0 {
		return -1
	}
	return 0
}
