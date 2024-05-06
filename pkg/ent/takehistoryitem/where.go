// Code generated by ent, DO NOT EDIT.

package takehistoryitem

import (
	"nursing_api/pkg/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldUserID, v))
}

// TakeHistoryID applies equality check predicate on the "take_history_id" field. It's identical to TakeHistoryIDEQ.
func TakeHistoryID(v uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldTakeHistoryID, v))
}

// PrescriptionItemID applies equality check predicate on the "prescription_item_id" field. It's identical to PrescriptionItemIDEQ.
func PrescriptionItemID(v uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldPrescriptionItemID, v))
}

// TakeStatus applies equality check predicate on the "take_status" field. It's identical to TakeStatusEQ.
func TakeStatus(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldTakeStatus, v))
}

// TakeAmount applies equality check predicate on the "take_amount" field. It's identical to TakeAmountEQ.
func TakeAmount(v float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldTakeAmount, v))
}

// RemainAmount applies equality check predicate on the "remain_amount" field. It's identical to RemainAmountEQ.
func RemainAmount(v float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldRemainAmount, v))
}

// TotalAmount applies equality check predicate on the "total_amount" field. It's identical to TotalAmountEQ.
func TotalAmount(v float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldTotalAmount, v))
}

// TakeUnit applies equality check predicate on the "take_unit" field. It's identical to TakeUnitEQ.
func TakeUnit(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldTakeUnit, v))
}

// Memo applies equality check predicate on the "memo" field. It's identical to MemoEQ.
func Memo(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldMemo, v))
}

// TakeDate applies equality check predicate on the "take_date" field. It's identical to TakeDateEQ.
func TakeDate(v time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldTakeDate, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLTE(FieldUserID, v))
}

// TakeHistoryIDEQ applies the EQ predicate on the "take_history_id" field.
func TakeHistoryIDEQ(v uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldTakeHistoryID, v))
}

// TakeHistoryIDNEQ applies the NEQ predicate on the "take_history_id" field.
func TakeHistoryIDNEQ(v uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNEQ(FieldTakeHistoryID, v))
}

// TakeHistoryIDIn applies the In predicate on the "take_history_id" field.
func TakeHistoryIDIn(vs ...uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldIn(FieldTakeHistoryID, vs...))
}

// TakeHistoryIDNotIn applies the NotIn predicate on the "take_history_id" field.
func TakeHistoryIDNotIn(vs ...uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNotIn(FieldTakeHistoryID, vs...))
}

// TakeHistoryIDGT applies the GT predicate on the "take_history_id" field.
func TakeHistoryIDGT(v uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGT(FieldTakeHistoryID, v))
}

// TakeHistoryIDGTE applies the GTE predicate on the "take_history_id" field.
func TakeHistoryIDGTE(v uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGTE(FieldTakeHistoryID, v))
}

// TakeHistoryIDLT applies the LT predicate on the "take_history_id" field.
func TakeHistoryIDLT(v uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLT(FieldTakeHistoryID, v))
}

// TakeHistoryIDLTE applies the LTE predicate on the "take_history_id" field.
func TakeHistoryIDLTE(v uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLTE(FieldTakeHistoryID, v))
}

// PrescriptionItemIDEQ applies the EQ predicate on the "prescription_item_id" field.
func PrescriptionItemIDEQ(v uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldPrescriptionItemID, v))
}

// PrescriptionItemIDNEQ applies the NEQ predicate on the "prescription_item_id" field.
func PrescriptionItemIDNEQ(v uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNEQ(FieldPrescriptionItemID, v))
}

// PrescriptionItemIDIn applies the In predicate on the "prescription_item_id" field.
func PrescriptionItemIDIn(vs ...uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldIn(FieldPrescriptionItemID, vs...))
}

// PrescriptionItemIDNotIn applies the NotIn predicate on the "prescription_item_id" field.
func PrescriptionItemIDNotIn(vs ...uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNotIn(FieldPrescriptionItemID, vs...))
}

