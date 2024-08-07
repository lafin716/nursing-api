// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"nursing_api/pkg/ent/predicate"
	"nursing_api/pkg/ent/takehistorymemo"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TakeHistoryMemoUpdate is the builder for updating TakeHistoryMemo entities.
type TakeHistoryMemoUpdate struct {
	config
	hooks    []Hook
	mutation *TakeHistoryMemoMutation
}

// Where appends a list predicates to the TakeHistoryMemoUpdate builder.
func (thmu *TakeHistoryMemoUpdate) Where(ps ...predicate.TakeHistoryMemo) *TakeHistoryMemoUpdate {
	thmu.mutation.Where(ps...)
	return thmu
}

// SetUserID sets the "user_id" field.
func (thmu *TakeHistoryMemoUpdate) SetUserID(u uuid.UUID) *TakeHistoryMemoUpdate {
	thmu.mutation.SetUserID(u)
	return thmu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (thmu *TakeHistoryMemoUpdate) SetNillableUserID(u *uuid.UUID) *TakeHistoryMemoUpdate {
	if u != nil {
		thmu.SetUserID(*u)
	}
	return thmu
}

// SetTimezoneID sets the "timezone_id" field.
func (thmu *TakeHistoryMemoUpdate) SetTimezoneID(u uuid.UUID) *TakeHistoryMemoUpdate {
	thmu.mutation.SetTimezoneID(u)
	return thmu
}

// SetNillableTimezoneID sets the "timezone_id" field if the given value is not nil.
func (thmu *TakeHistoryMemoUpdate) SetNillableTimezoneID(u *uuid.UUID) *TakeHistoryMemoUpdate {
	if u != nil {
		thmu.SetTimezoneID(*u)
	}
	return thmu
}

// SetTakeDate sets the "take_date" field.
func (thmu *TakeHistoryMemoUpdate) SetTakeDate(s string) *TakeHistoryMemoUpdate {
	thmu.mutation.SetTakeDate(s)
	return thmu
}

// SetNillableTakeDate sets the "take_date" field if the given value is not nil.
func (thmu *TakeHistoryMemoUpdate) SetNillableTakeDate(s *string) *TakeHistoryMemoUpdate {
	if s != nil {
		thmu.SetTakeDate(*s)
	}
	return thmu
}

// SetMemo sets the "memo" field.
func (thmu *TakeHistoryMemoUpdate) SetMemo(s string) *TakeHistoryMemoUpdate {
	thmu.mutation.SetMemo(s)
	return thmu
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (thmu *TakeHistoryMemoUpdate) SetNillableMemo(s *string) *TakeHistoryMemoUpdate {
	if s != nil {
		thmu.SetMemo(*s)
	}
	return thmu
}

// ClearMemo clears the value of the "memo" field.
func (thmu *TakeHistoryMemoUpdate) ClearMemo() *TakeHistoryMemoUpdate {
	thmu.mutation.ClearMemo()
	return thmu
}

// SetCreatedAt sets the "created_at" field.
func (thmu *TakeHistoryMemoUpdate) SetCreatedAt(t time.Time) *TakeHistoryMemoUpdate {
	thmu.mutation.SetCreatedAt(t)
	return thmu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (thmu *TakeHistoryMemoUpdate) SetNillableCreatedAt(t *time.Time) *TakeHistoryMemoUpdate {
	if t != nil {
		thmu.SetCreatedAt(*t)
	}
	return thmu
}

// SetUpdatedAt sets the "updated_at" field.
func (thmu *TakeHistoryMemoUpdate) SetUpdatedAt(t time.Time) *TakeHistoryMemoUpdate {
	thmu.mutation.SetUpdatedAt(t)
	return thmu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (thmu *TakeHistoryMemoUpdate) SetNillableUpdatedAt(t *time.Time) *TakeHistoryMemoUpdate {
	if t != nil {
		thmu.SetUpdatedAt(*t)
	}
	return thmu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (thmu *TakeHistoryMemoUpdate) ClearUpdatedAt() *TakeHistoryMemoUpdate {
	thmu.mutation.ClearUpdatedAt()
	return thmu
}

// Mutation returns the TakeHistoryMemoMutation object of the builder.
func (thmu *TakeHistoryMemoUpdate) Mutation() *TakeHistoryMemoMutation {
	return thmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (thmu *TakeHistoryMemoUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, thmu.sqlSave, thmu.mutation, thmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (thmu *TakeHistoryMemoUpdate) SaveX(ctx context.Context) int {
	affected, err := thmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (thmu *TakeHistoryMemoUpdate) Exec(ctx context.Context) error {
	_, err := thmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (thmu *TakeHistoryMemoUpdate) ExecX(ctx context.Context) {
	if err := thmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (thmu *TakeHistoryMemoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(takehistorymemo.Table, takehistorymemo.Columns, sqlgraph.NewFieldSpec(takehistorymemo.FieldID, field.TypeUUID))
	if ps := thmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := thmu.mutation.UserID(); ok {
		_spec.SetField(takehistorymemo.FieldUserID, field.TypeUUID, value)
	}
	if value, ok := thmu.mutation.TimezoneID(); ok {
		_spec.SetField(takehistorymemo.FieldTimezoneID, field.TypeUUID, value)
	}
	if value, ok := thmu.mutation.TakeDate(); ok {
		_spec.SetField(takehistorymemo.FieldTakeDate, field.TypeString, value)
	}
	if value, ok := thmu.mutation.Memo(); ok {
		_spec.SetField(takehistorymemo.FieldMemo, field.TypeString, value)
	}
	if thmu.mutation.MemoCleared() {
		_spec.ClearField(takehistorymemo.FieldMemo, field.TypeString)
	}
	if value, ok := thmu.mutation.CreatedAt(); ok {
		_spec.SetField(takehistorymemo.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := thmu.mutation.UpdatedAt(); ok {
		_spec.SetField(takehistorymemo.FieldUpdatedAt, field.TypeTime, value)
	}
	if thmu.mutation.UpdatedAtCleared() {
		_spec.ClearField(takehistorymemo.FieldUpdatedAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, thmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{takehistorymemo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	thmu.mutation.done = true
	return n, nil
}

// TakeHistoryMemoUpdateOne is the builder for updating a single TakeHistoryMemo entity.
type TakeHistoryMemoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TakeHistoryMemoMutation
}

// SetUserID sets the "user_id" field.
func (thmuo *TakeHistoryMemoUpdateOne) SetUserID(u uuid.UUID) *TakeHistoryMemoUpdateOne {
	thmuo.mutation.SetUserID(u)
	return thmuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (thmuo *TakeHistoryMemoUpdateOne) SetNillableUserID(u *uuid.UUID) *TakeHistoryMemoUpdateOne {
	if u != nil {
		thmuo.SetUserID(*u)
	}
	return thmuo
}

// SetTimezoneID sets the "timezone_id" field.
func (thmuo *TakeHistoryMemoUpdateOne) SetTimezoneID(u uuid.UUID) *TakeHistoryMemoUpdateOne {
	thmuo.mutation.SetTimezoneID(u)
	return thmuo
}

// SetNillableTimezoneID sets the "timezone_id" field if the given value is not nil.
func (thmuo *TakeHistoryMemoUpdateOne) SetNillableTimezoneID(u *uuid.UUID) *TakeHistoryMemoUpdateOne {
	if u != nil {
		thmuo.SetTimezoneID(*u)
	}
	return thmuo
}

// SetTakeDate sets the "take_date" field.
func (thmuo *TakeHistoryMemoUpdateOne) SetTakeDate(s string) *TakeHistoryMemoUpdateOne {
	thmuo.mutation.SetTakeDate(s)
	return thmuo
}

// SetNillableTakeDate sets the "take_date" field if the given value is not nil.
func (thmuo *TakeHistoryMemoUpdateOne) SetNillableTakeDate(s *string) *TakeHistoryMemoUpdateOne {
	if s != nil {
		thmuo.SetTakeDate(*s)
	}
	return thmuo
}

// SetMemo sets the "memo" field.
func (thmuo *TakeHistoryMemoUpdateOne) SetMemo(s string) *TakeHistoryMemoUpdateOne {
	thmuo.mutation.SetMemo(s)
	return thmuo
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (thmuo *TakeHistoryMemoUpdateOne) SetNillableMemo(s *string) *TakeHistoryMemoUpdateOne {
	if s != nil {
		thmuo.SetMemo(*s)
	}
	return thmuo
}

// ClearMemo clears the value of the "memo" field.
func (thmuo *TakeHistoryMemoUpdateOne) ClearMemo() *TakeHistoryMemoUpdateOne {
	thmuo.mutation.ClearMemo()
	return thmuo
}

// SetCreatedAt sets the "created_at" field.
func (thmuo *TakeHistoryMemoUpdateOne) SetCreatedAt(t time.Time) *TakeHistoryMemoUpdateOne {
	thmuo.mutation.SetCreatedAt(t)
	return thmuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (thmuo *TakeHistoryMemoUpdateOne) SetNillableCreatedAt(t *time.Time) *TakeHistoryMemoUpdateOne {
	if t != nil {
		thmuo.SetCreatedAt(*t)
	}
	return thmuo
}

// SetUpdatedAt sets the "updated_at" field.
func (thmuo *TakeHistoryMemoUpdateOne) SetUpdatedAt(t time.Time) *TakeHistoryMemoUpdateOne {
	thmuo.mutation.SetUpdatedAt(t)
	return thmuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (thmuo *TakeHistoryMemoUpdateOne) SetNillableUpdatedAt(t *time.Time) *TakeHistoryMemoUpdateOne {
	if t != nil {
		thmuo.SetUpdatedAt(*t)
	}
	return thmuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (thmuo *TakeHistoryMemoUpdateOne) ClearUpdatedAt() *TakeHistoryMemoUpdateOne {
	thmuo.mutation.ClearUpdatedAt()
	return thmuo
}

// Mutation returns the TakeHistoryMemoMutation object of the builder.
func (thmuo *TakeHistoryMemoUpdateOne) Mutation() *TakeHistoryMemoMutation {
	return thmuo.mutation
}

// Where appends a list predicates to the TakeHistoryMemoUpdate builder.
func (thmuo *TakeHistoryMemoUpdateOne) Where(ps ...predicate.TakeHistoryMemo) *TakeHistoryMemoUpdateOne {
	thmuo.mutation.Where(ps...)
	return thmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (thmuo *TakeHistoryMemoUpdateOne) Select(field string, fields ...string) *TakeHistoryMemoUpdateOne {
	thmuo.fields = append([]string{field}, fields...)
	return thmuo
}

// Save executes the query and returns the updated TakeHistoryMemo entity.
func (thmuo *TakeHistoryMemoUpdateOne) Save(ctx context.Context) (*TakeHistoryMemo, error) {
	return withHooks(ctx, thmuo.sqlSave, thmuo.mutation, thmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (thmuo *TakeHistoryMemoUpdateOne) SaveX(ctx context.Context) *TakeHistoryMemo {
	node, err := thmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (thmuo *TakeHistoryMemoUpdateOne) Exec(ctx context.Context) error {
	_, err := thmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (thmuo *TakeHistoryMemoUpdateOne) ExecX(ctx context.Context) {
	if err := thmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (thmuo *TakeHistoryMemoUpdateOne) sqlSave(ctx context.Context) (_node *TakeHistoryMemo, err error) {
	_spec := sqlgraph.NewUpdateSpec(takehistorymemo.Table, takehistorymemo.Columns, sqlgraph.NewFieldSpec(takehistorymemo.FieldID, field.TypeUUID))
	id, ok := thmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TakeHistoryMemo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := thmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, takehistorymemo.FieldID)
		for _, f := range fields {
			if !takehistorymemo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != takehistorymemo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := thmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := thmuo.mutation.UserID(); ok {
		_spec.SetField(takehistorymemo.FieldUserID, field.TypeUUID, value)
	}
	if value, ok := thmuo.mutation.TimezoneID(); ok {
		_spec.SetField(takehistorymemo.FieldTimezoneID, field.TypeUUID, value)
	}
	if value, ok := thmuo.mutation.TakeDate(); ok {
		_spec.SetField(takehistorymemo.FieldTakeDate, field.TypeString, value)
	}
	if value, ok := thmuo.mutation.Memo(); ok {
		_spec.SetField(takehistorymemo.FieldMemo, field.TypeString, value)
	}
	if thmuo.mutation.MemoCleared() {
		_spec.ClearField(takehistorymemo.FieldMemo, field.TypeString)
	}
	if value, ok := thmuo.mutation.CreatedAt(); ok {
		_spec.SetField(takehistorymemo.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := thmuo.mutation.UpdatedAt(); ok {
		_spec.SetField(takehistorymemo.FieldUpdatedAt, field.TypeTime, value)
	}
	if thmuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(takehistorymemo.FieldUpdatedAt, field.TypeTime)
	}
	_node = &TakeHistoryMemo{config: thmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, thmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{takehistorymemo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	thmuo.mutation.done = true
	return _node, nil
}
