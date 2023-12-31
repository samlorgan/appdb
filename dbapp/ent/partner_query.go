// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"dbapp/ent/application"
	"dbapp/ent/partner"
	"dbapp/ent/predicate"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PartnerQuery is the builder for querying Partner entities.
type PartnerQuery struct {
	config
	ctx             *QueryContext
	order           []partner.OrderOption
	inters          []Interceptor
	predicates      []predicate.Partner
	withApplication *ApplicationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PartnerQuery builder.
func (pq *PartnerQuery) Where(ps ...predicate.Partner) *PartnerQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PartnerQuery) Limit(limit int) *PartnerQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PartnerQuery) Offset(offset int) *PartnerQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PartnerQuery) Unique(unique bool) *PartnerQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PartnerQuery) Order(o ...partner.OrderOption) *PartnerQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryApplication chains the current query on the "application" edge.
func (pq *PartnerQuery) QueryApplication() *ApplicationQuery {
	query := (&ApplicationClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(partner.Table, partner.FieldID, selector),
			sqlgraph.To(application.Table, application.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, partner.ApplicationTable, partner.ApplicationPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Partner entity from the query.
// Returns a *NotFoundError when no Partner was found.
func (pq *PartnerQuery) First(ctx context.Context) (*Partner, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{partner.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PartnerQuery) FirstX(ctx context.Context) *Partner {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Partner ID from the query.
// Returns a *NotFoundError when no Partner ID was found.
func (pq *PartnerQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{partner.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PartnerQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Partner entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Partner entity is found.
// Returns a *NotFoundError when no Partner entities are found.
func (pq *PartnerQuery) Only(ctx context.Context) (*Partner, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{partner.Label}
	default:
		return nil, &NotSingularError{partner.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PartnerQuery) OnlyX(ctx context.Context) *Partner {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Partner ID in the query.
// Returns a *NotSingularError when more than one Partner ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PartnerQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{partner.Label}
	default:
		err = &NotSingularError{partner.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PartnerQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Partners.
func (pq *PartnerQuery) All(ctx context.Context) ([]*Partner, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Partner, *PartnerQuery]()
	return withInterceptors[[]*Partner](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PartnerQuery) AllX(ctx context.Context) []*Partner {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Partner IDs.
func (pq *PartnerQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(partner.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PartnerQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PartnerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PartnerQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PartnerQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PartnerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PartnerQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PartnerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PartnerQuery) Clone() *PartnerQuery {
	if pq == nil {
		return nil
	}
	return &PartnerQuery{
		config:          pq.config,
		ctx:             pq.ctx.Clone(),
		order:           append([]partner.OrderOption{}, pq.order...),
		inters:          append([]Interceptor{}, pq.inters...),
		predicates:      append([]predicate.Partner{}, pq.predicates...),
		withApplication: pq.withApplication.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithApplication tells the query-builder to eager-load the nodes that are connected to
// the "application" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PartnerQuery) WithApplication(opts ...func(*ApplicationQuery)) *PartnerQuery {
	query := (&ApplicationClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withApplication = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		WhamSiteID int `json:"wham_site_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Partner.Query().
//		GroupBy(partner.FieldWhamSiteID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PartnerQuery) GroupBy(field string, fields ...string) *PartnerGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PartnerGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = partner.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		WhamSiteID int `json:"wham_site_id,omitempty"`
//	}
//
//	client.Partner.Query().
//		Select(partner.FieldWhamSiteID).
//		Scan(ctx, &v)
func (pq *PartnerQuery) Select(fields ...string) *PartnerSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PartnerSelect{PartnerQuery: pq}
	sbuild.label = partner.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PartnerSelect configured with the given aggregations.
func (pq *PartnerQuery) Aggregate(fns ...AggregateFunc) *PartnerSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PartnerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !partner.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PartnerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Partner, error) {
	var (
		nodes       = []*Partner{}
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withApplication != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Partner).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Partner{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withApplication; query != nil {
		if err := pq.loadApplication(ctx, query, nodes,
			func(n *Partner) { n.Edges.Application = []*Application{} },
			func(n *Partner, e *Application) { n.Edges.Application = append(n.Edges.Application, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PartnerQuery) loadApplication(ctx context.Context, query *ApplicationQuery, nodes []*Partner, init func(*Partner), assign func(*Partner, *Application)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Partner)
	nids := make(map[uuid.UUID]map[*Partner]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(partner.ApplicationTable)
		s.Join(joinT).On(s.C(application.FieldID), joinT.C(partner.ApplicationPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(partner.ApplicationPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(partner.ApplicationPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Partner]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Application](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "application" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (pq *PartnerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PartnerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(partner.Table, partner.Columns, sqlgraph.NewFieldSpec(partner.FieldID, field.TypeUUID))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, partner.FieldID)
		for i := range fields {
			if fields[i] != partner.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PartnerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(partner.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = partner.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PartnerGroupBy is the group-by builder for Partner entities.
type PartnerGroupBy struct {
	selector
	build *PartnerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PartnerGroupBy) Aggregate(fns ...AggregateFunc) *PartnerGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PartnerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PartnerQuery, *PartnerGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PartnerGroupBy) sqlScan(ctx context.Context, root *PartnerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PartnerSelect is the builder for selecting fields of Partner entities.
type PartnerSelect struct {
	*PartnerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PartnerSelect) Aggregate(fns ...AggregateFunc) *PartnerSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PartnerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PartnerQuery, *PartnerSelect](ctx, ps.PartnerQuery, ps, ps.inters, v)
}

func (ps *PartnerSelect) sqlScan(ctx context.Context, root *PartnerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
