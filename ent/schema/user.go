package schema

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// 테이블명 지정
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").
			Unique().
			Immutable(),
		field.Time("createdAt").
			Default(time.Now).
			Immutable().
			StorageKey("createdAt"), // 이렇게 안하면 created_at 으로 생성되어버림
		field.Time("updatedAt").
			Default(time.Now).
			UpdateDefault(time.Now).
			StorageKey("updatedAt"),
		field.String("username").
			NotEmpty().
			Unique(),
		field.String("password").
			NotEmpty(),
		field.String("realname").
			NotEmpty(),
		field.String("phone").
			NotEmpty(),
		field.String("email").
			NotEmpty(),
		field.Enum("role").
			Values("Reader", "Writer", "Admin").
			Default("Reader"),
		field.Time("lastAccessed").StorageKey("lastAccessed"),
		field.Int("userGroupId").StorageKey("userGroupId"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
