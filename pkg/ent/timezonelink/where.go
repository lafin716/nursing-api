// Code generated by ent, DO NOT EDIT.

package timezonelink

import (
	"nursing_api/pkg/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldLTE(FieldID, id))
}

// PrescriptionID applies equality check predicate on the "prescription_id" field. It's identical to PrescriptionIDEQ.
func PrescriptionID(v uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEQ(FieldPrescriptionID, v))
}

// TimezoneID applies equality check predicate on the "timezone_id" field. It's identical to TimezoneIDEQ.
func TimezoneID(v uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEQ(FieldTimezoneID, v))
}

// TimezoneName applies equality check predicate on the "timezone_name" field. It's identical to TimezoneNameEQ.
func TimezoneName(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEQ(FieldTimezoneName, v))
}

// UseAlert applies equality check predicate on the "use_alert" field. It's identical to UseAlertEQ.
func UseAlert(v bool) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEQ(FieldUseAlert, v))
}

// Midday applies equality check predicate on the "midday" field. It's identical to MiddayEQ.
func Midday(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEQ(FieldMidday, v))
}

// Hour applies equality check predicate on the "hour" field. It's identical to HourEQ.
func Hour(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEQ(FieldHour, v))
}

// Minute applies equality check predicate on the "minute" field. It's identical to MinuteEQ.
func Minute(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEQ(FieldMinute, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEQ(FieldUpdatedAt, v))
}

// PrescriptionIDEQ applies the EQ predicate on the "prescription_id" field.
func PrescriptionIDEQ(v uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEQ(FieldPrescriptionID, v))
}

// PrescriptionIDNEQ applies the NEQ predicate on the "prescription_id" field.
func PrescriptionIDNEQ(v uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldNEQ(FieldPrescriptionID, v))
}

// PrescriptionIDIn applies the In predicate on the "prescription_id" field.
func PrescriptionIDIn(vs ...uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldIn(FieldPrescriptionID, vs...))
}

// PrescriptionIDNotIn applies the NotIn predicate on the "prescription_id" field.
func PrescriptionIDNotIn(vs ...uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldNotIn(FieldPrescriptionID, vs...))
}

// PrescriptionIDGT applies the GT predicate on the "prescription_id" field.
func PrescriptionIDGT(v uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldGT(FieldPrescriptionID, v))
}

// PrescriptionIDGTE applies the GTE predicate on the "prescription_id" field.
func PrescriptionIDGTE(v uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldGTE(FieldPrescriptionID, v))
}

// PrescriptionIDLT applies the LT predicate on the "prescription_id" field.
func PrescriptionIDLT(v uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldLT(FieldPrescriptionID, v))
}

// PrescriptionIDLTE applies the LTE predicate on the "prescription_id" field.
func PrescriptionIDLTE(v uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldLTE(FieldPrescriptionID, v))
}

// TimezoneIDEQ applies the EQ predicate on the "timezone_id" field.
func TimezoneIDEQ(v uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEQ(FieldTimezoneID, v))
}

// TimezoneIDNEQ applies the NEQ predicate on the "timezone_id" field.
func TimezoneIDNEQ(v uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldNEQ(FieldTimezoneID, v))
}

// TimezoneIDIn applies the In predicate on the "timezone_id" field.
func TimezoneIDIn(vs ...uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldIn(FieldTimezoneID, vs...))
}

// TimezoneIDNotIn applies the NotIn predicate on the "timezone_id" field.
func TimezoneIDNotIn(vs ...uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldNotIn(FieldTimezoneID, vs...))
}

// TimezoneIDGT applies the GT predicate on the "timezone_id" field.
func TimezoneIDGT(v uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldGT(FieldTimezoneID, v))
}

// TimezoneIDGTE applies the GTE predicate on the "timezone_id" field.
func TimezoneIDGTE(v uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldGTE(FieldTimezoneID, v))
}

// TimezoneIDLT applies the LT predicate on the "timezone_id" field.
func TimezoneIDLT(v uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldLT(FieldTimezoneID, v))
}

// TimezoneIDLTE applies the LTE predicate on the "timezone_id" field.
func TimezoneIDLTE(v uuid.UUID) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldLTE(FieldTimezoneID, v))
}

