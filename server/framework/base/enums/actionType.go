package enums

import (
	"errors"
)

type ActionType int8
type ActionTypeString string //provides a union with string and thus lets us extend the enum

const (
	UndefinedActionType = iota - 1 //while we can use << iota to make it bitwise, will prefer ActionType[] in the model instead
	CreateActionType
	ReadActionType
	UpdateActionType
	DeleteActionType
	ListActionType
	AddActionType
	RemoveActionType
)

//Just use the Enum directly instead of Enum.ordinal()

func ListActionTypes() [7]string {
	l := [...]string{"Create", "Read", "Update", "Delete", "List", "Add", "Remove"}
	return l
}

func (s ActionType) String() string { //This is same as Enum.name
	switch s {
	case CreateActionType:
		return "Create"
	case ReadActionType:
		return "Read"
	case UpdateActionType:
		return "Update"
	case DeleteActionType:
		return "Delete"
	case ListActionType:
		return "List"
	case AddActionType:
		return "Add"
	case RemoveActionType:
		return "Remove"
	}
	return "unknown"
}

func ParseActionType(s string) (ActionType, error) {
	switch s {
	case "Create":
		return CreateActionType, nil
	case "Read":
		return ReadActionType, nil
	case "Update":
		return UpdateActionType, nil
	case "Delete":
		return DeleteActionType, nil
	case "List":
		return ListActionType, nil
	case "Add":
		return AddActionType, nil
	case "Remove":
		return RemoveActionType, nil
	}
	return UndefinedActionType, errors.New("unknown ActionType")
}
