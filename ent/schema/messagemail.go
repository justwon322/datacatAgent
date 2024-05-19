package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// MessageMail holds the schema definition for the MessageMail entity.
type MessageMail struct {
	ent.Schema
}

// Fields of the MessageMail.
func (MessageMail) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("seq").
			Unique().
			Immutable().
			StorageKey("seq"), // Maps to SQL column name "seq"
		field.String("mailSubject").
			StorageKey("mail_subject"), // Maps to SQL column name "mail_subject"
		field.String("mailContents").
			StorageKey("mail_contents"), // Maps to SQL column name "mail_contents"
		field.Int("mailGb").
			StorageKey("mail_gb"), // Maps to SQL column name "mail_gb"
		field.String("mailRecvGroup").
			StorageKey("mail_recv_group"), // Maps to SQL column name "mail_recv_group"
		field.Int("sent"),
		field.Time("sentAt").
			Default(time.Now).
			UpdateDefault(time.Now).
			StorageKey("sentAt"), // Maps to SQL column name "sentAt"
		field.Int("createdById"),
		field.Time("createdAt").
			Default(time.Now).
			Immutable().
			StorageKey("createdAt"), // Maps to SQL column name "createdAt"
		field.Int("updatedById"),
		field.Time("updatedAt").
			Default(time.Now).
			UpdateDefault(time.Now).
			StorageKey("updatedAt"), // Maps to SQL column name "updatedAt"
	}
}

// Edges of the MessageMail.
func (MessageMail) Edges() []ent.Edge {
	return nil
}
