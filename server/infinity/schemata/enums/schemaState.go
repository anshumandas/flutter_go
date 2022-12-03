package schemata

import (
	"errors"
)

type SchemaState int8
type SchemaStateString string //provides a union with string and thus lets us extend the enum

const (
	UndefinedSchemaState = iota - 1 //while we can use << iota to make it bitwise, will prefer SchemaState[] in the model instead
	ActiveSchemaState
	ProposedSchemaState
	RejectedSchemaState
	DeprecatedSchemaState
	DroppedSchemaState
)

//Just use the Enum directly instead of Enum.ordinal()

func ListSchemaStates() [5]string {
	l := [...]string{"Active", "Proposed", "Rejected", "Deprecated", "Dropped"}
	return l
}

func (s SchemaState) String() string { //This is same as Enum.name
	switch s {
	case ActiveSchemaState:
		return "Active"
	case ProposedSchemaState:
		return "Proposed"
	case RejectedSchemaState:
		return "Rejected"
	case DeprecatedSchemaState:
		return "Deprecated"
	case DroppedSchemaState:
		return "Dropped"
	}
	return "unknown"
}

func ParseSchemaState(s string) (SchemaState, error) {
	switch s {
	case "Active":
		return ActiveSchemaState, nil
	case "Proposed":
		return ProposedSchemaState, nil
	case "Rejected":
		return RejectedSchemaState, nil
	case "Deprecated":
		return DeprecatedSchemaState, nil
	case "Dropped":
		return DroppedSchemaState, nil
	}
	return UndefinedSchemaState, errors.New("unknown SchemaState")
}
