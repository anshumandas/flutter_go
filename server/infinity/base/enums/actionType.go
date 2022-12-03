package enums

import (
    "errors"
)

type ActionType int8
type ActionTypeString string //provides a union with string and thus lets us extend the enum

const (
	UndefinedActionType = iota - 1 //while we can use << iota to make it bitwise, will prefer ActionType[] in the model instead
	ApproveActionType 
	EditApproveActionType 
	RejectActionType 
	ForwardActionType 
	ReturnActionType 
	ReturnToSourceActionType 
	AcceptActionType 
	AssignActionType 
	UnassignActionType 
	UnacceptActionType 
	SplitActionType 
	MergeActionType 
	ErrorActionType 
	TaskActionType 
	TaskCompletedActionType 
)

//Just use the Enum directly instead of Enum.ordinal()

func ListActionTypes() [15]string {
	l := [...]string{ "Approve", "EditApprove", "Reject", "Forward", "Return", "ReturnToSource", "Accept", "Assign", "Unassign", "Unaccept", "Split", "Merge", "Error", "Task", "TaskCompleted",  }
	return l
}

func (s ActionType) String() string { //This is same as Enum.name 
	switch s {
	case ApproveActionType:
		return "Approve"
	case EditApproveActionType:
		return "EditApprove"
	case RejectActionType:
		return "Reject"
	case ForwardActionType:
		return "Forward"
	case ReturnActionType:
		return "Return"
	case ReturnToSourceActionType:
		return "ReturnToSource"
	case AcceptActionType:
		return "Accept"
	case AssignActionType:
		return "Assign"
	case UnassignActionType:
		return "Unassign"
	case UnacceptActionType:
		return "Unaccept"
	case SplitActionType:
		return "Split"
	case MergeActionType:
		return "Merge"
	case ErrorActionType:
		return "Error"
	case TaskActionType:
		return "Task"
	case TaskCompletedActionType:
		return "TaskCompleted"
	}
	return "unknown"
}

func ParseActionType(s string) (ActionType, error) {
	switch s {
	case "Approve":
		return ApproveActionType, nil
	case "EditApprove":
		return EditApproveActionType, nil
	case "Reject":
		return RejectActionType, nil
	case "Forward":
		return ForwardActionType, nil
	case "Return":
		return ReturnActionType, nil
	case "ReturnToSource":
		return ReturnToSourceActionType, nil
	case "Accept":
		return AcceptActionType, nil
	case "Assign":
		return AssignActionType, nil
	case "Unassign":
		return UnassignActionType, nil
	case "Unaccept":
		return UnacceptActionType, nil
	case "Split":
		return SplitActionType, nil
	case "Merge":
		return MergeActionType, nil
	case "Error":
		return ErrorActionType, nil
	case "Task":
		return TaskActionType, nil
	case "TaskCompleted":
		return TaskCompletedActionType, nil
	}
	return UndefinedActionType, errors.New("unknown ActionType")
}


