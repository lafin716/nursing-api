// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"nursing_api/pkg/ent/timezone"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TimeZoneCreate is the builder for creating a TimeZone entity.
type TimeZoneCreate struct {
	config
	mutation *TimeZoneMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (tzc *TimeZoneCreate) SetUserID(u uuid.UUID) *TimeZoneCreate {
	tzc.mutation.SetUserID(u)
	return tzc
}

// SetTimezoneName sets the "timezone_name" field.
func (tzc *TimeZoneCreate) SetTimezoneName(s string) *TimeZoneCreate {
	tzc.mutation.SetTimezoneName(s)
	return tzc
}

// SetNillableTimezoneName sets the "timezone_name" field if the given value is not nil.
func (tzc *TimeZoneCreate) SetNillableTimezoneName(s *string) *TimeZoneCreate {
	if s != nil {
		tzc.SetTimezoneName(*s)
	}
	return tzc
}

// SetIsDefault sets the "is_default" field.
func (tzc *TimeZoneCreate) SetIsDefault(b bool) *TimeZoneCreate {
	tzc.mutation.SetIsDefault(b)
	return tzc
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (tzc *TimeZoneCreate) SetNillableIsDefault(b *bool) *TimeZoneCreate {
	if b != nil {
		tzc.SetIsDefault(*b)
	}
	return tzc
}

// SetMidday sets the "midday" field.
func (tzc *TimeZoneCreate) SetMidday(s string) *TimeZoneCreate {
	tzc.mutation.SetMidday(s)
	return tzc
}

// SetHour sets the "hour" field.
func (tzc *TimeZoneCreate) SetHour(s string) *TimeZoneCreate {
	tzc.mutation.SetHour(s)
	return tzc
}

// SetMinute sets the "minute" field.
func (tzc *TimeZoneCreate) SetMinute(s string) *TimeZoneCreate {
	tzc.mutation.SetMinute(s)
	return tzc
}

// SetCreatedAt sets the "created_at" field.
func (tzc *TimeZoneCreate) SetCreatedAt(t time.Time) *TimeZoneCreate {
	tzc.mutation.SetCreatedAt(t)
	return tzc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tzc *TimeZoneCreate) SetNillableCreatedAt(t *time.Time) *TimeZoneCreate {
	if t != nil {
		tzc.SetCreatedAt(*t)
	}
	return tzc
}

// SetUpdatedAt sets the "updated_at" field.
func (tzc *TimeZoneCreate) SetUpdatedAt(t time.Time) *TimeZoneCreate {
	tzc.mutation.SetUpdatedAt(t)
	return tzc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tzc *TimeZoneCreate) SetNillableUpdatedAt(t *time.Time) *TimeZoneCreate {
	if t != nil {
		tzc.SetUpdatedAt(*t)
	}
	return tzc
}

// SetID sets the "id" field.
func (tzc *TimeZoneCreate) SetID(u uuid.UUID) *TimeZoneCreate {
	tzc.mutation.SetID(u)
	return tzc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tzc *TimeZoneCreate) SetNillableID(u *uuid.UUID) *TimeZoneCreate {
	if u != nil {
		tzc.SetID(*u)
	}
	return tzc
}

// Mutation returns the TimeZoneMutation object of the builder.
func (tzc *TimeZoneCreate) Mutation() *TimeZoneMutation {
	return tzc.mutation
}

// Save creates the TimeZone in the database.
func (tzc *TimeZoneCreate) Save(ctx context.Context) (*TimeZone, error) {
	tzc.defaults()
	return withHooks(ctx, tzc.sqlSave, tzc.mutation, tzc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tzc *TimeZoneCreate) SaveX(ctx context.Context) *TimeZone {
	v, err := tzc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tzc *TimeZoneCreate) Exec(ctx context.Context) error {
	_, err := tzc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tzc *TimeZoneCreate) ExecX(ctx context.Context) {
	if err := tzc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tzc *TimeZoneCreate) defaults() {
	if _, ok := tzc.mutation.IsDefault(); !ok {
		v := timezone.DefaultIsDefault
		tzc.mutation.SetIsDefault(v)
	}
	if _, ok := tzc.mutation.CreatedAt(); !ok {
		v := timezone.DefaultCreatedAt()
		tzc.mutation.SetCreatedAt(v)
	}
	if _, ok := tzc.mutation.ID(); !ok {
		v := timezone.DefaultID()
		tzc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tzc *TimeZoneCreate) check() error {
	if _, ok := tzc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "TimeZone.user_id"`)}
	}
	if _, ok := tzc.mutation.IsDefault(); !ok {
		return &ValidationError{Name: "is_default", err: errors.New(`ent: missing required field "TimeZone.is_default"`)}
	}
	if _, ok := tzc.mutation.Midday(); !ok {
		return &ValidationError{Name: "midday", err: errors.New(`ent: missing required field "TimeZone.midday"`)}
	}
	if _, ok := tzc.mutation.Hour(); !ok {
		return &ValidationError{Name: "hour", err: errors.New(`ent: missing required field "TimeZone.hour"`)}
	}
	if _, ok := tzc.mutation.Minute(); !ok {
		return &ValidationError{Name: "minute", err: errors.New(`ent: missing required field "TimeZone.minute"`)}
	}
	if _, ok := tzc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TimeZone.created_at"`)}
	}
	return nil
}

func (tzc *TimeZoneCreate) sqlSave(ctx context.Context) (*TimeZone, error) {
	if err := tzc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tzc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tzc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	tzc.mutation.id = &_node.ID
	tzc.mutation.done = true
	return _node, nil
}

func (tzc *TimeZoneCreate) createSpec() (*TimeZone, *sqlgraph.CreateSpec) {
	var (
		_node = &TimeZone{config: tzc.config}
		_spec = sqlgraph.NewCreateSpec(timezone.Table, sqlgraph.NewFieldSpec(timezone.FieldID, field.TypeUUID))
	)
	if id, ok := tzc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tzc.mutation.UserID(); ok {
		_spec.SetField(timezone.FieldUserID, field.TypeUUID, value)
		_node.UserID = value
	}
	if value, ok := tzc.mutation.TimezoneName(); ok {
		_spec.SetField(timezone.FieldTimezoneName, field.TypeString, value)
		_node.TimezoneName = value
	}
	if value, ok := tzc.mutation.IsDefault(); ok {
		_spec.SetField(timezone.FieldIsDefault, field.TypeBool, value)
		_node.IsDefault = value
	}
	if value, ok := tzc.mutation.Midday(); ok {
		_spec.SetField(timezone.FieldMidday, field.TypeString, value)
		_node.Midday = value
	}
	if value, ok := tzc.mutation.Hour(); ok {
		_spec.SetField(timezone.FieldHour, field.TypeString, value)
		_node.Hour = value
	}
	if value, ok := tzc.mutation.Minute(); ok {
		_spec.SetField(timezone.FieldMinute, field.TypeString, value)
		_node.Minute = value
	}
	if value, ok := tzc.mutation.CreatedAt(); ok {
		_spec.SetField(timezone.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tzc.mutation.UpdatedAt(); ok {
		_spec.SetField(timezone.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// TimeZoneCreateBulk is the builder for creating many TimeZone entities in bulk.
type TimeZoneCreateBulk struct {
	config
	err      error
	builders []*TimeZoneCreate
}

// Save creates the TimeZone entities in the database.
func (tzcb *TimeZoneCreateBulk) Save(ctx context.Context) ([]*TimeZone, error) {
	if tzcb.err != nil {
		return nil, tzcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tzcb.builders))
	nodes := make([]*TimeZone, len(tzcb.builders))
	mutators := make([]Mutator, len(tzcb.builders))
	for i := range tzcb.builders {
		func(i int, root context.Context) {
			builder := tzcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TimeZoneMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tzcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tzcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tzcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tzcb *TimeZoneCreateBulk) SaveX(ctx context.Context) []*TimeZone {
	v, err := tzcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tzcb *TimeZoneCreateBulk) Exec(ctx context.Context) error {
	_, err := tzcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tzcb *TimeZoneCreateBulk) ExecX(ctx context.Context) {
	if err := tzcb.Exec(ctx); err != nil {
		panic(err)
	}
}
