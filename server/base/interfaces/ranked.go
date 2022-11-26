package interfaces

type Ranked struct {
	Rank0 int `json:"rk0"` //used to rank
}

func (a Ranked) Compare(b Ranked) int {
	if a.Rank0 > b.Rank0 {
		return 1
	}
	if a.Rank0 < b.Rank0 {
		return -1
	}
	return 0
}