// PrescriptionItemIDGT applies the GT predicate on the "prescription_item_id" field.
func PrescriptionItemIDGT(v uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGT(FieldPrescriptionItemID, v))
}

// PrescriptionItemIDGTE applies the GTE predicate on the "prescription_item_id" field.
func PrescriptionItemIDGTE(v uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGTE(FieldPrescriptionItemID, v))
}

// PrescriptionItemIDLT applies the LT predicate on the "prescription_item_id" field.
func PrescriptionItemIDLT(v uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLT(FieldPrescriptionItemID, v))
}

// PrescriptionItemIDLTE applies the LTE predicate on the "prescription_item_id" field.
func PrescriptionItemIDLTE(v uuid.UUID) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLTE(FieldPrescriptionItemID, v))
}

// TakeStatusEQ applies the EQ predicate on the "take_status" field.
func TakeStatusEQ(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldTakeStatus, v))
}

// TakeStatusNEQ applies the NEQ predicate on the "take_status" field.
func TakeStatusNEQ(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNEQ(FieldTakeStatus, v))
}

// TakeStatusIn applies the In predicate on the "take_status" field.
func TakeStatusIn(vs ...string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldIn(FieldTakeStatus, vs...))
}

// TakeStatusNotIn applies the NotIn predicate on the "take_status" field.
func TakeStatusNotIn(vs ...string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNotIn(FieldTakeStatus, vs...))
}

// TakeStatusGT applies the GT predicate on the "take_status" field.
func TakeStatusGT(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGT(FieldTakeStatus, v))
}

// TakeStatusGTE applies the GTE predicate on the "take_status" field.
func TakeStatusGTE(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGTE(FieldTakeStatus, v))
}

// TakeStatusLT applies the LT predicate on the "take_status" field.
func TakeStatusLT(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLT(FieldTakeStatus, v))
}

// TakeStatusLTE applies the LTE predicate on the "take_status" field.
func TakeStatusLTE(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLTE(FieldTakeStatus, v))
}

// TakeStatusContains applies the Contains predicate on the "take_status" field.
func TakeStatusContains(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldContains(FieldTakeStatus, v))
}

// TakeStatusHasPrefix applies the HasPrefix predicate on the "take_status" field.
func TakeStatusHasPrefix(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldHasPrefix(FieldTakeStatus, v))
}

// TakeStatusHasSuffix applies the HasSuffix predicate on the "take_status" field.
func TakeStatusHasSuffix(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldHasSuffix(FieldTakeStatus, v))
}

// TakeStatusEqualFold applies the EqualFold predicate on the "take_status" field.
func TakeStatusEqualFold(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEqualFold(FieldTakeStatus, v))
}

// TakeStatusContainsFold applies the ContainsFold predicate on the "take_status" field.
func TakeStatusContainsFold(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldContainsFold(FieldTakeStatus, v))
}

// TakeAmountEQ applies the EQ predicate on the "take_amount" field.
func TakeAmountEQ(v float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldTakeAmount, v))
}

// TakeAmountNEQ applies the NEQ predicate on the "take_amount" field.
func TakeAmountNEQ(v float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNEQ(FieldTakeAmount, v))
}

// TakeAmountIn applies the In predicate on the "take_amount" field.
func TakeAmountIn(vs ...float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldIn(FieldTakeAmount, vs...))
}

// TakeAmountNotIn applies the NotIn predicate on the "take_amount" field.
func TakeAmountNotIn(vs ...float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNotIn(FieldTakeAmount, vs...))
}

// TakeAmountGT applies the GT predicate on the "take_amount" field.
func TakeAmountGT(v float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGT(FieldTakeAmount, v))
}

// TakeAmountGTE applies the GTE predicate on the "take_amount" field.
func TakeAmountGTE(v float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGTE(FieldTakeAmount, v))
}

// TakeAmountLT applies the LT predicate on the "take_amount" field.
func TakeAmountLT(v float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLT(FieldTakeAmount, v))
}

// TakeAmountLTE applies the LTE predicate on the "take_amount" field.
func TakeAmountLTE(v float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLTE(FieldTakeAmount, v))
}

