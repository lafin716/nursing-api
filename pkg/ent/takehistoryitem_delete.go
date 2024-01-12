// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"nursing_api/pkg/ent/predicate"
	"nursing_api/pkg/ent/takehistoryitem"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TakeHistoryItemDelete is the builder for deleting a TakeHistoryItem entity.
type TakeHistoryItemDelete struct {
	config
	hooks    []Hook
	mutation *TakeHistoryItemMutation
}

// Where appends a list predicates to the TakeHistoryItemDelete builder.
func (thid *TakeHistoryItemDelete) Where(ps ...predicate.TakeHistoryItem) *TakeHistoryItemDelete {
	thid.mutation.Where(ps...)
	return thid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (thid *TakeHistoryItemDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, thid.sqlExec, thid.mutation, thid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (thid *TakeHistoryItemDelete) ExecX(ctx context.Context) int {
	n, err := thid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (thid *TakeHistoryItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(takehistoryitem.Table, sqlgraph.NewFieldSpec(takehistoryitem.FieldID, field.TypeInt))
	if ps := thid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, thid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	thid.mutation.done = true
	return affected, err
}

// TakeHistoryItemDeleteOne is the builder for deleting a single TakeHistoryItem entity.
type TakeHistoryItemDeleteOne struct {
	thid *TakeHistoryItemDelete
}

// Where appends a list predicates to the TakeHistoryItemDelete builder.
func (thido *TakeHistoryItemDeleteOne) Where(ps ...predicate.TakeHistoryItem) *TakeHistoryItemDeleteOne {
	thido.thid.mutation.Where(ps...)
	return thido
}

// Exec executes the deletion query.
func (thido *TakeHistoryItemDeleteOne) Exec(ctx context.Context) error {
	n, err := thido.thid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{takehistoryitem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (thido *TakeHistoryItemDeleteOne) ExecX(ctx context.Context) {
	if err := thido.Exec(ctx); err != nil {
		panic(err)
	}
}
