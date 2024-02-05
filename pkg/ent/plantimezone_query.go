// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"nursing_api/pkg/ent/plantimezone"
	"nursing_api/pkg/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PlanTimeZoneQuery is the builder for querying PlanTimeZone entities.
type PlanTimeZoneQuery struct {
	config
	ctx        *QueryContext
	order      []plantimezone.OrderOption
	inters     []Interceptor
	predicates []predicate.PlanTimeZone
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PlanTimeZoneQuery builder.
func (ptzq *PlanTimeZoneQuery) Where(ps ...predicate.PlanTimeZone) *PlanTimeZoneQuery {
	ptzq.predicates = append(ptzq.predicates, ps...)
	return ptzq
}

// Limit the number of records to be returned by this query.
func (ptzq *PlanTimeZoneQuery) Limit(limit int) *PlanTimeZoneQuery {
	ptzq.ctx.Limit = &limit
	return ptzq
}

// Offset to start from.
func (ptzq *PlanTimeZoneQuery) Offset(offset int) *PlanTimeZoneQuery {
	ptzq.ctx.Offset = &offset
	return ptzq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ptzq *PlanTimeZoneQuery) Unique(unique bool) *PlanTimeZoneQuery {
	ptzq.ctx.Unique = &unique
	return ptzq
}

// Order specifies how the records should be ordered.
func (ptzq *PlanTimeZoneQuery) Order(o ...plantimezone.OrderOption) *PlanTimeZoneQuery {
	ptzq.order = append(ptzq.order, o...)
	return ptzq
}

// First returns the first PlanTimeZone entity from the query.
// Returns a *NotFoundError when no PlanTimeZone was found.
func (ptzq *PlanTimeZoneQuery) First(ctx context.Context) (*PlanTimeZone, error) {
	nodes, err := ptzq.Limit(1).All(setContextOp(ctx, ptzq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{plantimezone.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ptzq *PlanTimeZoneQuery) FirstX(ctx context.Context) *PlanTimeZone {
	node, err := ptzq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PlanTimeZone ID from the query.
// Returns a *NotFoundError when no PlanTimeZone ID was found.
func (ptzq *PlanTimeZoneQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ptzq.Limit(1).IDs(setContextOp(ctx, ptzq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{plantimezone.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ptzq *PlanTimeZoneQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ptzq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PlanTimeZone entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PlanTimeZone entity is found.
// Returns a *NotFoundError when no PlanTimeZone entities are found.
func (ptzq *PlanTimeZoneQuery) Only(ctx context.Context) (*PlanTimeZone, error) {
	nodes, err := ptzq.Limit(2).All(setContextOp(ctx, ptzq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{plantimezone.Label}
	default:
		return nil, &NotSingularError{plantimezone.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ptzq *PlanTimeZoneQuery) OnlyX(ctx context.Context) *PlanTimeZone {
	node, err := ptzq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PlanTimeZone ID in the query.
// Returns a *NotSingularError when more than one PlanTimeZone ID is found.
// Returns a *NotFoundError when no entities are found.
func (ptzq *PlanTimeZoneQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ptzq.Limit(2).IDs(setContextOp(ctx, ptzq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{plantimezone.Label}
	default:
		err = &NotSingularError{plantimezone.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ptzq *PlanTimeZoneQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ptzq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PlanTimeZones.
func (ptzq *PlanTimeZoneQuery) All(ctx context.Context) ([]*PlanTimeZone, error) {
	ctx = setContextOp(ctx, ptzq.ctx, "All")
	if err := ptzq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PlanTimeZone, *PlanTimeZoneQuery]()
	return withInterceptors[[]*PlanTimeZone](ctx, ptzq, qr, ptzq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ptzq *PlanTimeZoneQuery) AllX(ctx context.Context) []*PlanTimeZone {
	nodes, err := ptzq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PlanTimeZone IDs.
func (ptzq *PlanTimeZoneQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ptzq.ctx.Unique == nil && ptzq.path != nil {
		ptzq.Unique(true)
	}
	ctx = setContextOp(ctx, ptzq.ctx, "IDs")
	if err = ptzq.Select(plantimezone.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ptzq *PlanTimeZoneQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ptzq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ptzq *PlanTimeZoneQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ptzq.ctx, "Count")
	if err := ptzq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ptzq, querierCount[*PlanTimeZoneQuery](), ptzq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ptzq *PlanTimeZoneQuery) CountX(ctx context.Context) int {
	count, err := ptzq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ptzq *PlanTimeZoneQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ptzq.ctx, "Exist")
	switch _, err := ptzq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ptzq *PlanTimeZoneQuery) ExistX(ctx context.Context) bool {
	exist, err := ptzq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PlanTimeZoneQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ptzq *PlanTimeZoneQuery) Clone() *PlanTimeZoneQuery {
	if ptzq == nil {
		return nil
	}
	return &PlanTimeZoneQuery{
		config:     ptzq.config,
		ctx:        ptzq.ctx.Clone(),
		order:      append([]plantimezone.OrderOption{}, ptzq.order...),
		inters:     append([]Interceptor{}, ptzq.inters...),
		predicates: append([]predicate.PlanTimeZone{}, ptzq.predicates...),
		// clone intermediate query.
		sql:  ptzq.sql.Clone(),
		path: ptzq.path,
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
//	client.PlanTimeZone.Query().
//		GroupBy(plantimezone.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ptzq *PlanTimeZoneQuery) GroupBy(field string, fields ...string) *PlanTimeZoneGroupBy {
	ptzq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PlanTimeZoneGroupBy{build: ptzq}
	grbuild.flds = &ptzq.ctx.Fields
	grbuild.label = plantimezone.Label
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
//	client.PlanTimeZone.Query().
//		Select(plantimezone.FieldUserID).
//		Scan(ctx, &v)
func (ptzq *PlanTimeZoneQuery) Select(fields ...string) *PlanTimeZoneSelect {
	ptzq.ctx.Fields = append(ptzq.ctx.Fields, fields...)
	sbuild := &PlanTimeZoneSelect{PlanTimeZoneQuery: ptzq}
	sbuild.label = plantimezone.Label
	sbuild.flds, sbuild.scan = &ptzq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PlanTimeZoneSelect configured with the given aggregations.
func (ptzq *PlanTimeZoneQuery) Aggregate(fns ...AggregateFunc) *PlanTimeZoneSelect {
	return ptzq.Select().Aggregate(fns...)
}

func (ptzq *PlanTimeZoneQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ptzq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ptzq); err != nil {
				return err
			}
		}
	}
	for _, f := range ptzq.ctx.Fields {
		if !plantimezone.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ptzq.path != nil {
		prev, err := ptzq.path(ctx)
		if err != nil {
			return err
		}
		ptzq.sql = prev
	}
	return nil
}

func (ptzq *PlanTimeZoneQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PlanTimeZone, error) {
	var (
		nodes = []*PlanTimeZone{}
		_spec = ptzq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PlanTimeZone).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PlanTimeZone{config: ptzq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ptzq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ptzq *PlanTimeZoneQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ptzq.querySpec()
	_spec.Node.Columns = ptzq.ctx.Fields
	if len(ptzq.ctx.Fields) > 0 {
		_spec.Unique = ptzq.ctx.Unique != nil && *ptzq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ptzq.driver, _spec)
}

func (ptzq *PlanTimeZoneQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(plantimezone.Table, plantimezone.Columns, sqlgraph.NewFieldSpec(plantimezone.FieldID, field.TypeUUID))
	_spec.From = ptzq.sql
	if unique := ptzq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ptzq.path != nil {
		_spec.Unique = true
	}
	if fields := ptzq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, plantimezone.FieldID)
		for i := range fields {
			if fields[i] != plantimezone.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ptzq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ptzq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ptzq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ptzq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ptzq *PlanTimeZoneQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ptzq.driver.Dialect())
	t1 := builder.Table(plantimezone.Table)
	columns := ptzq.ctx.Fields
	if len(columns) == 0 {
		columns = plantimezone.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ptzq.sql != nil {
		selector = ptzq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ptzq.ctx.Unique != nil && *ptzq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ptzq.predicates {
		p(selector)
	}
	for _, p := range ptzq.order {
		p(selector)
	}
	if offset := ptzq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ptzq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PlanTimeZoneGroupBy is the group-by builder for PlanTimeZone entities.
type PlanTimeZoneGroupBy struct {
	selector
	build *PlanTimeZoneQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ptzgb *PlanTimeZoneGroupBy) Aggregate(fns ...AggregateFunc) *PlanTimeZoneGroupBy {
	ptzgb.fns = append(ptzgb.fns, fns...)
	return ptzgb
}

// Scan applies the selector query and scans the result into the given value.
func (ptzgb *PlanTimeZoneGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ptzgb.build.ctx, "GroupBy")
	if err := ptzgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PlanTimeZoneQuery, *PlanTimeZoneGroupBy](ctx, ptzgb.build, ptzgb, ptzgb.build.inters, v)
}

func (ptzgb *PlanTimeZoneGroupBy) sqlScan(ctx context.Context, root *PlanTimeZoneQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ptzgb.fns))
	for _, fn := range ptzgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ptzgb.flds)+len(ptzgb.fns))
		for _, f := range *ptzgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ptzgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ptzgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PlanTimeZoneSelect is the builder for selecting fields of PlanTimeZone entities.
type PlanTimeZoneSelect struct {
	*PlanTimeZoneQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ptzs *PlanTimeZoneSelect) Aggregate(fns ...AggregateFunc) *PlanTimeZoneSelect {
	ptzs.fns = append(ptzs.fns, fns...)
	return ptzs
}

// Scan applies the selector query and scans the result into the given value.
func (ptzs *PlanTimeZoneSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ptzs.ctx, "Select")
	if err := ptzs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PlanTimeZoneQuery, *PlanTimeZoneSelect](ctx, ptzs.PlanTimeZoneQuery, ptzs, ptzs.inters, v)
}

func (ptzs *PlanTimeZoneSelect) sqlScan(ctx context.Context, root *PlanTimeZoneQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ptzs.fns))
	for _, fn := range ptzs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ptzs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ptzs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}