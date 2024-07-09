// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"nursing_api/pkg/ent/predicate"
	"nursing_api/pkg/ent/prescription"
	"nursing_api/pkg/ent/prescriptionitem"
	"nursing_api/pkg/ent/takehistoryitem"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PrescriptionItemUpdate is the builder for updating PrescriptionItem entities.
type PrescriptionItemUpdate struct {
	config
	hooks    []Hook
	mutation *PrescriptionItemMutation
}

// Where appends a list predicates to the PrescriptionItemUpdate builder.
func (piu *PrescriptionItemUpdate) Where(ps ...predicate.PrescriptionItem) *PrescriptionItemUpdate {
	piu.mutation.Where(ps...)
	return piu
}

// SetPrescriptionID sets the "prescription_id" field.
func (piu *PrescriptionItemUpdate) SetPrescriptionID(u uuid.UUID) *PrescriptionItemUpdate {
	piu.mutation.SetPrescriptionID(u)
	return piu
}

// SetNillablePrescriptionID sets the "prescription_id" field if the given value is not nil.
func (piu *PrescriptionItemUpdate) SetNillablePrescriptionID(u *uuid.UUID) *PrescriptionItemUpdate {
	if u != nil {
		piu.SetPrescriptionID(*u)
	}
	return piu
}

// SetTimezoneID sets the "timezone_id" field.
func (piu *PrescriptionItemUpdate) SetTimezoneID(u uuid.UUID) *PrescriptionItemUpdate {
	piu.mutation.SetTimezoneID(u)
	return piu
}

// SetNillableTimezoneID sets the "timezone_id" field if the given value is not nil.
func (piu *PrescriptionItemUpdate) SetNillableTimezoneID(u *uuid.UUID) *PrescriptionItemUpdate {
	if u != nil {
		piu.SetTimezoneID(*u)
	}
	return piu
}

// SetMedicineID sets the "medicine_id" field.
func (piu *PrescriptionItemUpdate) SetMedicineID(u uuid.UUID) *PrescriptionItemUpdate {
	piu.mutation.SetMedicineID(u)
	return piu
}

// SetNillableMedicineID sets the "medicine_id" field if the given value is not nil.
func (piu *PrescriptionItemUpdate) SetNillableMedicineID(u *uuid.UUID) *PrescriptionItemUpdate {
	if u != nil {
		piu.SetMedicineID(*u)
	}
	return piu
}

// SetMedicineName sets the "medicine_name" field.
func (piu *PrescriptionItemUpdate) SetMedicineName(s string) *PrescriptionItemUpdate {
	piu.mutation.SetMedicineName(s)
	return piu
}

// SetNillableMedicineName sets the "medicine_name" field if the given value is not nil.
func (piu *PrescriptionItemUpdate) SetNillableMedicineName(s *string) *PrescriptionItemUpdate {
	if s != nil {
		piu.SetMedicineName(*s)
	}
	return piu
}

// SetTimezoneName sets the "timezone_name" field.
func (piu *PrescriptionItemUpdate) SetTimezoneName(s string) *PrescriptionItemUpdate {
	piu.mutation.SetTimezoneName(s)
	return piu
}

// SetNillableTimezoneName sets the "timezone_name" field if the given value is not nil.
func (piu *PrescriptionItemUpdate) SetNillableTimezoneName(s *string) *PrescriptionItemUpdate {
	if s != nil {
		piu.SetTimezoneName(*s)
	}
	return piu
}

// ClearTimezoneName clears the value of the "timezone_name" field.
func (piu *PrescriptionItemUpdate) ClearTimezoneName() *PrescriptionItemUpdate {
	piu.mutation.ClearTimezoneName()
	return piu
}

// SetMidday sets the "midday" field.
func (piu *PrescriptionItemUpdate) SetMidday(s string) *PrescriptionItemUpdate {
	piu.mutation.SetMidday(s)
	return piu
}

// SetNillableMidday sets the "midday" field if the given value is not nil.
func (piu *PrescriptionItemUpdate) SetNillableMidday(s *string) *PrescriptionItemUpdate {
	if s != nil {
		piu.SetMidday(*s)
	}
	return piu
}

// SetHour sets the "hour" field.
func (piu *PrescriptionItemUpdate) SetHour(s string) *PrescriptionItemUpdate {
	piu.mutation.SetHour(s)
	return piu
}

