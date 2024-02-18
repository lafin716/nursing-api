// Code generated by ent, DO NOT EDIT.

package takehistory

import (
	"nursing_api/pkg/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldEQ(FieldUserID, v))
}

// PrescriptionID applies equality check predicate on the "prescription_id" field. It's identical to PrescriptionIDEQ.
func PrescriptionID(v uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldEQ(FieldPrescriptionID, v))
}

// TimezoneID applies equality check predicate on the "timezone_id" field. It's identical to TimezoneIDEQ.
func TimezoneID(v uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldEQ(FieldTimezoneID, v))
}

// TakeDate applies equality check predicate on the "take_date" field. It's identical to TakeDateEQ.
func TakeDate(v time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldEQ(FieldTakeDate, v))
}

// TakeStatus applies equality check predicate on the "take_status" field. It's identical to TakeStatusEQ.
func TakeStatus(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldEQ(FieldTakeStatus, v))
}

// Memo applies equality check predicate on the "memo" field. It's identical to MemoEQ.
func Memo(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldEQ(FieldMemo, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldLTE(FieldUserID, v))
}

// PrescriptionIDEQ applies the EQ predicate on the "prescription_id" field.
func PrescriptionIDEQ(v uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldEQ(FieldPrescriptionID, v))
}

// PrescriptionIDNEQ applies the NEQ predicate on the "prescription_id" field.
func PrescriptionIDNEQ(v uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldNEQ(FieldPrescriptionID, v))
}

// PrescriptionIDIn applies the In predicate on the "prescription_id" field.
func PrescriptionIDIn(vs ...uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldIn(FieldPrescriptionID, vs...))
}

// PrescriptionIDNotIn applies the NotIn predicate on the "prescription_id" field.
func PrescriptionIDNotIn(vs ...uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldNotIn(FieldPrescriptionID, vs...))
}

// PrescriptionIDGT applies the GT predicate on the "prescription_id" field.
func PrescriptionIDGT(v uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldGT(FieldPrescriptionID, v))
}

// PrescriptionIDGTE applies the GTE predicate on the "prescription_id" field.
func PrescriptionIDGTE(v uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldGTE(FieldPrescriptionID, v))
}

// PrescriptionIDLT applies the LT predicate on the "prescription_id" field.
func PrescriptionIDLT(v uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldLT(FieldPrescriptionID, v))
}

// PrescriptionIDLTE applies the LTE predicate on the "prescription_id" field.
func PrescriptionIDLTE(v uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldLTE(FieldPrescriptionID, v))
}

// TimezoneIDEQ applies the EQ predicate on the "timezone_id" field.
func TimezoneIDEQ(v uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldEQ(FieldTimezoneID, v))
}

// TimezoneIDNEQ applies the NEQ predicate on the "timezone_id" field.
func TimezoneIDNEQ(v uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldNEQ(FieldTimezoneID, v))
}

// TimezoneIDIn applies the In predicate on the "timezone_id" field.
func TimezoneIDIn(vs ...uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldIn(FieldTimezoneID, vs...))
}

// TimezoneIDNotIn applies the NotIn predicate on the "timezone_id" field.
func TimezoneIDNotIn(vs ...uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldNotIn(FieldTimezoneID, vs...))
}

// TimezoneIDGT applies the GT predicate on the "timezone_id" field.
func TimezoneIDGT(v uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldGT(FieldTimezoneID, v))
}

// TimezoneIDGTE applies the GTE predicate on the "timezone_id" field.
func TimezoneIDGTE(v uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldGTE(FieldTimezoneID, v))
}

// TimezoneIDLT applies the LT predicate on the "timezone_id" field.
func TimezoneIDLT(v uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldLT(FieldTimezoneID, v))
}

// TimezoneIDLTE applies the LTE predicate on the "timezone_id" field.
func TimezoneIDLTE(v uuid.UUID) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldLTE(FieldTimezoneID, v))
}

// TakeDateEQ applies the EQ predicate on the "take_date" field.
func TakeDateEQ(v time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldEQ(FieldTakeDate, v))
}

// TakeDateNEQ applies the NEQ predicate on the "take_date" field.
func TakeDateNEQ(v time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldNEQ(FieldTakeDate, v))
}

// TakeDateIn applies the In predicate on the "take_date" field.
func TakeDateIn(vs ...time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldIn(FieldTakeDate, vs...))
}

// TakeDateNotIn applies the NotIn predicate on the "take_date" field.
func TakeDateNotIn(vs ...time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldNotIn(FieldTakeDate, vs...))
}

// TakeDateGT applies the GT predicate on the "take_date" field.
func TakeDateGT(v time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldGT(FieldTakeDate, v))
}

// TakeDateGTE applies the GTE predicate on the "take_date" field.
func TakeDateGTE(v time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldGTE(FieldTakeDate, v))
}

// TakeDateLT applies the LT predicate on the "take_date" field.
func TakeDateLT(v time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldLT(FieldTakeDate, v))
}

