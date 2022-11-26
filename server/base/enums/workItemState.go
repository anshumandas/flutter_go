package enums

import (
    "errors"
)

type WorkItemState int8
type WorkItemStateString string //provides a union with string and thus lets us extend the enum

const (
	UndefinedWorkItemState = iota - 1 //while we can use << iota to make it bitwise, will prefer WorkItemState[] in the model instead
	ActiveWorkItemState //in the stage
	EditedWorkItemState //edited and approved by the stage
	ApprovedWorkItemState //approved by the stage
	RejectedWorkItemState 
	RecalledWorkItemState 
	ReturnedWorkItemState 
	ErrorWorkItemState 
	SplitWorkItemState 
	MergeWorkItemState 
	AssignedWorkItemState 
)

//Just use the Enum directly instead of Enum.ordinal()

func ListWorkItemStates() [10]string {
	l := [...]string{ "Active", "Edited", "Approved", "Rejected", "Recalled", "Returned", "Error", "Split", "Merge", "Assigned",  }
	return l
}

func (s WorkItemState) String() string { //This is same as Enum.name 
	switch s {
	case ActiveWorkItemState:
		return "Active"
	case EditedWorkItemState:
		return "Edited"
	case ApprovedWorkItemState:
		return "Approved"
	case RejectedWorkItemState:
		return "Rejected"
	case RecalledWorkItemState:
		return "Recalled"
	case ReturnedWorkItemState:
		return "Returned"
	case ErrorWorkItemState:
		return "Error"
	case SplitWorkItemState:
		return "Split"
	case MergeWorkItemState:
		return "Merge"
	case AssignedWorkItemState:
		return "Assigned"
	}
	return "unknown"
}

func ParseWorkItemState(s string) (WorkItemState, error) {
	switch s {
	case "Active":
		return ActiveWorkItemState, nil
	case "Edited":
		return EditedWorkItemState, nil
	case "Approved":
		return ApprovedWorkItemState, nil
	case "Rejected":
		return RejectedWorkItemState, nil
	case "Recalled":
		return RecalledWorkItemState, nil
	case "Returned":
		return ReturnedWorkItemState, nil
	case "Error":
		return ErrorWorkItemState, nil
	case "Split":
		return SplitWorkItemState, nil
	case "Merge":
		return MergeWorkItemState, nil
	case "Assigned":
		return AssignedWorkItemState, nil
	}
	return UndefinedWorkItemState, errors.New("unknown WorkItemState")
}


