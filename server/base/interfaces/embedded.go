package interfaces

type Embedded struct {
	Data
	// Ranked
	// Named
	SecondaryKey0 string `json:"sk0"` //self
}

// func (a Embedded) Compare(b Embedded) int {
// 	var r = a.Ranked.Compare(b.Ranked)
// 	if r == 0 {
// 		r = a.Named.Compare(b.Named)
// 	}
// 	return r
// }
