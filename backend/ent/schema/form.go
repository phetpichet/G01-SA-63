package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field" 
	"github.com/facebookincubator/ent/schema/edge"
)

// Form holds the schema definition for the Form entity.
type Form struct {
	ent.Schema
}

// Fields of the Form.
func (Form) Fields() []ent.Field {
	return []ent.Field {
		field.String("FormType").NotEmpty(), 
		
	}
}

// Edges of the Form.
func (Form) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("drugs",Drug.Type).StorageKey(edge.Column("form_id")),
	}
}


