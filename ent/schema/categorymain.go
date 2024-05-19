package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// CategoryMain holds the schema definition for the CategoryMain entity.
type CategoryMain struct {
	ent.Schema
}

// 테이블명 지정
func (CategoryMain) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "category_main"},
	}
}

// Fields of the CategoryMain.
func (CategoryMain) Fields() []ent.Field {

	return []ent.Field{
		field.Int32("id").
			Unique().
			Immutable(),
		field.String("program"),
		field.Time("updatedAt").
			Default(time.Now).
			UpdateDefault(time.Now).StorageKey("updatedAt"),
		field.Time("createdAt").
			Default(time.Now).StorageKey("createdAt"),
		field.String("name"),
		field.Int("priority").
			Default(0),
		field.Int("updatedById").
			Default(0).StorageKey("updatedById"),
		field.Int("createdById").
			Default(0).StorageKey("createdById"),
	}
}

// Edges of the CategoryMain.
func (CategoryMain) Edges() []ent.Edge {
	return nil
}