// SetNillableHour sets the "hour" field if the given value is not nil.
func (piu *PrescriptionItemUpdate) SetNillableHour(s *string) *PrescriptionItemUpdate {
	if s != nil {
		piu.SetHour(*s)
	}
	return piu
}

// SetMinute sets the "minute" field.
func (piu *PrescriptionItemUpdate) SetMinute(s string) *PrescriptionItemUpdate {
	piu.mutation.SetMinute(s)
	return piu
}

// SetNillableMinute sets the "minute" field if the given value is not nil.
func (piu *PrescriptionItemUpdate) SetNillableMinute(s *string) *PrescriptionItemUpdate {
	if s != nil {
		piu.SetMinute(*s)
	}
	return piu
}

// SetTotalAmount sets the "total_amount" field.
func (piu *PrescriptionItemUpdate) SetTotalAmount(f float64) *PrescriptionItemUpdate {
	piu.mutation.ResetTotalAmount()
	piu.mutation.SetTotalAmount(f)
	return piu
}

// SetNillableTotalAmount sets the "total_amount" field if the given value is not nil.
func (piu *PrescriptionItemUpdate) SetNillableTotalAmount(f *float64) *PrescriptionItemUpdate {
	if f != nil {
		piu.SetTotalAmount(*f)
	}
	return piu
}

// AddTotalAmount adds f to the "total_amount" field.
func (piu *PrescriptionItemUpdate) AddTotalAmount(f float64) *PrescriptionItemUpdate {
	piu.mutation.AddTotalAmount(f)
	return piu
}

// SetRemainAmount sets the "remain_amount" field.
func (piu *PrescriptionItemUpdate) SetRemainAmount(f float64) *PrescriptionItemUpdate {
	piu.mutation.ResetRemainAmount()
	piu.mutation.SetRemainAmount(f)
	return piu
}

// SetNillableRemainAmount sets the "remain_amount" field if the given value is not nil.
func (piu *PrescriptionItemUpdate) SetNillableRemainAmount(f *float64) *PrescriptionItemUpdate {
	if f != nil {
		piu.SetRemainAmount(*f)
	}
	return piu
}

// AddRemainAmount adds f to the "remain_amount" field.
func (piu *PrescriptionItemUpdate) AddRemainAmount(f float64) *PrescriptionItemUpdate {
	piu.mutation.AddRemainAmount(f)
	return piu
}

// SetTakeAmount sets the "take_amount" field.
func (piu *PrescriptionItemUpdate) SetTakeAmount(f float64) *PrescriptionItemUpdate {
	piu.mutation.ResetTakeAmount()
	piu.mutation.SetTakeAmount(f)
	return piu
}

// SetNillableTakeAmount sets the "take_amount" field if the given value is not nil.
func (piu *PrescriptionItemUpdate) SetNillableTakeAmount(f *float64) *PrescriptionItemUpdate {
	if f != nil {
		piu.SetTakeAmount(*f)
	}
	return piu
}

// AddTakeAmount adds f to the "take_amount" field.
func (piu *PrescriptionItemUpdate) AddTakeAmount(f float64) *PrescriptionItemUpdate {
	piu.mutation.AddTakeAmount(f)
	return piu
}

// SetMedicineUnit sets the "medicine_unit" field.
func (piu *PrescriptionItemUpdate) SetMedicineUnit(s string) *PrescriptionItemUpdate {
	piu.mutation.SetMedicineUnit(s)
	return piu
}

// SetNillableMedicineUnit sets the "medicine_unit" field if the given value is not nil.
func (piu *PrescriptionItemUpdate) SetNillableMedicineUnit(s *string) *PrescriptionItemUpdate {
	if s != nil {
		piu.SetMedicineUnit(*s)
	}
	return piu
}

// ClearMedicineUnit clears the value of the "medicine_unit" field.
func (piu *PrescriptionItemUpdate) ClearMedicineUnit() *PrescriptionItemUpdate {
	piu.mutation.ClearMedicineUnit()
	return piu
}

