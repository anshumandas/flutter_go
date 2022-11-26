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
	InCreateWorkflowDataState 
	ActiveDataState 
	InModifyWorkflowDataState 
	InDeleteWorkflowDataState 
	InactiveDataState 
)

//Just use the Enum directly instead of Enum.ordinal()

func ListDataStates() [5]string {
	l := [...]string{ "InCreateWorkflow", "Active", "InModifyWorkflow", "InDeleteWorkflow", "Inactive",  }
	return l
}

func (s DataState) String() string { //This is same as Enum.name 
	switch s {
	case InCreateWorkflowDataState:
		return "InCreateWorkflow"
	case ActiveDataState:
		return "Active"
	case InModifyWorkflowDataState:
		return "InModifyWorkflow"
	case InDeleteWorkflowDataState:
		return "InDeleteWorkflow"
	case InactiveDataState:
		return "Inactive"
	}
	return "unknown"
}

func ParseDataState(s string) (DataState, error) {
	switch s {
	case "InCreateWorkflow":
		return InCreateWorkflowDataState, nil
	case "Active":
		return ActiveDataState, nil
	case "InModifyWorkflow":
		return InModifyWorkflowDataState, nil
	case "InDeleteWorkflow":
		return InDeleteWorkflowDataState, nil
	case "Inactive":
		return InactiveDataState, nil
	}
	return UndefinedDataState, errors.New("unknown DataState")
}


func (s DataState) Boolean() bool {
	switch s {
	case InCreateWorkflowDataState:
		return true
	case InModifyWorkflowDataState:
		return true
	case InDeleteWorkflowDataState:
		return true
	}
	return false
}
