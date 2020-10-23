package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field" 
	"github.com/facebookincubator/ent/schema/edge"
)

// Drug holds the schema definition for the Drug entity.
type Drug struct {
	ent.Schema
}

// Fields of the Drug.
func (Drug) Fields() []ent.Field {
	return []ent.Field {
		field.String("DrugType").NotEmpty(),
		field.String("Strength").Default("unknown"),
		field.String("Information").Default("unknown"),
	}
}

// Edges of the Drug.
func (Drug) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("form", Form.Type).Ref("drugs").Unique(),
		edge.From("unit", Unit.Type).Ref("drugs").Unique(),
		edge.From("volume", Volume.Type).Ref("drugs").Unique(),
		edge.To("users", User.Type).StorageKey(edge.Column("drug_id")),
		edge.To("dispenses", Dispense.Type).StorageKey(edge.Column("drug_id")),
	}
}


