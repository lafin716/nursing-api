// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"nursing_api/pkg/ent/predicate"
	"nursing_api/pkg/ent/takehistorymemo"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TakeHistoryMemoQuery is the builder for querying TakeHistoryMemo entities.
type TakeHistoryMemoQuery struct {
	config
	ctx        *QueryContext
	order      []takehistorymemo.OrderOption
	inters     []Interceptor
	predicates []predicate.TakeHistoryMemo
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TakeHistoryMemoQuery builder.
func (thmq *TakeHistoryMemoQuery) Where(ps ...predicate.TakeHistoryMemo) *TakeHistoryMemoQuery {
	thmq.predicates = append(thmq.predicates, ps...)
	return thmq
}

// Limit the number of records to be returned by this query.
func (thmq *TakeHistoryMemoQuery) Limit(limit int) *TakeHistoryMemoQuery {
	thmq.ctx.Limit = &limit
	return thmq
}

// Offset to start from.
func (thmq *TakeHistoryMemoQuery) Offset(offset int) *TakeHistoryMemoQuery {
	thmq.ctx.Offset = &offset
	return thmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (thmq *TakeHistoryMemoQuery) Unique(unique bool) *TakeHistoryMemoQuery {
	thmq.ctx.Unique = &unique
	return thmq
}

// Order specifies how the records should be ordered.
func (thmq *TakeHistoryMemoQuery) Order(o ...takehistorymemo.OrderOption) *TakeHistoryMemoQuery {
	thmq.order = append(thmq.order, o...)
	return thmq
}

// First returns the first TakeHistoryMemo entity from the query.
// Returns a *NotFoundError when no TakeHistoryMemo was found.
func (thmq *TakeHistoryMemoQuery) First(ctx context.Context) (*TakeHistoryMemo, error) {
	nodes, err := thmq.Limit(1).All(setContextOp(ctx, thmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{takehistorymemo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (thmq *TakeHistoryMemoQuery) FirstX(ctx context.Context) *TakeHistoryMemo {
	node, err := thmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TakeHistoryMemo ID from the query.
// Returns a *NotFoundError when no TakeHistoryMemo ID was found.
func (thmq *TakeHistoryMemoQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = thmq.Limit(1).IDs(setContextOp(ctx, thmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{takehistorymemo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (thmq *TakeHistoryMemoQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := thmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TakeHistoryMemo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TakeHistoryMemo entity is found.
// Returns a *NotFoundError when no TakeHistoryMemo entities are found.
func (thmq *TakeHistoryMemoQuery) Only(ctx context.Context) (*TakeHistoryMemo, error) {
	nodes, err := thmq.Limit(2).All(setContextOp(ctx, thmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{takehistorymemo.Label}
	default:
		return nil, &NotSingularError{takehistorymemo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (thmq *TakeHistoryMemoQuery) OnlyX(ctx context.Context) *TakeHistoryMemo {
	node, err := thmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TakeHistoryMemo ID in the query.
// Returns a *NotSingularError when more than one TakeHistoryMemo ID is found.
// Returns a *NotFoundError when no entities are found.
func (thmq *TakeHistoryMemoQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = thmq.Limit(2).IDs(setContextOp(ctx, thmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{takehistorymemo.Label}
	default:
		err = &NotSingularError{takehistorymemo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (thmq *TakeHistoryMemoQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := thmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TakeHistoryMemos.
func (thmq *TakeHistoryMemoQuery) All(ctx context.Context) ([]*TakeHistoryMemo, error) {
	ctx = setContextOp(ctx, thmq.ctx, "All")
	if err := thmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TakeHistoryMemo, *TakeHistoryMemoQuery]()
	return withInterceptors[[]*TakeHistoryMemo](ctx, thmq, qr, thmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (thmq *TakeHistoryMemoQuery) AllX(ctx context.Context) []*TakeHistoryMemo {
	nodes, err := thmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TakeHistoryMemo IDs.
func (thmq *TakeHistoryMemoQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if thmq.ctx.Unique == nil && thmq.path != nil {
		thmq.Unique(true)
	}
	ctx = setContextOp(ctx, thmq.ctx, "IDs")
	if err = thmq.Select(takehistorymemo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (thmq *TakeHistoryMemoQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := thmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (thmq *TakeHistoryMemoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, thmq.ctx, "Count")
	if err := thmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, thmq, querierCount[*TakeHistoryMemoQuery](), thmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (thmq *TakeHistoryMemoQuery) CountX(ctx context.Context) int {
	count, err := thmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (thmq *TakeHistoryMemoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, thmq.ctx, "Exist")
	switch _, err := thmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (thmq *TakeHistoryMemoQuery) ExistX(ctx context.Context) bool {
	exist, err := thmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TakeHistoryMemoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (thmq *TakeHistoryMemoQuery) Clone() *TakeHistoryMemoQuery {
	if thmq == nil {
		return nil
	}
	return &TakeHistoryMemoQuery{
		config:     thmq.config,
		ctx:        thmq.ctx.Clone(),
		order:      append([]takehistorymemo.OrderOption{}, thmq.order...),
		inters:     append([]Interceptor{}, thmq.inters...),
		predicates: append([]predicate.TakeHistoryMemo{}, thmq.predicates...),
		// clone intermediate query.
		sql:  thmq.sql.Clone(),
		path: thmq.path,
	}
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
//	client.TakeHistoryMemo.Query().
//		GroupBy(takehistorymemo.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (thmq *TakeHistoryMemoQuery) GroupBy(field string, fields ...string) *TakeHistoryMemoGroupBy {
	thmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TakeHistoryMemoGroupBy{build: thmq}
	grbuild.flds = &thmq.ctx.Fields
	grbuild.label = takehistorymemo.Label
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
//	client.TakeHistoryMemo.Query().
//		Select(takehistorymemo.FieldUserID).
//		Scan(ctx, &v)
func (thmq *TakeHistoryMemoQuery) Select(fields ...string) *TakeHistoryMemoSelect {
	thmq.ctx.Fields = append(thmq.ctx.Fields, fields...)
	sbuild := &TakeHistoryMemoSelect{TakeHistoryMemoQuery: thmq}
	sbuild.label = takehistorymemo.Label
	sbuild.flds, sbuild.scan = &thmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TakeHistoryMemoSelect configured with the given aggregations.
func (thmq *TakeHistoryMemoQuery) Aggregate(fns ...AggregateFunc) *TakeHistoryMemoSelect {
	return thmq.Select().Aggregate(fns...)
}

func (thmq *TakeHistoryMemoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range thmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, thmq); err != nil {
				return err
			}
		}
	}
	for _, f := range thmq.ctx.Fields {
		if !takehistorymemo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if thmq.path != nil {
		prev, err := thmq.path(ctx)
		if err != nil {
			return err
		}
		thmq.sql = prev
	}
	return nil
}

func (thmq *TakeHistoryMemoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TakeHistoryMemo, error) {
	var (
		nodes = []*TakeHistoryMemo{}
		_spec = thmq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TakeHistoryMemo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TakeHistoryMemo{config: thmq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, thmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (thmq *TakeHistoryMemoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := thmq.querySpec()
	_spec.Node.Columns = thmq.ctx.Fields
	if len(thmq.ctx.Fields) > 0 {
		_spec.Unique = thmq.ctx.Unique != nil && *thmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, thmq.driver, _spec)
}

func (thmq *TakeHistoryMemoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(takehistorymemo.Table, takehistorymemo.Columns, sqlgraph.NewFieldSpec(takehistorymemo.FieldID, field.TypeUUID))
	_spec.From = thmq.sql
	if unique := thmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if thmq.path != nil {
		_spec.Unique = true
	}
	if fields := thmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, takehistorymemo.FieldID)
		for i := range fields {
			if fields[i] != takehistorymemo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := thmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := thmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := thmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := thmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (thmq *TakeHistoryMemoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(thmq.driver.Dialect())
	t1 := builder.Table(takehistorymemo.Table)
	columns := thmq.ctx.Fields
	if len(columns) == 0 {
		columns = takehistorymemo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if thmq.sql != nil {
		selector = thmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if thmq.ctx.Unique != nil && *thmq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range thmq.predicates {
		p(selector)
	}
	for _, p := range thmq.order {
		p(selector)
	}
	if offset := thmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := thmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TakeHistoryMemoGroupBy is the group-by builder for TakeHistoryMemo entities.
type TakeHistoryMemoGroupBy struct {
	selector
	build *TakeHistoryMemoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (thmgb *TakeHistoryMemoGroupBy) Aggregate(fns ...AggregateFunc) *TakeHistoryMemoGroupBy {
	thmgb.fns = append(thmgb.fns, fns...)
	return thmgb
}

// Scan applies the selector query and scans the result into the given value.
func (thmgb *TakeHistoryMemoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, thmgb.build.ctx, "GroupBy")
	if err := thmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TakeHistoryMemoQuery, *TakeHistoryMemoGroupBy](ctx, thmgb.build, thmgb, thmgb.build.inters, v)
}

func (thmgb *TakeHistoryMemoGroupBy) sqlScan(ctx context.Context, root *TakeHistoryMemoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(thmgb.fns))
	for _, fn := range thmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*thmgb.flds)+len(thmgb.fns))
		for _, f := range *thmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*thmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := thmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TakeHistoryMemoSelect is the builder for selecting fields of TakeHistoryMemo entities.
type TakeHistoryMemoSelect struct {
	*TakeHistoryMemoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (thms *TakeHistoryMemoSelect) Aggregate(fns ...AggregateFunc) *TakeHistoryMemoSelect {
	thms.fns = append(thms.fns, fns...)
	return thms
}

// Scan applies the selector query and scans the result into the given value.
func (thms *TakeHistoryMemoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, thms.ctx, "Select")
	if err := thms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TakeHistoryMemoQuery, *TakeHistoryMemoSelect](ctx, thms.TakeHistoryMemoQuery, thms, thms.inters, v)
}

func (thms *TakeHistoryMemoSelect) sqlScan(ctx context.Context, root *TakeHistoryMemoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(thms.fns))
	for _, fn := range thms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*thms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := thms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
