package schemata

// This is used for workflow branching.
// It is a schema that enables combination of schemas with different workflows
// Thus you can have specific fields that initiate a workflow
// The document gets stored as multiple Data entrys

type Document struct {
	Schema                       //Document itself is a Schema
	Sections map[string]Document `json:"sections,omitempty"` //contains sections of Documents. For parallel flow must create sections that can be worked in parallel
}

// The generated class section variables (map key as names) of defined schema
