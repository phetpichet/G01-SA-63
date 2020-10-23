package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field" 
	"github.com/facebookincubator/ent/schema/edge"
)

// Volume holds the schema definition for the User entity.
type Volume struct {
	ent.Schema
}

// Fields of the Volume.
func (Volume) Fields() []ent.Field {
	return []ent.Field {
		field.String("VolumeType").NotEmpty(), 
		
	}
}

// Edges of the Volume.
func (Volume) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("drugs",Drug.Type).StorageKey(edge.Column("volume_id")),
	}
}



