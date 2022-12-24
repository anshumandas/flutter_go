package enums

import (
	"errors"
)

type WorkflowActionType int8
type WorkflowActionTypeString string //provides a union with string and thus lets us extend the enum

const (
	UndefinedWorkflowActionType = iota - 1 //while we can use << iota to make it bitwise, will prefer WorkflowActionType[] in the model instead
	ApproveWorkflowActionType
	EditApproveWorkflowActionType
	RejectWorkflowActionType
	ForwardWorkflowActionType
	ReturnWorkflowActionType
	ReturnToSourceWorkflowActionType
	AcceptWorkflowActionType
	AssignWorkflowActionType
	UnassignWorkflowActionType
	UnacceptWorkflowActionType
	SplitWorkflowActionType
	MergeWorkflowActionType
	ErrorWorkflowActionType
	TaskWorkflowActionType
	TaskCompletedWorkflowActionType
)

//Just use the Enum directly instead of Enum.ordinal()

func ListWorkflowActionTypes() [15]string {
	l := [...]string{"Approve", "EditApprove", "Reject", "Forward", "Return", "ReturnToSource", "Accept", "Assign", "Unassign", "Unaccept", "Split", "Merge", "Error", "Task", "TaskCompleted"}
	return l
}

func (s WorkflowActionType) String() string { //This is same as Enum.name
	switch s {
	case ApproveWorkflowActionType:
		return "Approve"
	case EditApproveWorkflowActionType:
		return "EditApprove"
	case RejectWorkflowActionType:
		return "Reject"
	case ForwardWorkflowActionType:
		return "Forward"
	case ReturnWorkflowActionType:
		return "Return"
	case ReturnToSourceWorkflowActionType:
		return "ReturnToSource"
	case AcceptWorkflowActionType:
		return "Accept"
	case AssignWorkflowActionType:
		return "Assign"
	case UnassignWorkflowActionType:
		return "Unassign"
	case UnacceptWorkflowActionType:
		return "Unaccept"
	case SplitWorkflowActionType:
		return "Split"
	case MergeWorkflowActionType:
		return "Merge"
	case ErrorWorkflowActionType:
		return "Error"
	case TaskWorkflowActionType:
		return "Task"
	case TaskCompletedWorkflowActionType:
		return "TaskCompleted"
	}
	return "unknown"
}

func ParseWorkflowActionType(s string) (WorkflowActionType, error) {
	switch s {
	case "Approve":
		return ApproveWorkflowActionType, nil
	case "EditApprove":
		return EditApproveWorkflowActionType, nil
	case "Reject":
		return RejectWorkflowActionType, nil
	case "Forward":
		return ForwardWorkflowActionType, nil
	case "Return":
		return ReturnWorkflowActionType, nil
	case "ReturnToSource":
		return ReturnToSourceWorkflowActionType, nil
	case "Accept":
		return AcceptWorkflowActionType, nil
	case "Assign":
		return AssignWorkflowActionType, nil
	case "Unassign":
		return UnassignWorkflowActionType, nil
	case "Unaccept":
		return UnacceptWorkflowActionType, nil
	case "Split":
		return SplitWorkflowActionType, nil
	case "Merge":
		return MergeWorkflowActionType, nil
	case "Error":
		return ErrorWorkflowActionType, nil
	case "Task":
		return TaskWorkflowActionType, nil
	case "TaskCompleted":
		return TaskCompletedWorkflowActionType, nil
	}
	return UndefinedWorkflowActionType, errors.New("unknown WorkflowActionType")
}
