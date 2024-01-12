// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"nursing_api/pkg/ent/predicate"
	"nursing_api/pkg/ent/prescriptionitem"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PrescriptionItemQuery is the builder for querying PrescriptionItem entities.
type PrescriptionItemQuery struct {
	config
	ctx        *QueryContext
	order      []prescriptionitem.OrderOption
	inters     []Interceptor
	predicates []predicate.PrescriptionItem
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PrescriptionItemQuery builder.
func (piq *PrescriptionItemQuery) Where(ps ...predicate.PrescriptionItem) *PrescriptionItemQuery {
	piq.predicates = append(piq.predicates, ps...)
	return piq
}

// Limit the number of records to be returned by this query.
func (piq *PrescriptionItemQuery) Limit(limit int) *PrescriptionItemQuery {
	piq.ctx.Limit = &limit
	return piq
}

// Offset to start from.
func (piq *PrescriptionItemQuery) Offset(offset int) *PrescriptionItemQuery {
	piq.ctx.Offset = &offset
	return piq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (piq *PrescriptionItemQuery) Unique(unique bool) *PrescriptionItemQuery {
	piq.ctx.Unique = &unique
	return piq
}

// Order specifies how the records should be ordered.
func (piq *PrescriptionItemQuery) Order(o ...prescriptionitem.OrderOption) *PrescriptionItemQuery {
	piq.order = append(piq.order, o...)
	return piq
}

// First returns the first PrescriptionItem entity from the query.
// Returns a *NotFoundError when no PrescriptionItem was found.
func (piq *PrescriptionItemQuery) First(ctx context.Context) (*PrescriptionItem, error) {
	nodes, err := piq.Limit(1).All(setContextOp(ctx, piq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{prescriptionitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (piq *PrescriptionItemQuery) FirstX(ctx context.Context) *PrescriptionItem {
	node, err := piq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PrescriptionItem ID from the query.
// Returns a *NotFoundError when no PrescriptionItem ID was found.
func (piq *PrescriptionItemQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = piq.Limit(1).IDs(setContextOp(ctx, piq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{prescriptionitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (piq *PrescriptionItemQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := piq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PrescriptionItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PrescriptionItem entity is found.
// Returns a *NotFoundError when no PrescriptionItem entities are found.
func (piq *PrescriptionItemQuery) Only(ctx context.Context) (*PrescriptionItem, error) {
	nodes, err := piq.Limit(2).All(setContextOp(ctx, piq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{prescriptionitem.Label}
	default:
		return nil, &NotSingularError{prescriptionitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (piq *PrescriptionItemQuery) OnlyX(ctx context.Context) *PrescriptionItem {
	node, err := piq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PrescriptionItem ID in the query.
// Returns a *NotSingularError when more than one PrescriptionItem ID is found.
// Returns a *NotFoundError when no entities are found.
func (piq *PrescriptionItemQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = piq.Limit(2).IDs(setContextOp(ctx, piq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{prescriptionitem.Label}
	default:
		err = &NotSingularError{prescriptionitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (piq *PrescriptionItemQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := piq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PrescriptionItems.
func (piq *PrescriptionItemQuery) All(ctx context.Context) ([]*PrescriptionItem, error) {
	ctx = setContextOp(ctx, piq.ctx, "All")
	if err := piq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PrescriptionItem, *PrescriptionItemQuery]()
	return withInterceptors[[]*PrescriptionItem](ctx, piq, qr, piq.inters)
}

// AllX is like All, but panics if an error occurs.
func (piq *PrescriptionItemQuery) AllX(ctx context.Context) []*PrescriptionItem {
	nodes, err := piq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PrescriptionItem IDs.
func (piq *PrescriptionItemQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if piq.ctx.Unique == nil && piq.path != nil {
		piq.Unique(true)
	}
	ctx = setContextOp(ctx, piq.ctx, "IDs")
	if err = piq.Select(prescriptionitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (piq *PrescriptionItemQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := piq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (piq *PrescriptionItemQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, piq.ctx, "Count")
	if err := piq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, piq, querierCount[*PrescriptionItemQuery](), piq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (piq *PrescriptionItemQuery) CountX(ctx context.Context) int {
	count, err := piq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (piq *PrescriptionItemQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, piq.ctx, "Exist")
	switch _, err := piq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (piq *PrescriptionItemQuery) ExistX(ctx context.Context) bool {
	exist, err := piq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PrescriptionItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (piq *PrescriptionItemQuery) Clone() *PrescriptionItemQuery {
	if piq == nil {
		return nil
	}
	return &PrescriptionItemQuery{
		config:     piq.config,
		ctx:        piq.ctx.Clone(),
		order:      append([]prescriptionitem.OrderOption{}, piq.order...),
		inters:     append([]Interceptor{}, piq.inters...),
		predicates: append([]predicate.PrescriptionItem{}, piq.predicates...),
		// clone intermediate query.
		sql:  piq.sql.Clone(),
		path: piq.path,
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
//	client.PrescriptionItem.Query().
//		GroupBy(prescriptionitem.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (piq *PrescriptionItemQuery) GroupBy(field string, fields ...string) *PrescriptionItemGroupBy {
	piq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PrescriptionItemGroupBy{build: piq}
	grbuild.flds = &piq.ctx.Fields
	grbuild.label = prescriptionitem.Label
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
//	client.PrescriptionItem.Query().
//		Select(prescriptionitem.FieldUserID).
//		Scan(ctx, &v)
func (piq *PrescriptionItemQuery) Select(fields ...string) *PrescriptionItemSelect {
	piq.ctx.Fields = append(piq.ctx.Fields, fields...)
	sbuild := &PrescriptionItemSelect{PrescriptionItemQuery: piq}
	sbuild.label = prescriptionitem.Label
	sbuild.flds, sbuild.scan = &piq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PrescriptionItemSelect configured with the given aggregations.
func (piq *PrescriptionItemQuery) Aggregate(fns ...AggregateFunc) *PrescriptionItemSelect {
	return piq.Select().Aggregate(fns...)
}

func (piq *PrescriptionItemQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range piq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, piq); err != nil {
				return err
			}
		}
	}
	for _, f := range piq.ctx.Fields {
		if !prescriptionitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if piq.path != nil {
		prev, err := piq.path(ctx)
		if err != nil {
			return err
		}
		piq.sql = prev
	}
	return nil
}

func (piq *PrescriptionItemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PrescriptionItem, error) {
	var (
		nodes = []*PrescriptionItem{}
		_spec = piq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PrescriptionItem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PrescriptionItem{config: piq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, piq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (piq *PrescriptionItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := piq.querySpec()
	_spec.Node.Columns = piq.ctx.Fields
	if len(piq.ctx.Fields) > 0 {
		_spec.Unique = piq.ctx.Unique != nil && *piq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, piq.driver, _spec)
}

func (piq *PrescriptionItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(prescriptionitem.Table, prescriptionitem.Columns, sqlgraph.NewFieldSpec(prescriptionitem.FieldID, field.TypeUUID))
	_spec.From = piq.sql
	if unique := piq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if piq.path != nil {
		_spec.Unique = true
	}
	if fields := piq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, prescriptionitem.FieldID)
		for i := range fields {
			if fields[i] != prescriptionitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := piq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := piq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := piq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := piq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (piq *PrescriptionItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(piq.driver.Dialect())
	t1 := builder.Table(prescriptionitem.Table)
	columns := piq.ctx.Fields
	if len(columns) == 0 {
		columns = prescriptionitem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if piq.sql != nil {
		selector = piq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if piq.ctx.Unique != nil && *piq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range piq.predicates {
		p(selector)
	}
	for _, p := range piq.order {
		p(selector)
	}
	if offset := piq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := piq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PrescriptionItemGroupBy is the group-by builder for PrescriptionItem entities.
type PrescriptionItemGroupBy struct {
	selector
	build *PrescriptionItemQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pigb *PrescriptionItemGroupBy) Aggregate(fns ...AggregateFunc) *PrescriptionItemGroupBy {
	pigb.fns = append(pigb.fns, fns...)
	return pigb
}

// Scan applies the selector query and scans the result into the given value.
func (pigb *PrescriptionItemGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pigb.build.ctx, "GroupBy")
	if err := pigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PrescriptionItemQuery, *PrescriptionItemGroupBy](ctx, pigb.build, pigb, pigb.build.inters, v)
}

func (pigb *PrescriptionItemGroupBy) sqlScan(ctx context.Context, root *PrescriptionItemQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pigb.fns))
	for _, fn := range pigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pigb.flds)+len(pigb.fns))
		for _, f := range *pigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PrescriptionItemSelect is the builder for selecting fields of PrescriptionItem entities.
type PrescriptionItemSelect struct {
	*PrescriptionItemQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pis *PrescriptionItemSelect) Aggregate(fns ...AggregateFunc) *PrescriptionItemSelect {
	pis.fns = append(pis.fns, fns...)
	return pis
}

// Scan applies the selector query and scans the result into the given value.
func (pis *PrescriptionItemSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pis.ctx, "Select")
	if err := pis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PrescriptionItemQuery, *PrescriptionItemSelect](ctx, pis.PrescriptionItemQuery, pis, pis.inters, v)
}

func (pis *PrescriptionItemSelect) sqlScan(ctx context.Context, root *PrescriptionItemQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pis.fns))
	for _, fn := range pis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
