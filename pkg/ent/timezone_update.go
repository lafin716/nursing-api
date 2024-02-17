// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"nursing_api/pkg/ent/predicate"
	"nursing_api/pkg/ent/timezone"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TimeZoneUpdate is the builder for updating TimeZone entities.
type TimeZoneUpdate struct {
	config
	hooks    []Hook
	mutation *TimeZoneMutation
}

// Where appends a list predicates to the TimeZoneUpdate builder.
func (tzu *TimeZoneUpdate) Where(ps ...predicate.TimeZone) *TimeZoneUpdate {
	tzu.mutation.Where(ps...)
	return tzu
}

// SetUserID sets the "user_id" field.
func (tzu *TimeZoneUpdate) SetUserID(u uuid.UUID) *TimeZoneUpdate {
	tzu.mutation.SetUserID(u)
	return tzu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tzu *TimeZoneUpdate) SetNillableUserID(u *uuid.UUID) *TimeZoneUpdate {
	if u != nil {
		tzu.SetUserID(*u)
	}
	return tzu
}

// SetTimezoneName sets the "timezone_name" field.
func (tzu *TimeZoneUpdate) SetTimezoneName(s string) *TimeZoneUpdate {
	tzu.mutation.SetTimezoneName(s)
	return tzu
}

// SetNillableTimezoneName sets the "timezone_name" field if the given value is not nil.
func (tzu *TimeZoneUpdate) SetNillableTimezoneName(s *string) *TimeZoneUpdate {
	if s != nil {
		tzu.SetTimezoneName(*s)
	}
	return tzu
}

// ClearTimezoneName clears the value of the "timezone_name" field.
func (tzu *TimeZoneUpdate) ClearTimezoneName() *TimeZoneUpdate {
	tzu.mutation.ClearTimezoneName()
	return tzu
}

