package interfaces

type Typed struct {
	Type0 string `json:"t0"` //used as the discriminator
}

func (a Typed) Compare(b Typed) bool {
	return a.Type0 == b.Type0
}
