// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"nursing_api/pkg/ent/takehistoryitem"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TakeHistoryItemCreate is the builder for creating a TakeHistoryItem entity.
type TakeHistoryItemCreate struct {
	config
	mutation *TakeHistoryItemMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (thic *TakeHistoryItemCreate) SetUserID(u uuid.UUID) *TakeHistoryItemCreate {
	thic.mutation.SetUserID(u)
	return thic
}

// SetTakeHistoryID sets the "take_history_id" field.
func (thic *TakeHistoryItemCreate) SetTakeHistoryID(u uuid.UUID) *TakeHistoryItemCreate {
	thic.mutation.SetTakeHistoryID(u)
	return thic
}

// SetPrescriptionItemID sets the "prescription_item_id" field.
func (thic *TakeHistoryItemCreate) SetPrescriptionItemID(u uuid.UUID) *TakeHistoryItemCreate {
	thic.mutation.SetPrescriptionItemID(u)
	return thic
}

// SetTakeStatus sets the "take_status" field.
func (thic *TakeHistoryItemCreate) SetTakeStatus(s string) *TakeHistoryItemCreate {
	thic.mutation.SetTakeStatus(s)
	return thic
}

// SetNillableTakeStatus sets the "take_status" field if the given value is not nil.
func (thic *TakeHistoryItemCreate) SetNillableTakeStatus(s *string) *TakeHistoryItemCreate {
	if s != nil {
		thic.SetTakeStatus(*s)
	}
	return thic
}

// SetTakeAmount sets the "take_amount" field.
func (thic *TakeHistoryItemCreate) SetTakeAmount(f float64) *TakeHistoryItemCreate {
	thic.mutation.SetTakeAmount(f)
	return thic
}

// SetNillableTakeAmount sets the "take_amount" field if the given value is not nil.
func (thic *TakeHistoryItemCreate) SetNillableTakeAmount(f *float64) *TakeHistoryItemCreate {
	if f != nil {
		thic.SetTakeAmount(*f)
	}
	return thic
}

// SetTakeUnit sets the "take_unit" field.
func (thic *TakeHistoryItemCreate) SetTakeUnit(s string) *TakeHistoryItemCreate {
	thic.mutation.SetTakeUnit(s)
	return thic
}

// SetMemo sets the "memo" field.
func (thic *TakeHistoryItemCreate) SetMemo(s string) *TakeHistoryItemCreate {
	thic.mutation.SetMemo(s)
	return thic
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (thic *TakeHistoryItemCreate) SetNillableMemo(s *string) *TakeHistoryItemCreate {
	if s != nil {
		thic.SetMemo(*s)
	}
	return thic
}

// SetTakeDate sets the "take_date" field.
func (thic *TakeHistoryItemCreate) SetTakeDate(t time.Time) *TakeHistoryItemCreate {
	thic.mutation.SetTakeDate(t)
	return thic
}

// SetCreatedAt sets the "created_at" field.
func (thic *TakeHistoryItemCreate) SetCreatedAt(t time.Time) *TakeHistoryItemCreate {
	thic.mutation.SetCreatedAt(t)
	return thic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (thic *TakeHistoryItemCreate) SetNillableCreatedAt(t *time.Time) *TakeHistoryItemCreate {
	if t != nil {
		thic.SetCreatedAt(*t)
	}
	return thic
}

// SetUpdatedAt sets the "updated_at" field.
func (thic *TakeHistoryItemCreate) SetUpdatedAt(t time.Time) *TakeHistoryItemCreate {
	thic.mutation.SetUpdatedAt(t)
	return thic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (thic *TakeHistoryItemCreate) SetNillableUpdatedAt(t *time.Time) *TakeHistoryItemCreate {
	if t != nil {
		thic.SetUpdatedAt(*t)
	}
	return thic
}

// SetID sets the "id" field.
func (thic *TakeHistoryItemCreate) SetID(u uuid.UUID) *TakeHistoryItemCreate {
	thic.mutation.SetID(u)
	return thic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (thic *TakeHistoryItemCreate) SetNillableID(u *uuid.UUID) *TakeHistoryItemCreate {
	if u != nil {
		thic.SetID(*u)
	}
	return thic
}

// Mutation returns the TakeHistoryItemMutation object of the builder.
func (thic *TakeHistoryItemCreate) Mutation() *TakeHistoryItemMutation {
	return thic.mutation
}

// Save creates the TakeHistoryItem in the database.
func (thic *TakeHistoryItemCreate) Save(ctx context.Context) (*TakeHistoryItem, error) {
	thic.defaults()
	return withHooks(ctx, thic.sqlSave, thic.mutation, thic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (thic *TakeHistoryItemCreate) SaveX(ctx context.Context) *TakeHistoryItem {
	v, err := thic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (thic *TakeHistoryItemCreate) Exec(ctx context.Context) error {
	_, err := thic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (thic *TakeHistoryItemCreate) ExecX(ctx context.Context) {
	if err := thic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (thic *TakeHistoryItemCreate) defaults() {
	if _, ok := thic.mutation.TakeStatus(); !ok {
		v := takehistoryitem.DefaultTakeStatus
		thic.mutation.SetTakeStatus(v)
	}
	if _, ok := thic.mutation.TakeAmount(); !ok {
		v := takehistoryitem.DefaultTakeAmount
		thic.mutation.SetTakeAmount(v)
	}
	if _, ok := thic.mutation.CreatedAt(); !ok {
		v := takehistoryitem.DefaultCreatedAt()
		thic.mutation.SetCreatedAt(v)
	}
	if _, ok := thic.mutation.ID(); !ok {
		v := takehistoryitem.DefaultID()
		thic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (thic *TakeHistoryItemCreate) check() error {
	if _, ok := thic.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "TakeHistoryItem.user_id"`)}
	}
	if _, ok := thic.mutation.TakeHistoryID(); !ok {
		return &ValidationError{Name: "take_history_id", err: errors.New(`ent: missing required field "TakeHistoryItem.take_history_id"`)}
	}
	if _, ok := thic.mutation.PrescriptionItemID(); !ok {
		return &ValidationError{Name: "prescription_item_id", err: errors.New(`ent: missing required field "TakeHistoryItem.prescription_item_id"`)}
	}
	if _, ok := thic.mutation.TakeStatus(); !ok {
		return &ValidationError{Name: "take_status", err: errors.New(`ent: missing required field "TakeHistoryItem.take_status"`)}
	}
	if _, ok := thic.mutation.TakeAmount(); !ok {
		return &ValidationError{Name: "take_amount", err: errors.New(`ent: missing required field "TakeHistoryItem.take_amount"`)}
	}
	if _, ok := thic.mutation.TakeUnit(); !ok {
		return &ValidationError{Name: "take_unit", err: errors.New(`ent: missing required field "TakeHistoryItem.take_unit"`)}
	}
	if _, ok := thic.mutation.TakeDate(); !ok {
		return &ValidationError{Name: "take_date", err: errors.New(`ent: missing required field "TakeHistoryItem.take_date"`)}
	}
	if _, ok := thic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TakeHistoryItem.created_at"`)}
	}
	return nil
}

func (thic *TakeHistoryItemCreate) sqlSave(ctx context.Context) (*TakeHistoryItem, error) {
	if err := thic.check(); err != nil {
		return nil, err
	}
	_node, _spec := thic.createSpec()
	if err := sqlgraph.CreateNode(ctx, thic.driver, _spec); err != nil {
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
	thic.mutation.id = &_node.ID
	thic.mutation.done = true
	return _node, nil
}

func (thic *TakeHistoryItemCreate) createSpec() (*TakeHistoryItem, *sqlgraph.CreateSpec) {
	var (
		_node = &TakeHistoryItem{config: thic.config}
		_spec = sqlgraph.NewCreateSpec(takehistoryitem.Table, sqlgraph.NewFieldSpec(takehistoryitem.FieldID, field.TypeUUID))
	)
	if id, ok := thic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := thic.mutation.UserID(); ok {
		_spec.SetField(takehistoryitem.FieldUserID, field.TypeUUID, value)
		_node.UserID = value
	}
	if value, ok := thic.mutation.TakeHistoryID(); ok {
		_spec.SetField(takehistoryitem.FieldTakeHistoryID, field.TypeUUID, value)
		_node.TakeHistoryID = value
	}
	if value, ok := thic.mutation.PrescriptionItemID(); ok {
		_spec.SetField(takehistoryitem.FieldPrescriptionItemID, field.TypeUUID, value)
		_node.PrescriptionItemID = value
	}
	if value, ok := thic.mutation.TakeStatus(); ok {
		_spec.SetField(takehistoryitem.FieldTakeStatus, field.TypeString, value)
		_node.TakeStatus = value
	}
	if value, ok := thic.mutation.TakeAmount(); ok {
		_spec.SetField(takehistoryitem.FieldTakeAmount, field.TypeFloat64, value)
		_node.TakeAmount = value
	}
	if value, ok := thic.mutation.TakeUnit(); ok {
		_spec.SetField(takehistoryitem.FieldTakeUnit, field.TypeString, value)
		_node.TakeUnit = value
	}
	if value, ok := thic.mutation.Memo(); ok {
		_spec.SetField(takehistoryitem.FieldMemo, field.TypeString, value)
		_node.Memo = value
	}
	if value, ok := thic.mutation.TakeDate(); ok {
		_spec.SetField(takehistoryitem.FieldTakeDate, field.TypeTime, value)
		_node.TakeDate = value
	}
	if value, ok := thic.mutation.CreatedAt(); ok {
		_spec.SetField(takehistoryitem.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := thic.mutation.UpdatedAt(); ok {
		_spec.SetField(takehistoryitem.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// TakeHistoryItemCreateBulk is the builder for creating many TakeHistoryItem entities in bulk.
type TakeHistoryItemCreateBulk struct {
	config
	err      error
	builders []*TakeHistoryItemCreate
}

// Save creates the TakeHistoryItem entities in the database.
func (thicb *TakeHistoryItemCreateBulk) Save(ctx context.Context) ([]*TakeHistoryItem, error) {
	if thicb.err != nil {
		return nil, thicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(thicb.builders))
	nodes := make([]*TakeHistoryItem, len(thicb.builders))
	mutators := make([]Mutator, len(thicb.builders))
	for i := range thicb.builders {
		func(i int, root context.Context) {
			builder := thicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TakeHistoryItemMutation)
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
					_, err = mutators[i+1].Mutate(root, thicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, thicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, thicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (thicb *TakeHistoryItemCreateBulk) SaveX(ctx context.Context) []*TakeHistoryItem {
	v, err := thicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (thicb *TakeHistoryItemCreateBulk) Exec(ctx context.Context) error {
	_, err := thicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (thicb *TakeHistoryItemCreateBulk) ExecX(ctx context.Context) {
	if err := thicb.Exec(ctx); err != nil {
		panic(err)
	}
}
