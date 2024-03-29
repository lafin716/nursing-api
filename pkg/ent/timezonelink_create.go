// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"nursing_api/pkg/ent/timezonelink"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TimeZoneLinkCreate is the builder for creating a TimeZoneLink entity.
type TimeZoneLinkCreate struct {
	config
	mutation *TimeZoneLinkMutation
	hooks    []Hook
}

// SetPrescriptionID sets the "prescription_id" field.
func (tzlc *TimeZoneLinkCreate) SetPrescriptionID(u uuid.UUID) *TimeZoneLinkCreate {
	tzlc.mutation.SetPrescriptionID(u)
	return tzlc
}

// SetTimezoneID sets the "timezone_id" field.
func (tzlc *TimeZoneLinkCreate) SetTimezoneID(u uuid.UUID) *TimeZoneLinkCreate {
	tzlc.mutation.SetTimezoneID(u)
	return tzlc
}

// SetTimezoneName sets the "timezone_name" field.
func (tzlc *TimeZoneLinkCreate) SetTimezoneName(s string) *TimeZoneLinkCreate {
	tzlc.mutation.SetTimezoneName(s)
	return tzlc
}

// SetNillableTimezoneName sets the "timezone_name" field if the given value is not nil.
func (tzlc *TimeZoneLinkCreate) SetNillableTimezoneName(s *string) *TimeZoneLinkCreate {
	if s != nil {
		tzlc.SetTimezoneName(*s)
	}
	return tzlc
}

// SetUseAlert sets the "use_alert" field.
func (tzlc *TimeZoneLinkCreate) SetUseAlert(b bool) *TimeZoneLinkCreate {
	tzlc.mutation.SetUseAlert(b)
	return tzlc
}

// SetNillableUseAlert sets the "use_alert" field if the given value is not nil.
func (tzlc *TimeZoneLinkCreate) SetNillableUseAlert(b *bool) *TimeZoneLinkCreate {
	if b != nil {
		tzlc.SetUseAlert(*b)
	}
	return tzlc
}

// SetMidday sets the "midday" field.
func (tzlc *TimeZoneLinkCreate) SetMidday(s string) *TimeZoneLinkCreate {
	tzlc.mutation.SetMidday(s)
	return tzlc
}

// SetHour sets the "hour" field.
func (tzlc *TimeZoneLinkCreate) SetHour(s string) *TimeZoneLinkCreate {
	tzlc.mutation.SetHour(s)
	return tzlc
}

// SetMinute sets the "minute" field.
func (tzlc *TimeZoneLinkCreate) SetMinute(s string) *TimeZoneLinkCreate {
	tzlc.mutation.SetMinute(s)
	return tzlc
}

// SetCreatedAt sets the "created_at" field.
func (tzlc *TimeZoneLinkCreate) SetCreatedAt(t time.Time) *TimeZoneLinkCreate {
	tzlc.mutation.SetCreatedAt(t)
	return tzlc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tzlc *TimeZoneLinkCreate) SetNillableCreatedAt(t *time.Time) *TimeZoneLinkCreate {
	if t != nil {
		tzlc.SetCreatedAt(*t)
	}
	return tzlc
}

// SetUpdatedAt sets the "updated_at" field.
func (tzlc *TimeZoneLinkCreate) SetUpdatedAt(t time.Time) *TimeZoneLinkCreate {
	tzlc.mutation.SetUpdatedAt(t)
	return tzlc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tzlc *TimeZoneLinkCreate) SetNillableUpdatedAt(t *time.Time) *TimeZoneLinkCreate {
	if t != nil {
		tzlc.SetUpdatedAt(*t)
	}
	return tzlc
}

// SetID sets the "id" field.
func (tzlc *TimeZoneLinkCreate) SetID(u uuid.UUID) *TimeZoneLinkCreate {
	tzlc.mutation.SetID(u)
	return tzlc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tzlc *TimeZoneLinkCreate) SetNillableID(u *uuid.UUID) *TimeZoneLinkCreate {
	if u != nil {
		tzlc.SetID(*u)
	}
	return tzlc
}

// Mutation returns the TimeZoneLinkMutation object of the builder.
func (tzlc *TimeZoneLinkCreate) Mutation() *TimeZoneLinkMutation {
	return tzlc.mutation
}

