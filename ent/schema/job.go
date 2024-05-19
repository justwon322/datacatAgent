package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Job holds the schema definition for the Job entity.
type Job struct {
	ent.Schema
}

// 테이블명 지정
func (Job) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "job"},
	}
}

// Fields of the Job.
func (Job) Fields() []ent.Field {
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
		field.Int("categorySubId").StorageKey("categorySubId"),
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return nil
}
