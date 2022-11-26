package schemata

import (
	"errors"
)

type SchemaType int8
type SchemaTypeString string //provides a union with string and thus lets us extend the enum

const (
	UndefinedSchemaType = iota - 1 //while we can use << iota to make it bitwise, will prefer SchemaType[] in the model instead
	SortedSetSchemaType
	ReferrableSchemaType //Referrable can be both Form and Timeseries. Must have an id which is referred
	EmbeddedSchemaType   //part of the main referable entry. In RDBMS may have an internal id which need not be sent in the objectjust name of already defined field
	EnumSchemaType       //not necessary to be in DB as is part of the code
	HashSchemaType       //is a simple key value pair entry
	SetSchemaType        //is a single field entry in the DB
	InterfaceSchemaType  //is just a code gen construct and no db significance
	AbstractSchemaType   //same as interface but means a separate entry in DB with the type field as discriminator.
	EventSchemaType
	AggregatorSchemaType
	ReportSchemaType
	ParserSchemaType
	MappingSchemaType //used for transformations
	WrapperSchemaType
	SchemaSchemaType
	TimeseriesSchemaType //Referrable which hasTimeseries
	FormSchemaType       //must have version and state fields
)

//Just use the Enum directly instead of Enum.ordinal()

func ListSchemaTypes() [17]string {
	l := [...]string{"SortedSet", "Referrable", "Embedded", "Enum", "Hash", "Set", "Interface", "Abstract", "Event", "Aggregator", "Report", "Parser", "Mapping", "Wrapper", "Schema", "Timeseries", "Form"}
	return l
}

func (s SchemaType) String() string { //This is same as Enum.name
	switch s {
	case SortedSetSchemaType:
		return "SortedSet"
	case ReferrableSchemaType:
		return "Referrable"
	case EmbeddedSchemaType:
		return "Embedded"
	case EnumSchemaType:
		return "Enum"
	case HashSchemaType:
		return "Hash"
	case SetSchemaType:
		return "Set"
	case InterfaceSchemaType:
		return "Interface"
	case AbstractSchemaType:
		return "Abstract"
	case EventSchemaType:
		return "Event"
	case AggregatorSchemaType:
		return "Aggregator"
	case ReportSchemaType:
		return "Report"
	case ParserSchemaType:
		return "Parser"
	case MappingSchemaType:
		return "Mapping"
	case WrapperSchemaType:
		return "Wrapper"
	case SchemaSchemaType:
		return "Schema"
	case TimeseriesSchemaType:
		return "Timeseries"
	case FormSchemaType:
		return "Form"
	}
	return "unknown"
}

func ParseSchemaType(s string) (SchemaType, error) {
	switch s {
	case "SortedSet":
		return SortedSetSchemaType, nil
	case "Referrable":
		return ReferrableSchemaType, nil
	case "Embedded":
		return EmbeddedSchemaType, nil
	case "Enum":
		return EnumSchemaType, nil
	case "Hash":
		return HashSchemaType, nil
	case "Set":
		return SetSchemaType, nil
	case "Interface":
		return InterfaceSchemaType, nil
	case "Abstract":
		return AbstractSchemaType, nil
	case "Event":
		return EventSchemaType, nil
	case "Aggregator":
		return AggregatorSchemaType, nil
	case "Report":
		return ReportSchemaType, nil
	case "Parser":
		return ParserSchemaType, nil
	case "Mapping":
		return MappingSchemaType, nil
	case "Wrapper":
		return WrapperSchemaType, nil
	case "Schema":
		return SchemaSchemaType, nil
	case "Timeseries":
		return TimeseriesSchemaType, nil
	case "Form":
		return FormSchemaType, nil
	}
	return UndefinedSchemaType, errors.New("unknown SchemaType")
}