// RemainAmountEQ applies the EQ predicate on the "remain_amount" field.
func RemainAmountEQ(v float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldRemainAmount, v))
}

// RemainAmountNEQ applies the NEQ predicate on the "remain_amount" field.
func RemainAmountNEQ(v float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNEQ(FieldRemainAmount, v))
}

// RemainAmountIn applies the In predicate on the "remain_amount" field.
func RemainAmountIn(vs ...float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldIn(FieldRemainAmount, vs...))
}

// RemainAmountNotIn applies the NotIn predicate on the "remain_amount" field.
func RemainAmountNotIn(vs ...float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNotIn(FieldRemainAmount, vs...))
}

// RemainAmountGT applies the GT predicate on the "remain_amount" field.
func RemainAmountGT(v float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGT(FieldRemainAmount, v))
}

// RemainAmountGTE applies the GTE predicate on the "remain_amount" field.
func RemainAmountGTE(v float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGTE(FieldRemainAmount, v))
}

// RemainAmountLT applies the LT predicate on the "remain_amount" field.
func RemainAmountLT(v float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLT(FieldRemainAmount, v))
}

// RemainAmountLTE applies the LTE predicate on the "remain_amount" field.
func RemainAmountLTE(v float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLTE(FieldRemainAmount, v))
}

// TotalAmountEQ applies the EQ predicate on the "total_amount" field.
func TotalAmountEQ(v float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldTotalAmount, v))
}

// TotalAmountNEQ applies the NEQ predicate on the "total_amount" field.
func TotalAmountNEQ(v float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNEQ(FieldTotalAmount, v))
}

// TotalAmountIn applies the In predicate on the "total_amount" field.
func TotalAmountIn(vs ...float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldIn(FieldTotalAmount, vs...))
}

// TotalAmountNotIn applies the NotIn predicate on the "total_amount" field.
func TotalAmountNotIn(vs ...float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNotIn(FieldTotalAmount, vs...))
}

// TotalAmountGT applies the GT predicate on the "total_amount" field.
func TotalAmountGT(v float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGT(FieldTotalAmount, v))
}

// TotalAmountGTE applies the GTE predicate on the "total_amount" field.
func TotalAmountGTE(v float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGTE(FieldTotalAmount, v))
}

// TotalAmountLT applies the LT predicate on the "total_amount" field.
func TotalAmountLT(v float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLT(FieldTotalAmount, v))
}

// TotalAmountLTE applies the LTE predicate on the "total_amount" field.
func TotalAmountLTE(v float64) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLTE(FieldTotalAmount, v))
}

// TakeUnitEQ applies the EQ predicate on the "take_unit" field.
func TakeUnitEQ(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldTakeUnit, v))
}

// TakeUnitNEQ applies the NEQ predicate on the "take_unit" field.
func TakeUnitNEQ(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNEQ(FieldTakeUnit, v))
}

// TakeUnitIn applies the In predicate on the "take_unit" field.
func TakeUnitIn(vs ...string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldIn(FieldTakeUnit, vs...))
}

// TakeUnitNotIn applies the NotIn predicate on the "take_unit" field.
func TakeUnitNotIn(vs ...string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNotIn(FieldTakeUnit, vs...))
}

// TakeUnitGT applies the GT predicate on the "take_unit" field.
func TakeUnitGT(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGT(FieldTakeUnit, v))
}

// TakeUnitGTE applies the GTE predicate on the "take_unit" field.
func TakeUnitGTE(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGTE(FieldTakeUnit, v))
}

// TakeUnitLT applies the LT predicate on the "take_unit" field.
func TakeUnitLT(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLT(FieldTakeUnit, v))
}

// TakeUnitLTE applies the LTE predicate on the "take_unit" field.
func TakeUnitLTE(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLTE(FieldTakeUnit, v))
}

// TakeUnitContains applies the Contains predicate on the "take_unit" field.
func TakeUnitContains(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldContains(FieldTakeUnit, v))
}

// TakeUnitHasPrefix applies the HasPrefix predicate on the "take_unit" field.
func TakeUnitHasPrefix(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldHasPrefix(FieldTakeUnit, v))
}

