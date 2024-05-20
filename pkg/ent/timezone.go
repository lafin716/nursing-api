// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"nursing_api/pkg/ent/timezone"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// TimeZone is the model entity for the TimeZone schema.
type TimeZone struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// TimezoneName holds the value of the "timezone_name" field.
	TimezoneName string `json:"timezone_name,omitempty"`
	// IsDefault holds the value of the "is_default" field.
	IsDefault bool `json:"is_default,omitempty"`
	// Midday holds the value of the "midday" field.
	Midday string `json:"midday,omitempty"`
	// Hour holds the value of the "hour" field.
	Hour string `json:"hour,omitempty"`
	// Minute holds the value of the "minute" field.
	Minute string `json:"minute,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TimeZoneQuery when eager-loading is set.
	Edges        TimeZoneEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TimeZoneEdges holds the relations/edges for other nodes in the graph.
type TimeZoneEdges struct {
	// TakeHistory holds the value of the take_history edge.
	TakeHistory []*TakeHistory `json:"take_history,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TakeHistoryOrErr returns the TakeHistory value or an error if the edge
// was not loaded in eager-loading.
func (e TimeZoneEdges) TakeHistoryOrErr() ([]*TakeHistory, error) {
	if e.loadedTypes[0] {
		return e.TakeHistory, nil
	}
	return nil, &NotLoadedError{edge: "take_history"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TimeZone) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case timezone.FieldIsDefault:
			values[i] = new(sql.NullBool)
		case timezone.FieldTimezoneName, timezone.FieldMidday, timezone.FieldHour, timezone.FieldMinute:
			values[i] = new(sql.NullString)
		case timezone.FieldCreatedAt, timezone.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case timezone.FieldID, timezone.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TimeZone fields.
func (tz *TimeZone) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case timezone.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				tz.ID = *value
			}
		case timezone.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				tz.UserID = *value
			}
		case timezone.FieldTimezoneName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field timezone_name", values[i])
			} else if value.Valid {
				tz.TimezoneName = value.String
			}
		case timezone.FieldIsDefault:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_default", values[i])
			} else if value.Valid {
				tz.IsDefault = value.Bool
			}
		case timezone.FieldMidday:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field midday", values[i])
			} else if value.Valid {
				tz.Midday = value.String
			}
		case timezone.FieldHour:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hour", values[i])
			} else if value.Valid {
				tz.Hour = value.String
			}
		case timezone.FieldMinute:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field minute", values[i])
			} else if value.Valid {
				tz.Minute = value.String
			}
		case timezone.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tz.CreatedAt = value.Time
			}
		case timezone.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tz.UpdatedAt = value.Time
			}
		default:
			tz.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TimeZone.
// This includes values selected through modifiers, order, etc.
func (tz *TimeZone) Value(name string) (ent.Value, error) {
	return tz.selectValues.Get(name)
}

// QueryTakeHistory queries the "take_history" edge of the TimeZone entity.
func (tz *TimeZone) QueryTakeHistory() *TakeHistoryQuery {
	return NewTimeZoneClient(tz.config).QueryTakeHistory(tz)
}

// Update returns a builder for updating this TimeZone.
// Note that you need to call TimeZone.Unwrap() before calling this method if this TimeZone
// was returned from a transaction, and the transaction was committed or rolled back.
func (tz *TimeZone) Update() *TimeZoneUpdateOne {
	return NewTimeZoneClient(tz.config).UpdateOne(tz)
}

// Unwrap unwraps the TimeZone entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tz *TimeZone) Unwrap() *TimeZone {
	_tx, ok := tz.config.driver.(*txDriver)
	if !ok {
		panic("ent: TimeZone is not a transactional entity")
	}
	tz.config.driver = _tx.drv
	return tz
}

// String implements the fmt.Stringer.
func (tz *TimeZone) String() string {
	var builder strings.Builder
	builder.WriteString("TimeZone(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tz.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", tz.UserID))
	builder.WriteString(", ")
	builder.WriteString("timezone_name=")
	builder.WriteString(tz.TimezoneName)
	builder.WriteString(", ")
	builder.WriteString("is_default=")
	builder.WriteString(fmt.Sprintf("%v", tz.IsDefault))
	builder.WriteString(", ")
	builder.WriteString("midday=")
	builder.WriteString(tz.Midday)
	builder.WriteString(", ")
	builder.WriteString("hour=")
	builder.WriteString(tz.Hour)
	builder.WriteString(", ")
	builder.WriteString("minute=")
	builder.WriteString(tz.Minute)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(tz.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(tz.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TimeZones is a parsable slice of TimeZone.
type TimeZones []*TimeZone
