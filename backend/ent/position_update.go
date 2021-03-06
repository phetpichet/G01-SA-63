// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/Piichet-3-/app/ent/position"
	"github.com/Piichet-3-/app/ent/predicate"
	"github.com/Piichet-3-/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// PositionUpdate is the builder for updating Position entities.
type PositionUpdate struct {
	config
	hooks      []Hook
	mutation   *PositionMutation
	predicates []predicate.Position
}

// Where adds a new predicate for the builder.
func (pu *PositionUpdate) Where(ps ...predicate.Position) *PositionUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetPosition sets the position field.
func (pu *PositionUpdate) SetPosition(s string) *PositionUpdate {
	pu.mutation.SetPosition(s)
	return pu
}

// AddUserIDs adds the users edge to User by ids.
func (pu *PositionUpdate) AddUserIDs(ids ...int) *PositionUpdate {
	pu.mutation.AddUserIDs(ids...)
	return pu
}

// AddUsers adds the users edges to User.
func (pu *PositionUpdate) AddUsers(u ...*User) *PositionUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.AddUserIDs(ids...)
}

// Mutation returns the PositionMutation object of the builder.
func (pu *PositionUpdate) Mutation() *PositionMutation {
	return pu.mutation
}

// RemoveUserIDs removes the users edge to User by ids.
func (pu *PositionUpdate) RemoveUserIDs(ids ...int) *PositionUpdate {
	pu.mutation.RemoveUserIDs(ids...)
	return pu
}

// RemoveUsers removes users edges to User.
func (pu *PositionUpdate) RemoveUsers(u ...*User) *PositionUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PositionUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.Position(); ok {
		if err := position.PositionValidator(v); err != nil {
			return 0, &ValidationError{Name: "position", err: fmt.Errorf("ent: validator failed for field \"position\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PositionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PositionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PositionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PositionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PositionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   position.Table,
			Columns: position.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: position.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Position(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: position.FieldPosition,
		})
	}
	if nodes := pu.mutation.RemovedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.UsersTable,
			Columns: []string{position.UsersColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.UsersTable,
			Columns: []string{position.UsersColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{position.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PositionUpdateOne is the builder for updating a single Position entity.
type PositionUpdateOne struct {
	config
	hooks    []Hook
	mutation *PositionMutation
}

// SetPosition sets the position field.
func (puo *PositionUpdateOne) SetPosition(s string) *PositionUpdateOne {
	puo.mutation.SetPosition(s)
	return puo
}

// AddUserIDs adds the users edge to User by ids.
func (puo *PositionUpdateOne) AddUserIDs(ids ...int) *PositionUpdateOne {
	puo.mutation.AddUserIDs(ids...)
	return puo
}

// AddUsers adds the users edges to User.
func (puo *PositionUpdateOne) AddUsers(u ...*User) *PositionUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.AddUserIDs(ids...)
}

// Mutation returns the PositionMutation object of the builder.
func (puo *PositionUpdateOne) Mutation() *PositionMutation {
	return puo.mutation
}

// RemoveUserIDs removes the users edge to User by ids.
func (puo *PositionUpdateOne) RemoveUserIDs(ids ...int) *PositionUpdateOne {
	puo.mutation.RemoveUserIDs(ids...)
	return puo
}

// RemoveUsers removes users edges to User.
func (puo *PositionUpdateOne) RemoveUsers(u ...*User) *PositionUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.RemoveUserIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *PositionUpdateOne) Save(ctx context.Context) (*Position, error) {
	if v, ok := puo.mutation.Position(); ok {
		if err := position.PositionValidator(v); err != nil {
			return nil, &ValidationError{Name: "position", err: fmt.Errorf("ent: validator failed for field \"position\": %w", err)}
		}
	}

	var (
		err  error
		node *Position
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PositionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PositionUpdateOne) SaveX(ctx context.Context) *Position {
	po, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return po
}

// Exec executes the query on the entity.
func (puo *PositionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PositionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PositionUpdateOne) sqlSave(ctx context.Context) (po *Position, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   position.Table,
			Columns: position.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: position.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Position.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Position(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: position.FieldPosition,
		})
	}
	if nodes := puo.mutation.RemovedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.UsersTable,
			Columns: []string{position.UsersColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.UsersTable,
			Columns: []string{position.UsersColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	po = &Position{config: puo.config}
	_spec.Assign = po.assignValues
	_spec.ScanValues = po.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{position.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return po, nil
}
