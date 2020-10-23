// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/Piichet-3-/app/ent/dispense"
	"github.com/Piichet-3-/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// DispenseDelete is the builder for deleting a Dispense entity.
type DispenseDelete struct {
	config
	hooks      []Hook
	mutation   *DispenseMutation
	predicates []predicate.Dispense
}

// Where adds a new predicate to the delete builder.
func (dd *DispenseDelete) Where(ps ...predicate.Dispense) *DispenseDelete {
	dd.predicates = append(dd.predicates, ps...)
	return dd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dd *DispenseDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dd.hooks) == 0 {
		affected, err = dd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DispenseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dd.mutation = mutation
			affected, err = dd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dd.hooks) - 1; i >= 0; i-- {
			mut = dd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dd *DispenseDelete) ExecX(ctx context.Context) int {
	n, err := dd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dd *DispenseDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: dispense.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dispense.FieldID,
			},
		},
	}
	if ps := dd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, dd.driver, _spec)
}

// DispenseDeleteOne is the builder for deleting a single Dispense entity.
type DispenseDeleteOne struct {
	dd *DispenseDelete
}

// Exec executes the deletion query.
func (ddo *DispenseDeleteOne) Exec(ctx context.Context) error {
	n, err := ddo.dd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{dispense.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ddo *DispenseDeleteOne) ExecX(ctx context.Context) {
	ddo.dd.ExecX(ctx)
}
