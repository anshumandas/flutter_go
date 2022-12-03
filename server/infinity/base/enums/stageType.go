package enums

import (
    "errors"
)

type StageType int8
type StageTypeString string //provides a union with string and thus lets us extend the enum

const (
	UndefinedStageType = iota - 1 //while we can use << iota to make it bitwise, will prefer StageType[] in the model instead
	PersonStageType 
	TeamStageType 
	SplitStageType //use to open a parallel flow. increments the same split counter for each split at each split stage
	JoinStageType //use to close a parallel flow. checks if all parts have reached to push forward
	DelayStageType 
	TaskStageType //create a task and send to Task Action if no active task of same cause and collate all items coming to this stage till active task completed. once done approve all together
	ServiceStageType 
)

//Just use the Enum directly instead of Enum.ordinal()

func ListStageTypes() [7]string {
	l := [...]string{ "Person", "Team", "Split", "Join", "Delay", "Task", "Service",  }
	return l
}

func (s StageType) String() string { //This is same as Enum.name 
	switch s {
	case PersonStageType:
		return "Person"
	case TeamStageType:
		return "Team"
	case SplitStageType:
		return "Split"
	case JoinStageType:
		return "Join"
	case DelayStageType:
		return "Delay"
	case TaskStageType:
		return "Task"
	case ServiceStageType:
		return "Service"
	}
	return "unknown"
}

func ParseStageType(s string) (StageType, error) {
	switch s {
	case "Person":
		return PersonStageType, nil
	case "Team":
		return TeamStageType, nil
	case "Split":
		return SplitStageType, nil
	case "Join":
		return JoinStageType, nil
	case "Delay":
		return DelayStageType, nil
	case "Task":
		return TaskStageType, nil
	case "Service":
		return ServiceStageType, nil
	}
	return UndefinedStageType, errors.New("unknown StageType")
}


