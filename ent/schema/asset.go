package schema

import "entgo.io/ent"

// Asset holds the schema definition for the Asset entity.
type Asset struct {
	ent.Schema
}

// Fields of the Asset.
func (Asset) Fields() []ent.Field {
	return nil
}

// Edges of the Asset.
func (Asset) Edges() []ent.Edge {
	return nil
}
