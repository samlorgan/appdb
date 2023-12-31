// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"dbapp/ent/adgroup"
	"dbapp/ent/community"
	"dbapp/ent/predicate"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ADGroupUpdate is the builder for updating ADGroup entities.
type ADGroupUpdate struct {
	config
	hooks    []Hook
	mutation *ADGroupMutation
}

// Where appends a list predicates to the ADGroupUpdate builder.
func (agu *ADGroupUpdate) Where(ps ...predicate.ADGroup) *ADGroupUpdate {
	agu.mutation.Where(ps...)
	return agu
}

// SetName sets the "name" field.
func (agu *ADGroupUpdate) SetName(s string) *ADGroupUpdate {
	agu.mutation.SetName(s)
	return agu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (agu *ADGroupUpdate) SetNillableName(s *string) *ADGroupUpdate {
	if s != nil {
		agu.SetName(*s)
	}
	return agu
}

// AddCommunityIDs adds the "community" edge to the Community entity by IDs.
func (agu *ADGroupUpdate) AddCommunityIDs(ids ...uuid.UUID) *ADGroupUpdate {
	agu.mutation.AddCommunityIDs(ids...)
	return agu
}

// AddCommunity adds the "community" edges to the Community entity.
func (agu *ADGroupUpdate) AddCommunity(c ...*Community) *ADGroupUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return agu.AddCommunityIDs(ids...)
}

// Mutation returns the ADGroupMutation object of the builder.
func (agu *ADGroupUpdate) Mutation() *ADGroupMutation {
	return agu.mutation
}

// ClearCommunity clears all "community" edges to the Community entity.
func (agu *ADGroupUpdate) ClearCommunity() *ADGroupUpdate {
	agu.mutation.ClearCommunity()
	return agu
}

// RemoveCommunityIDs removes the "community" edge to Community entities by IDs.
func (agu *ADGroupUpdate) RemoveCommunityIDs(ids ...uuid.UUID) *ADGroupUpdate {
	agu.mutation.RemoveCommunityIDs(ids...)
	return agu
}

// RemoveCommunity removes "community" edges to Community entities.
func (agu *ADGroupUpdate) RemoveCommunity(c ...*Community) *ADGroupUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return agu.RemoveCommunityIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (agu *ADGroupUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, agu.sqlSave, agu.mutation, agu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (agu *ADGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := agu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (agu *ADGroupUpdate) Exec(ctx context.Context) error {
	_, err := agu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (agu *ADGroupUpdate) ExecX(ctx context.Context) {
	if err := agu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (agu *ADGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(adgroup.Table, adgroup.Columns, sqlgraph.NewFieldSpec(adgroup.FieldID, field.TypeUUID))
	if ps := agu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := agu.mutation.Name(); ok {
		_spec.SetField(adgroup.FieldName, field.TypeString, value)
	}
	if agu.mutation.CommunityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   adgroup.CommunityTable,
			Columns: adgroup.CommunityPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(community.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := agu.mutation.RemovedCommunityIDs(); len(nodes) > 0 && !agu.mutation.CommunityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   adgroup.CommunityTable,
			Columns: adgroup.CommunityPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(community.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := agu.mutation.CommunityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   adgroup.CommunityTable,
			Columns: adgroup.CommunityPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(community.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, agu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	agu.mutation.done = true
	return n, nil
}

// ADGroupUpdateOne is the builder for updating a single ADGroup entity.
type ADGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ADGroupMutation
}

// SetName sets the "name" field.
func (aguo *ADGroupUpdateOne) SetName(s string) *ADGroupUpdateOne {
	aguo.mutation.SetName(s)
	return aguo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (aguo *ADGroupUpdateOne) SetNillableName(s *string) *ADGroupUpdateOne {
	if s != nil {
		aguo.SetName(*s)
	}
	return aguo
}

// AddCommunityIDs adds the "community" edge to the Community entity by IDs.
func (aguo *ADGroupUpdateOne) AddCommunityIDs(ids ...uuid.UUID) *ADGroupUpdateOne {
	aguo.mutation.AddCommunityIDs(ids...)
	return aguo
}

// AddCommunity adds the "community" edges to the Community entity.
func (aguo *ADGroupUpdateOne) AddCommunity(c ...*Community) *ADGroupUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return aguo.AddCommunityIDs(ids...)
}

// Mutation returns the ADGroupMutation object of the builder.
func (aguo *ADGroupUpdateOne) Mutation() *ADGroupMutation {
	return aguo.mutation
}

// ClearCommunity clears all "community" edges to the Community entity.
func (aguo *ADGroupUpdateOne) ClearCommunity() *ADGroupUpdateOne {
	aguo.mutation.ClearCommunity()
	return aguo
}

// RemoveCommunityIDs removes the "community" edge to Community entities by IDs.
func (aguo *ADGroupUpdateOne) RemoveCommunityIDs(ids ...uuid.UUID) *ADGroupUpdateOne {
	aguo.mutation.RemoveCommunityIDs(ids...)
	return aguo
}

// RemoveCommunity removes "community" edges to Community entities.
func (aguo *ADGroupUpdateOne) RemoveCommunity(c ...*Community) *ADGroupUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return aguo.RemoveCommunityIDs(ids...)
}

// Where appends a list predicates to the ADGroupUpdate builder.
func (aguo *ADGroupUpdateOne) Where(ps ...predicate.ADGroup) *ADGroupUpdateOne {
	aguo.mutation.Where(ps...)
	return aguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aguo *ADGroupUpdateOne) Select(field string, fields ...string) *ADGroupUpdateOne {
	aguo.fields = append([]string{field}, fields...)
	return aguo
}

// Save executes the query and returns the updated ADGroup entity.
func (aguo *ADGroupUpdateOne) Save(ctx context.Context) (*ADGroup, error) {
	return withHooks(ctx, aguo.sqlSave, aguo.mutation, aguo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aguo *ADGroupUpdateOne) SaveX(ctx context.Context) *ADGroup {
	node, err := aguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aguo *ADGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := aguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aguo *ADGroupUpdateOne) ExecX(ctx context.Context) {
	if err := aguo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aguo *ADGroupUpdateOne) sqlSave(ctx context.Context) (_node *ADGroup, err error) {
	_spec := sqlgraph.NewUpdateSpec(adgroup.Table, adgroup.Columns, sqlgraph.NewFieldSpec(adgroup.FieldID, field.TypeUUID))
	id, ok := aguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ADGroup.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adgroup.FieldID)
		for _, f := range fields {
			if !adgroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != adgroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aguo.mutation.Name(); ok {
		_spec.SetField(adgroup.FieldName, field.TypeString, value)
	}
	if aguo.mutation.CommunityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   adgroup.CommunityTable,
			Columns: adgroup.CommunityPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(community.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aguo.mutation.RemovedCommunityIDs(); len(nodes) > 0 && !aguo.mutation.CommunityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   adgroup.CommunityTable,
			Columns: adgroup.CommunityPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(community.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aguo.mutation.CommunityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   adgroup.CommunityTable,
			Columns: adgroup.CommunityPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(community.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ADGroup{config: aguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aguo.mutation.done = true
	return _node, nil
}
