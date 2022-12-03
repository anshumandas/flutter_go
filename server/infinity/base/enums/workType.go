package enums

import (
    "errors"
)

type WorkType int8
type WorkTypeString string //provides a union with string and thus lets us extend the enum

const (
	UndefinedWorkType = iota - 1 //while we can use << iota to make it bitwise, will prefer WorkType[] in the model instead
	ReadWorkType 
	CreateWorkType 
	UpdateWorkType 
	DeactivateWorkType 
	ReactivateWorkType 
	UploadWorkType 
	DownloadWorkType 
	PrintWorkType 
	ShareWorkType 
	AllWorkType 
)

//Just use the Enum directly instead of Enum.ordinal()

func ListWorkTypes() [10]string {
	l := [...]string{ "Read", "Create", "Update", "Deactivate", "Reactivate", "Upload", "Download", "Print", "Share", "All",  }
	return l
}

func (s WorkType) String() string { //This is same as Enum.name 
	switch s {
	case ReadWorkType:
		return "Read"
	case CreateWorkType:
		return "Create"
	case UpdateWorkType:
		return "Update"
	case DeactivateWorkType:
		return "Deactivate"
	case ReactivateWorkType:
		return "Reactivate"
	case UploadWorkType:
		return "Upload"
	case DownloadWorkType:
		return "Download"
	case PrintWorkType:
		return "Print"
	case ShareWorkType:
		return "Share"
	case AllWorkType:
		return "All"
	}
	return "unknown"
}

func ParseWorkType(s string) (WorkType, error) {
	switch s {
	case "Read":
		return ReadWorkType, nil
	case "Create":
		return CreateWorkType, nil
	case "Update":
		return UpdateWorkType, nil
	case "Deactivate":
		return DeactivateWorkType, nil
	case "Reactivate":
		return ReactivateWorkType, nil
	case "Upload":
		return UploadWorkType, nil
	case "Download":
		return DownloadWorkType, nil
	case "Print":
		return PrintWorkType, nil
	case "Share":
		return ShareWorkType, nil
	case "All":
		return AllWorkType, nil
	}
	return UndefinedWorkType, errors.New("unknown WorkType")
}