// SetIsDefault sets the "is_default" field.
func (tzu *TimeZoneUpdate) SetIsDefault(b bool) *TimeZoneUpdate {
	tzu.mutation.SetIsDefault(b)
	return tzu
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (tzu *TimeZoneUpdate) SetNillableIsDefault(b *bool) *TimeZoneUpdate {
	if b != nil {
		tzu.SetIsDefault(*b)
	}
	return tzu
}

// SetMidday sets the "midday" field.
func (tzu *TimeZoneUpdate) SetMidday(s string) *TimeZoneUpdate {
	tzu.mutation.SetMidday(s)
	return tzu
}

// SetNillableMidday sets the "midday" field if the given value is not nil.
func (tzu *TimeZoneUpdate) SetNillableMidday(s *string) *TimeZoneUpdate {
	if s != nil {
		tzu.SetMidday(*s)
	}
	return tzu
}

// SetHour sets the "hour" field.
func (tzu *TimeZoneUpdate) SetHour(s string) *TimeZoneUpdate {
	tzu.mutation.SetHour(s)
	return tzu
}

// SetNillableHour sets the "hour" field if the given value is not nil.
func (tzu *TimeZoneUpdate) SetNillableHour(s *string) *TimeZoneUpdate {
	if s != nil {
		tzu.SetHour(*s)
	}
	return tzu
}

// SetMinute sets the "minute" field.
func (tzu *TimeZoneUpdate) SetMinute(s string) *TimeZoneUpdate {
	tzu.mutation.SetMinute(s)
	return tzu
}

// SetNillableMinute sets the "minute" field if the given value is not nil.
func (tzu *TimeZoneUpdate) SetNillableMinute(s *string) *TimeZoneUpdate {
	if s != nil {
		tzu.SetMinute(*s)
	}
	return tzu
}

// SetCreatedAt sets the "created_at" field.
func (tzu *TimeZoneUpdate) SetCreatedAt(t time.Time) *TimeZoneUpdate {
	tzu.mutation.SetCreatedAt(t)
	return tzu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tzu *TimeZoneUpdate) SetNillableCreatedAt(t *time.Time) *TimeZoneUpdate {
	if t != nil {
		tzu.SetCreatedAt(*t)
	}
	return tzu
}

// SetUpdatedAt sets the "updated_at" field.
func (tzu *TimeZoneUpdate) SetUpdatedAt(t time.Time) *TimeZoneUpdate {
	tzu.mutation.SetUpdatedAt(t)
	return tzu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tzu *TimeZoneUpdate) SetNillableUpdatedAt(t *time.Time) *TimeZoneUpdate {
	if t != nil {
		tzu.SetUpdatedAt(*t)
	}
	return tzu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tzu *TimeZoneUpdate) ClearUpdatedAt() *TimeZoneUpdate {
	tzu.mutation.ClearUpdatedAt()
	return tzu
}

// Mutation returns the TimeZoneMutation object of the builder.
func (tzu *TimeZoneUpdate) Mutation() *TimeZoneMutation {
	return tzu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tzu *TimeZoneUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tzu.sqlSave, tzu.mutation, tzu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tzu *TimeZoneUpdate) SaveX(ctx context.Context) int {
	affected, err := tzu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tzu *TimeZoneUpdate) Exec(ctx context.Context) error {
	_, err := tzu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tzu *TimeZoneUpdate) ExecX(ctx context.Context) {
	if err := tzu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tzu *TimeZoneUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(timezone.Table, timezone.Columns, sqlgraph.NewFieldSpec(timezone.FieldID, field.TypeUUID))
	if ps := tzu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tzu.mutation.UserID(); ok {
		_spec.SetField(timezone.FieldUserID, field.TypeUUID, value)
	}
	if value, ok := tzu.mutation.TimezoneName(); ok {
		_spec.SetField(timezone.FieldTimezoneName, field.TypeString, value)
	}
	if tzu.mutation.TimezoneNameCleared() {
		_spec.ClearField(timezone.FieldTimezoneName, field.TypeString)
	}
	if value, ok := tzu.mutation.IsDefault(); ok {
		_spec.SetField(timezone.FieldIsDefault, field.TypeBool, value)
	}
	if value, ok := tzu.mutation.Midday(); ok {
		_spec.SetField(timezone.FieldMidday, field.TypeString, value)
	}
	if value, ok := tzu.mutation.Hour(); ok {
		_spec.SetField(timezone.FieldHour, field.TypeString, value)
	}
	if value, ok := tzu.mutation.Minute(); ok {
		_spec.SetField(timezone.FieldMinute, field.TypeString, value)
	}
	if value, ok := tzu.mutation.CreatedAt(); ok {
		_spec.SetField(timezone.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tzu.mutation.UpdatedAt(); ok {
		_spec.SetField(timezone.FieldUpdatedAt, field.TypeTime, value)
	}
	if tzu.mutation.UpdatedAtCleared() {
		_spec.ClearField(timezone.FieldUpdatedAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tzu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{timezone.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tzu.mutation.done = true
	return n, nil
}

// TimeZoneUpdateOne is the builder for updating a single TimeZone entity.
type TimeZoneUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TimeZoneMutation
}

// SetUserID sets the "user_id" field.
func (tzuo *TimeZoneUpdateOne) SetUserID(u uuid.UUID) *TimeZoneUpdateOne {
	tzuo.mutation.SetUserID(u)
	return tzuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tzuo *TimeZoneUpdateOne) SetNillableUserID(u *uuid.UUID) *TimeZoneUpdateOne {
	if u != nil {
		tzuo.SetUserID(*u)
	}
	return tzuo
}

// SetTimezoneName sets the "timezone_name" field.
func (tzuo *TimeZoneUpdateOne) SetTimezoneName(s string) *TimeZoneUpdateOne {
	tzuo.mutation.SetTimezoneName(s)
	return tzuo
}

// SetNillableTimezoneName sets the "timezone_name" field if the given value is not nil.
func (tzuo *TimeZoneUpdateOne) SetNillableTimezoneName(s *string) *TimeZoneUpdateOne {
	if s != nil {
		tzuo.SetTimezoneName(*s)
	}
	return tzuo
}

// ClearTimezoneName clears the value of the "timezone_name" field.
func (tzuo *TimeZoneUpdateOne) ClearTimezoneName() *TimeZoneUpdateOne {
	tzuo.mutation.ClearTimezoneName()
	return tzuo
}

// SetIsDefault sets the "is_default" field.
func (tzuo *TimeZoneUpdateOne) SetIsDefault(b bool) *TimeZoneUpdateOne {
	tzuo.mutation.SetIsDefault(b)
	return tzuo
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (tzuo *TimeZoneUpdateOne) SetNillableIsDefault(b *bool) *TimeZoneUpdateOne {
	if b != nil {
		tzuo.SetIsDefault(*b)
	}
	return tzuo
}

// SetMidday sets the "midday" field.
func (tzuo *TimeZoneUpdateOne) SetMidday(s string) *TimeZoneUpdateOne {
	tzuo.mutation.SetMidday(s)
	return tzuo
}

// SetNillableMidday sets the "midday" field if the given value is not nil.
func (tzuo *TimeZoneUpdateOne) SetNillableMidday(s *string) *TimeZoneUpdateOne {
	if s != nil {
		tzuo.SetMidday(*s)
	}
	return tzuo
}

// SetHour sets the "hour" field.
func (tzuo *TimeZoneUpdateOne) SetHour(s string) *TimeZoneUpdateOne {
	tzuo.mutation.SetHour(s)
	return tzuo
}

// SetNillableHour sets the "hour" field if the given value is not nil.
func (tzuo *TimeZoneUpdateOne) SetNillableHour(s *string) *TimeZoneUpdateOne {
	if s != nil {
		tzuo.SetHour(*s)
	}
	return tzuo
}

// SetMinute sets the "minute" field.
func (tzuo *TimeZoneUpdateOne) SetMinute(s string) *TimeZoneUpdateOne {
	tzuo.mutation.SetMinute(s)
	return tzuo
}

// SetNillableMinute sets the "minute" field if the given value is not nil.
func (tzuo *TimeZoneUpdateOne) SetNillableMinute(s *string) *TimeZoneUpdateOne {
	if s != nil {
		tzuo.SetMinute(*s)
	}
	return tzuo
}

// SetCreatedAt sets the "created_at" field.
func (tzuo *TimeZoneUpdateOne) SetCreatedAt(t time.Time) *TimeZoneUpdateOne {
	tzuo.mutation.SetCreatedAt(t)
	return tzuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tzuo *TimeZoneUpdateOne) SetNillableCreatedAt(t *time.Time) *TimeZoneUpdateOne {
	if t != nil {
		tzuo.SetCreatedAt(*t)
	}
	return tzuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tzuo *TimeZoneUpdateOne) SetUpdatedAt(t time.Time) *TimeZoneUpdateOne {
	tzuo.mutation.SetUpdatedAt(t)
	return tzuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tzuo *TimeZoneUpdateOne) SetNillableUpdatedAt(t *time.Time) *TimeZoneUpdateOne {
	if t != nil {
		tzuo.SetUpdatedAt(*t)
	}
	return tzuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tzuo *TimeZoneUpdateOne) ClearUpdatedAt() *TimeZoneUpdateOne {
	tzuo.mutation.ClearUpdatedAt()
	return tzuo
}

// Mutation returns the TimeZoneMutation object of the builder.
func (tzuo *TimeZoneUpdateOne) Mutation() *TimeZoneMutation {
	return tzuo.mutation
}

// Where appends a list predicates to the TimeZoneUpdate builder.
func (tzuo *TimeZoneUpdateOne) Where(ps ...predicate.TimeZone) *TimeZoneUpdateOne {
	tzuo.mutation.Where(ps...)
	return tzuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tzuo *TimeZoneUpdateOne) Select(field string, fields ...string) *TimeZoneUpdateOne {
	tzuo.fields = append([]string{field}, fields...)
	return tzuo
}

// Save executes the query and returns the updated TimeZone entity.
func (tzuo *TimeZoneUpdateOne) Save(ctx context.Context) (*TimeZone, error) {
	return withHooks(ctx, tzuo.sqlSave, tzuo.mutation, tzuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tzuo *TimeZoneUpdateOne) SaveX(ctx context.Context) *TimeZone {
	node, err := tzuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tzuo *TimeZoneUpdateOne) Exec(ctx context.Context) error {
	_, err := tzuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tzuo *TimeZoneUpdateOne) ExecX(ctx context.Context) {
	if err := tzuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tzuo *TimeZoneUpdateOne) sqlSave(ctx context.Context) (_node *TimeZone, err error) {
	_spec := sqlgraph.NewUpdateSpec(timezone.Table, timezone.Columns, sqlgraph.NewFieldSpec(timezone.FieldID, field.TypeUUID))
	id, ok := tzuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TimeZone.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tzuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, timezone.FieldID)
		for _, f := range fields {
			if !timezone.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != timezone.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tzuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tzuo.mutation.UserID(); ok {
		_spec.SetField(timezone.FieldUserID, field.TypeUUID, value)
	}
	if value, ok := tzuo.mutation.TimezoneName(); ok {
		_spec.SetField(timezone.FieldTimezoneName, field.TypeString, value)
	}
	if tzuo.mutation.TimezoneNameCleared() {
		_spec.ClearField(timezone.FieldTimezoneName, field.TypeString)
	}
	if value, ok := tzuo.mutation.IsDefault(); ok {
		_spec.SetField(timezone.FieldIsDefault, field.TypeBool, value)
	}
	if value, ok := tzuo.mutation.Midday(); ok {
		_spec.SetField(timezone.FieldMidday, field.TypeString, value)
	}
	if value, ok := tzuo.mutation.Hour(); ok {
		_spec.SetField(timezone.FieldHour, field.TypeString, value)
	}
	if value, ok := tzuo.mutation.Minute(); ok {
		_spec.SetField(timezone.FieldMinute, field.TypeString, value)
	}
	if value, ok := tzuo.mutation.CreatedAt(); ok {
		_spec.SetField(timezone.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tzuo.mutation.UpdatedAt(); ok {
		_spec.SetField(timezone.FieldUpdatedAt, field.TypeTime, value)
	}
	if tzuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(timezone.FieldUpdatedAt, field.TypeTime)
	}
	_node = &TimeZone{config: tzuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tzuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{timezone.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tzuo.mutation.done = true
	return _node, nil
}