// TakeUnitHasSuffix applies the HasSuffix predicate on the "take_unit" field.
func TakeUnitHasSuffix(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldHasSuffix(FieldTakeUnit, v))
}

// TakeUnitEqualFold applies the EqualFold predicate on the "take_unit" field.
func TakeUnitEqualFold(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEqualFold(FieldTakeUnit, v))
}

// TakeUnitContainsFold applies the ContainsFold predicate on the "take_unit" field.
func TakeUnitContainsFold(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldContainsFold(FieldTakeUnit, v))
}

// MemoEQ applies the EQ predicate on the "memo" field.
func MemoEQ(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldMemo, v))
}

// MemoNEQ applies the NEQ predicate on the "memo" field.
func MemoNEQ(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNEQ(FieldMemo, v))
}

// MemoIn applies the In predicate on the "memo" field.
func MemoIn(vs ...string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldIn(FieldMemo, vs...))
}

// MemoNotIn applies the NotIn predicate on the "memo" field.
func MemoNotIn(vs ...string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNotIn(FieldMemo, vs...))
}

// MemoGT applies the GT predicate on the "memo" field.
func MemoGT(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGT(FieldMemo, v))
}

// MemoGTE applies the GTE predicate on the "memo" field.
func MemoGTE(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGTE(FieldMemo, v))
}

// MemoLT applies the LT predicate on the "memo" field.
func MemoLT(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLT(FieldMemo, v))
}

// MemoLTE applies the LTE predicate on the "memo" field.
func MemoLTE(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLTE(FieldMemo, v))
}

// MemoContains applies the Contains predicate on the "memo" field.
func MemoContains(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldContains(FieldMemo, v))
}

// MemoHasPrefix applies the HasPrefix predicate on the "memo" field.
func MemoHasPrefix(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldHasPrefix(FieldMemo, v))
}

// MemoHasSuffix applies the HasSuffix predicate on the "memo" field.
func MemoHasSuffix(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldHasSuffix(FieldMemo, v))
}

// MemoIsNil applies the IsNil predicate on the "memo" field.
func MemoIsNil() predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldIsNull(FieldMemo))
}

// MemoNotNil applies the NotNil predicate on the "memo" field.
func MemoNotNil() predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNotNull(FieldMemo))
}

// MemoEqualFold applies the EqualFold predicate on the "memo" field.
func MemoEqualFold(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEqualFold(FieldMemo, v))
}

// MemoContainsFold applies the ContainsFold predicate on the "memo" field.
func MemoContainsFold(v string) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldContainsFold(FieldMemo, v))
}

// TakeDateEQ applies the EQ predicate on the "take_date" field.
func TakeDateEQ(v time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldTakeDate, v))
}

// TakeDateNEQ applies the NEQ predicate on the "take_date" field.
func TakeDateNEQ(v time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNEQ(FieldTakeDate, v))
}

// TakeDateIn applies the In predicate on the "take_date" field.
func TakeDateIn(vs ...time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldIn(FieldTakeDate, vs...))
}

// TakeDateNotIn applies the NotIn predicate on the "take_date" field.
func TakeDateNotIn(vs ...time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNotIn(FieldTakeDate, vs...))
}

// TakeDateGT applies the GT predicate on the "take_date" field.
func TakeDateGT(v time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGT(FieldTakeDate, v))
}

// TakeDateGTE applies the GTE predicate on the "take_date" field.
func TakeDateGTE(v time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGTE(FieldTakeDate, v))
}

// TakeDateLT applies the LT predicate on the "take_date" field.
func TakeDateLT(v time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLT(FieldTakeDate, v))
}

// TakeDateLTE applies the LTE predicate on the "take_date" field.
func TakeDateLTE(v time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLTE(FieldTakeDate, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.FieldNotNull(FieldUpdatedAt))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TakeHistoryItem) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TakeHistoryItem) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TakeHistoryItem) predicate.TakeHistoryItem {
	return predicate.TakeHistoryItem(sql.NotPredicates(p))
}