// Save creates the TimeZoneLink in the database.
func (tzlc *TimeZoneLinkCreate) Save(ctx context.Context) (*TimeZoneLink, error) {
	tzlc.defaults()
	return withHooks(ctx, tzlc.sqlSave, tzlc.mutation, tzlc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tzlc *TimeZoneLinkCreate) SaveX(ctx context.Context) *TimeZoneLink {
	v, err := tzlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tzlc *TimeZoneLinkCreate) Exec(ctx context.Context) error {
	_, err := tzlc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tzlc *TimeZoneLinkCreate) ExecX(ctx context.Context) {
	if err := tzlc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tzlc *TimeZoneLinkCreate) defaults() {
	if _, ok := tzlc.mutation.UseAlert(); !ok {
		v := timezonelink.DefaultUseAlert
		tzlc.mutation.SetUseAlert(v)
	}
	if _, ok := tzlc.mutation.CreatedAt(); !ok {
		v := timezonelink.DefaultCreatedAt()
		tzlc.mutation.SetCreatedAt(v)
	}
	if _, ok := tzlc.mutation.ID(); !ok {
		v := timezonelink.DefaultID()
		tzlc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tzlc *TimeZoneLinkCreate) check() error {
	if _, ok := tzlc.mutation.PrescriptionID(); !ok {
		return &ValidationError{Name: "prescription_id", err: errors.New(`ent: missing required field "TimeZoneLink.prescription_id"`)}
	}
	if _, ok := tzlc.mutation.TimezoneID(); !ok {
		return &ValidationError{Name: "timezone_id", err: errors.New(`ent: missing required field "TimeZoneLink.timezone_id"`)}
	}
	if _, ok := tzlc.mutation.UseAlert(); !ok {
		return &ValidationError{Name: "use_alert", err: errors.New(`ent: missing required field "TimeZoneLink.use_alert"`)}
	}
	if _, ok := tzlc.mutation.Midday(); !ok {
		return &ValidationError{Name: "midday", err: errors.New(`ent: missing required field "TimeZoneLink.midday"`)}
	}
	if _, ok := tzlc.mutation.Hour(); !ok {
		return &ValidationError{Name: "hour", err: errors.New(`ent: missing required field "TimeZoneLink.hour"`)}
	}
	if _, ok := tzlc.mutation.Minute(); !ok {
		return &ValidationError{Name: "minute", err: errors.New(`ent: missing required field "TimeZoneLink.minute"`)}
	}
	if _, ok := tzlc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TimeZoneLink.created_at"`)}
	}
	return nil
}

func (tzlc *TimeZoneLinkCreate) sqlSave(ctx context.Context) (*TimeZoneLink, error) {
	if err := tzlc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tzlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tzlc.driver, _spec); err != nil {
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
	tzlc.mutation.id = &_node.ID
	tzlc.mutation.done = true
	return _node, nil
}

func (tzlc *TimeZoneLinkCreate) createSpec() (*TimeZoneLink, *sqlgraph.CreateSpec) {
	var (
		_node = &TimeZoneLink{config: tzlc.config}
		_spec = sqlgraph.NewCreateSpec(timezonelink.Table, sqlgraph.NewFieldSpec(timezonelink.FieldID, field.TypeUUID))
	)
	if id, ok := tzlc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tzlc.mutation.PrescriptionID(); ok {
		_spec.SetField(timezonelink.FieldPrescriptionID, field.TypeUUID, value)
		_node.PrescriptionID = value
	}
	if value, ok := tzlc.mutation.TimezoneID(); ok {
		_spec.SetField(timezonelink.FieldTimezoneID, field.TypeUUID, value)
		_node.TimezoneID = value
	}
	if value, ok := tzlc.mutation.TimezoneName(); ok {
		_spec.SetField(timezonelink.FieldTimezoneName, field.TypeString, value)
		_node.TimezoneName = value
	}
	if value, ok := tzlc.mutation.UseAlert(); ok {
		_spec.SetField(timezonelink.FieldUseAlert, field.TypeBool, value)
		_node.UseAlert = value
	}
	if value, ok := tzlc.mutation.Midday(); ok {
		_spec.SetField(timezonelink.FieldMidday, field.TypeString, value)
		_node.Midday = value
	}
	if value, ok := tzlc.mutation.Hour(); ok {
		_spec.SetField(timezonelink.FieldHour, field.TypeString, value)
		_node.Hour = value
	}
	if value, ok := tzlc.mutation.Minute(); ok {
		_spec.SetField(timezonelink.FieldMinute, field.TypeString, value)
		_node.Minute = value
	}
	if value, ok := tzlc.mutation.CreatedAt(); ok {
		_spec.SetField(timezonelink.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tzlc.mutation.UpdatedAt(); ok {
		_spec.SetField(timezonelink.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// TimeZoneLinkCreateBulk is the builder for creating many TimeZoneLink entities in bulk.
type TimeZoneLinkCreateBulk struct {
	config
	err      error
	builders []*TimeZoneLinkCreate
}

// Save creates the TimeZoneLink entities in the database.
func (tzlcb *TimeZoneLinkCreateBulk) Save(ctx context.Context) ([]*TimeZoneLink, error) {
	if tzlcb.err != nil {
		return nil, tzlcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tzlcb.builders))
	nodes := make([]*TimeZoneLink, len(tzlcb.builders))
	mutators := make([]Mutator, len(tzlcb.builders))
	for i := range tzlcb.builders {
		func(i int, root context.Context) {
			builder := tzlcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TimeZoneLinkMutation)
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
					_, err = mutators[i+1].Mutate(root, tzlcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tzlcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tzlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tzlcb *TimeZoneLinkCreateBulk) SaveX(ctx context.Context) []*TimeZoneLink {
	v, err := tzlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tzlcb *TimeZoneLinkCreateBulk) Exec(ctx context.Context) error {
	_, err := tzlcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tzlcb *TimeZoneLinkCreateBulk) ExecX(ctx context.Context) {
	if err := tzlcb.Exec(ctx); err != nil {
		panic(err)
	}
}
