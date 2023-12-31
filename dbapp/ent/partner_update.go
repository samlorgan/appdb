// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"dbapp/ent/application"
	"dbapp/ent/partner"
	"dbapp/ent/predicate"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PartnerUpdate is the builder for updating Partner entities.
type PartnerUpdate struct {
	config
	hooks    []Hook
	mutation *PartnerMutation
}

// Where appends a list predicates to the PartnerUpdate builder.
func (pu *PartnerUpdate) Where(ps ...predicate.Partner) *PartnerUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetWhamSiteID sets the "wham_site_id" field.
func (pu *PartnerUpdate) SetWhamSiteID(i int) *PartnerUpdate {
	pu.mutation.ResetWhamSiteID()
	pu.mutation.SetWhamSiteID(i)
	return pu
}

// SetNillableWhamSiteID sets the "wham_site_id" field if the given value is not nil.
func (pu *PartnerUpdate) SetNillableWhamSiteID(i *int) *PartnerUpdate {
	if i != nil {
		pu.SetWhamSiteID(*i)
	}
	return pu
}

// AddWhamSiteID adds i to the "wham_site_id" field.
func (pu *PartnerUpdate) AddWhamSiteID(i int) *PartnerUpdate {
	pu.mutation.AddWhamSiteID(i)
	return pu
}

// SetWhamTitle sets the "wham_title" field.
func (pu *PartnerUpdate) SetWhamTitle(s string) *PartnerUpdate {
	pu.mutation.SetWhamTitle(s)
	return pu
}

// SetNillableWhamTitle sets the "wham_title" field if the given value is not nil.
func (pu *PartnerUpdate) SetNillableWhamTitle(s *string) *PartnerUpdate {
	if s != nil {
		pu.SetWhamTitle(*s)
	}
	return pu
}

// SetWhamDescription sets the "wham_description" field.
func (pu *PartnerUpdate) SetWhamDescription(s string) *PartnerUpdate {
	pu.mutation.SetWhamDescription(s)
	return pu
}

// SetNillableWhamDescription sets the "wham_description" field if the given value is not nil.
func (pu *PartnerUpdate) SetNillableWhamDescription(s *string) *PartnerUpdate {
	if s != nil {
		pu.SetWhamDescription(*s)
	}
	return pu
}

// SetKeycloakOrganisation sets the "keycloak_organisation" field.
func (pu *PartnerUpdate) SetKeycloakOrganisation(s string) *PartnerUpdate {
	pu.mutation.SetKeycloakOrganisation(s)
	return pu
}

// SetNillableKeycloakOrganisation sets the "keycloak_organisation" field if the given value is not nil.
func (pu *PartnerUpdate) SetNillableKeycloakOrganisation(s *string) *PartnerUpdate {
	if s != nil {
		pu.SetKeycloakOrganisation(*s)
	}
	return pu
}

// SetWhamPartnerURL sets the "wham_partner_url" field.
func (pu *PartnerUpdate) SetWhamPartnerURL(s string) *PartnerUpdate {
	pu.mutation.SetWhamPartnerURL(s)
	return pu
}

// SetNillableWhamPartnerURL sets the "wham_partner_url" field if the given value is not nil.
func (pu *PartnerUpdate) SetNillableWhamPartnerURL(s *string) *PartnerUpdate {
	if s != nil {
		pu.SetWhamPartnerURL(*s)
	}
	return pu
}

// SetWhamCreated sets the "wham_created" field.
func (pu *PartnerUpdate) SetWhamCreated(t time.Time) *PartnerUpdate {
	pu.mutation.SetWhamCreated(t)
	return pu
}

// SetNillableWhamCreated sets the "wham_created" field if the given value is not nil.
func (pu *PartnerUpdate) SetNillableWhamCreated(t *time.Time) *PartnerUpdate {
	if t != nil {
		pu.SetWhamCreated(*t)
	}
	return pu
}

// SetWhamUpdated sets the "wham_updated" field.
func (pu *PartnerUpdate) SetWhamUpdated(t time.Time) *PartnerUpdate {
	pu.mutation.SetWhamUpdated(t)
	return pu
}

// SetNillableWhamUpdated sets the "wham_updated" field if the given value is not nil.
func (pu *PartnerUpdate) SetNillableWhamUpdated(t *time.Time) *PartnerUpdate {
	if t != nil {
		pu.SetWhamUpdated(*t)
	}
	return pu
}

// AddApplicationIDs adds the "application" edge to the Application entity by IDs.
func (pu *PartnerUpdate) AddApplicationIDs(ids ...uuid.UUID) *PartnerUpdate {
	pu.mutation.AddApplicationIDs(ids...)
	return pu
}

// AddApplication adds the "application" edges to the Application entity.
func (pu *PartnerUpdate) AddApplication(a ...*Application) *PartnerUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.AddApplicationIDs(ids...)
}

// Mutation returns the PartnerMutation object of the builder.
func (pu *PartnerUpdate) Mutation() *PartnerMutation {
	return pu.mutation
}

