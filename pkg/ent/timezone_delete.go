// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"nursing_api/pkg/ent/predicate"
	"nursing_api/pkg/ent/timezone"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TimeZoneDelete is the builder for deleting a TimeZone entity.
type TimeZoneDelete struct {
	config
	hooks    []Hook
	mutation *TimeZoneMutation
}

// Where appends a list predicates to the TimeZoneDelete builder.
func (tzd *TimeZoneDelete) Where(ps ...predicate.TimeZone) *TimeZoneDelete {
	tzd.mutation.Where(ps...)
	return tzd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tzd *TimeZoneDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tzd.sqlExec, tzd.mutation, tzd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tzd *TimeZoneDelete) ExecX(ctx context.Context) int {
	n, err := tzd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tzd *TimeZoneDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(timezone.Table, sqlgraph.NewFieldSpec(timezone.FieldID, field.TypeUUID))
	if ps := tzd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tzd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tzd.mutation.done = true
	return affected, err
}

// TimeZoneDeleteOne is the builder for deleting a single TimeZone entity.
type TimeZoneDeleteOne struct {
	tzd *TimeZoneDelete
}

// Where appends a list predicates to the TimeZoneDelete builder.
func (tzdo *TimeZoneDeleteOne) Where(ps ...predicate.TimeZone) *TimeZoneDeleteOne {
	tzdo.tzd.mutation.Where(ps...)
	return tzdo
}

// Exec executes the deletion query.
func (tzdo *TimeZoneDeleteOne) Exec(ctx context.Context) error {
	n, err := tzdo.tzd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{timezone.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tzdo *TimeZoneDeleteOne) ExecX(ctx context.Context) {
	if err := tzdo.Exec(ctx); err != nil {
		panic(err)
	}
}
