package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Hstest holds the schema definition for the Hstest entity.
type Hstest struct {
	ent.Schema
}

// Fields of the Hstest.
func (Hstest) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("age"),
		field.String("rowname"),
		field.String("value"),
		field.Time("updatetime").
			Default(time.Now),
		field.String("body"),
		field.Float("floa"),
		field.Int("level")}
}

// Edges of the Hstest.
func (Hstest) Edges() []ent.Edge {
	return nil
}