// TimezoneNameEQ applies the EQ predicate on the "timezone_name" field.
func TimezoneNameEQ(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEQ(FieldTimezoneName, v))
}

// TimezoneNameNEQ applies the NEQ predicate on the "timezone_name" field.
func TimezoneNameNEQ(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldNEQ(FieldTimezoneName, v))
}

// TimezoneNameIn applies the In predicate on the "timezone_name" field.
func TimezoneNameIn(vs ...string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldIn(FieldTimezoneName, vs...))
}

// TimezoneNameNotIn applies the NotIn predicate on the "timezone_name" field.
func TimezoneNameNotIn(vs ...string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldNotIn(FieldTimezoneName, vs...))
}

// TimezoneNameGT applies the GT predicate on the "timezone_name" field.
func TimezoneNameGT(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldGT(FieldTimezoneName, v))
}

// TimezoneNameGTE applies the GTE predicate on the "timezone_name" field.
func TimezoneNameGTE(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldGTE(FieldTimezoneName, v))
}

// TimezoneNameLT applies the LT predicate on the "timezone_name" field.
func TimezoneNameLT(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldLT(FieldTimezoneName, v))
}

// TimezoneNameLTE applies the LTE predicate on the "timezone_name" field.
func TimezoneNameLTE(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldLTE(FieldTimezoneName, v))
}

// TimezoneNameContains applies the Contains predicate on the "timezone_name" field.
func TimezoneNameContains(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldContains(FieldTimezoneName, v))
}

// TimezoneNameHasPrefix applies the HasPrefix predicate on the "timezone_name" field.
func TimezoneNameHasPrefix(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldHasPrefix(FieldTimezoneName, v))
}

// TimezoneNameHasSuffix applies the HasSuffix predicate on the "timezone_name" field.
func TimezoneNameHasSuffix(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldHasSuffix(FieldTimezoneName, v))
}

// TimezoneNameIsNil applies the IsNil predicate on the "timezone_name" field.
func TimezoneNameIsNil() predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldIsNull(FieldTimezoneName))
}

// TimezoneNameNotNil applies the NotNil predicate on the "timezone_name" field.
func TimezoneNameNotNil() predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldNotNull(FieldTimezoneName))
}

// TimezoneNameEqualFold applies the EqualFold predicate on the "timezone_name" field.
func TimezoneNameEqualFold(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEqualFold(FieldTimezoneName, v))
}

// TimezoneNameContainsFold applies the ContainsFold predicate on the "timezone_name" field.
func TimezoneNameContainsFold(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldContainsFold(FieldTimezoneName, v))
}

// UseAlertEQ applies the EQ predicate on the "use_alert" field.
func UseAlertEQ(v bool) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEQ(FieldUseAlert, v))
}

// UseAlertNEQ applies the NEQ predicate on the "use_alert" field.
func UseAlertNEQ(v bool) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldNEQ(FieldUseAlert, v))
}

// MiddayEQ applies the EQ predicate on the "midday" field.
func MiddayEQ(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEQ(FieldMidday, v))
}

// MiddayNEQ applies the NEQ predicate on the "midday" field.
func MiddayNEQ(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldNEQ(FieldMidday, v))
}

// MiddayIn applies the In predicate on the "midday" field.
func MiddayIn(vs ...string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldIn(FieldMidday, vs...))
}

// MiddayNotIn applies the NotIn predicate on the "midday" field.
func MiddayNotIn(vs ...string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldNotIn(FieldMidday, vs...))
}

// MiddayGT applies the GT predicate on the "midday" field.
func MiddayGT(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldGT(FieldMidday, v))
}

// MiddayGTE applies the GTE predicate on the "midday" field.
func MiddayGTE(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldGTE(FieldMidday, v))
}

// MiddayLT applies the LT predicate on the "midday" field.
func MiddayLT(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldLT(FieldMidday, v))
}

// MiddayLTE applies the LTE predicate on the "midday" field.
func MiddayLTE(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldLTE(FieldMidday, v))
}

// MiddayContains applies the Contains predicate on the "midday" field.
func MiddayContains(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldContains(FieldMidday, v))
}

// MiddayHasPrefix applies the HasPrefix predicate on the "midday" field.
func MiddayHasPrefix(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldHasPrefix(FieldMidday, v))
}

