package schemata

import (
	"errors"
)

type Persistence int8
type PersistenceString string //provides a union with string and thus lets us extend the enum

const (
	UndefinedPersistence    = iota - 1 //while we can use << iota to make it bitwise, will prefer Persistence[] in the model instead
	LRUPersistence                     //Last Recently Used
	SessionPersistence                 //till the session is active for the day. Login can be across days but this data will persist only for that session
	ProfilePersistence                 //downloaded on first login. Stays till the user stays logged in. If someone else logins, then data gets overwritten completely
	AppFirstLoadPersistence            //downloaded on first use of the app een before login
	InAppPersistence                   //is part of the app download
)

//Just use the Enum directly instead of Enum.ordinal()

func ListPersistences() [5]string {
	l := [...]string{"LRU", "Session", "Profile", "AppFirstLoad", "InApp"}
	return l
}

func (s Persistence) String() string { //This is same as Enum.name
	switch s {
	case LRUPersistence:
		return "LRU"
	case SessionPersistence:
		return "Session"
	case ProfilePersistence:
		return "Profile"
	case AppFirstLoadPersistence:
		return "AppFirstLoad"
	case InAppPersistence:
		return "InApp"
	}
	return "unknown"
}

func ParsePersistence(s string) (Persistence, error) {
	switch s {
	case "LRU":
		return LRUPersistence, nil
	case "Session":
		return SessionPersistence, nil
	case "Profile":
		return ProfilePersistence, nil
	case "AppFirstLoad":
		return AppFirstLoadPersistence, nil
	case "InApp":
		return InAppPersistence, nil
	}
	return UndefinedPersistence, errors.New("unknown Persistence")
}
