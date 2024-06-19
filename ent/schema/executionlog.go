package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// ExecutionLog holds the schema definition for the ExecutionLog entity.
type ExecutionLog struct {
	ent.Schema
}

// 테이블명 지정
func (ExecutionLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "execution_log"},
	}
}

// Fields of the ExecutionLog.
func (ExecutionLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Unique().
			Immutable().
			StorageKey("id"), // Maps to SQL column name "id"
		field.Int("status"),
		field.String("result"),
		field.Time("executedAt").
			Default(time.Now).
			Immutable().
			StorageKey("executedAt"), // Maps to SQL column name "executedAt"
		field.Int("scriptId").StorageKey("scriptId"),
	}
}

// Edges of the ExecutionLog.
func (ExecutionLog) Edges() []ent.Edge {
	return nil
}