// ClearApplication clears all "application" edges to the Application entity.
func (pu *PartnerUpdate) ClearApplication() *PartnerUpdate {
	pu.mutation.ClearApplication()
	return pu
}

// RemoveApplicationIDs removes the "application" edge to Application entities by IDs.
func (pu *PartnerUpdate) RemoveApplicationIDs(ids ...uuid.UUID) *PartnerUpdate {
	pu.mutation.RemoveApplicationIDs(ids...)
	return pu
}

// RemoveApplication removes "application" edges to Application entities.
func (pu *PartnerUpdate) RemoveApplication(a ...*Application) *PartnerUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.RemoveApplicationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PartnerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PartnerUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PartnerUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PartnerUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PartnerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(partner.Table, partner.Columns, sqlgraph.NewFieldSpec(partner.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.WhamSiteID(); ok {
		_spec.SetField(partner.FieldWhamSiteID, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedWhamSiteID(); ok {
		_spec.AddField(partner.FieldWhamSiteID, field.TypeInt, value)
	}
	if value, ok := pu.mutation.WhamTitle(); ok {
		_spec.SetField(partner.FieldWhamTitle, field.TypeString, value)
	}
	if value, ok := pu.mutation.WhamDescription(); ok {
		_spec.SetField(partner.FieldWhamDescription, field.TypeString, value)
	}
	if value, ok := pu.mutation.KeycloakOrganisation(); ok {
		_spec.SetField(partner.FieldKeycloakOrganisation, field.TypeString, value)
	}
	if value, ok := pu.mutation.WhamPartnerURL(); ok {
		_spec.SetField(partner.FieldWhamPartnerURL, field.TypeString, value)
	}
	if value, ok := pu.mutation.WhamCreated(); ok {
		_spec.SetField(partner.FieldWhamCreated, field.TypeTime, value)
	}
	if value, ok := pu.mutation.WhamUpdated(); ok {
		_spec.SetField(partner.FieldWhamUpdated, field.TypeTime, value)
	}
	if pu.mutation.ApplicationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   partner.ApplicationTable,
			Columns: partner.ApplicationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(application.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedApplicationIDs(); len(nodes) > 0 && !pu.mutation.ApplicationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   partner.ApplicationTable,
			Columns: partner.ApplicationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(application.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ApplicationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   partner.ApplicationTable,
			Columns: partner.ApplicationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(application.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{partner.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PartnerUpdateOne is the builder for updating a single Partner entity.
type PartnerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PartnerMutation
}

// SetWhamSiteID sets the "wham_site_id" field.
func (puo *PartnerUpdateOne) SetWhamSiteID(i int) *PartnerUpdateOne {
	puo.mutation.ResetWhamSiteID()
	puo.mutation.SetWhamSiteID(i)
	return puo
}

// SetNillableWhamSiteID sets the "wham_site_id" field if the given value is not nil.
func (puo *PartnerUpdateOne) SetNillableWhamSiteID(i *int) *PartnerUpdateOne {
	if i != nil {
		puo.SetWhamSiteID(*i)
	}
	return puo
}

// AddWhamSiteID adds i to the "wham_site_id" field.
func (puo *PartnerUpdateOne) AddWhamSiteID(i int) *PartnerUpdateOne {
	puo.mutation.AddWhamSiteID(i)
	return puo
}

// SetWhamTitle sets the "wham_title" field.
func (puo *PartnerUpdateOne) SetWhamTitle(s string) *PartnerUpdateOne {
	puo.mutation.SetWhamTitle(s)
	return puo
}

// SetNillableWhamTitle sets the "wham_title" field if the given value is not nil.
func (puo *PartnerUpdateOne) SetNillableWhamTitle(s *string) *PartnerUpdateOne {
	if s != nil {
		puo.SetWhamTitle(*s)
	}
	return puo
}

// SetWhamDescription sets the "wham_description" field.
func (puo *PartnerUpdateOne) SetWhamDescription(s string) *PartnerUpdateOne {
	puo.mutation.SetWhamDescription(s)
	return puo
}

// SetNillableWhamDescription sets the "wham_description" field if the given value is not nil.
func (puo *PartnerUpdateOne) SetNillableWhamDescription(s *string) *PartnerUpdateOne {
	if s != nil {
		puo.SetWhamDescription(*s)
	}
	return puo
}

// SetKeycloakOrganisation sets the "keycloak_organisation" field.
func (puo *PartnerUpdateOne) SetKeycloakOrganisation(s string) *PartnerUpdateOne {
	puo.mutation.SetKeycloakOrganisation(s)
	return puo
}

// SetNillableKeycloakOrganisation sets the "keycloak_organisation" field if the given value is not nil.
func (puo *PartnerUpdateOne) SetNillableKeycloakOrganisation(s *string) *PartnerUpdateOne {
	if s != nil {
		puo.SetKeycloakOrganisation(*s)
	}
	return puo
}

// SetWhamPartnerURL sets the "wham_partner_url" field.
func (puo *PartnerUpdateOne) SetWhamPartnerURL(s string) *PartnerUpdateOne {
	puo.mutation.SetWhamPartnerURL(s)
	return puo
}

// SetNillableWhamPartnerURL sets the "wham_partner_url" field if the given value is not nil.
func (puo *PartnerUpdateOne) SetNillableWhamPartnerURL(s *string) *PartnerUpdateOne {
	if s != nil {
		puo.SetWhamPartnerURL(*s)
	}
	return puo
}

// SetWhamCreated sets the "wham_created" field.
func (puo *PartnerUpdateOne) SetWhamCreated(t time.Time) *PartnerUpdateOne {
	puo.mutation.SetWhamCreated(t)
	return puo
}

// SetNillableWhamCreated sets the "wham_created" field if the given value is not nil.
func (puo *PartnerUpdateOne) SetNillableWhamCreated(t *time.Time) *PartnerUpdateOne {
	if t != nil {
		puo.SetWhamCreated(*t)
	}
	return puo
}

// SetWhamUpdated sets the "wham_updated" field.
func (puo *PartnerUpdateOne) SetWhamUpdated(t time.Time) *PartnerUpdateOne {
	puo.mutation.SetWhamUpdated(t)
	return puo
}

// SetNillableWhamUpdated sets the "wham_updated" field if the given value is not nil.
func (puo *PartnerUpdateOne) SetNillableWhamUpdated(t *time.Time) *PartnerUpdateOne {
	if t != nil {
		puo.SetWhamUpdated(*t)
	}
	return puo
}

// AddApplicationIDs adds the "application" edge to the Application entity by IDs.
func (puo *PartnerUpdateOne) AddApplicationIDs(ids ...uuid.UUID) *PartnerUpdateOne {
	puo.mutation.AddApplicationIDs(ids...)
	return puo
}

// AddApplication adds the "application" edges to the Application entity.
func (puo *PartnerUpdateOne) AddApplication(a ...*Application) *PartnerUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.AddApplicationIDs(ids...)
}

// Mutation returns the PartnerMutation object of the builder.
func (puo *PartnerUpdateOne) Mutation() *PartnerMutation {
	return puo.mutation
}

// ClearApplication clears all "application" edges to the Application entity.
func (puo *PartnerUpdateOne) ClearApplication() *PartnerUpdateOne {
	puo.mutation.ClearApplication()
	return puo
}

// RemoveApplicationIDs removes the "application" edge to Application entities by IDs.
func (puo *PartnerUpdateOne) RemoveApplicationIDs(ids ...uuid.UUID) *PartnerUpdateOne {
	puo.mutation.RemoveApplicationIDs(ids...)
	return puo
}

// RemoveApplication removes "application" edges to Application entities.
func (puo *PartnerUpdateOne) RemoveApplication(a ...*Application) *PartnerUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.RemoveApplicationIDs(ids...)
}

// Where appends a list predicates to the PartnerUpdate builder.
func (puo *PartnerUpdateOne) Where(ps ...predicate.Partner) *PartnerUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PartnerUpdateOne) Select(field string, fields ...string) *PartnerUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Partner entity.
func (puo *PartnerUpdateOne) Save(ctx context.Context) (*Partner, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PartnerUpdateOne) SaveX(ctx context.Context) *Partner {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PartnerUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PartnerUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PartnerUpdateOne) sqlSave(ctx context.Context) (_node *Partner, err error) {
	_spec := sqlgraph.NewUpdateSpec(partner.Table, partner.Columns, sqlgraph.NewFieldSpec(partner.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Partner.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, partner.FieldID)
		for _, f := range fields {
			if !partner.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != partner.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.WhamSiteID(); ok {
		_spec.SetField(partner.FieldWhamSiteID, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedWhamSiteID(); ok {
		_spec.AddField(partner.FieldWhamSiteID, field.TypeInt, value)
	}
	if value, ok := puo.mutation.WhamTitle(); ok {
		_spec.SetField(partner.FieldWhamTitle, field.TypeString, value)
	}
	if value, ok := puo.mutation.WhamDescription(); ok {
		_spec.SetField(partner.FieldWhamDescription, field.TypeString, value)
	}
	if value, ok := puo.mutation.KeycloakOrganisation(); ok {
		_spec.SetField(partner.FieldKeycloakOrganisation, field.TypeString, value)
	}
	if value, ok := puo.mutation.WhamPartnerURL(); ok {
		_spec.SetField(partner.FieldWhamPartnerURL, field.TypeString, value)
	}
	if value, ok := puo.mutation.WhamCreated(); ok {
		_spec.SetField(partner.FieldWhamCreated, field.TypeTime, value)
	}
	if value, ok := puo.mutation.WhamUpdated(); ok {
		_spec.SetField(partner.FieldWhamUpdated, field.TypeTime, value)
	}
	if puo.mutation.ApplicationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   partner.ApplicationTable,
			Columns: partner.ApplicationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(application.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedApplicationIDs(); len(nodes) > 0 && !puo.mutation.ApplicationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   partner.ApplicationTable,
			Columns: partner.ApplicationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(application.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ApplicationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   partner.ApplicationTable,
			Columns: partner.ApplicationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(application.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Partner{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{partner.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
