package schemata

import (
	"errors"
)

type Indexed int8
type IndexedString string //provides a union with string and thus lets us extend the enum

/**
nullable : Id is always indexed
*/

/**
Unique combined key, if this is PK then | is prefix else suffix
*/

const (
	UndefinedIndexed = iota - 1 //while we can use << iota to make it bitwise, will prefer Indexed[] in the model instead
	AscendingIndexed
	DescendingIndexed
	UniqueAscendingIndexed
	UniqueDescendingIndexed
	TextIndexed //this is used for complex text search, may mean add to another DB like elastiSearch
)

//Just use the Enum directly instead of Enum.ordinal()

func ListIndexeds() [5]string {
	l := [...]string{"Ascending", "Descending", "UniqueAscending", "UniqueDescending", "Text"}
	return l
}

func (s Indexed) String() string { //This is same as Enum.name
	switch s {
	case AscendingIndexed:
		return "Ascending"
	case DescendingIndexed:
		return "Descending"
	case UniqueAscendingIndexed:
		return "UniqueAscending"
	case UniqueDescendingIndexed:
		return "UniqueDescending"
	case TextIndexed:
		return "Text"
	}
	return "unknown"
}

func ParseIndexed(s string) (Indexed, error) {
	switch s {
	case "Ascending":
		return AscendingIndexed, nil
	case "Descending":
		return DescendingIndexed, nil
	case "UniqueAscending":
		return UniqueAscendingIndexed, nil
	case "UniqueDescending":
		return UniqueDescendingIndexed, nil
	case "Text":
		return TextIndexed, nil
	}
	return UndefinedIndexed, errors.New("unknown Indexed")
}
