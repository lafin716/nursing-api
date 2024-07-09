// Code generated by ent, DO NOT EDIT.

package prescription

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
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
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgePrescriptionItems holds the string denoting the prescription_items edge name in mutations.
	EdgePrescriptionItems = "prescription_items"
	// Table holds the table name of the prescription in the database.
	Table = "prescriptions"
	// PrescriptionItemsTable is the table that holds the prescription_items relation/edge.
	PrescriptionItemsTable = "prescription_items"
	// PrescriptionItemsInverseTable is the table name for the PrescriptionItem entity.
	// It exists in this package in order to avoid circular dependency with the "prescriptionitem" package.
	PrescriptionItemsInverseTable = "prescription_items"
	// PrescriptionItemsColumn is the table column denoting the prescription_items relation/edge.
	PrescriptionItemsColumn = "prescription_id"
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

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByPrescriptionItemsCount orders the results by prescription_items count.
func ByPrescriptionItemsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPrescriptionItemsStep(), opts...)
	}
}

// ByPrescriptionItems orders the results by prescription_items terms.
func ByPrescriptionItems(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPrescriptionItemsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPrescriptionItemsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PrescriptionItemsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PrescriptionItemsTable, PrescriptionItemsColumn),
	)
}