// SetMemo sets the "memo" field.
func (piu *PrescriptionItemUpdate) SetMemo(s string) *PrescriptionItemUpdate {
	piu.mutation.SetMemo(s)
	return piu
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (piu *PrescriptionItemUpdate) SetNillableMemo(s *string) *PrescriptionItemUpdate {
	if s != nil {
		piu.SetMemo(*s)
	}
	return piu
}

// ClearMemo clears the value of the "memo" field.
func (piu *PrescriptionItemUpdate) ClearMemo() *PrescriptionItemUpdate {
	piu.mutation.ClearMemo()
	return piu
}

// SetCreatedAt sets the "created_at" field.
func (piu *PrescriptionItemUpdate) SetCreatedAt(t time.Time) *PrescriptionItemUpdate {
	piu.mutation.SetCreatedAt(t)
	return piu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (piu *PrescriptionItemUpdate) SetNillableCreatedAt(t *time.Time) *PrescriptionItemUpdate {
	if t != nil {
		piu.SetCreatedAt(*t)
	}
	return piu
}

// SetUpdatedAt sets the "updated_at" field.
func (piu *PrescriptionItemUpdate) SetUpdatedAt(t time.Time) *PrescriptionItemUpdate {
	piu.mutation.SetUpdatedAt(t)
	return piu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (piu *PrescriptionItemUpdate) SetNillableUpdatedAt(t *time.Time) *PrescriptionItemUpdate {
	if t != nil {
		piu.SetUpdatedAt(*t)
	}
	return piu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (piu *PrescriptionItemUpdate) ClearUpdatedAt() *PrescriptionItemUpdate {
	piu.mutation.ClearUpdatedAt()
	return piu
}

// SetPrescription sets the "prescription" edge to the Prescription entity.
func (piu *PrescriptionItemUpdate) SetPrescription(p *Prescription) *PrescriptionItemUpdate {
	return piu.SetPrescriptionID(p.ID)
}

// SetTakeHistoryItemID sets the "take_history_item" edge to the TakeHistoryItem entity by ID.
func (piu *PrescriptionItemUpdate) SetTakeHistoryItemID(id uuid.UUID) *PrescriptionItemUpdate {
	piu.mutation.SetTakeHistoryItemID(id)
	return piu
}

// SetNillableTakeHistoryItemID sets the "take_history_item" edge to the TakeHistoryItem entity by ID if the given value is not nil.
func (piu *PrescriptionItemUpdate) SetNillableTakeHistoryItemID(id *uuid.UUID) *PrescriptionItemUpdate {
	if id != nil {
		piu = piu.SetTakeHistoryItemID(*id)
	}
	return piu
}

// SetTakeHistoryItem sets the "take_history_item" edge to the TakeHistoryItem entity.
func (piu *PrescriptionItemUpdate) SetTakeHistoryItem(t *TakeHistoryItem) *PrescriptionItemUpdate {
	return piu.SetTakeHistoryItemID(t.ID)
}

// Mutation returns the PrescriptionItemMutation object of the builder.
func (piu *PrescriptionItemUpdate) Mutation() *PrescriptionItemMutation {
	return piu.mutation
}

// ClearPrescription clears the "prescription" edge to the Prescription entity.
func (piu *PrescriptionItemUpdate) ClearPrescription() *PrescriptionItemUpdate {
	piu.mutation.ClearPrescription()
	return piu
}

// ClearTakeHistoryItem clears the "take_history_item" edge to the TakeHistoryItem entity.
func (piu *PrescriptionItemUpdate) ClearTakeHistoryItem() *PrescriptionItemUpdate {
	piu.mutation.ClearTakeHistoryItem()
	return piu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (piu *PrescriptionItemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, piu.sqlSave, piu.mutation, piu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (piu *PrescriptionItemUpdate) SaveX(ctx context.Context) int {
	affected, err := piu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (piu *PrescriptionItemUpdate) Exec(ctx context.Context) error {
	_, err := piu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piu *PrescriptionItemUpdate) ExecX(ctx context.Context) {
	if err := piu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (piu *PrescriptionItemUpdate) check() error {
	if _, ok := piu.mutation.PrescriptionID(); piu.mutation.PrescriptionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PrescriptionItem.prescription"`)
	}
	return nil
}

func (piu *PrescriptionItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := piu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(prescriptionitem.Table, prescriptionitem.Columns, sqlgraph.NewFieldSpec(prescriptionitem.FieldID, field.TypeUUID))
	if ps := piu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piu.mutation.TimezoneID(); ok {
		_spec.SetField(prescriptionitem.FieldTimezoneID, field.TypeUUID, value)
	}
	if value, ok := piu.mutation.MedicineID(); ok {
		_spec.SetField(prescriptionitem.FieldMedicineID, field.TypeUUID, value)
	}
	if value, ok := piu.mutation.MedicineName(); ok {
		_spec.SetField(prescriptionitem.FieldMedicineName, field.TypeString, value)
	}
	if value, ok := piu.mutation.TimezoneName(); ok {
		_spec.SetField(prescriptionitem.FieldTimezoneName, field.TypeString, value)
	}
	if piu.mutation.TimezoneNameCleared() {
		_spec.ClearField(prescriptionitem.FieldTimezoneName, field.TypeString)
	}
	if value, ok := piu.mutation.Midday(); ok {
		_spec.SetField(prescriptionitem.FieldMidday, field.TypeString, value)
	}
	if value, ok := piu.mutation.Hour(); ok {
		_spec.SetField(prescriptionitem.FieldHour, field.TypeString, value)
	}
	if value, ok := piu.mutation.Minute(); ok {
		_spec.SetField(prescriptionitem.FieldMinute, field.TypeString, value)
	}
	if value, ok := piu.mutation.TotalAmount(); ok {
		_spec.SetField(prescriptionitem.FieldTotalAmount, field.TypeFloat64, value)
	}
	if value, ok := piu.mutation.AddedTotalAmount(); ok {
		_spec.AddField(prescriptionitem.FieldTotalAmount, field.TypeFloat64, value)
	}
	if value, ok := piu.mutation.RemainAmount(); ok {
		_spec.SetField(prescriptionitem.FieldRemainAmount, field.TypeFloat64, value)
	}
	if value, ok := piu.mutation.AddedRemainAmount(); ok {
		_spec.AddField(prescriptionitem.FieldRemainAmount, field.TypeFloat64, value)
	}
	if value, ok := piu.mutation.TakeAmount(); ok {
		_spec.SetField(prescriptionitem.FieldTakeAmount, field.TypeFloat64, value)
	}
	if value, ok := piu.mutation.AddedTakeAmount(); ok {
		_spec.AddField(prescriptionitem.FieldTakeAmount, field.TypeFloat64, value)
	}
	if value, ok := piu.mutation.MedicineUnit(); ok {
		_spec.SetField(prescriptionitem.FieldMedicineUnit, field.TypeString, value)
	}
	if piu.mutation.MedicineUnitCleared() {
		_spec.ClearField(prescriptionitem.FieldMedicineUnit, field.TypeString)
	}
	if value, ok := piu.mutation.Memo(); ok {
		_spec.SetField(prescriptionitem.FieldMemo, field.TypeString, value)
	}
	if piu.mutation.MemoCleared() {
		_spec.ClearField(prescriptionitem.FieldMemo, field.TypeString)
	}
	if value, ok := piu.mutation.CreatedAt(); ok {
		_spec.SetField(prescriptionitem.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := piu.mutation.UpdatedAt(); ok {
		_spec.SetField(prescriptionitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if piu.mutation.UpdatedAtCleared() {
		_spec.ClearField(prescriptionitem.FieldUpdatedAt, field.TypeTime)
	}
	if piu.mutation.PrescriptionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescriptionitem.PrescriptionTable,
			Columns: []string{prescriptionitem.PrescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(prescription.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piu.mutation.PrescriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescriptionitem.PrescriptionTable,
			Columns: []string{prescriptionitem.PrescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(prescription.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if piu.mutation.TakeHistoryItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   prescriptionitem.TakeHistoryItemTable,
			Columns: []string{prescriptionitem.TakeHistoryItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(takehistoryitem.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piu.mutation.TakeHistoryItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   prescriptionitem.TakeHistoryItemTable,
			Columns: []string{prescriptionitem.TakeHistoryItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(takehistoryitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, piu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prescriptionitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	piu.mutation.done = true
	return n, nil
}

// PrescriptionItemUpdateOne is the builder for updating a single PrescriptionItem entity.
type PrescriptionItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PrescriptionItemMutation
}

// SetPrescriptionID sets the "prescription_id" field.
func (piuo *PrescriptionItemUpdateOne) SetPrescriptionID(u uuid.UUID) *PrescriptionItemUpdateOne {
	piuo.mutation.SetPrescriptionID(u)
	return piuo
}

// SetNillablePrescriptionID sets the "prescription_id" field if the given value is not nil.
func (piuo *PrescriptionItemUpdateOne) SetNillablePrescriptionID(u *uuid.UUID) *PrescriptionItemUpdateOne {
	if u != nil {
		piuo.SetPrescriptionID(*u)
	}
	return piuo
}

// SetTimezoneID sets the "timezone_id" field.
func (piuo *PrescriptionItemUpdateOne) SetTimezoneID(u uuid.UUID) *PrescriptionItemUpdateOne {
	piuo.mutation.SetTimezoneID(u)
	return piuo
}

// SetNillableTimezoneID sets the "timezone_id" field if the given value is not nil.
func (piuo *PrescriptionItemUpdateOne) SetNillableTimezoneID(u *uuid.UUID) *PrescriptionItemUpdateOne {
	if u != nil {
		piuo.SetTimezoneID(*u)
	}
	return piuo
}

// SetMedicineID sets the "medicine_id" field.
func (piuo *PrescriptionItemUpdateOne) SetMedicineID(u uuid.UUID) *PrescriptionItemUpdateOne {
	piuo.mutation.SetMedicineID(u)
	return piuo
}

// SetNillableMedicineID sets the "medicine_id" field if the given value is not nil.
func (piuo *PrescriptionItemUpdateOne) SetNillableMedicineID(u *uuid.UUID) *PrescriptionItemUpdateOne {
	if u != nil {
		piuo.SetMedicineID(*u)
	}
	return piuo
}

// SetMedicineName sets the "medicine_name" field.
func (piuo *PrescriptionItemUpdateOne) SetMedicineName(s string) *PrescriptionItemUpdateOne {
	piuo.mutation.SetMedicineName(s)
	return piuo
}

// SetNillableMedicineName sets the "medicine_name" field if the given value is not nil.
func (piuo *PrescriptionItemUpdateOne) SetNillableMedicineName(s *string) *PrescriptionItemUpdateOne {
	if s != nil {
		piuo.SetMedicineName(*s)
	}
	return piuo
}

// SetTimezoneName sets the "timezone_name" field.
func (piuo *PrescriptionItemUpdateOne) SetTimezoneName(s string) *PrescriptionItemUpdateOne {
	piuo.mutation.SetTimezoneName(s)
	return piuo
}

// SetNillableTimezoneName sets the "timezone_name" field if the given value is not nil.
func (piuo *PrescriptionItemUpdateOne) SetNillableTimezoneName(s *string) *PrescriptionItemUpdateOne {
	if s != nil {
		piuo.SetTimezoneName(*s)
	}
	return piuo
}

// ClearTimezoneName clears the value of the "timezone_name" field.
func (piuo *PrescriptionItemUpdateOne) ClearTimezoneName() *PrescriptionItemUpdateOne {
	piuo.mutation.ClearTimezoneName()
	return piuo
}

// SetMidday sets the "midday" field.
func (piuo *PrescriptionItemUpdateOne) SetMidday(s string) *PrescriptionItemUpdateOne {
	piuo.mutation.SetMidday(s)
	return piuo
}

// SetNillableMidday sets the "midday" field if the given value is not nil.
func (piuo *PrescriptionItemUpdateOne) SetNillableMidday(s *string) *PrescriptionItemUpdateOne {
	if s != nil {
		piuo.SetMidday(*s)
	}
	return piuo
}

// SetHour sets the "hour" field.
func (piuo *PrescriptionItemUpdateOne) SetHour(s string) *PrescriptionItemUpdateOne {
	piuo.mutation.SetHour(s)
	return piuo
}

// SetNillableHour sets the "hour" field if the given value is not nil.
func (piuo *PrescriptionItemUpdateOne) SetNillableHour(s *string) *PrescriptionItemUpdateOne {
	if s != nil {
		piuo.SetHour(*s)
	}
	return piuo
}

// SetMinute sets the "minute" field.
func (piuo *PrescriptionItemUpdateOne) SetMinute(s string) *PrescriptionItemUpdateOne {
	piuo.mutation.SetMinute(s)
	return piuo
}

// SetNillableMinute sets the "minute" field if the given value is not nil.
func (piuo *PrescriptionItemUpdateOne) SetNillableMinute(s *string) *PrescriptionItemUpdateOne {
	if s != nil {
		piuo.SetMinute(*s)
	}
	return piuo
}

// SetTotalAmount sets the "total_amount" field.
func (piuo *PrescriptionItemUpdateOne) SetTotalAmount(f float64) *PrescriptionItemUpdateOne {
	piuo.mutation.ResetTotalAmount()
	piuo.mutation.SetTotalAmount(f)
	return piuo
}

// SetNillableTotalAmount sets the "total_amount" field if the given value is not nil.
func (piuo *PrescriptionItemUpdateOne) SetNillableTotalAmount(f *float64) *PrescriptionItemUpdateOne {
	if f != nil {
		piuo.SetTotalAmount(*f)
	}
	return piuo
}

// AddTotalAmount adds f to the "total_amount" field.
func (piuo *PrescriptionItemUpdateOne) AddTotalAmount(f float64) *PrescriptionItemUpdateOne {
	piuo.mutation.AddTotalAmount(f)
	return piuo
}

// SetRemainAmount sets the "remain_amount" field.
func (piuo *PrescriptionItemUpdateOne) SetRemainAmount(f float64) *PrescriptionItemUpdateOne {
	piuo.mutation.ResetRemainAmount()
	piuo.mutation.SetRemainAmount(f)
	return piuo
}

// SetNillableRemainAmount sets the "remain_amount" field if the given value is not nil.
func (piuo *PrescriptionItemUpdateOne) SetNillableRemainAmount(f *float64) *PrescriptionItemUpdateOne {
	if f != nil {
		piuo.SetRemainAmount(*f)
	}
	return piuo
}

// AddRemainAmount adds f to the "remain_amount" field.
func (piuo *PrescriptionItemUpdateOne) AddRemainAmount(f float64) *PrescriptionItemUpdateOne {
	piuo.mutation.AddRemainAmount(f)
	return piuo
}

// SetTakeAmount sets the "take_amount" field.
func (piuo *PrescriptionItemUpdateOne) SetTakeAmount(f float64) *PrescriptionItemUpdateOne {
	piuo.mutation.ResetTakeAmount()
	piuo.mutation.SetTakeAmount(f)
	return piuo
}

// SetNillableTakeAmount sets the "take_amount" field if the given value is not nil.
func (piuo *PrescriptionItemUpdateOne) SetNillableTakeAmount(f *float64) *PrescriptionItemUpdateOne {
	if f != nil {
		piuo.SetTakeAmount(*f)
	}
	return piuo
}

// AddTakeAmount adds f to the "take_amount" field.
func (piuo *PrescriptionItemUpdateOne) AddTakeAmount(f float64) *PrescriptionItemUpdateOne {
	piuo.mutation.AddTakeAmount(f)
	return piuo
}

// SetMedicineUnit sets the "medicine_unit" field.
func (piuo *PrescriptionItemUpdateOne) SetMedicineUnit(s string) *PrescriptionItemUpdateOne {
	piuo.mutation.SetMedicineUnit(s)
	return piuo
}

// SetNillableMedicineUnit sets the "medicine_unit" field if the given value is not nil.
func (piuo *PrescriptionItemUpdateOne) SetNillableMedicineUnit(s *string) *PrescriptionItemUpdateOne {
	if s != nil {
		piuo.SetMedicineUnit(*s)
	}
	return piuo
}

// ClearMedicineUnit clears the value of the "medicine_unit" field.
func (piuo *PrescriptionItemUpdateOne) ClearMedicineUnit() *PrescriptionItemUpdateOne {
	piuo.mutation.ClearMedicineUnit()
	return piuo
}

// SetMemo sets the "memo" field.
func (piuo *PrescriptionItemUpdateOne) SetMemo(s string) *PrescriptionItemUpdateOne {
	piuo.mutation.SetMemo(s)
	return piuo
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (piuo *PrescriptionItemUpdateOne) SetNillableMemo(s *string) *PrescriptionItemUpdateOne {
	if s != nil {
		piuo.SetMemo(*s)
	}
	return piuo
}

// ClearMemo clears the value of the "memo" field.
func (piuo *PrescriptionItemUpdateOne) ClearMemo() *PrescriptionItemUpdateOne {
	piuo.mutation.ClearMemo()
	return piuo
}

// SetCreatedAt sets the "created_at" field.
func (piuo *PrescriptionItemUpdateOne) SetCreatedAt(t time.Time) *PrescriptionItemUpdateOne {
	piuo.mutation.SetCreatedAt(t)
	return piuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (piuo *PrescriptionItemUpdateOne) SetNillableCreatedAt(t *time.Time) *PrescriptionItemUpdateOne {
	if t != nil {
		piuo.SetCreatedAt(*t)
	}
	return piuo
}

// SetUpdatedAt sets the "updated_at" field.
func (piuo *PrescriptionItemUpdateOne) SetUpdatedAt(t time.Time) *PrescriptionItemUpdateOne {
	piuo.mutation.SetUpdatedAt(t)
	return piuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (piuo *PrescriptionItemUpdateOne) SetNillableUpdatedAt(t *time.Time) *PrescriptionItemUpdateOne {
	if t != nil {
		piuo.SetUpdatedAt(*t)
	}
	return piuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (piuo *PrescriptionItemUpdateOne) ClearUpdatedAt() *PrescriptionItemUpdateOne {
	piuo.mutation.ClearUpdatedAt()
	return piuo
}

// SetPrescription sets the "prescription" edge to the Prescription entity.
func (piuo *PrescriptionItemUpdateOne) SetPrescription(p *Prescription) *PrescriptionItemUpdateOne {
	return piuo.SetPrescriptionID(p.ID)
}

// SetTakeHistoryItemID sets the "take_history_item" edge to the TakeHistoryItem entity by ID.
func (piuo *PrescriptionItemUpdateOne) SetTakeHistoryItemID(id uuid.UUID) *PrescriptionItemUpdateOne {
	piuo.mutation.SetTakeHistoryItemID(id)
	return piuo
}

// SetNillableTakeHistoryItemID sets the "take_history_item" edge to the TakeHistoryItem entity by ID if the given value is not nil.
func (piuo *PrescriptionItemUpdateOne) SetNillableTakeHistoryItemID(id *uuid.UUID) *PrescriptionItemUpdateOne {
	if id != nil {
		piuo = piuo.SetTakeHistoryItemID(*id)
	}
	return piuo
}

// SetTakeHistoryItem sets the "take_history_item" edge to the TakeHistoryItem entity.
func (piuo *PrescriptionItemUpdateOne) SetTakeHistoryItem(t *TakeHistoryItem) *PrescriptionItemUpdateOne {
	return piuo.SetTakeHistoryItemID(t.ID)
}

// Mutation returns the PrescriptionItemMutation object of the builder.
func (piuo *PrescriptionItemUpdateOne) Mutation() *PrescriptionItemMutation {
	return piuo.mutation
}

// ClearPrescription clears the "prescription" edge to the Prescription entity.
func (piuo *PrescriptionItemUpdateOne) ClearPrescription() *PrescriptionItemUpdateOne {
	piuo.mutation.ClearPrescription()
	return piuo
}

// ClearTakeHistoryItem clears the "take_history_item" edge to the TakeHistoryItem entity.
func (piuo *PrescriptionItemUpdateOne) ClearTakeHistoryItem() *PrescriptionItemUpdateOne {
	piuo.mutation.ClearTakeHistoryItem()
	return piuo
}

// Where appends a list predicates to the PrescriptionItemUpdate builder.
func (piuo *PrescriptionItemUpdateOne) Where(ps ...predicate.PrescriptionItem) *PrescriptionItemUpdateOne {
	piuo.mutation.Where(ps...)
	return piuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (piuo *PrescriptionItemUpdateOne) Select(field string, fields ...string) *PrescriptionItemUpdateOne {
	piuo.fields = append([]string{field}, fields...)
	return piuo
}

// Save executes the query and returns the updated PrescriptionItem entity.
func (piuo *PrescriptionItemUpdateOne) Save(ctx context.Context) (*PrescriptionItem, error) {
	return withHooks(ctx, piuo.sqlSave, piuo.mutation, piuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (piuo *PrescriptionItemUpdateOne) SaveX(ctx context.Context) *PrescriptionItem {
	node, err := piuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (piuo *PrescriptionItemUpdateOne) Exec(ctx context.Context) error {
	_, err := piuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piuo *PrescriptionItemUpdateOne) ExecX(ctx context.Context) {
	if err := piuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (piuo *PrescriptionItemUpdateOne) check() error {
	if _, ok := piuo.mutation.PrescriptionID(); piuo.mutation.PrescriptionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PrescriptionItem.prescription"`)
	}
	return nil
}

func (piuo *PrescriptionItemUpdateOne) sqlSave(ctx context.Context) (_node *PrescriptionItem, err error) {
	if err := piuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(prescriptionitem.Table, prescriptionitem.Columns, sqlgraph.NewFieldSpec(prescriptionitem.FieldID, field.TypeUUID))
	id, ok := piuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PrescriptionItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := piuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, prescriptionitem.FieldID)
		for _, f := range fields {
			if !prescriptionitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != prescriptionitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := piuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piuo.mutation.TimezoneID(); ok {
		_spec.SetField(prescriptionitem.FieldTimezoneID, field.TypeUUID, value)
	}
	if value, ok := piuo.mutation.MedicineID(); ok {
		_spec.SetField(prescriptionitem.FieldMedicineID, field.TypeUUID, value)
	}
	if value, ok := piuo.mutation.MedicineName(); ok {
		_spec.SetField(prescriptionitem.FieldMedicineName, field.TypeString, value)
	}
	if value, ok := piuo.mutation.TimezoneName(); ok {
		_spec.SetField(prescriptionitem.FieldTimezoneName, field.TypeString, value)
	}
	if piuo.mutation.TimezoneNameCleared() {
		_spec.ClearField(prescriptionitem.FieldTimezoneName, field.TypeString)
	}
	if value, ok := piuo.mutation.Midday(); ok {
		_spec.SetField(prescriptionitem.FieldMidday, field.TypeString, value)
	}
	if value, ok := piuo.mutation.Hour(); ok {
		_spec.SetField(prescriptionitem.FieldHour, field.TypeString, value)
	}
	if value, ok := piuo.mutation.Minute(); ok {
		_spec.SetField(prescriptionitem.FieldMinute, field.TypeString, value)
	}
	if value, ok := piuo.mutation.TotalAmount(); ok {
		_spec.SetField(prescriptionitem.FieldTotalAmount, field.TypeFloat64, value)
	}
	if value, ok := piuo.mutation.AddedTotalAmount(); ok {
		_spec.AddField(prescriptionitem.FieldTotalAmount, field.TypeFloat64, value)
	}
	if value, ok := piuo.mutation.RemainAmount(); ok {
		_spec.SetField(prescriptionitem.FieldRemainAmount, field.TypeFloat64, value)
	}
	if value, ok := piuo.mutation.AddedRemainAmount(); ok {
		_spec.AddField(prescriptionitem.FieldRemainAmount, field.TypeFloat64, value)
	}
	if value, ok := piuo.mutation.TakeAmount(); ok {
		_spec.SetField(prescriptionitem.FieldTakeAmount, field.TypeFloat64, value)
	}
	if value, ok := piuo.mutation.AddedTakeAmount(); ok {
		_spec.AddField(prescriptionitem.FieldTakeAmount, field.TypeFloat64, value)
	}
	if value, ok := piuo.mutation.MedicineUnit(); ok {
		_spec.SetField(prescriptionitem.FieldMedicineUnit, field.TypeString, value)
	}
	if piuo.mutation.MedicineUnitCleared() {
		_spec.ClearField(prescriptionitem.FieldMedicineUnit, field.TypeString)
	}
	if value, ok := piuo.mutation.Memo(); ok {
		_spec.SetField(prescriptionitem.FieldMemo, field.TypeString, value)
	}
	if piuo.mutation.MemoCleared() {
		_spec.ClearField(prescriptionitem.FieldMemo, field.TypeString)
	}
	if value, ok := piuo.mutation.CreatedAt(); ok {
		_spec.SetField(prescriptionitem.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := piuo.mutation.UpdatedAt(); ok {
		_spec.SetField(prescriptionitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if piuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(prescriptionitem.FieldUpdatedAt, field.TypeTime)
	}
	if piuo.mutation.PrescriptionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescriptionitem.PrescriptionTable,
			Columns: []string{prescriptionitem.PrescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(prescription.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piuo.mutation.PrescriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescriptionitem.PrescriptionTable,
			Columns: []string{prescriptionitem.PrescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(prescription.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if piuo.mutation.TakeHistoryItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   prescriptionitem.TakeHistoryItemTable,
			Columns: []string{prescriptionitem.TakeHistoryItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(takehistoryitem.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piuo.mutation.TakeHistoryItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   prescriptionitem.TakeHistoryItemTable,
			Columns: []string{prescriptionitem.TakeHistoryItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(takehistoryitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PrescriptionItem{config: piuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, piuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prescriptionitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	piuo.mutation.done = true
	return _node, nil
}
