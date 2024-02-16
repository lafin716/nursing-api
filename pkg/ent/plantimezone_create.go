// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"nursing_api/pkg/ent/plantimezone"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PlanTimeZoneCreate is the builder for creating a PlanTimeZone entity.
type PlanTimeZoneCreate struct {
	config
	mutation *PlanTimeZoneMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (ptzc *PlanTimeZoneCreate) SetUserID(u uuid.UUID) *PlanTimeZoneCreate {
	ptzc.mutation.SetUserID(u)
	return ptzc
}

// SetTimezoneName sets the "timezone_name" field.
func (ptzc *PlanTimeZoneCreate) SetTimezoneName(s string) *PlanTimeZoneCreate {
	ptzc.mutation.SetTimezoneName(s)
	return ptzc
}

// SetNillableTimezoneName sets the "timezone_name" field if the given value is not nil.
func (ptzc *PlanTimeZoneCreate) SetNillableTimezoneName(s *string) *PlanTimeZoneCreate {
	if s != nil {
		ptzc.SetTimezoneName(*s)
	}
	return ptzc
}

// SetIsDefault sets the "is_default" field.
func (ptzc *PlanTimeZoneCreate) SetIsDefault(b bool) *PlanTimeZoneCreate {
	ptzc.mutation.SetIsDefault(b)
	return ptzc
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (ptzc *PlanTimeZoneCreate) SetNillableIsDefault(b *bool) *PlanTimeZoneCreate {
	if b != nil {
		ptzc.SetIsDefault(*b)
	}
	return ptzc
}

// SetMidday sets the "midday" field.
func (ptzc *PlanTimeZoneCreate) SetMidday(s string) *PlanTimeZoneCreate {
	ptzc.mutation.SetMidday(s)
	return ptzc
}

// SetHour sets the "hour" field.
func (ptzc *PlanTimeZoneCreate) SetHour(s string) *PlanTimeZoneCreate {
	ptzc.mutation.SetHour(s)
	return ptzc
}

// SetMinute sets the "minute" field.
func (ptzc *PlanTimeZoneCreate) SetMinute(s string) *PlanTimeZoneCreate {
	ptzc.mutation.SetMinute(s)
	return ptzc
}

// SetCreatedAt sets the "created_at" field.
func (ptzc *PlanTimeZoneCreate) SetCreatedAt(t time.Time) *PlanTimeZoneCreate {
	ptzc.mutation.SetCreatedAt(t)
	return ptzc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ptzc *PlanTimeZoneCreate) SetNillableCreatedAt(t *time.Time) *PlanTimeZoneCreate {
	if t != nil {
		ptzc.SetCreatedAt(*t)
	}
	return ptzc
}

// SetUpdatedAt sets the "updated_at" field.
func (ptzc *PlanTimeZoneCreate) SetUpdatedAt(t time.Time) *PlanTimeZoneCreate {
	ptzc.mutation.SetUpdatedAt(t)
	return ptzc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ptzc *PlanTimeZoneCreate) SetNillableUpdatedAt(t *time.Time) *PlanTimeZoneCreate {
	if t != nil {
		ptzc.SetUpdatedAt(*t)
	}
	return ptzc
}

// SetID sets the "id" field.
func (ptzc *PlanTimeZoneCreate) SetID(u uuid.UUID) *PlanTimeZoneCreate {
	ptzc.mutation.SetID(u)
	return ptzc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ptzc *PlanTimeZoneCreate) SetNillableID(u *uuid.UUID) *PlanTimeZoneCreate {
	if u != nil {
		ptzc.SetID(*u)
	}
	return ptzc
}

// Mutation returns the PlanTimeZoneMutation object of the builder.
func (ptzc *PlanTimeZoneCreate) Mutation() *PlanTimeZoneMutation {
	return ptzc.mutation
}

// Save creates the PlanTimeZone in the database.
func (ptzc *PlanTimeZoneCreate) Save(ctx context.Context) (*PlanTimeZone, error) {
	ptzc.defaults()
	return withHooks(ctx, ptzc.sqlSave, ptzc.mutation, ptzc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ptzc *PlanTimeZoneCreate) SaveX(ctx context.Context) *PlanTimeZone {
	v, err := ptzc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptzc *PlanTimeZoneCreate) Exec(ctx context.Context) error {
	_, err := ptzc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptzc *PlanTimeZoneCreate) ExecX(ctx context.Context) {
	if err := ptzc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ptzc *PlanTimeZoneCreate) defaults() {
	if _, ok := ptzc.mutation.IsDefault(); !ok {
		v := plantimezone.DefaultIsDefault
		ptzc.mutation.SetIsDefault(v)
	}
	if _, ok := ptzc.mutation.CreatedAt(); !ok {
		v := plantimezone.DefaultCreatedAt()
		ptzc.mutation.SetCreatedAt(v)
	}
	if _, ok := ptzc.mutation.ID(); !ok {
		v := plantimezone.DefaultID()
		ptzc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptzc *PlanTimeZoneCreate) check() error {
	if _, ok := ptzc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "PlanTimeZone.user_id"`)}
	}
	if _, ok := ptzc.mutation.IsDefault(); !ok {
		return &ValidationError{Name: "is_default", err: errors.New(`ent: missing required field "PlanTimeZone.is_default"`)}
	}
	if _, ok := ptzc.mutation.Midday(); !ok {
		return &ValidationError{Name: "midday", err: errors.New(`ent: missing required field "PlanTimeZone.midday"`)}
	}
	if _, ok := ptzc.mutation.Hour(); !ok {
		return &ValidationError{Name: "hour", err: errors.New(`ent: missing required field "PlanTimeZone.hour"`)}
	}
	if _, ok := ptzc.mutation.Minute(); !ok {
		return &ValidationError{Name: "minute", err: errors.New(`ent: missing required field "PlanTimeZone.minute"`)}
	}
	if _, ok := ptzc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "PlanTimeZone.created_at"`)}
	}
	return nil
}

func (ptzc *PlanTimeZoneCreate) sqlSave(ctx context.Context) (*PlanTimeZone, error) {
	if err := ptzc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ptzc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ptzc.driver, _spec); err != nil {
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
	ptzc.mutation.id = &_node.ID
	ptzc.mutation.done = true
	return _node, nil
}

func (ptzc *PlanTimeZoneCreate) createSpec() (*PlanTimeZone, *sqlgraph.CreateSpec) {
	var (
		_node = &PlanTimeZone{config: ptzc.config}
		_spec = sqlgraph.NewCreateSpec(plantimezone.Table, sqlgraph.NewFieldSpec(plantimezone.FieldID, field.TypeUUID))
	)
	if id, ok := ptzc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ptzc.mutation.UserID(); ok {
		_spec.SetField(plantimezone.FieldUserID, field.TypeUUID, value)
		_node.UserID = value
	}
	if value, ok := ptzc.mutation.TimezoneName(); ok {
		_spec.SetField(plantimezone.FieldTimezoneName, field.TypeString, value)
		_node.TimezoneName = value
	}
	if value, ok := ptzc.mutation.IsDefault(); ok {
		_spec.SetField(plantimezone.FieldIsDefault, field.TypeBool, value)
		_node.IsDefault = value
	}
	if value, ok := ptzc.mutation.Midday(); ok {
		_spec.SetField(plantimezone.FieldMidday, field.TypeString, value)
		_node.Midday = value
	}
	if value, ok := ptzc.mutation.Hour(); ok {
		_spec.SetField(plantimezone.FieldHour, field.TypeString, value)
		_node.Hour = value
	}
	if value, ok := ptzc.mutation.Minute(); ok {
		_spec.SetField(plantimezone.FieldMinute, field.TypeString, value)
		_node.Minute = value
	}
	if value, ok := ptzc.mutation.CreatedAt(); ok {
		_spec.SetField(plantimezone.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ptzc.mutation.UpdatedAt(); ok {
		_spec.SetField(plantimezone.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// PlanTimeZoneCreateBulk is the builder for creating many PlanTimeZone entities in bulk.
type PlanTimeZoneCreateBulk struct {
	config
	err      error
	builders []*PlanTimeZoneCreate
}

// Save creates the PlanTimeZone entities in the database.
func (ptzcb *PlanTimeZoneCreateBulk) Save(ctx context.Context) ([]*PlanTimeZone, error) {
	if ptzcb.err != nil {
		return nil, ptzcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ptzcb.builders))
	nodes := make([]*PlanTimeZone, len(ptzcb.builders))
	mutators := make([]Mutator, len(ptzcb.builders))
	for i := range ptzcb.builders {
		func(i int, root context.Context) {
			builder := ptzcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlanTimeZoneMutation)
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
					_, err = mutators[i+1].Mutate(root, ptzcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ptzcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ptzcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ptzcb *PlanTimeZoneCreateBulk) SaveX(ctx context.Context) []*PlanTimeZone {
	v, err := ptzcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptzcb *PlanTimeZoneCreateBulk) Exec(ctx context.Context) error {
	_, err := ptzcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptzcb *PlanTimeZoneCreateBulk) ExecX(ctx context.Context) {
	if err := ptzcb.Exec(ctx); err != nil {
		panic(err)
	}
}
