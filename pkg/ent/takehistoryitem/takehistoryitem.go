// Code generated by ent, DO NOT EDIT.

package takehistoryitem

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the takehistoryitem type in the database.
	Label = "take_history_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldTakeHistoryID holds the string denoting the take_history_id field in the database.
	FieldTakeHistoryID = "take_history_id"
	// FieldPrescriptionItemID holds the string denoting the prescription_item_id field in the database.
	FieldPrescriptionItemID = "prescription_item_id"
	// FieldTakeStatus holds the string denoting the take_status field in the database.
	FieldTakeStatus = "take_status"
	// FieldTakeAmount holds the string denoting the take_amount field in the database.
	FieldTakeAmount = "take_amount"
	// FieldTakeUnit holds the string denoting the take_unit field in the database.
	FieldTakeUnit = "take_unit"
	// FieldMemo holds the string denoting the memo field in the database.
	FieldMemo = "memo"
	// FieldTakeDate holds the string denoting the take_date field in the database.
	FieldTakeDate = "take_date"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the takehistoryitem in the database.
	Table = "take_history_items"
)

// Columns holds all SQL columns for takehistoryitem fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldTakeHistoryID,
	FieldPrescriptionItemID,
	FieldTakeStatus,
	FieldTakeAmount,
	FieldTakeUnit,
	FieldMemo,
	FieldTakeDate,
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
	// DefaultTakeStatus holds the default value on creation for the "take_status" field.
	DefaultTakeStatus string
	// DefaultTakeAmount holds the default value on creation for the "take_amount" field.
	DefaultTakeAmount float64
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the TakeHistoryItem queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByTakeHistoryID orders the results by the take_history_id field.
func ByTakeHistoryID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTakeHistoryID, opts...).ToFunc()
}

// ByPrescriptionItemID orders the results by the prescription_item_id field.
func ByPrescriptionItemID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrescriptionItemID, opts...).ToFunc()
}

// ByTakeStatus orders the results by the take_status field.
func ByTakeStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTakeStatus, opts...).ToFunc()
}

// ByTakeAmount orders the results by the take_amount field.
func ByTakeAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTakeAmount, opts...).ToFunc()
}

// ByTakeUnit orders the results by the take_unit field.
func ByTakeUnit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTakeUnit, opts...).ToFunc()
}

// ByMemo orders the results by the memo field.
func ByMemo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemo, opts...).ToFunc()
}

// ByTakeDate orders the results by the take_date field.
func ByTakeDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTakeDate, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}
