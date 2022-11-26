package schemata

import (
	"errors"
)

type Editability int8
type EditabilityString string //provides a union with string and thus lets us extend the enum

const (
	UndefinedEditability          = iota - 1 //while we can use << iota to make it bitwise, will prefer Editability[] in the model instead
	IdEditability                            //editable only on creation, always required, default option
	EditableIfRequiredEditability            //editable only if required on creation
	ImmutableEditability                     //only if not an Id field
	UpdateIfNullEditability                  //only if not a required field, only set if null
	MutableEditability                       //always editable but versioned, default option
	MutableInWorkflowEditability
	UnversionedEditability //always editable and not versioned, should be for system only and not user edited
)

//Just use the Enum directly instead of Enum.ordinal()

func ListEditabilities() [7]string {
	l := [...]string{"Id", "EditableIfRequired", "Immutable", "UpdateIfNull", "Mutable", "MutableInWorkflow", "Unversioned"}
	return l
}

func (s Editability) String() string { //This is same as Enum.name
	switch s {
	case IdEditability:
		return "Id"
	case EditableIfRequiredEditability:
		return "EditableIfRequired"
	case ImmutableEditability:
		return "Immutable"
	case UpdateIfNullEditability:
		return "UpdateIfNull"
	case MutableEditability:
		return "Mutable"
	case MutableInWorkflowEditability:
		return "MutableInWorkflow"
	case UnversionedEditability:
		return "Unversioned"
	}
	return "unknown"
}

func ParseEditability(s string) (Editability, error) {
	switch s {
	case "Id":
		return IdEditability, nil
	case "EditableIfRequired":
		return EditableIfRequiredEditability, nil
	case "Immutable":
		return ImmutableEditability, nil
	case "UpdateIfNull":
		return UpdateIfNullEditability, nil
	case "Mutable":
		return MutableEditability, nil
	case "MutableInWorkflow":
		return MutableInWorkflowEditability, nil
	case "Unversioned":
		return UnversionedEditability, nil
	}
	return UndefinedEditability, errors.New("unknown Editability")
}
