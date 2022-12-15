package interfaces

import "encoding/json"

type Typed struct {
	type0 string //used as the discriminator
}

func (w Typed) MarshalJSON() ([]byte, error) {
	//TODO: This method must be overridden
	return json.Marshal(struct {
		Typed
		Type0 string `json:"t0"`
	}{
		//add all other fields
		Type0: w.Type0(),
	})
}

func (p Typed) Type0() string {
	return p.type0 //We can also use reflect.TypeOf(). However this may be faster and easier especially if code is generated
}

//This is not needed as every sub class should do a class type setting in their declaration
// func (p *Typed) SetType0(type0 string) {
// 	p.type0 = type0
// }

func (a Typed) Compare(b Typed) bool {
	return a.Type0() == b.Type0()
}
