package schemata

import "github.com/flutter_go/base/interfaces"

//This is used for workflow branching

type Document struct {
	interfaces.Referred            //Document itself is a Data
	Sections            []Document `json:"_sctn,omitempty"` //contains sections of Documents. For parallel flow must create sections that can be worked in parallel
}
