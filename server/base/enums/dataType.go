package enums

import (
    "errors"
)

type DataType int8
type DataTypeString string //provides a union with string and thus lets us extend the enum

/**
nullable : @ means reference (full PK), # means tag set name; !# means not in the set; Name without any symbol can be enum or embedded; $xyz shows format; | is either of two or more types. Just adding | means String or the type
*/

const (
	UndefinedDataType = iota - 1 //while we can use << iota to make it bitwise, will prefer DataType[] in the model instead
	BooleanDataType 
	StringDataType 
	ShortDataType //16 bit
	IntDataType //32 bit
	WholeNumDataType 
	NaturalNumDataType 
	LongDataType //64 bit
	FloatDataType //32 bit
	DoubleDataType //64 bit
	EnumDataType 
	ObjectDataType //denotes any complex type with schema, stored same as hash
	DateTimeDataType 
	DateDataType 
	TimeDataType //used for recurring cron etc
	TextDataType 
	RichTextDataType 
	TupleDataType //(label
	ConditionDataType //same as boolean|string|[set of string that computes to boolean with and]. just a field name means null check
	LinkDataType //is string that symbolises a link from the containing class
	ListDataType //linked list; denoted as [$type]; can have duplicate values
	SetDataType //unsorted set, denoted as &lt;$type&gt;; cannot have duplicate values
	OsetDataType //ordered set or array list, denoted as &lt;[$type]&gt;; cannot have duplicate values; has index
	SsetDataType //sorted set denoted as &lt;&lt;$type&gt;&gt;; can have duplicate values
	WsetDataType //weighted&#x2F;scored set denoted as &lt;($type, $score_field_in_type)&gt;; can have duplicate values; only type_id and score_field values are saved
	HashDataType //map&#x2F;dictionary&#x2F;hash of {$value_type, $key_type (omit for string) } pairs
	AnyDataType //denotes any type above. stored as a (JSON) string
)

//Just use the Enum directly instead of Enum.ordinal()

func ListDataTypes() [26]string {
	l := [...]string{ "boolean", "string", "short", "int", "whole_num", "natural_num", "long", "float", "double", "enum", "object", "dateTime", "date", "time", "text", "richText", "tuple", "condition", "link", "list", "set", "oset", "sset", "wset", "hash", "any",  }
	return l
}

func (s DataType) String() string { //This is same as Enum.name 
	switch s {
	case BooleanDataType:
		return "boolean"
	case StringDataType:
		return "string"
	case ShortDataType:
		return "short"
	case IntDataType:
		return "int"
	case WholeNumDataType:
		return "whole_num"
	case NaturalNumDataType:
		return "natural_num"
	case LongDataType:
		return "long"
	case FloatDataType:
		return "float"
	case DoubleDataType:
		return "double"
	case EnumDataType:
		return "enum"
	case ObjectDataType:
		return "object"
	case DateTimeDataType:
		return "dateTime"
	case DateDataType:
		return "date"
	case TimeDataType:
		return "time"
	case TextDataType:
		return "text"
	case RichTextDataType:
		return "richText"
	case TupleDataType:
		return "tuple"
	case ConditionDataType:
		return "condition"
	case LinkDataType:
		return "link"
	case ListDataType:
		return "list"
	case SetDataType:
		return "set"
	case OsetDataType:
		return "oset"
	case SsetDataType:
		return "sset"
	case WsetDataType:
		return "wset"
	case HashDataType:
		return "hash"
	case AnyDataType:
		return "any"
	}
	return "unknown"
}

func ParseDataType(s string) (DataType, error) {
	switch s {
	case "boolean":
		return BooleanDataType, nil
	case "string":
		return StringDataType, nil
	case "short":
		return ShortDataType, nil
	case "int":
		return IntDataType, nil
	case "whole_num":
		return WholeNumDataType, nil
	case "natural_num":
		return NaturalNumDataType, nil
	case "long":
		return LongDataType, nil
	case "float":
		return FloatDataType, nil
	case "double":
		return DoubleDataType, nil
	case "enum":
		return EnumDataType, nil
	case "object":
		return ObjectDataType, nil
	case "dateTime":
		return DateTimeDataType, nil
	case "date":
		return DateDataType, nil
	case "time":
		return TimeDataType, nil
	case "text":
		return TextDataType, nil
	case "richText":
		return RichTextDataType, nil
	case "tuple":
		return TupleDataType, nil
	case "condition":
		return ConditionDataType, nil
	case "link":
		return LinkDataType, nil
	case "list":
		return ListDataType, nil
	case "set":
		return SetDataType, nil
	case "oset":
		return OsetDataType, nil
	case "sset":
		return SsetDataType, nil
	case "wset":
		return WsetDataType, nil
	case "hash":
		return HashDataType, nil
	case "any":
		return AnyDataType, nil
	}
	return UndefinedDataType, errors.New("unknown DataType")
}


