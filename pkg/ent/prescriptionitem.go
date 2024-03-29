// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"nursing_api/pkg/ent/prescriptionitem"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// PrescriptionItem is the model entity for the PrescriptionItem schema.
type PrescriptionItem struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// TimezoneLinkID holds the value of the "timezone_link_id" field.
	TimezoneLinkID uuid.UUID `json:"timezone_link_id,omitempty"`
	// MedicineID holds the value of the "medicine_id" field.
	MedicineID uuid.UUID `json:"medicine_id,omitempty"`
	// MedicineName holds the value of the "medicine_name" field.
	MedicineName string `json:"medicine_name,omitempty"`
	// TakeAmount holds the value of the "take_amount" field.
	TakeAmount float64 `json:"take_amount,omitempty"`
	// MedicineUnit holds the value of the "medicine_unit" field.
	MedicineUnit string `json:"medicine_unit,omitempty"`
	// Memo holds the value of the "memo" field.
	Memo string `json:"memo,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PrescriptionItem) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case prescriptionitem.FieldTakeAmount:
			values[i] = new(sql.NullFloat64)
		case prescriptionitem.FieldMedicineName, prescriptionitem.FieldMedicineUnit, prescriptionitem.FieldMemo:
			values[i] = new(sql.NullString)
		case prescriptionitem.FieldCreatedAt, prescriptionitem.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case prescriptionitem.FieldID, prescriptionitem.FieldTimezoneLinkID, prescriptionitem.FieldMedicineID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PrescriptionItem fields.
func (pi *PrescriptionItem) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case prescriptionitem.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pi.ID = *value
			}
		case prescriptionitem.FieldTimezoneLinkID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field timezone_link_id", values[i])
			} else if value != nil {
				pi.TimezoneLinkID = *value
			}
		case prescriptionitem.FieldMedicineID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field medicine_id", values[i])
			} else if value != nil {
				pi.MedicineID = *value
			}
		case prescriptionitem.FieldMedicineName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field medicine_name", values[i])
			} else if value.Valid {
				pi.MedicineName = value.String
			}
		case prescriptionitem.FieldTakeAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field take_amount", values[i])
			} else if value.Valid {
				pi.TakeAmount = value.Float64
			}
		case prescriptionitem.FieldMedicineUnit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field medicine_unit", values[i])
			} else if value.Valid {
				pi.MedicineUnit = value.String
			}
		case prescriptionitem.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				pi.Memo = value.String
			}
		case prescriptionitem.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pi.CreatedAt = value.Time
			}
		case prescriptionitem.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pi.UpdatedAt = value.Time
			}
		default:
			pi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PrescriptionItem.
// This includes values selected through modifiers, order, etc.
func (pi *PrescriptionItem) Value(name string) (ent.Value, error) {
	return pi.selectValues.Get(name)
}

// Update returns a builder for updating this PrescriptionItem.
// Note that you need to call PrescriptionItem.Unwrap() before calling this method if this PrescriptionItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (pi *PrescriptionItem) Update() *PrescriptionItemUpdateOne {
	return NewPrescriptionItemClient(pi.config).UpdateOne(pi)
}

// Unwrap unwraps the PrescriptionItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pi *PrescriptionItem) Unwrap() *PrescriptionItem {
	_tx, ok := pi.config.driver.(*txDriver)
	if !ok {
		panic("ent: PrescriptionItem is not a transactional entity")
	}
	pi.config.driver = _tx.drv
	return pi
}

// String implements the fmt.Stringer.
func (pi *PrescriptionItem) String() string {
	var builder strings.Builder
	builder.WriteString("PrescriptionItem(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pi.ID))
	builder.WriteString("timezone_link_id=")
	builder.WriteString(fmt.Sprintf("%v", pi.TimezoneLinkID))
	builder.WriteString(", ")
	builder.WriteString("medicine_id=")
	builder.WriteString(fmt.Sprintf("%v", pi.MedicineID))
	builder.WriteString(", ")
	builder.WriteString("medicine_name=")
	builder.WriteString(pi.MedicineName)
	builder.WriteString(", ")
	builder.WriteString("take_amount=")
	builder.WriteString(fmt.Sprintf("%v", pi.TakeAmount))
	builder.WriteString(", ")
	builder.WriteString("medicine_unit=")
	builder.WriteString(pi.MedicineUnit)
	builder.WriteString(", ")
	builder.WriteString("memo=")
	builder.WriteString(pi.Memo)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pi.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pi.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// PrescriptionItems is a parsable slice of PrescriptionItem.
type PrescriptionItems []*PrescriptionItem
