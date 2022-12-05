package schemata

import (
	"errors"
)

type Editability int8

const (
	UndefinedEditability         = iota - 1 //while we can use << iota to make it bitwise, will prefer Editability[] in the model instead
	KeyFieldEditability                     //editable only on creation, always required, default option
	SummaryFieldEditability                 //editable only if required on creation
	ImmutableEditability                    //only if not an Id field
	UpdateIfNullEditability                 //only if not a required field, only set if null
	MutableEditability                      //always editable but versioned, default option
	MutableInWorkflowEditability            //allows workflow editability
	SystemFieldEditability                  //always editable and not versioned, should be for system only and not user edited
)

//Just use the Enum directly instead of Enum.ordinal()

func ListEditabilities() [7]string {
	l := [...]string{"Id", "EditableIfRequired", "Immutable", "UpdateIfNull", "Mutable", "MutableInWorkflow", "Unversioned"}
	return l
}

func (s Editability) String() string { //This is same as Enum.name
	switch s {
	case KeyFieldEditability:
		return "Id"
	case SummaryFieldEditability:
		return "EditableIfRequired"
	case ImmutableEditability:
		return "Immutable"
	case UpdateIfNullEditability:
		return "UpdateIfNull"
	case MutableEditability:
		return "Mutable"
	case MutableInWorkflowEditability:
		return "MutableInWorkflow"
	case SystemFieldEditability:
		return "Unversioned"
	}
	return "unknown"
}

func ParseEditability(s string) (Editability, error) {
	switch s {
	case "Id":
		return KeyFieldEditability, nil
	case "EditableIfRequired":
		return SummaryFieldEditability, nil
	case "Immutable":
		return ImmutableEditability, nil
	case "UpdateIfNull":
		return UpdateIfNullEditability, nil
	case "Mutable":
		return MutableEditability, nil
	case "MutableInWorkflow":
		return MutableInWorkflowEditability, nil
	case "Unversioned":
		return SystemFieldEditability, nil
	}
	return UndefinedEditability, errors.New("unknown Editability")
}
