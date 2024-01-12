// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"nursing_api/pkg/ent/predicate"
	"nursing_api/pkg/ent/prescriptionitem"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PrescriptionItemDelete is the builder for deleting a PrescriptionItem entity.
type PrescriptionItemDelete struct {
	config
	hooks    []Hook
	mutation *PrescriptionItemMutation
}

// Where appends a list predicates to the PrescriptionItemDelete builder.
func (pid *PrescriptionItemDelete) Where(ps ...predicate.PrescriptionItem) *PrescriptionItemDelete {
	pid.mutation.Where(ps...)
	return pid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pid *PrescriptionItemDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pid.sqlExec, pid.mutation, pid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pid *PrescriptionItemDelete) ExecX(ctx context.Context) int {
	n, err := pid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pid *PrescriptionItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(prescriptionitem.Table, sqlgraph.NewFieldSpec(prescriptionitem.FieldID, field.TypeUUID))
	if ps := pid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pid.mutation.done = true
	return affected, err
}

// PrescriptionItemDeleteOne is the builder for deleting a single PrescriptionItem entity.
type PrescriptionItemDeleteOne struct {
	pid *PrescriptionItemDelete
}

// Where appends a list predicates to the PrescriptionItemDelete builder.
func (pido *PrescriptionItemDeleteOne) Where(ps ...predicate.PrescriptionItem) *PrescriptionItemDeleteOne {
	pido.pid.mutation.Where(ps...)
	return pido
}

// Exec executes the deletion query.
func (pido *PrescriptionItemDeleteOne) Exec(ctx context.Context) error {
	n, err := pido.pid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{prescriptionitem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pido *PrescriptionItemDeleteOne) ExecX(ctx context.Context) {
	if err := pido.Exec(ctx); err != nil {
		panic(err)
	}
}
