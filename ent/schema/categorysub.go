package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// CategorySub holds the schema definition for the CategorySub entity.
type CategorySub struct {
	ent.Schema
}

// 테이블명 지정
func (CategorySub) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "category_sub"},
	}
}

// Fields of the CategorySub.
func (CategorySub) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").
			Unique().
			Immutable().
			StorageKey("id"), // Maps to SQL column name "id"
		field.String("program"),
		field.Time("updatedAt").
			Default(time.Now).
			UpdateDefault(time.Now).
			StorageKey("updatedAt"), // Maps to SQL column name "updatedAt"
		field.Time("createdAt").
			Default(time.Now).
			Immutable().
			StorageKey("createdAt"), // Maps to SQL column name "createdAt"
		field.String("name"),
		field.Int("updatedById").StorageKey("updatedById"),
		field.Int("createdById").StorageKey("createdById"),
		field.Int32("categoryMainId").StorageKey("categoryMainId"),
	}
}

// Edges of the CategorySub.
func (CategorySub) Edges() []ent.Edge {
	return nil
}
