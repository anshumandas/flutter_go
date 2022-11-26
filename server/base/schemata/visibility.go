package schemata

import (
	"errors"
)

type Visibility int8
type VisibilityString string //provides a union with string and thus lets us extend the enum

const (
	UndefinedVisibility         = iota - 1 //while we can use << iota to make it bitwise, will prefer Visibility[] in the model instead
	VisibleIfRequiredVisibility            //visible only if required, default option
	VisibleVisibility
	VisibleIfNullVisibility
	VisibleIfNotNullVisibility
	HiddenVisibility
	MaskedVisibility
)

//Just use the Enum directly instead of Enum.ordinal()

func ListVisibilities() [6]string {
	l := [...]string{"VisibleIfRequired", "Visible", "VisibleIfNull", "VisibleIfNotNull", "Hidden", "Masked"}
	return l
}

func (s Visibility) String() string { //This is same as Enum.name
	switch s {
	case VisibleIfRequiredVisibility:
		return "VisibleIfRequired"
	case VisibleVisibility:
		return "Visible"
	case VisibleIfNullVisibility:
		return "VisibleIfNull"
	case VisibleIfNotNullVisibility:
		return "VisibleIfNotNull"
	case HiddenVisibility:
		return "Hidden"
	case MaskedVisibility:
		return "Masked"
	}
	return "unknown"
}

func ParseVisibility(s string) (Visibility, error) {
	switch s {
	case "VisibleIfRequired":
		return VisibleIfRequiredVisibility, nil
	case "Visible":
		return VisibleVisibility, nil
	case "VisibleIfNull":
		return VisibleIfNullVisibility, nil
	case "VisibleIfNotNull":
		return VisibleIfNotNullVisibility, nil
	case "Hidden":
		return HiddenVisibility, nil
	case "Masked":
		return MaskedVisibility, nil
	}
	return UndefinedVisibility, errors.New("unknown Visibility")
}
