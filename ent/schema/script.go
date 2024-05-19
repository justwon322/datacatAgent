package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Script holds the schema definition for the Script entity.
type Script struct {
	ent.Schema
}

// 테이블명 지정
func (Script) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "script"},
	}
}

// Fields of the Script.
func (Script) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Unique().
			Immutable().
			StorageKey("id"), // Maps to SQL column name "id"
		field.String("program"),
		field.Time("updatedAt").
			Default(time.Now).
			UpdateDefault(time.Now).
			StorageKey("updatedAt"), // Automatically update with current time on update
		field.Time("createdAt").
			Default(time.Now).
			Immutable().
			StorageKey("createdAt"), // Automatically set on create
		field.Int("jobId").
			StorageKey("jobId"),
		field.String("hostname"),
		field.String("command"),
		field.String("comment"),
		field.Int("usable"),
		field.Int("repeatInterval").
			StorageKey("repeatInterval"),
		field.Int("updatedById").
			StorageKey("updatedById"),
		field.Int("createdById").
			StorageKey("createdById"),
		field.Int("regionId").
			StorageKey("regionId"),
		field.Int("managerId").
			StorageKey("managerId"),
		field.Int("managerGroupId").
			StorageKey("managerGroupId"),
		field.Time("startTime").
			StorageKey("startTime"),
		field.Time("endTime").
			StorageKey("endTime"),
	}
}

// Edges of the Script.
func (Script) Edges() []ent.Edge {
	return nil
}
