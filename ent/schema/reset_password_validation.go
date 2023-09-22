package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ResetPasswordValidation holds the schema definition for the ResetPasswordValidation entity.
type ResetPasswordValidation struct {
	ent.Schema
}

// Fields of the ResetPasswordValidation.
func (ResetPasswordValidation) Fields() []ent.Field {
	return []ent.Field{
		field.String("email"),
		field.Time("expire_date"),
		field.String("code"),
	}
}

// Edges of the ResetPasswordValidation.
func (ResetPasswordValidation) Edges() []ent.Edge {
	return nil
}
