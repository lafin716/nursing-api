// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"nursing_api/pkg/ent/plantimezonelink"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// PlanTimeZoneLink is the model entity for the PlanTimeZoneLink schema.
type PlanTimeZoneLink struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// PrescriptionID holds the value of the "prescription_id" field.
	PrescriptionID uuid.UUID `json:"prescription_id,omitempty"`
	// TimezoneID holds the value of the "timezone_id" field.
	TimezoneID uuid.UUID `json:"timezone_id,omitempty"`
	// TimezoneName holds the value of the "timezone_name" field.
	TimezoneName string `json:"timezone_name,omitempty"`
	// UseAlert holds the value of the "use_alert" field.
	UseAlert bool `json:"use_alert,omitempty"`
	// Midday holds the value of the "midday" field.
	Midday string `json:"midday,omitempty"`
	// Hour holds the value of the "hour" field.
	Hour string `json:"hour,omitempty"`
	// Minute holds the value of the "minute" field.
	Minute string `json:"minute,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PlanTimeZoneLink) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case plantimezonelink.FieldUseAlert:
			values[i] = new(sql.NullBool)
		case plantimezonelink.FieldTimezoneName, plantimezonelink.FieldMidday, plantimezonelink.FieldHour, plantimezonelink.FieldMinute:
			values[i] = new(sql.NullString)
		case plantimezonelink.FieldCreatedAt, plantimezonelink.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case plantimezonelink.FieldID, plantimezonelink.FieldPrescriptionID, plantimezonelink.FieldTimezoneID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PlanTimeZoneLink fields.
func (ptzl *PlanTimeZoneLink) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case plantimezonelink.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ptzl.ID = *value
			}
		case plantimezonelink.FieldPrescriptionID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field prescription_id", values[i])
			} else if value != nil {
				ptzl.PrescriptionID = *value
			}
		case plantimezonelink.FieldTimezoneID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field timezone_id", values[i])
			} else if value != nil {
				ptzl.TimezoneID = *value
			}
		case plantimezonelink.FieldTimezoneName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field timezone_name", values[i])
			} else if value.Valid {
				ptzl.TimezoneName = value.String
			}
		case plantimezonelink.FieldUseAlert:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field use_alert", values[i])
			} else if value.Valid {
				ptzl.UseAlert = value.Bool
			}
		case plantimezonelink.FieldMidday:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field midday", values[i])
			} else if value.Valid {
				ptzl.Midday = value.String
			}
		case plantimezonelink.FieldHour:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hour", values[i])
			} else if value.Valid {
				ptzl.Hour = value.String
			}
		case plantimezonelink.FieldMinute:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field minute", values[i])
			} else if value.Valid {
				ptzl.Minute = value.String
			}
		case plantimezonelink.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ptzl.CreatedAt = value.Time
			}
		case plantimezonelink.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ptzl.UpdatedAt = value.Time
			}
		default:
			ptzl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PlanTimeZoneLink.
// This includes values selected through modifiers, order, etc.
func (ptzl *PlanTimeZoneLink) Value(name string) (ent.Value, error) {
	return ptzl.selectValues.Get(name)
}

// Update returns a builder for updating this PlanTimeZoneLink.
// Note that you need to call PlanTimeZoneLink.Unwrap() before calling this method if this PlanTimeZoneLink
// was returned from a transaction, and the transaction was committed or rolled back.
func (ptzl *PlanTimeZoneLink) Update() *PlanTimeZoneLinkUpdateOne {
	return NewPlanTimeZoneLinkClient(ptzl.config).UpdateOne(ptzl)
}

// Unwrap unwraps the PlanTimeZoneLink entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ptzl *PlanTimeZoneLink) Unwrap() *PlanTimeZoneLink {
	_tx, ok := ptzl.config.driver.(*txDriver)
	if !ok {
		panic("ent: PlanTimeZoneLink is not a transactional entity")
	}
	ptzl.config.driver = _tx.drv
	return ptzl
}

// String implements the fmt.Stringer.
func (ptzl *PlanTimeZoneLink) String() string {
	var builder strings.Builder
	builder.WriteString("PlanTimeZoneLink(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ptzl.ID))
	builder.WriteString("prescription_id=")
	builder.WriteString(fmt.Sprintf("%v", ptzl.PrescriptionID))
	builder.WriteString(", ")
	builder.WriteString("timezone_id=")
	builder.WriteString(fmt.Sprintf("%v", ptzl.TimezoneID))
	builder.WriteString(", ")
	builder.WriteString("timezone_name=")
	builder.WriteString(ptzl.TimezoneName)
	builder.WriteString(", ")
	builder.WriteString("use_alert=")
	builder.WriteString(fmt.Sprintf("%v", ptzl.UseAlert))
	builder.WriteString(", ")
	builder.WriteString("midday=")
	builder.WriteString(ptzl.Midday)
	builder.WriteString(", ")
	builder.WriteString("hour=")
	builder.WriteString(ptzl.Hour)
	builder.WriteString(", ")
	builder.WriteString("minute=")
	builder.WriteString(ptzl.Minute)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ptzl.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ptzl.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// PlanTimeZoneLinks is a parsable slice of PlanTimeZoneLink.
type PlanTimeZoneLinks []*PlanTimeZoneLink
