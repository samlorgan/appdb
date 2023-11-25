// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"dbapp/ent/partnerpagelinkfragment"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PartnerPageLinkFragmentCreate is the builder for creating a PartnerPageLinkFragment entity.
type PartnerPageLinkFragmentCreate struct {
	config
	mutation *PartnerPageLinkFragmentMutation
	hooks    []Hook
}

// SetLinkText sets the "link_text" field.
func (pplfc *PartnerPageLinkFragmentCreate) SetLinkText(s string) *PartnerPageLinkFragmentCreate {
	pplfc.mutation.SetLinkText(s)
	return pplfc
}

// SetWhamPartnerURLSuffix sets the "wham_partner_url_suffix" field.
func (pplfc *PartnerPageLinkFragmentCreate) SetWhamPartnerURLSuffix(s string) *PartnerPageLinkFragmentCreate {
	pplfc.mutation.SetWhamPartnerURLSuffix(s)
	return pplfc
}

// SetNillableWhamPartnerURLSuffix sets the "wham_partner_url_suffix" field if the given value is not nil.
func (pplfc *PartnerPageLinkFragmentCreate) SetNillableWhamPartnerURLSuffix(s *string) *PartnerPageLinkFragmentCreate {
	if s != nil {
		pplfc.SetWhamPartnerURLSuffix(*s)
	}
	return pplfc
}

// SetDisplayOrder sets the "display_order" field.
func (pplfc *PartnerPageLinkFragmentCreate) SetDisplayOrder(i int) *PartnerPageLinkFragmentCreate {
	pplfc.mutation.SetDisplayOrder(i)
	return pplfc
}

// SetID sets the "id" field.
func (pplfc *PartnerPageLinkFragmentCreate) SetID(u uuid.UUID) *PartnerPageLinkFragmentCreate {
	pplfc.mutation.SetID(u)
	return pplfc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pplfc *PartnerPageLinkFragmentCreate) SetNillableID(u *uuid.UUID) *PartnerPageLinkFragmentCreate {
	if u != nil {
		pplfc.SetID(*u)
	}
	return pplfc
}

// Mutation returns the PartnerPageLinkFragmentMutation object of the builder.
func (pplfc *PartnerPageLinkFragmentCreate) Mutation() *PartnerPageLinkFragmentMutation {
	return pplfc.mutation
}

// Save creates the PartnerPageLinkFragment in the database.
func (pplfc *PartnerPageLinkFragmentCreate) Save(ctx context.Context) (*PartnerPageLinkFragment, error) {
	pplfc.defaults()
	return withHooks(ctx, pplfc.sqlSave, pplfc.mutation, pplfc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pplfc *PartnerPageLinkFragmentCreate) SaveX(ctx context.Context) *PartnerPageLinkFragment {
	v, err := pplfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pplfc *PartnerPageLinkFragmentCreate) Exec(ctx context.Context) error {
	_, err := pplfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pplfc *PartnerPageLinkFragmentCreate) ExecX(ctx context.Context) {
	if err := pplfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pplfc *PartnerPageLinkFragmentCreate) defaults() {
	if _, ok := pplfc.mutation.ID(); !ok {
		v := partnerpagelinkfragment.DefaultID()
		pplfc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pplfc *PartnerPageLinkFragmentCreate) check() error {
	if _, ok := pplfc.mutation.LinkText(); !ok {
		return &ValidationError{Name: "link_text", err: errors.New(`ent: missing required field "PartnerPageLinkFragment.link_text"`)}
	}
	if _, ok := pplfc.mutation.DisplayOrder(); !ok {
		return &ValidationError{Name: "display_order", err: errors.New(`ent: missing required field "PartnerPageLinkFragment.display_order"`)}
	}
	return nil
}

func (pplfc *PartnerPageLinkFragmentCreate) sqlSave(ctx context.Context) (*PartnerPageLinkFragment, error) {
	if err := pplfc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pplfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pplfc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	pplfc.mutation.id = &_node.ID
	pplfc.mutation.done = true
	return _node, nil
}

func (pplfc *PartnerPageLinkFragmentCreate) createSpec() (*PartnerPageLinkFragment, *sqlgraph.CreateSpec) {
	var (
		_node = &PartnerPageLinkFragment{config: pplfc.config}
		_spec = sqlgraph.NewCreateSpec(partnerpagelinkfragment.Table, sqlgraph.NewFieldSpec(partnerpagelinkfragment.FieldID, field.TypeUUID))
	)
	if id, ok := pplfc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pplfc.mutation.LinkText(); ok {
		_spec.SetField(partnerpagelinkfragment.FieldLinkText, field.TypeString, value)
		_node.LinkText = value
	}
	if value, ok := pplfc.mutation.WhamPartnerURLSuffix(); ok {
		_spec.SetField(partnerpagelinkfragment.FieldWhamPartnerURLSuffix, field.TypeString, value)
		_node.WhamPartnerURLSuffix = value
	}
	if value, ok := pplfc.mutation.DisplayOrder(); ok {
		_spec.SetField(partnerpagelinkfragment.FieldDisplayOrder, field.TypeInt, value)
		_node.DisplayOrder = value
	}
	return _node, _spec
}

// PartnerPageLinkFragmentCreateBulk is the builder for creating many PartnerPageLinkFragment entities in bulk.
type PartnerPageLinkFragmentCreateBulk struct {
	config
	err      error
	builders []*PartnerPageLinkFragmentCreate
}

// Save creates the PartnerPageLinkFragment entities in the database.
func (pplfcb *PartnerPageLinkFragmentCreateBulk) Save(ctx context.Context) ([]*PartnerPageLinkFragment, error) {
	if pplfcb.err != nil {
		return nil, pplfcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pplfcb.builders))
	nodes := make([]*PartnerPageLinkFragment, len(pplfcb.builders))
	mutators := make([]Mutator, len(pplfcb.builders))
	for i := range pplfcb.builders {
		func(i int, root context.Context) {
			builder := pplfcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PartnerPageLinkFragmentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pplfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pplfcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pplfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pplfcb *PartnerPageLinkFragmentCreateBulk) SaveX(ctx context.Context) []*PartnerPageLinkFragment {
	v, err := pplfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pplfcb *PartnerPageLinkFragmentCreateBulk) Exec(ctx context.Context) error {
	_, err := pplfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pplfcb *PartnerPageLinkFragmentCreateBulk) ExecX(ctx context.Context) {
	if err := pplfcb.Exec(ctx); err != nil {
		panic(err)
	}
}
