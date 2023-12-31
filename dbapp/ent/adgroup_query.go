// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"dbapp/ent/adgroup"
	"dbapp/ent/community"
	"dbapp/ent/predicate"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ADGroupQuery is the builder for querying ADGroup entities.
type ADGroupQuery struct {
	config
	ctx           *QueryContext
	order         []adgroup.OrderOption
	inters        []Interceptor
	predicates    []predicate.ADGroup
	withCommunity *CommunityQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ADGroupQuery builder.
func (agq *ADGroupQuery) Where(ps ...predicate.ADGroup) *ADGroupQuery {
	agq.predicates = append(agq.predicates, ps...)
	return agq
}

// Limit the number of records to be returned by this query.
func (agq *ADGroupQuery) Limit(limit int) *ADGroupQuery {
	agq.ctx.Limit = &limit
	return agq
}

// Offset to start from.
func (agq *ADGroupQuery) Offset(offset int) *ADGroupQuery {
	agq.ctx.Offset = &offset
	return agq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (agq *ADGroupQuery) Unique(unique bool) *ADGroupQuery {
	agq.ctx.Unique = &unique
	return agq
}

// Order specifies how the records should be ordered.
func (agq *ADGroupQuery) Order(o ...adgroup.OrderOption) *ADGroupQuery {
	agq.order = append(agq.order, o...)
	return agq
}

// QueryCommunity chains the current query on the "community" edge.
func (agq *ADGroupQuery) QueryCommunity() *CommunityQuery {
	query := (&CommunityClient{config: agq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := agq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := agq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(adgroup.Table, adgroup.FieldID, selector),
			sqlgraph.To(community.Table, community.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, adgroup.CommunityTable, adgroup.CommunityPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(agq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ADGroup entity from the query.
// Returns a *NotFoundError when no ADGroup was found.
func (agq *ADGroupQuery) First(ctx context.Context) (*ADGroup, error) {
	nodes, err := agq.Limit(1).All(setContextOp(ctx, agq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{adgroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (agq *ADGroupQuery) FirstX(ctx context.Context) *ADGroup {
	node, err := agq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ADGroup ID from the query.
// Returns a *NotFoundError when no ADGroup ID was found.
func (agq *ADGroupQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = agq.Limit(1).IDs(setContextOp(ctx, agq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{adgroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (agq *ADGroupQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := agq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ADGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ADGroup entity is found.
// Returns a *NotFoundError when no ADGroup entities are found.
func (agq *ADGroupQuery) Only(ctx context.Context) (*ADGroup, error) {
	nodes, err := agq.Limit(2).All(setContextOp(ctx, agq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{adgroup.Label}
	default:
		return nil, &NotSingularError{adgroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (agq *ADGroupQuery) OnlyX(ctx context.Context) *ADGroup {
	node, err := agq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ADGroup ID in the query.
// Returns a *NotSingularError when more than one ADGroup ID is found.
// Returns a *NotFoundError when no entities are found.
func (agq *ADGroupQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = agq.Limit(2).IDs(setContextOp(ctx, agq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{adgroup.Label}
	default:
		err = &NotSingularError{adgroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (agq *ADGroupQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := agq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ADGroups.
func (agq *ADGroupQuery) All(ctx context.Context) ([]*ADGroup, error) {
	ctx = setContextOp(ctx, agq.ctx, "All")
	if err := agq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ADGroup, *ADGroupQuery]()
	return withInterceptors[[]*ADGroup](ctx, agq, qr, agq.inters)
}

// AllX is like All, but panics if an error occurs.
func (agq *ADGroupQuery) AllX(ctx context.Context) []*ADGroup {
	nodes, err := agq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ADGroup IDs.
func (agq *ADGroupQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if agq.ctx.Unique == nil && agq.path != nil {
		agq.Unique(true)
	}
	ctx = setContextOp(ctx, agq.ctx, "IDs")
	if err = agq.Select(adgroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (agq *ADGroupQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := agq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (agq *ADGroupQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, agq.ctx, "Count")
	if err := agq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, agq, querierCount[*ADGroupQuery](), agq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (agq *ADGroupQuery) CountX(ctx context.Context) int {
	count, err := agq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (agq *ADGroupQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, agq.ctx, "Exist")
	switch _, err := agq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (agq *ADGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := agq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ADGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (agq *ADGroupQuery) Clone() *ADGroupQuery {
	if agq == nil {
		return nil
	}
	return &ADGroupQuery{
		config:        agq.config,
		ctx:           agq.ctx.Clone(),
		order:         append([]adgroup.OrderOption{}, agq.order...),
		inters:        append([]Interceptor{}, agq.inters...),
		predicates:    append([]predicate.ADGroup{}, agq.predicates...),
		withCommunity: agq.withCommunity.Clone(),
		// clone intermediate query.
		sql:  agq.sql.Clone(),
		path: agq.path,
	}
}

// WithCommunity tells the query-builder to eager-load the nodes that are connected to
// the "community" edge. The optional arguments are used to configure the query builder of the edge.
func (agq *ADGroupQuery) WithCommunity(opts ...func(*CommunityQuery)) *ADGroupQuery {
	query := (&CommunityClient{config: agq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	agq.withCommunity = query
	return agq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ADGroup.Query().
//		GroupBy(adgroup.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (agq *ADGroupQuery) GroupBy(field string, fields ...string) *ADGroupGroupBy {
	agq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ADGroupGroupBy{build: agq}
	grbuild.flds = &agq.ctx.Fields
	grbuild.label = adgroup.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.ADGroup.Query().
//		Select(adgroup.FieldName).
//		Scan(ctx, &v)
func (agq *ADGroupQuery) Select(fields ...string) *ADGroupSelect {
	agq.ctx.Fields = append(agq.ctx.Fields, fields...)
	sbuild := &ADGroupSelect{ADGroupQuery: agq}
	sbuild.label = adgroup.Label
	sbuild.flds, sbuild.scan = &agq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ADGroupSelect configured with the given aggregations.
func (agq *ADGroupQuery) Aggregate(fns ...AggregateFunc) *ADGroupSelect {
	return agq.Select().Aggregate(fns...)
}

func (agq *ADGroupQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range agq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, agq); err != nil {
				return err
			}
		}
	}
	for _, f := range agq.ctx.Fields {
		if !adgroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if agq.path != nil {
		prev, err := agq.path(ctx)
		if err != nil {
			return err
		}
		agq.sql = prev
	}
	return nil
}

func (agq *ADGroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ADGroup, error) {
	var (
		nodes       = []*ADGroup{}
		_spec       = agq.querySpec()
		loadedTypes = [1]bool{
			agq.withCommunity != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ADGroup).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ADGroup{config: agq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, agq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := agq.withCommunity; query != nil {
		if err := agq.loadCommunity(ctx, query, nodes,
			func(n *ADGroup) { n.Edges.Community = []*Community{} },
			func(n *ADGroup, e *Community) { n.Edges.Community = append(n.Edges.Community, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (agq *ADGroupQuery) loadCommunity(ctx context.Context, query *CommunityQuery, nodes []*ADGroup, init func(*ADGroup), assign func(*ADGroup, *Community)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*ADGroup)
	nids := make(map[uuid.UUID]map[*ADGroup]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(adgroup.CommunityTable)
		s.Join(joinT).On(s.C(community.FieldID), joinT.C(adgroup.CommunityPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(adgroup.CommunityPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(adgroup.CommunityPrimaryKey[0]))
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
					nids[inValue] = map[*ADGroup]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Community](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "community" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (agq *ADGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := agq.querySpec()
	_spec.Node.Columns = agq.ctx.Fields
	if len(agq.ctx.Fields) > 0 {
		_spec.Unique = agq.ctx.Unique != nil && *agq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, agq.driver, _spec)
}

func (agq *ADGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(adgroup.Table, adgroup.Columns, sqlgraph.NewFieldSpec(adgroup.FieldID, field.TypeUUID))
	_spec.From = agq.sql
	if unique := agq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if agq.path != nil {
		_spec.Unique = true
	}
	if fields := agq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adgroup.FieldID)
		for i := range fields {
			if fields[i] != adgroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := agq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := agq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := agq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := agq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (agq *ADGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(agq.driver.Dialect())
	t1 := builder.Table(adgroup.Table)
	columns := agq.ctx.Fields
	if len(columns) == 0 {
		columns = adgroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if agq.sql != nil {
		selector = agq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if agq.ctx.Unique != nil && *agq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range agq.predicates {
		p(selector)
	}
	for _, p := range agq.order {
		p(selector)
	}
	if offset := agq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := agq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ADGroupGroupBy is the group-by builder for ADGroup entities.
type ADGroupGroupBy struct {
	selector
	build *ADGroupQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aggb *ADGroupGroupBy) Aggregate(fns ...AggregateFunc) *ADGroupGroupBy {
	aggb.fns = append(aggb.fns, fns...)
	return aggb
}

// Scan applies the selector query and scans the result into the given value.
func (aggb *ADGroupGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aggb.build.ctx, "GroupBy")
	if err := aggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ADGroupQuery, *ADGroupGroupBy](ctx, aggb.build, aggb, aggb.build.inters, v)
}

func (aggb *ADGroupGroupBy) sqlScan(ctx context.Context, root *ADGroupQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(aggb.fns))
	for _, fn := range aggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*aggb.flds)+len(aggb.fns))
		for _, f := range *aggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*aggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ADGroupSelect is the builder for selecting fields of ADGroup entities.
type ADGroupSelect struct {
	*ADGroupQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ags *ADGroupSelect) Aggregate(fns ...AggregateFunc) *ADGroupSelect {
	ags.fns = append(ags.fns, fns...)
	return ags
}

// Scan applies the selector query and scans the result into the given value.
func (ags *ADGroupSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ags.ctx, "Select")
	if err := ags.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ADGroupQuery, *ADGroupSelect](ctx, ags.ADGroupQuery, ags, ags.inters, v)
}

func (ags *ADGroupSelect) sqlScan(ctx context.Context, root *ADGroupQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ags.fns))
	for _, fn := range ags.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ags.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ags.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
