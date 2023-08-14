package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().StorageKey("id").Immutable(),
		field.Int("user_id"),
		field.Int("asset_id"),
		field.String("account_no").MaxLen(20).Unique().NotEmpty(),
		field.Bool("is_active").Default(true),
		field.Time("created_date").Default(time.Now).Immutable(),
		field.Time("updated_date").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}
