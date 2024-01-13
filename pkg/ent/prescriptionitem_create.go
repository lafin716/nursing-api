// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"nursing_api/pkg/ent/prescriptionitem"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PrescriptionItemCreate is the builder for creating a PrescriptionItem entity.
type PrescriptionItemCreate struct {
	config
	mutation *PrescriptionItemMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (pic *PrescriptionItemCreate) SetUserID(u uuid.UUID) *PrescriptionItemCreate {
	pic.mutation.SetUserID(u)
	return pic
}

// SetPrescriptionID sets the "prescription_id" field.
func (pic *PrescriptionItemCreate) SetPrescriptionID(u uuid.UUID) *PrescriptionItemCreate {
	pic.mutation.SetPrescriptionID(u)
	return pic
}

// SetMedicineName sets the "medicine_name" field.
func (pic *PrescriptionItemCreate) SetMedicineName(s string) *PrescriptionItemCreate {
	pic.mutation.SetMedicineName(s)
	return pic
}

// SetTakeTimeZone sets the "take_time_zone" field.
func (pic *PrescriptionItemCreate) SetTakeTimeZone(s string) *PrescriptionItemCreate {
	pic.mutation.SetTakeTimeZone(s)
	return pic
}

// SetNillableTakeTimeZone sets the "take_time_zone" field if the given value is not nil.
func (pic *PrescriptionItemCreate) SetNillableTakeTimeZone(s *string) *PrescriptionItemCreate {
	if s != nil {
		pic.SetTakeTimeZone(*s)
	}
	return pic
}

// SetTakeMoment sets the "take_moment" field.
func (pic *PrescriptionItemCreate) SetTakeMoment(s string) *PrescriptionItemCreate {
	pic.mutation.SetTakeMoment(s)
	return pic
}

// SetNillableTakeMoment sets the "take_moment" field if the given value is not nil.
func (pic *PrescriptionItemCreate) SetNillableTakeMoment(s *string) *PrescriptionItemCreate {
	if s != nil {
		pic.SetTakeMoment(*s)
	}
	return pic
}

// SetTakeEtc sets the "take_etc" field.
func (pic *PrescriptionItemCreate) SetTakeEtc(s string) *PrescriptionItemCreate {
	pic.mutation.SetTakeEtc(s)
	return pic
}

// SetNillableTakeEtc sets the "take_etc" field if the given value is not nil.
func (pic *PrescriptionItemCreate) SetNillableTakeEtc(s *string) *PrescriptionItemCreate {
	if s != nil {
		pic.SetTakeEtc(*s)
	}
	return pic
}

// SetTakeAmount sets the "take_amount" field.
func (pic *PrescriptionItemCreate) SetTakeAmount(f float64) *PrescriptionItemCreate {
	pic.mutation.SetTakeAmount(f)
	return pic
}

// SetNillableTakeAmount sets the "take_amount" field if the given value is not nil.
func (pic *PrescriptionItemCreate) SetNillableTakeAmount(f *float64) *PrescriptionItemCreate {
	if f != nil {
		pic.SetTakeAmount(*f)
	}
	return pic
}

// SetMedicineUnit sets the "medicine_unit" field.
func (pic *PrescriptionItemCreate) SetMedicineUnit(s string) *PrescriptionItemCreate {
	pic.mutation.SetMedicineUnit(s)
	return pic
}

// SetNillableMedicineUnit sets the "medicine_unit" field if the given value is not nil.
func (pic *PrescriptionItemCreate) SetNillableMedicineUnit(s *string) *PrescriptionItemCreate {
	if s != nil {
		pic.SetMedicineUnit(*s)
	}
	return pic
}

// SetMemo sets the "memo" field.
func (pic *PrescriptionItemCreate) SetMemo(s string) *PrescriptionItemCreate {
	pic.mutation.SetMemo(s)
	return pic
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (pic *PrescriptionItemCreate) SetNillableMemo(s *string) *PrescriptionItemCreate {
	if s != nil {
		pic.SetMemo(*s)
	}
	return pic
}

// SetCreatedAt sets the "created_at" field.
func (pic *PrescriptionItemCreate) SetCreatedAt(t time.Time) *PrescriptionItemCreate {
	pic.mutation.SetCreatedAt(t)
	return pic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pic *PrescriptionItemCreate) SetNillableCreatedAt(t *time.Time) *PrescriptionItemCreate {
	if t != nil {
		pic.SetCreatedAt(*t)
	}
	return pic
}

// SetUpdatedAt sets the "updated_at" field.
func (pic *PrescriptionItemCreate) SetUpdatedAt(t time.Time) *PrescriptionItemCreate {
	pic.mutation.SetUpdatedAt(t)
	return pic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pic *PrescriptionItemCreate) SetNillableUpdatedAt(t *time.Time) *PrescriptionItemCreate {
	if t != nil {
		pic.SetUpdatedAt(*t)
	}
	return pic
}

// SetID sets the "id" field.
func (pic *PrescriptionItemCreate) SetID(u uuid.UUID) *PrescriptionItemCreate {
	pic.mutation.SetID(u)
	return pic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pic *PrescriptionItemCreate) SetNillableID(u *uuid.UUID) *PrescriptionItemCreate {
	if u != nil {
		pic.SetID(*u)
	}
	return pic
}

// Mutation returns the PrescriptionItemMutation object of the builder.
func (pic *PrescriptionItemCreate) Mutation() *PrescriptionItemMutation {
	return pic.mutation
}

// Save creates the PrescriptionItem in the database.
func (pic *PrescriptionItemCreate) Save(ctx context.Context) (*PrescriptionItem, error) {
	pic.defaults()
	return withHooks(ctx, pic.sqlSave, pic.mutation, pic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pic *PrescriptionItemCreate) SaveX(ctx context.Context) *PrescriptionItem {
	v, err := pic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pic *PrescriptionItemCreate) Exec(ctx context.Context) error {
	_, err := pic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pic *PrescriptionItemCreate) ExecX(ctx context.Context) {
	if err := pic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pic *PrescriptionItemCreate) defaults() {
	if _, ok := pic.mutation.TakeAmount(); !ok {
		v := prescriptionitem.DefaultTakeAmount
		pic.mutation.SetTakeAmount(v)
	}
	if _, ok := pic.mutation.MedicineUnit(); !ok {
		v := prescriptionitem.DefaultMedicineUnit
		pic.mutation.SetMedicineUnit(v)
	}
	if _, ok := pic.mutation.CreatedAt(); !ok {
		v := prescriptionitem.DefaultCreatedAt()
		pic.mutation.SetCreatedAt(v)
	}
	if _, ok := pic.mutation.ID(); !ok {
		v := prescriptionitem.DefaultID()
		pic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pic *PrescriptionItemCreate) check() error {
	if _, ok := pic.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "PrescriptionItem.user_id"`)}
	}
	if _, ok := pic.mutation.PrescriptionID(); !ok {
		return &ValidationError{Name: "prescription_id", err: errors.New(`ent: missing required field "PrescriptionItem.prescription_id"`)}
	}
	if _, ok := pic.mutation.MedicineName(); !ok {
		return &ValidationError{Name: "medicine_name", err: errors.New(`ent: missing required field "PrescriptionItem.medicine_name"`)}
	}
	if _, ok := pic.mutation.TakeAmount(); !ok {
		return &ValidationError{Name: "take_amount", err: errors.New(`ent: missing required field "PrescriptionItem.take_amount"`)}
	}
	if _, ok := pic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "PrescriptionItem.created_at"`)}
	}
	return nil
}

func (pic *PrescriptionItemCreate) sqlSave(ctx context.Context) (*PrescriptionItem, error) {
	if err := pic.check(); err != nil {
		return nil, err
	}
	_node, _spec := pic.createSpec()
	if err := sqlgraph.CreateNode(ctx, pic.driver, _spec); err != nil {
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
	pic.mutation.id = &_node.ID
	pic.mutation.done = true
	return _node, nil
}

func (pic *PrescriptionItemCreate) createSpec() (*PrescriptionItem, *sqlgraph.CreateSpec) {
	var (
		_node = &PrescriptionItem{config: pic.config}
		_spec = sqlgraph.NewCreateSpec(prescriptionitem.Table, sqlgraph.NewFieldSpec(prescriptionitem.FieldID, field.TypeUUID))
	)
	if id, ok := pic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pic.mutation.UserID(); ok {
		_spec.SetField(prescriptionitem.FieldUserID, field.TypeUUID, value)
		_node.UserID = value
	}
	if value, ok := pic.mutation.PrescriptionID(); ok {
		_spec.SetField(prescriptionitem.FieldPrescriptionID, field.TypeUUID, value)
		_node.PrescriptionID = value
	}
	if value, ok := pic.mutation.MedicineName(); ok {
		_spec.SetField(prescriptionitem.FieldMedicineName, field.TypeString, value)
		_node.MedicineName = value
	}
	if value, ok := pic.mutation.TakeTimeZone(); ok {
		_spec.SetField(prescriptionitem.FieldTakeTimeZone, field.TypeString, value)
		_node.TakeTimeZone = value
	}
	if value, ok := pic.mutation.TakeMoment(); ok {
		_spec.SetField(prescriptionitem.FieldTakeMoment, field.TypeString, value)
		_node.TakeMoment = value
	}
	if value, ok := pic.mutation.TakeEtc(); ok {
		_spec.SetField(prescriptionitem.FieldTakeEtc, field.TypeString, value)
		_node.TakeEtc = value
	}
	if value, ok := pic.mutation.TakeAmount(); ok {
		_spec.SetField(prescriptionitem.FieldTakeAmount, field.TypeFloat64, value)
		_node.TakeAmount = value
	}
	if value, ok := pic.mutation.MedicineUnit(); ok {
		_spec.SetField(prescriptionitem.FieldMedicineUnit, field.TypeString, value)
		_node.MedicineUnit = value
	}
	if value, ok := pic.mutation.Memo(); ok {
		_spec.SetField(prescriptionitem.FieldMemo, field.TypeString, value)
		_node.Memo = value
	}
	if value, ok := pic.mutation.CreatedAt(); ok {
		_spec.SetField(prescriptionitem.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pic.mutation.UpdatedAt(); ok {
		_spec.SetField(prescriptionitem.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// PrescriptionItemCreateBulk is the builder for creating many PrescriptionItem entities in bulk.
type PrescriptionItemCreateBulk struct {
	config
	err      error
	builders []*PrescriptionItemCreate
}

// Save creates the PrescriptionItem entities in the database.
func (picb *PrescriptionItemCreateBulk) Save(ctx context.Context) ([]*PrescriptionItem, error) {
	if picb.err != nil {
		return nil, picb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(picb.builders))
	nodes := make([]*PrescriptionItem, len(picb.builders))
	mutators := make([]Mutator, len(picb.builders))
	for i := range picb.builders {
		func(i int, root context.Context) {
			builder := picb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PrescriptionItemMutation)
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
					_, err = mutators[i+1].Mutate(root, picb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, picb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, picb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (picb *PrescriptionItemCreateBulk) SaveX(ctx context.Context) []*PrescriptionItem {
	v, err := picb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (picb *PrescriptionItemCreateBulk) Exec(ctx context.Context) error {
	_, err := picb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (picb *PrescriptionItemCreateBulk) ExecX(ctx context.Context) {
	if err := picb.Exec(ctx); err != nil {
		panic(err)
	}
}