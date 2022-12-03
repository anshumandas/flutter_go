package interfaces

type Versioned struct {
	Version0 int `json:"v0,omitempty"` //if not defined then take latest
}

func (a Versioned) Compare(b Versioned) int {
	if a.Version0 == b.Version0 {
		return 0
	} else if a.Version0 > b.Version0 {
		return -1
	} else {
		return 1
	}
}
