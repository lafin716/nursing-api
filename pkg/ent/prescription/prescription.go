// Code generated by ent, DO NOT EDIT.

package prescription

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the prescription type in the database.
	Label = "prescription"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldPrescriptionName holds the string denoting the prescription_name field in the database.
	FieldPrescriptionName = "prescription_name"
	// FieldHospitalName holds the string denoting the hospital_name field in the database.
	FieldHospitalName = "hospital_name"
	// FieldTakeDays holds the string denoting the take_days field in the database.
	FieldTakeDays = "take_days"
	// FieldStartedAt holds the string denoting the started_at field in the database.
	FieldStartedAt = "started_at"
	// FieldFinishedAt holds the string denoting the finished_at field in the database.
	FieldFinishedAt = "finished_at"
	// FieldMemo holds the string denoting the memo field in the database.
	FieldMemo = "memo"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the prescription in the database.
	Table = "prescriptions"
)

// Columns holds all SQL columns for prescription fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldPrescriptionName,
	FieldHospitalName,
	FieldTakeDays,
	FieldStartedAt,
	FieldFinishedAt,
	FieldMemo,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultTakeDays holds the default value on creation for the "take_days" field.
	DefaultTakeDays int
	// DefaultStartedAt holds the default value on creation for the "started_at" field.
	DefaultStartedAt func() time.Time
	// DefaultFinishedAt holds the default value on creation for the "finished_at" field.
	DefaultFinishedAt func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Prescription queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByPrescriptionName orders the results by the prescription_name field.
func ByPrescriptionName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrescriptionName, opts...).ToFunc()
}

// ByHospitalName orders the results by the hospital_name field.
func ByHospitalName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHospitalName, opts...).ToFunc()
}

// ByTakeDays orders the results by the take_days field.
func ByTakeDays(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTakeDays, opts...).ToFunc()
}

// ByStartedAt orders the results by the started_at field.
func ByStartedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartedAt, opts...).ToFunc()
}

// ByFinishedAt orders the results by the finished_at field.
func ByFinishedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFinishedAt, opts...).ToFunc()
}

// ByMemo orders the results by the memo field.
func ByMemo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemo, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}