// MiddayHasSuffix applies the HasSuffix predicate on the "midday" field.
func MiddayHasSuffix(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldHasSuffix(FieldMidday, v))
}

// MiddayEqualFold applies the EqualFold predicate on the "midday" field.
func MiddayEqualFold(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEqualFold(FieldMidday, v))
}

// MiddayContainsFold applies the ContainsFold predicate on the "midday" field.
func MiddayContainsFold(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldContainsFold(FieldMidday, v))
}

// HourEQ applies the EQ predicate on the "hour" field.
func HourEQ(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEQ(FieldHour, v))
}

// HourNEQ applies the NEQ predicate on the "hour" field.
func HourNEQ(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldNEQ(FieldHour, v))
}

// HourIn applies the In predicate on the "hour" field.
func HourIn(vs ...string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldIn(FieldHour, vs...))
}

// HourNotIn applies the NotIn predicate on the "hour" field.
func HourNotIn(vs ...string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldNotIn(FieldHour, vs...))
}

// HourGT applies the GT predicate on the "hour" field.
func HourGT(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldGT(FieldHour, v))
}

// HourGTE applies the GTE predicate on the "hour" field.
func HourGTE(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldGTE(FieldHour, v))
}

// HourLT applies the LT predicate on the "hour" field.
func HourLT(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldLT(FieldHour, v))
}

// HourLTE applies the LTE predicate on the "hour" field.
func HourLTE(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldLTE(FieldHour, v))
}

// HourContains applies the Contains predicate on the "hour" field.
func HourContains(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldContains(FieldHour, v))
}

// HourHasPrefix applies the HasPrefix predicate on the "hour" field.
func HourHasPrefix(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldHasPrefix(FieldHour, v))
}

// HourHasSuffix applies the HasSuffix predicate on the "hour" field.
func HourHasSuffix(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldHasSuffix(FieldHour, v))
}

// HourEqualFold applies the EqualFold predicate on the "hour" field.
func HourEqualFold(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEqualFold(FieldHour, v))
}

// HourContainsFold applies the ContainsFold predicate on the "hour" field.
func HourContainsFold(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldContainsFold(FieldHour, v))
}

// MinuteEQ applies the EQ predicate on the "minute" field.
func MinuteEQ(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEQ(FieldMinute, v))
}

// MinuteNEQ applies the NEQ predicate on the "minute" field.
func MinuteNEQ(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldNEQ(FieldMinute, v))
}

// MinuteIn applies the In predicate on the "minute" field.
func MinuteIn(vs ...string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldIn(FieldMinute, vs...))
}

// MinuteNotIn applies the NotIn predicate on the "minute" field.
func MinuteNotIn(vs ...string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldNotIn(FieldMinute, vs...))
}

// MinuteGT applies the GT predicate on the "minute" field.
func MinuteGT(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldGT(FieldMinute, v))
}

// MinuteGTE applies the GTE predicate on the "minute" field.
func MinuteGTE(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldGTE(FieldMinute, v))
}

// MinuteLT applies the LT predicate on the "minute" field.
func MinuteLT(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldLT(FieldMinute, v))
}

// MinuteLTE applies the LTE predicate on the "minute" field.
func MinuteLTE(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldLTE(FieldMinute, v))
}

// MinuteContains applies the Contains predicate on the "minute" field.
func MinuteContains(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldContains(FieldMinute, v))
}

// MinuteHasPrefix applies the HasPrefix predicate on the "minute" field.
func MinuteHasPrefix(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldHasPrefix(FieldMinute, v))
}

// MinuteHasSuffix applies the HasSuffix predicate on the "minute" field.
func MinuteHasSuffix(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldHasSuffix(FieldMinute, v))
}

// MinuteEqualFold applies the EqualFold predicate on the "minute" field.
func MinuteEqualFold(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEqualFold(FieldMinute, v))
}

// MinuteContainsFold applies the ContainsFold predicate on the "minute" field.
func MinuteContainsFold(v string) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldContainsFold(FieldMinute, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.FieldNotNull(FieldUpdatedAt))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TimeZoneLink) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TimeZoneLink) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TimeZoneLink) predicate.TimeZoneLink {
	return predicate.TimeZoneLink(sql.NotPredicates(p))
}
