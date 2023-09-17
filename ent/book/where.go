// Code generated by ent, DO NOT EDIT.

package book

import (
	"library/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldTitle, v))
}

// Author applies equality check predicate on the "author" field. It's identical to AuthorEQ.
func Author(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldAuthor, v))
}

// Genre applies equality check predicate on the "genre" field. It's identical to GenreEQ.
func Genre(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldGenre, v))
}

// Height applies equality check predicate on the "height" field. It's identical to HeightEQ.
func Height(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldHeight, v))
}

// Publisher applies equality check predicate on the "publisher" field. It's identical to PublisherEQ.
func Publisher(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldPublisher, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldUpdatedAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldTitle, v))
}

// AuthorEQ applies the EQ predicate on the "author" field.
func AuthorEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldAuthor, v))
}

// AuthorNEQ applies the NEQ predicate on the "author" field.
func AuthorNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldAuthor, v))
}

// AuthorIn applies the In predicate on the "author" field.
func AuthorIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldAuthor, vs...))
}

// AuthorNotIn applies the NotIn predicate on the "author" field.
func AuthorNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldAuthor, vs...))
}

// AuthorGT applies the GT predicate on the "author" field.
func AuthorGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldAuthor, v))
}

// AuthorGTE applies the GTE predicate on the "author" field.
func AuthorGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldAuthor, v))
}

// AuthorLT applies the LT predicate on the "author" field.
func AuthorLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldAuthor, v))
}

// AuthorLTE applies the LTE predicate on the "author" field.
func AuthorLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldAuthor, v))
}

// AuthorContains applies the Contains predicate on the "author" field.
func AuthorContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldAuthor, v))
}

// AuthorHasPrefix applies the HasPrefix predicate on the "author" field.
func AuthorHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldAuthor, v))
}

// AuthorHasSuffix applies the HasSuffix predicate on the "author" field.
func AuthorHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldAuthor, v))
}

// AuthorEqualFold applies the EqualFold predicate on the "author" field.
func AuthorEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldAuthor, v))
}

// AuthorContainsFold applies the ContainsFold predicate on the "author" field.
func AuthorContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldAuthor, v))
}

// GenreEQ applies the EQ predicate on the "genre" field.
func GenreEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldGenre, v))
}

// GenreNEQ applies the NEQ predicate on the "genre" field.
func GenreNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldGenre, v))
}

// GenreIn applies the In predicate on the "genre" field.
func GenreIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldGenre, vs...))
}

// GenreNotIn applies the NotIn predicate on the "genre" field.
func GenreNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldGenre, vs...))
}

// GenreGT applies the GT predicate on the "genre" field.
func GenreGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldGenre, v))
}

// GenreGTE applies the GTE predicate on the "genre" field.
func GenreGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldGenre, v))
}

// GenreLT applies the LT predicate on the "genre" field.
func GenreLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldGenre, v))
}

// GenreLTE applies the LTE predicate on the "genre" field.
func GenreLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldGenre, v))
}

// GenreContains applies the Contains predicate on the "genre" field.
func GenreContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldGenre, v))
}

// GenreHasPrefix applies the HasPrefix predicate on the "genre" field.
func GenreHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldGenre, v))
}

// GenreHasSuffix applies the HasSuffix predicate on the "genre" field.
func GenreHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldGenre, v))
}

// GenreEqualFold applies the EqualFold predicate on the "genre" field.
func GenreEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldGenre, v))
}

// GenreContainsFold applies the ContainsFold predicate on the "genre" field.
func GenreContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldGenre, v))
}

// HeightEQ applies the EQ predicate on the "height" field.
func HeightEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldHeight, v))
}

// HeightNEQ applies the NEQ predicate on the "height" field.
func HeightNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldHeight, v))
}

// HeightIn applies the In predicate on the "height" field.
func HeightIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldHeight, vs...))
}

// HeightNotIn applies the NotIn predicate on the "height" field.
func HeightNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldHeight, vs...))
}

// HeightGT applies the GT predicate on the "height" field.
func HeightGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldHeight, v))
}

// HeightGTE applies the GTE predicate on the "height" field.
func HeightGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldHeight, v))
}

// HeightLT applies the LT predicate on the "height" field.
func HeightLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldHeight, v))
}

// HeightLTE applies the LTE predicate on the "height" field.
func HeightLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldHeight, v))
}

// HeightContains applies the Contains predicate on the "height" field.
func HeightContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldHeight, v))
}

// HeightHasPrefix applies the HasPrefix predicate on the "height" field.
func HeightHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldHeight, v))
}

// HeightHasSuffix applies the HasSuffix predicate on the "height" field.
func HeightHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldHeight, v))
}

// HeightEqualFold applies the EqualFold predicate on the "height" field.
func HeightEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldHeight, v))
}

// HeightContainsFold applies the ContainsFold predicate on the "height" field.
func HeightContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldHeight, v))
}

// PublisherEQ applies the EQ predicate on the "publisher" field.
func PublisherEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldPublisher, v))
}

// PublisherNEQ applies the NEQ predicate on the "publisher" field.
func PublisherNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldPublisher, v))
}

// PublisherIn applies the In predicate on the "publisher" field.
func PublisherIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldPublisher, vs...))
}

// PublisherNotIn applies the NotIn predicate on the "publisher" field.
func PublisherNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldPublisher, vs...))
}

// PublisherGT applies the GT predicate on the "publisher" field.
func PublisherGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldPublisher, v))
}

// PublisherGTE applies the GTE predicate on the "publisher" field.
func PublisherGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldPublisher, v))
}

// PublisherLT applies the LT predicate on the "publisher" field.
func PublisherLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldPublisher, v))
}

// PublisherLTE applies the LTE predicate on the "publisher" field.
func PublisherLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldPublisher, v))
}

// PublisherContains applies the Contains predicate on the "publisher" field.
func PublisherContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldPublisher, v))
}

// PublisherHasPrefix applies the HasPrefix predicate on the "publisher" field.
func PublisherHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldPublisher, v))
}

// PublisherHasSuffix applies the HasSuffix predicate on the "publisher" field.
func PublisherHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldPublisher, v))
}

// PublisherEqualFold applies the EqualFold predicate on the "publisher" field.
func PublisherEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldPublisher, v))
}

// PublisherContainsFold applies the ContainsFold predicate on the "publisher" field.
func PublisherContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldPublisher, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Book) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Book) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Book) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		p(s.Not())
	})
}