// TakeDateLTE applies the LTE predicate on the "take_date" field.
func TakeDateLTE(v time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldLTE(FieldTakeDate, v))
}

// TakeStatusEQ applies the EQ predicate on the "take_status" field.
func TakeStatusEQ(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldEQ(FieldTakeStatus, v))
}

// TakeStatusNEQ applies the NEQ predicate on the "take_status" field.
func TakeStatusNEQ(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldNEQ(FieldTakeStatus, v))
}

// TakeStatusIn applies the In predicate on the "take_status" field.
func TakeStatusIn(vs ...string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldIn(FieldTakeStatus, vs...))
}

// TakeStatusNotIn applies the NotIn predicate on the "take_status" field.
func TakeStatusNotIn(vs ...string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldNotIn(FieldTakeStatus, vs...))
}

// TakeStatusGT applies the GT predicate on the "take_status" field.
func TakeStatusGT(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldGT(FieldTakeStatus, v))
}

// TakeStatusGTE applies the GTE predicate on the "take_status" field.
func TakeStatusGTE(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldGTE(FieldTakeStatus, v))
}

// TakeStatusLT applies the LT predicate on the "take_status" field.
func TakeStatusLT(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldLT(FieldTakeStatus, v))
}

// TakeStatusLTE applies the LTE predicate on the "take_status" field.
func TakeStatusLTE(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldLTE(FieldTakeStatus, v))
}

// TakeStatusContains applies the Contains predicate on the "take_status" field.
func TakeStatusContains(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldContains(FieldTakeStatus, v))
}

// TakeStatusHasPrefix applies the HasPrefix predicate on the "take_status" field.
func TakeStatusHasPrefix(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldHasPrefix(FieldTakeStatus, v))
}

// TakeStatusHasSuffix applies the HasSuffix predicate on the "take_status" field.
func TakeStatusHasSuffix(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldHasSuffix(FieldTakeStatus, v))
}

// TakeStatusIsNil applies the IsNil predicate on the "take_status" field.
func TakeStatusIsNil() predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldIsNull(FieldTakeStatus))
}

// TakeStatusNotNil applies the NotNil predicate on the "take_status" field.
func TakeStatusNotNil() predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldNotNull(FieldTakeStatus))
}

// TakeStatusEqualFold applies the EqualFold predicate on the "take_status" field.
func TakeStatusEqualFold(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldEqualFold(FieldTakeStatus, v))
}

// TakeStatusContainsFold applies the ContainsFold predicate on the "take_status" field.
func TakeStatusContainsFold(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldContainsFold(FieldTakeStatus, v))
}

// MemoEQ applies the EQ predicate on the "memo" field.
func MemoEQ(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldEQ(FieldMemo, v))
}

// MemoNEQ applies the NEQ predicate on the "memo" field.
func MemoNEQ(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldNEQ(FieldMemo, v))
}

// MemoIn applies the In predicate on the "memo" field.
func MemoIn(vs ...string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldIn(FieldMemo, vs...))
}

// MemoNotIn applies the NotIn predicate on the "memo" field.
func MemoNotIn(vs ...string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldNotIn(FieldMemo, vs...))
}

// MemoGT applies the GT predicate on the "memo" field.
func MemoGT(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldGT(FieldMemo, v))
}

// MemoGTE applies the GTE predicate on the "memo" field.
func MemoGTE(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldGTE(FieldMemo, v))
}

// MemoLT applies the LT predicate on the "memo" field.
func MemoLT(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldLT(FieldMemo, v))
}

// MemoLTE applies the LTE predicate on the "memo" field.
func MemoLTE(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldLTE(FieldMemo, v))
}

// MemoContains applies the Contains predicate on the "memo" field.
func MemoContains(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldContains(FieldMemo, v))
}

// MemoHasPrefix applies the HasPrefix predicate on the "memo" field.
func MemoHasPrefix(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldHasPrefix(FieldMemo, v))
}

// MemoHasSuffix applies the HasSuffix predicate on the "memo" field.
func MemoHasSuffix(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldHasSuffix(FieldMemo, v))
}

// MemoIsNil applies the IsNil predicate on the "memo" field.
func MemoIsNil() predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldIsNull(FieldMemo))
}

// MemoNotNil applies the NotNil predicate on the "memo" field.
func MemoNotNil() predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldNotNull(FieldMemo))
}

// MemoEqualFold applies the EqualFold predicate on the "memo" field.
func MemoEqualFold(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldEqualFold(FieldMemo, v))
}

// MemoContainsFold applies the ContainsFold predicate on the "memo" field.
func MemoContainsFold(v string) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldContainsFold(FieldMemo, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.TakeHistory {
	return predicate.TakeHistory(sql.FieldNotNull(FieldUpdatedAt))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TakeHistory) predicate.TakeHistory {
	return predicate.TakeHistory(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TakeHistory) predicate.TakeHistory {
	return predicate.TakeHistory(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TakeHistory) predicate.TakeHistory {
	return predicate.TakeHistory(sql.NotPredicates(p))
}
