package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().StorageKey("id").Immutable(),
		field.String("username").MaxLen(255).Unique().NotEmpty(),
		field.String("password").MaxLen(512).NotEmpty(),
		field.String("first_name").MaxLen(255).NotEmpty(),
		field.String("last_name").MaxLen(255).NotEmpty(),
		field.Time("birth_date"),
		field.Bool("is_active").Default(true),
		field.Time("created_date").Default(time.Now).Immutable(),
		field.Time("updated_date").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
