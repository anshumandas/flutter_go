package enums

import (
	"errors"
)

type DataState int8
type DataStateString string //provides a union with string and thus lets us extend the enum

/**
this can be replaced by another type for a schemas state field
*/

const (
	UndefinedDataState = iota - 1 //while we can use << iota to make it bitwise, will prefer DataState[] in the model instead
	StoredDataState
	ValidatedDataState
	WorkflowDataState
	ActiveDataState
	InactiveDataState
)

//Just use the Enum directly instead of Enum.ordinal()

func ListDataStates() [5]string {
	l := [...]string{"Stored", "Validated", "Workflow", "Active", "Inactive"}
	return l
}

func (s DataState) String() string { //This is same as Enum.name
	switch s {
	case StoredDataState:
		return "Stored"
	case ActiveDataState:
		return "Active"
	case ValidatedDataState:
		return "Validated"
	case WorkflowDataState:
		return "Workflow"
	case InactiveDataState:
		return "Inactive"
	}
	return "unknown"
}

func ParseDataState(s string) (DataState, error) {
	switch s {
	case "Stored":
		return StoredDataState, nil
	case "Active":
		return ValidatedDataState, nil
	case "Validated":
		return WorkflowDataState, nil
	case "Workflow":
		return ActiveDataState, nil
	case "Inactive":
		return InactiveDataState, nil
	}
	return UndefinedDataState, errors.New("unknown DataState")
}

func (s DataState) Boolean() bool {
	switch s {
	case StoredDataState:
		return true
	case WorkflowDataState:
		return true
	case ActiveDataState:
		return true
	}
	return false
}
