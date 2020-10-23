// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/Piichet-3-/app/ent/dispense"
	"github.com/Piichet-3-/app/ent/drug"
	"github.com/Piichet-3-/app/ent/form"
	"github.com/Piichet-3-/app/ent/unit"
	"github.com/Piichet-3-/app/ent/user"
	"github.com/Piichet-3-/app/ent/volume"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// DrugCreate is the builder for creating a Drug entity.
type DrugCreate struct {
	config
	mutation *DrugMutation
	hooks    []Hook
}

// SetDrugType sets the DrugType field.
func (dc *DrugCreate) SetDrugType(s string) *DrugCreate {
	dc.mutation.SetDrugType(s)
	return dc
}

// SetStrength sets the Strength field.
func (dc *DrugCreate) SetStrength(s string) *DrugCreate {
	dc.mutation.SetStrength(s)
	return dc
}

// SetNillableStrength sets the Strength field if the given value is not nil.
func (dc *DrugCreate) SetNillableStrength(s *string) *DrugCreate {
	if s != nil {
		dc.SetStrength(*s)
	}
	return dc
}

// SetInformation sets the Information field.
func (dc *DrugCreate) SetInformation(s string) *DrugCreate {
	dc.mutation.SetInformation(s)
	return dc
}

// SetNillableInformation sets the Information field if the given value is not nil.
func (dc *DrugCreate) SetNillableInformation(s *string) *DrugCreate {
	if s != nil {
		dc.SetInformation(*s)
	}
	return dc
}

// SetFormID sets the form edge to Form by id.
func (dc *DrugCreate) SetFormID(id int) *DrugCreate {
	dc.mutation.SetFormID(id)
	return dc
}

// SetNillableFormID sets the form edge to Form by id if the given value is not nil.
func (dc *DrugCreate) SetNillableFormID(id *int) *DrugCreate {
	if id != nil {
		dc = dc.SetFormID(*id)
	}
	return dc
}

// SetForm sets the form edge to Form.
func (dc *DrugCreate) SetForm(f *Form) *DrugCreate {
	return dc.SetFormID(f.ID)
}

// SetUnitID sets the unit edge to Unit by id.
func (dc *DrugCreate) SetUnitID(id int) *DrugCreate {
	dc.mutation.SetUnitID(id)
	return dc
}

// SetNillableUnitID sets the unit edge to Unit by id if the given value is not nil.
func (dc *DrugCreate) SetNillableUnitID(id *int) *DrugCreate {
	if id != nil {
		dc = dc.SetUnitID(*id)
	}
	return dc
}

// SetUnit sets the unit edge to Unit.
func (dc *DrugCreate) SetUnit(u *Unit) *DrugCreate {
	return dc.SetUnitID(u.ID)
}

// SetVolumeID sets the volume edge to Volume by id.
func (dc *DrugCreate) SetVolumeID(id int) *DrugCreate {
	dc.mutation.SetVolumeID(id)
	return dc
}

// SetNillableVolumeID sets the volume edge to Volume by id if the given value is not nil.
func (dc *DrugCreate) SetNillableVolumeID(id *int) *DrugCreate {
	if id != nil {
		dc = dc.SetVolumeID(*id)
	}
	return dc
}

// SetVolume sets the volume edge to Volume.
func (dc *DrugCreate) SetVolume(v *Volume) *DrugCreate {
	return dc.SetVolumeID(v.ID)
}

// AddUserIDs adds the users edge to User by ids.
func (dc *DrugCreate) AddUserIDs(ids ...int) *DrugCreate {
	dc.mutation.AddUserIDs(ids...)
	return dc
}

// AddUsers adds the users edges to User.
func (dc *DrugCreate) AddUsers(u ...*User) *DrugCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return dc.AddUserIDs(ids...)
}

// AddDispenseIDs adds the dispenses edge to Dispense by ids.
func (dc *DrugCreate) AddDispenseIDs(ids ...int) *DrugCreate {
	dc.mutation.AddDispenseIDs(ids...)
	return dc
}

// AddDispenses adds the dispenses edges to Dispense.
func (dc *DrugCreate) AddDispenses(d ...*Dispense) *DrugCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddDispenseIDs(ids...)
}

// Mutation returns the DrugMutation object of the builder.
func (dc *DrugCreate) Mutation() *DrugMutation {
	return dc.mutation
}

// Save creates the Drug in the database.
func (dc *DrugCreate) Save(ctx context.Context) (*Drug, error) {
	if _, ok := dc.mutation.DrugType(); !ok {
		return nil, &ValidationError{Name: "DrugType", err: errors.New("ent: missing required field \"DrugType\"")}
	}
	if v, ok := dc.mutation.DrugType(); ok {
		if err := drug.DrugTypeValidator(v); err != nil {
			return nil, &ValidationError{Name: "DrugType", err: fmt.Errorf("ent: validator failed for field \"DrugType\": %w", err)}
		}
	}
	if _, ok := dc.mutation.Strength(); !ok {
		v := drug.DefaultStrength
		dc.mutation.SetStrength(v)
	}
	if _, ok := dc.mutation.Information(); !ok {
		v := drug.DefaultInformation
		dc.mutation.SetInformation(v)
	}
	var (
		err  error
		node *Drug
	)
	if len(dc.hooks) == 0 {
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DrugMutation)
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
func (dc *DrugCreate) SaveX(ctx context.Context) *Drug {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dc *DrugCreate) sqlSave(ctx context.Context) (*Drug, error) {
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

func (dc *DrugCreate) createSpec() (*Drug, *sqlgraph.CreateSpec) {
	var (
		d     = &Drug{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: drug.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: drug.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.DrugType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: drug.FieldDrugType,
		})
		d.DrugType = value
	}
	if value, ok := dc.mutation.Strength(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: drug.FieldStrength,
		})
		d.Strength = value
	}
	if value, ok := dc.mutation.Information(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: drug.FieldInformation,
		})
		d.Information = value
	}
	if nodes := dc.mutation.FormIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   drug.FormTable,
			Columns: []string{drug.FormColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: form.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.UnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   drug.UnitTable,
			Columns: []string{drug.UnitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: unit.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.VolumeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   drug.VolumeTable,
			Columns: []string{drug.VolumeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: volume.FieldID,
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
			Table:   drug.UsersTable,
			Columns: []string{drug.UsersColumn},
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
	if nodes := dc.mutation.DispensesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   drug.DispensesTable,
			Columns: []string{drug.DispensesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dispense.FieldID,
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
