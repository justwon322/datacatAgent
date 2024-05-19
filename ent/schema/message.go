package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Unique().
			Immutable().
			StorageKey("id"), // Maps to SQL column name "id"
		field.String("payload"),
		field.Int("sent"),
		field.Int("resent"),
		field.Time("sentAt").
			Default(time.Now).
			UpdateDefault(time.Now).
			StorageKey("sentAt"), // Maps to SQL column name "sentAt"
		field.Int("jobId"),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return nil
}
