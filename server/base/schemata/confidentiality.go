package schemata

import (
	"errors"
)

type Confidentiality int8
type ConfidentialityString string //provides a union with string and thus lets us extend the enum

const (
	UndefinedConfidentiality = iota - 1 //while we can use << iota to make it bitwise, will prefer Confidentiality[] in the model instead
	PublicConfidentiality               //visible only if required, default option
	InternalConfidentiality
	ConfidentialConfidentiality
	RestrictedConfidentiality
	SecretConfidentiality
)

//Just use the Enum directly instead of Enum.ordinal()

func ListConfidentialities() [5]string {
	l := [...]string{"Public", "Internal", "Confidential", "Restricted", "Secret"}
	return l
}

func (s Confidentiality) String() string { //This is same as Enum.name
	switch s {
	case PublicConfidentiality:
		return "Public"
	case InternalConfidentiality:
		return "Internal"
	case ConfidentialConfidentiality:
		return "Confidential"
	case RestrictedConfidentiality:
		return "Restricted"
	case SecretConfidentiality:
		return "Secret"
	}
	return "unknown"
}

func ParseConfidentiality(s string) (Confidentiality, error) {
	switch s {
	case "Public":
		return PublicConfidentiality, nil
	case "Internal":
		return InternalConfidentiality, nil
	case "Confidential":
		return ConfidentialConfidentiality, nil
	case "Restricted":
		return RestrictedConfidentiality, nil
	case "Secret":
		return SecretConfidentiality, nil
	}
	return UndefinedConfidentiality, errors.New("unknown Confidentiality")
}
