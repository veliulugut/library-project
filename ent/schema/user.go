package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email"),
		field.String("password"),
		field.String("username").Unique(),
		field.Time("created_at"),
		field.Time("updated_at").Default(time.Now().UTC()),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
