// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"nursing_api/pkg/ent/predicate"
	"nursing_api/pkg/ent/timezone"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TimeZoneQuery is the builder for querying TimeZone entities.
type TimeZoneQuery struct {
	config
	ctx        *QueryContext
	order      []timezone.OrderOption
	inters     []Interceptor
	predicates []predicate.TimeZone
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TimeZoneQuery builder.
func (tzq *TimeZoneQuery) Where(ps ...predicate.TimeZone) *TimeZoneQuery {
	tzq.predicates = append(tzq.predicates, ps...)
	return tzq
}

// Limit the number of records to be returned by this query.
func (tzq *TimeZoneQuery) Limit(limit int) *TimeZoneQuery {
	tzq.ctx.Limit = &limit
	return tzq
}

// Offset to start from.
func (tzq *TimeZoneQuery) Offset(offset int) *TimeZoneQuery {
	tzq.ctx.Offset = &offset
	return tzq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tzq *TimeZoneQuery) Unique(unique bool) *TimeZoneQuery {
	tzq.ctx.Unique = &unique
	return tzq
}

// Order specifies how the records should be ordered.
func (tzq *TimeZoneQuery) Order(o ...timezone.OrderOption) *TimeZoneQuery {
	tzq.order = append(tzq.order, o...)
	return tzq
}

// First returns the first TimeZone entity from the query.
// Returns a *NotFoundError when no TimeZone was found.
func (tzq *TimeZoneQuery) First(ctx context.Context) (*TimeZone, error) {
	nodes, err := tzq.Limit(1).All(setContextOp(ctx, tzq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{timezone.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tzq *TimeZoneQuery) FirstX(ctx context.Context) *TimeZone {
	node, err := tzq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TimeZone ID from the query.
// Returns a *NotFoundError when no TimeZone ID was found.
func (tzq *TimeZoneQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tzq.Limit(1).IDs(setContextOp(ctx, tzq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{timezone.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tzq *TimeZoneQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := tzq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TimeZone entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TimeZone entity is found.
// Returns a *NotFoundError when no TimeZone entities are found.
func (tzq *TimeZoneQuery) Only(ctx context.Context) (*TimeZone, error) {
	nodes, err := tzq.Limit(2).All(setContextOp(ctx, tzq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{timezone.Label}
	default:
		return nil, &NotSingularError{timezone.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tzq *TimeZoneQuery) OnlyX(ctx context.Context) *TimeZone {
	node, err := tzq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TimeZone ID in the query.
// Returns a *NotSingularError when more than one TimeZone ID is found.
// Returns a *NotFoundError when no entities are found.
func (tzq *TimeZoneQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tzq.Limit(2).IDs(setContextOp(ctx, tzq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{timezone.Label}
	default:
		err = &NotSingularError{timezone.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tzq *TimeZoneQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := tzq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TimeZones.
func (tzq *TimeZoneQuery) All(ctx context.Context) ([]*TimeZone, error) {
	ctx = setContextOp(ctx, tzq.ctx, "All")
	if err := tzq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TimeZone, *TimeZoneQuery]()
	return withInterceptors[[]*TimeZone](ctx, tzq, qr, tzq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tzq *TimeZoneQuery) AllX(ctx context.Context) []*TimeZone {
	nodes, err := tzq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TimeZone IDs.
func (tzq *TimeZoneQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if tzq.ctx.Unique == nil && tzq.path != nil {
		tzq.Unique(true)
	}
	ctx = setContextOp(ctx, tzq.ctx, "IDs")
	if err = tzq.Select(timezone.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tzq *TimeZoneQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := tzq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tzq *TimeZoneQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tzq.ctx, "Count")
	if err := tzq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tzq, querierCount[*TimeZoneQuery](), tzq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tzq *TimeZoneQuery) CountX(ctx context.Context) int {
	count, err := tzq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tzq *TimeZoneQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tzq.ctx, "Exist")
	switch _, err := tzq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tzq *TimeZoneQuery) ExistX(ctx context.Context) bool {
	exist, err := tzq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TimeZoneQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tzq *TimeZoneQuery) Clone() *TimeZoneQuery {
	if tzq == nil {
		return nil
	}
	return &TimeZoneQuery{
		config:     tzq.config,
		ctx:        tzq.ctx.Clone(),
		order:      append([]timezone.OrderOption{}, tzq.order...),
		inters:     append([]Interceptor{}, tzq.inters...),
		predicates: append([]predicate.TimeZone{}, tzq.predicates...),
		// clone intermediate query.
		sql:  tzq.sql.Clone(),
		path: tzq.path,
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
//	client.TimeZone.Query().
//		GroupBy(timezone.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tzq *TimeZoneQuery) GroupBy(field string, fields ...string) *TimeZoneGroupBy {
	tzq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TimeZoneGroupBy{build: tzq}
	grbuild.flds = &tzq.ctx.Fields
	grbuild.label = timezone.Label
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
//	client.TimeZone.Query().
//		Select(timezone.FieldUserID).
//		Scan(ctx, &v)
func (tzq *TimeZoneQuery) Select(fields ...string) *TimeZoneSelect {
	tzq.ctx.Fields = append(tzq.ctx.Fields, fields...)
	sbuild := &TimeZoneSelect{TimeZoneQuery: tzq}
	sbuild.label = timezone.Label
	sbuild.flds, sbuild.scan = &tzq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TimeZoneSelect configured with the given aggregations.
func (tzq *TimeZoneQuery) Aggregate(fns ...AggregateFunc) *TimeZoneSelect {
	return tzq.Select().Aggregate(fns...)
}

func (tzq *TimeZoneQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tzq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tzq); err != nil {
				return err
			}
		}
	}
	for _, f := range tzq.ctx.Fields {
		if !timezone.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tzq.path != nil {
		prev, err := tzq.path(ctx)
		if err != nil {
			return err
		}
		tzq.sql = prev
	}
	return nil
}

func (tzq *TimeZoneQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TimeZone, error) {
	var (
		nodes = []*TimeZone{}
		_spec = tzq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TimeZone).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TimeZone{config: tzq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tzq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tzq *TimeZoneQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tzq.querySpec()
	_spec.Node.Columns = tzq.ctx.Fields
	if len(tzq.ctx.Fields) > 0 {
		_spec.Unique = tzq.ctx.Unique != nil && *tzq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tzq.driver, _spec)
}

func (tzq *TimeZoneQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(timezone.Table, timezone.Columns, sqlgraph.NewFieldSpec(timezone.FieldID, field.TypeUUID))
	_spec.From = tzq.sql
	if unique := tzq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tzq.path != nil {
		_spec.Unique = true
	}
	if fields := tzq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, timezone.FieldID)
		for i := range fields {
			if fields[i] != timezone.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tzq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tzq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tzq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tzq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tzq *TimeZoneQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tzq.driver.Dialect())
	t1 := builder.Table(timezone.Table)
	columns := tzq.ctx.Fields
	if len(columns) == 0 {
		columns = timezone.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tzq.sql != nil {
		selector = tzq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tzq.ctx.Unique != nil && *tzq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tzq.predicates {
		p(selector)
	}
	for _, p := range tzq.order {
		p(selector)
	}
	if offset := tzq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tzq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TimeZoneGroupBy is the group-by builder for TimeZone entities.
type TimeZoneGroupBy struct {
	selector
	build *TimeZoneQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tzgb *TimeZoneGroupBy) Aggregate(fns ...AggregateFunc) *TimeZoneGroupBy {
	tzgb.fns = append(tzgb.fns, fns...)
	return tzgb
}

// Scan applies the selector query and scans the result into the given value.
func (tzgb *TimeZoneGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tzgb.build.ctx, "GroupBy")
	if err := tzgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TimeZoneQuery, *TimeZoneGroupBy](ctx, tzgb.build, tzgb, tzgb.build.inters, v)
}

func (tzgb *TimeZoneGroupBy) sqlScan(ctx context.Context, root *TimeZoneQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tzgb.fns))
	for _, fn := range tzgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tzgb.flds)+len(tzgb.fns))
		for _, f := range *tzgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tzgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tzgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TimeZoneSelect is the builder for selecting fields of TimeZone entities.
type TimeZoneSelect struct {
	*TimeZoneQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tzs *TimeZoneSelect) Aggregate(fns ...AggregateFunc) *TimeZoneSelect {
	tzs.fns = append(tzs.fns, fns...)
	return tzs
}

// Scan applies the selector query and scans the result into the given value.
func (tzs *TimeZoneSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tzs.ctx, "Select")
	if err := tzs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TimeZoneQuery, *TimeZoneSelect](ctx, tzs.TimeZoneQuery, tzs, tzs.inters, v)
}

func (tzs *TimeZoneSelect) sqlScan(ctx context.Context, root *TimeZoneQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tzs.fns))
	for _, fn := range tzs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tzs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tzs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
