// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/Piichet-3-/app/ent/drug"
	"github.com/Piichet-3-/app/ent/form"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// FormCreate is the builder for creating a Form entity.
type FormCreate struct {
	config
	mutation *FormMutation
	hooks    []Hook
}

// SetFormType sets the FormType field.
func (fc *FormCreate) SetFormType(s string) *FormCreate {
	fc.mutation.SetFormType(s)
	return fc
}

// AddDrugIDs adds the drugs edge to Drug by ids.
func (fc *FormCreate) AddDrugIDs(ids ...int) *FormCreate {
	fc.mutation.AddDrugIDs(ids...)
	return fc
}

// AddDrugs adds the drugs edges to Drug.
func (fc *FormCreate) AddDrugs(d ...*Drug) *FormCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return fc.AddDrugIDs(ids...)
}

// Mutation returns the FormMutation object of the builder.
func (fc *FormCreate) Mutation() *FormMutation {
	return fc.mutation
}

// Save creates the Form in the database.
func (fc *FormCreate) Save(ctx context.Context) (*Form, error) {
	if _, ok := fc.mutation.FormType(); !ok {
		return nil, &ValidationError{Name: "FormType", err: errors.New("ent: missing required field \"FormType\"")}
	}
	if v, ok := fc.mutation.FormType(); ok {
		if err := form.FormTypeValidator(v); err != nil {
			return nil, &ValidationError{Name: "FormType", err: fmt.Errorf("ent: validator failed for field \"FormType\": %w", err)}
		}
	}
	var (
		err  error
		node *Form
	)
	if len(fc.hooks) == 0 {
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FormMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fc.mutation = mutation
			node, err = fc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FormCreate) SaveX(ctx context.Context) *Form {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fc *FormCreate) sqlSave(ctx context.Context) (*Form, error) {
	f, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	f.ID = int(id)
	return f, nil
}

func (fc *FormCreate) createSpec() (*Form, *sqlgraph.CreateSpec) {
	var (
		f     = &Form{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: form.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: form.FieldID,
			},
		}
	)
	if value, ok := fc.mutation.FormType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: form.FieldFormType,
		})
		f.FormType = value
	}
	if nodes := fc.mutation.DrugsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   form.DrugsTable,
			Columns: []string{form.DrugsColumn},
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
	return f, _spec
}
