// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/Piichet-3-/app/ent/dispense"
	"github.com/Piichet-3-/app/ent/drug"
	"github.com/Piichet-3-/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// DispenseCreate is the builder for creating a Dispense entity.
type DispenseCreate struct {
	config
	mutation *DispenseMutation
	hooks    []Hook
}

// SetNote sets the note field.
func (dc *DispenseCreate) SetNote(s string) *DispenseCreate {
	dc.mutation.SetNote(s)
	return dc
}

// SetDrugID sets the drug edge to Drug by id.
func (dc *DispenseCreate) SetDrugID(id int) *DispenseCreate {
	dc.mutation.SetDrugID(id)
	return dc
}

// SetNillableDrugID sets the drug edge to Drug by id if the given value is not nil.
func (dc *DispenseCreate) SetNillableDrugID(id *int) *DispenseCreate {
	if id != nil {
		dc = dc.SetDrugID(*id)
	}
	return dc
}

// SetDrug sets the drug edge to Drug.
func (dc *DispenseCreate) SetDrug(d *Drug) *DispenseCreate {
	return dc.SetDrugID(d.ID)
}

// AddUserIDs adds the users edge to User by ids.
func (dc *DispenseCreate) AddUserIDs(ids ...int) *DispenseCreate {
	dc.mutation.AddUserIDs(ids...)
	return dc
}

// AddUsers adds the users edges to User.
func (dc *DispenseCreate) AddUsers(u ...*User) *DispenseCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return dc.AddUserIDs(ids...)
}

// AddDrugIDs adds the drugs edge to Drug by ids.
func (dc *DispenseCreate) AddDrugIDs(ids ...int) *DispenseCreate {
	dc.mutation.AddDrugIDs(ids...)
	return dc
}

// AddDrugs adds the drugs edges to Drug.
func (dc *DispenseCreate) AddDrugs(d ...*Drug) *DispenseCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddDrugIDs(ids...)
}

// Mutation returns the DispenseMutation object of the builder.
func (dc *DispenseCreate) Mutation() *DispenseMutation {
	return dc.mutation
}

// Save creates the Dispense in the database.
func (dc *DispenseCreate) Save(ctx context.Context) (*Dispense, error) {
	if _, ok := dc.mutation.Note(); !ok {
		return nil, &ValidationError{Name: "note", err: errors.New("ent: missing required field \"note\"")}
	}
	if v, ok := dc.mutation.Note(); ok {
		if err := dispense.NoteValidator(v); err != nil {
			return nil, &ValidationError{Name: "note", err: fmt.Errorf("ent: validator failed for field \"note\": %w", err)}
		}
	}
	var (
		err  error
		node *Dispense
	)
	if len(dc.hooks) == 0 {
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DispenseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dc.mutation = mutation
			node, err = dc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DispenseCreate) SaveX(ctx context.Context) *Dispense {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dc *DispenseCreate) sqlSave(ctx context.Context) (*Dispense, error) {
	d, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	d.ID = int(id)
	return d, nil
}

func (dc *DispenseCreate) createSpec() (*Dispense, *sqlgraph.CreateSpec) {
	var (
		d     = &Dispense{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: dispense.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dispense.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.Note(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dispense.FieldNote,
		})
		d.Note = value
	}
	if nodes := dc.mutation.DrugIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dispense.DrugTable,
			Columns: []string{dispense.DrugColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: drug.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dispense.UsersTable,
			Columns: []string{dispense.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.DrugsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dispense.DrugsTable,
			Columns: []string{dispense.DrugsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: drug.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return d, _spec
}
