// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"nursing_api/pkg/ent/predicate"
	"nursing_api/pkg/ent/takehistory"
	"nursing_api/pkg/ent/timezone"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TakeHistoryQuery is the builder for querying TakeHistory entities.
type TakeHistoryQuery struct {
	config
	ctx          *QueryContext
	order        []takehistory.OrderOption
	inters       []Interceptor
	predicates   []predicate.TakeHistory
	withTimezone *TimeZoneQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TakeHistoryQuery builder.
func (thq *TakeHistoryQuery) Where(ps ...predicate.TakeHistory) *TakeHistoryQuery {
	thq.predicates = append(thq.predicates, ps...)
	return thq
}

// Limit the number of records to be returned by this query.
func (thq *TakeHistoryQuery) Limit(limit int) *TakeHistoryQuery {
	thq.ctx.Limit = &limit
	return thq
}

// Offset to start from.
func (thq *TakeHistoryQuery) Offset(offset int) *TakeHistoryQuery {
	thq.ctx.Offset = &offset
	return thq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (thq *TakeHistoryQuery) Unique(unique bool) *TakeHistoryQuery {
	thq.ctx.Unique = &unique
	return thq
}

// Order specifies how the records should be ordered.
func (thq *TakeHistoryQuery) Order(o ...takehistory.OrderOption) *TakeHistoryQuery {
	thq.order = append(thq.order, o...)
	return thq
}

// QueryTimezone chains the current query on the "timezone" edge.
func (thq *TakeHistoryQuery) QueryTimezone() *TimeZoneQuery {
	query := (&TimeZoneClient{config: thq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := thq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := thq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(takehistory.Table, takehistory.FieldID, selector),
			sqlgraph.To(timezone.Table, timezone.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, takehistory.TimezoneTable, takehistory.TimezoneColumn),
		)
		fromU = sqlgraph.SetNeighbors(thq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TakeHistory entity from the query.
// Returns a *NotFoundError when no TakeHistory was found.
func (thq *TakeHistoryQuery) First(ctx context.Context) (*TakeHistory, error) {
	nodes, err := thq.Limit(1).All(setContextOp(ctx, thq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{takehistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (thq *TakeHistoryQuery) FirstX(ctx context.Context) *TakeHistory {
	node, err := thq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TakeHistory ID from the query.
// Returns a *NotFoundError when no TakeHistory ID was found.
func (thq *TakeHistoryQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = thq.Limit(1).IDs(setContextOp(ctx, thq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{takehistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (thq *TakeHistoryQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := thq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TakeHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TakeHistory entity is found.
// Returns a *NotFoundError when no TakeHistory entities are found.
func (thq *TakeHistoryQuery) Only(ctx context.Context) (*TakeHistory, error) {
	nodes, err := thq.Limit(2).All(setContextOp(ctx, thq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{takehistory.Label}
	default:
		return nil, &NotSingularError{takehistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (thq *TakeHistoryQuery) OnlyX(ctx context.Context) *TakeHistory {
	node, err := thq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TakeHistory ID in the query.
// Returns a *NotSingularError when more than one TakeHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (thq *TakeHistoryQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = thq.Limit(2).IDs(setContextOp(ctx, thq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{takehistory.Label}
	default:
		err = &NotSingularError{takehistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (thq *TakeHistoryQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := thq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TakeHistories.
func (thq *TakeHistoryQuery) All(ctx context.Context) ([]*TakeHistory, error) {
	ctx = setContextOp(ctx, thq.ctx, "All")
	if err := thq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TakeHistory, *TakeHistoryQuery]()
	return withInterceptors[[]*TakeHistory](ctx, thq, qr, thq.inters)
}

// AllX is like All, but panics if an error occurs.
func (thq *TakeHistoryQuery) AllX(ctx context.Context) []*TakeHistory {
	nodes, err := thq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TakeHistory IDs.
func (thq *TakeHistoryQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if thq.ctx.Unique == nil && thq.path != nil {
		thq.Unique(true)
	}
	ctx = setContextOp(ctx, thq.ctx, "IDs")
	if err = thq.Select(takehistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (thq *TakeHistoryQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := thq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (thq *TakeHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, thq.ctx, "Count")
	if err := thq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, thq, querierCount[*TakeHistoryQuery](), thq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (thq *TakeHistoryQuery) CountX(ctx context.Context) int {
	count, err := thq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (thq *TakeHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, thq.ctx, "Exist")
	switch _, err := thq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (thq *TakeHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := thq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TakeHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (thq *TakeHistoryQuery) Clone() *TakeHistoryQuery {
	if thq == nil {
		return nil
	}
	return &TakeHistoryQuery{
		config:       thq.config,
		ctx:          thq.ctx.Clone(),
		order:        append([]takehistory.OrderOption{}, thq.order...),
		inters:       append([]Interceptor{}, thq.inters...),
		predicates:   append([]predicate.TakeHistory{}, thq.predicates...),
		withTimezone: thq.withTimezone.Clone(),
		// clone intermediate query.
		sql:  thq.sql.Clone(),
		path: thq.path,
	}
}

// WithTimezone tells the query-builder to eager-load the nodes that are connected to
// the "timezone" edge. The optional arguments are used to configure the query builder of the edge.
func (thq *TakeHistoryQuery) WithTimezone(opts ...func(*TimeZoneQuery)) *TakeHistoryQuery {
	query := (&TimeZoneClient{config: thq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	thq.withTimezone = query
	return thq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TakeHistory.Query().
//		GroupBy(takehistory.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (thq *TakeHistoryQuery) GroupBy(field string, fields ...string) *TakeHistoryGroupBy {
	thq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TakeHistoryGroupBy{build: thq}
	grbuild.flds = &thq.ctx.Fields
	grbuild.label = takehistory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//	}
//
//	client.TakeHistory.Query().
//		Select(takehistory.FieldUserID).
//		Scan(ctx, &v)
func (thq *TakeHistoryQuery) Select(fields ...string) *TakeHistorySelect {
	thq.ctx.Fields = append(thq.ctx.Fields, fields...)
	sbuild := &TakeHistorySelect{TakeHistoryQuery: thq}
	sbuild.label = takehistory.Label
	sbuild.flds, sbuild.scan = &thq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TakeHistorySelect configured with the given aggregations.
func (thq *TakeHistoryQuery) Aggregate(fns ...AggregateFunc) *TakeHistorySelect {
	return thq.Select().Aggregate(fns...)
}

func (thq *TakeHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range thq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, thq); err != nil {
				return err
			}
		}
	}
	for _, f := range thq.ctx.Fields {
		if !takehistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if thq.path != nil {
		prev, err := thq.path(ctx)
		if err != nil {
			return err
		}
		thq.sql = prev
	}
	return nil
}

func (thq *TakeHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TakeHistory, error) {
	var (
		nodes       = []*TakeHistory{}
		_spec       = thq.querySpec()
		loadedTypes = [1]bool{
			thq.withTimezone != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TakeHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TakeHistory{config: thq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, thq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := thq.withTimezone; query != nil {
		if err := thq.loadTimezone(ctx, query, nodes, nil,
			func(n *TakeHistory, e *TimeZone) { n.Edges.Timezone = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (thq *TakeHistoryQuery) loadTimezone(ctx context.Context, query *TimeZoneQuery, nodes []*TakeHistory, init func(*TakeHistory), assign func(*TakeHistory, *TimeZone)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*TakeHistory)
	for i := range nodes {
		fk := nodes[i].TimezoneID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(timezone.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "timezone_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (thq *TakeHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := thq.querySpec()
	_spec.Node.Columns = thq.ctx.Fields
	if len(thq.ctx.Fields) > 0 {
		_spec.Unique = thq.ctx.Unique != nil && *thq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, thq.driver, _spec)
}

func (thq *TakeHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(takehistory.Table, takehistory.Columns, sqlgraph.NewFieldSpec(takehistory.FieldID, field.TypeUUID))
	_spec.From = thq.sql
	if unique := thq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if thq.path != nil {
		_spec.Unique = true
	}
	if fields := thq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, takehistory.FieldID)
		for i := range fields {
			if fields[i] != takehistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if thq.withTimezone != nil {
			_spec.Node.AddColumnOnce(takehistory.FieldTimezoneID)
		}
	}
	if ps := thq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := thq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := thq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := thq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (thq *TakeHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(thq.driver.Dialect())
	t1 := builder.Table(takehistory.Table)
	columns := thq.ctx.Fields
	if len(columns) == 0 {
		columns = takehistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if thq.sql != nil {
		selector = thq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if thq.ctx.Unique != nil && *thq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range thq.predicates {
		p(selector)
	}
	for _, p := range thq.order {
		p(selector)
	}
	if offset := thq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := thq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TakeHistoryGroupBy is the group-by builder for TakeHistory entities.
type TakeHistoryGroupBy struct {
	selector
	build *TakeHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (thgb *TakeHistoryGroupBy) Aggregate(fns ...AggregateFunc) *TakeHistoryGroupBy {
	thgb.fns = append(thgb.fns, fns...)
	return thgb
}

// Scan applies the selector query and scans the result into the given value.
func (thgb *TakeHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, thgb.build.ctx, "GroupBy")
	if err := thgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TakeHistoryQuery, *TakeHistoryGroupBy](ctx, thgb.build, thgb, thgb.build.inters, v)
}

func (thgb *TakeHistoryGroupBy) sqlScan(ctx context.Context, root *TakeHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(thgb.fns))
	for _, fn := range thgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*thgb.flds)+len(thgb.fns))
		for _, f := range *thgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*thgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := thgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TakeHistorySelect is the builder for selecting fields of TakeHistory entities.
type TakeHistorySelect struct {
	*TakeHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ths *TakeHistorySelect) Aggregate(fns ...AggregateFunc) *TakeHistorySelect {
	ths.fns = append(ths.fns, fns...)
	return ths
}

// Scan applies the selector query and scans the result into the given value.
func (ths *TakeHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ths.ctx, "Select")
	if err := ths.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TakeHistoryQuery, *TakeHistorySelect](ctx, ths.TakeHistoryQuery, ths, ths.inters, v)
}

func (ths *TakeHistorySelect) sqlScan(ctx context.Context, root *TakeHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ths.fns))
	for _, fn := range ths.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ths.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ths.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
