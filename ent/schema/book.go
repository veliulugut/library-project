package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("author"),
		field.String("genre"),
		field.String("height"),
		field.String("publisher"),
		field.Time("created_at").Default(time.Now().UTC()),
		field.Time("updated_at").Default(time.Now().UTC()),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return nil
}
