// Code generated by ent, DO NOT EDIT.

package takehistorymemo

import (
	"nursing_api/pkg/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldEQ(FieldUserID, v))
}

// TimezoneID applies equality check predicate on the "timezone_id" field. It's identical to TimezoneIDEQ.
func TimezoneID(v uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldEQ(FieldTimezoneID, v))
}

// TakeDate applies equality check predicate on the "take_date" field. It's identical to TakeDateEQ.
func TakeDate(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldEQ(FieldTakeDate, v))
}

// Memo applies equality check predicate on the "memo" field. It's identical to MemoEQ.
func Memo(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldEQ(FieldMemo, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldLTE(FieldUserID, v))
}

// TimezoneIDEQ applies the EQ predicate on the "timezone_id" field.
func TimezoneIDEQ(v uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldEQ(FieldTimezoneID, v))
}

// TimezoneIDNEQ applies the NEQ predicate on the "timezone_id" field.
func TimezoneIDNEQ(v uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldNEQ(FieldTimezoneID, v))
}

// TimezoneIDIn applies the In predicate on the "timezone_id" field.
func TimezoneIDIn(vs ...uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldIn(FieldTimezoneID, vs...))
}

// TimezoneIDNotIn applies the NotIn predicate on the "timezone_id" field.
func TimezoneIDNotIn(vs ...uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldNotIn(FieldTimezoneID, vs...))
}

// TimezoneIDGT applies the GT predicate on the "timezone_id" field.
func TimezoneIDGT(v uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldGT(FieldTimezoneID, v))
}

// TimezoneIDGTE applies the GTE predicate on the "timezone_id" field.
func TimezoneIDGTE(v uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldGTE(FieldTimezoneID, v))
}

// TimezoneIDLT applies the LT predicate on the "timezone_id" field.
func TimezoneIDLT(v uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldLT(FieldTimezoneID, v))
}

// TimezoneIDLTE applies the LTE predicate on the "timezone_id" field.
func TimezoneIDLTE(v uuid.UUID) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldLTE(FieldTimezoneID, v))
}

// TakeDateEQ applies the EQ predicate on the "take_date" field.
func TakeDateEQ(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldEQ(FieldTakeDate, v))
}

// TakeDateNEQ applies the NEQ predicate on the "take_date" field.
func TakeDateNEQ(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldNEQ(FieldTakeDate, v))
}

// TakeDateIn applies the In predicate on the "take_date" field.
func TakeDateIn(vs ...string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldIn(FieldTakeDate, vs...))
}

// TakeDateNotIn applies the NotIn predicate on the "take_date" field.
func TakeDateNotIn(vs ...string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldNotIn(FieldTakeDate, vs...))
}

// TakeDateGT applies the GT predicate on the "take_date" field.
func TakeDateGT(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldGT(FieldTakeDate, v))
}

// TakeDateGTE applies the GTE predicate on the "take_date" field.
func TakeDateGTE(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldGTE(FieldTakeDate, v))
}

// TakeDateLT applies the LT predicate on the "take_date" field.
func TakeDateLT(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldLT(FieldTakeDate, v))
}

// TakeDateLTE applies the LTE predicate on the "take_date" field.
func TakeDateLTE(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldLTE(FieldTakeDate, v))
}

// TakeDateContains applies the Contains predicate on the "take_date" field.
func TakeDateContains(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldContains(FieldTakeDate, v))
}

// TakeDateHasPrefix applies the HasPrefix predicate on the "take_date" field.
func TakeDateHasPrefix(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldHasPrefix(FieldTakeDate, v))
}

// TakeDateHasSuffix applies the HasSuffix predicate on the "take_date" field.
func TakeDateHasSuffix(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldHasSuffix(FieldTakeDate, v))
}

// TakeDateEqualFold applies the EqualFold predicate on the "take_date" field.
func TakeDateEqualFold(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldEqualFold(FieldTakeDate, v))
}

// TakeDateContainsFold applies the ContainsFold predicate on the "take_date" field.
func TakeDateContainsFold(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldContainsFold(FieldTakeDate, v))
}

// MemoEQ applies the EQ predicate on the "memo" field.
func MemoEQ(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldEQ(FieldMemo, v))
}

// MemoNEQ applies the NEQ predicate on the "memo" field.
func MemoNEQ(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldNEQ(FieldMemo, v))
}

// MemoIn applies the In predicate on the "memo" field.
func MemoIn(vs ...string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldIn(FieldMemo, vs...))
}

// MemoNotIn applies the NotIn predicate on the "memo" field.
func MemoNotIn(vs ...string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldNotIn(FieldMemo, vs...))
}

// MemoGT applies the GT predicate on the "memo" field.
func MemoGT(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldGT(FieldMemo, v))
}

// MemoGTE applies the GTE predicate on the "memo" field.
func MemoGTE(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldGTE(FieldMemo, v))
}

// MemoLT applies the LT predicate on the "memo" field.
func MemoLT(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldLT(FieldMemo, v))
}

// MemoLTE applies the LTE predicate on the "memo" field.
func MemoLTE(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldLTE(FieldMemo, v))
}

// MemoContains applies the Contains predicate on the "memo" field.
func MemoContains(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldContains(FieldMemo, v))
}

// MemoHasPrefix applies the HasPrefix predicate on the "memo" field.
func MemoHasPrefix(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldHasPrefix(FieldMemo, v))
}

// MemoHasSuffix applies the HasSuffix predicate on the "memo" field.
func MemoHasSuffix(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldHasSuffix(FieldMemo, v))
}

// MemoIsNil applies the IsNil predicate on the "memo" field.
func MemoIsNil() predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldIsNull(FieldMemo))
}

// MemoNotNil applies the NotNil predicate on the "memo" field.
func MemoNotNil() predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldNotNull(FieldMemo))
}

// MemoEqualFold applies the EqualFold predicate on the "memo" field.
func MemoEqualFold(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldEqualFold(FieldMemo, v))
}

// MemoContainsFold applies the ContainsFold predicate on the "memo" field.
func MemoContainsFold(v string) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldContainsFold(FieldMemo, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.FieldNotNull(FieldUpdatedAt))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TakeHistoryMemo) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TakeHistoryMemo) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TakeHistoryMemo) predicate.TakeHistoryMemo {
	return predicate.TakeHistoryMemo(sql.NotPredicates(p))
}
