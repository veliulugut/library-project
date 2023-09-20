package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Reset_Password_Validation holds the schema definition for the Reset_Password_Validation entity.
type Reset_Password_Validation struct {
	ent.Schema
}

// Fields of the Reset_Password_Validation.
func (Reset_Password_Validation) Fields() []ent.Field {
	return []ent.Field{
		field.String("email"),
		field.Time("expire_date"),
		field.String("code"),
	}
}

// Edges of the Reset_Password_Validation.
func (Reset_Password_Validation) Edges() []ent.Edge {
	return nil
}
