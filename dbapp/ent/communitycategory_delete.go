// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"dbapp/ent/communitycategory"
	"dbapp/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommunityCategoryDelete is the builder for deleting a CommunityCategory entity.
type CommunityCategoryDelete struct {
	config
	hooks    []Hook
	mutation *CommunityCategoryMutation
}

// Where appends a list predicates to the CommunityCategoryDelete builder.
func (ccd *CommunityCategoryDelete) Where(ps ...predicate.CommunityCategory) *CommunityCategoryDelete {
	ccd.mutation.Where(ps...)
	return ccd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ccd *CommunityCategoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ccd.sqlExec, ccd.mutation, ccd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ccd *CommunityCategoryDelete) ExecX(ctx context.Context) int {
	n, err := ccd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ccd *CommunityCategoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(communitycategory.Table, sqlgraph.NewFieldSpec(communitycategory.FieldID, field.TypeUUID))
	if ps := ccd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ccd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ccd.mutation.done = true
	return affected, err
}

// CommunityCategoryDeleteOne is the builder for deleting a single CommunityCategory entity.
type CommunityCategoryDeleteOne struct {
	ccd *CommunityCategoryDelete
}

// Where appends a list predicates to the CommunityCategoryDelete builder.
func (ccdo *CommunityCategoryDeleteOne) Where(ps ...predicate.CommunityCategory) *CommunityCategoryDeleteOne {
	ccdo.ccd.mutation.Where(ps...)
	return ccdo
}

// Exec executes the deletion query.
func (ccdo *CommunityCategoryDeleteOne) Exec(ctx context.Context) error {
	n, err := ccdo.ccd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{communitycategory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ccdo *CommunityCategoryDeleteOne) ExecX(ctx context.Context) {
	if err := ccdo.Exec(ctx); err != nil {
		panic(err)
	}
}
