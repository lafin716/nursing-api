// Code generated by ent, DO NOT EDIT.

package user

import (
	"nursing_api/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// UserEmail applies equality check predicate on the "user_email" field. It's identical to UserEmailEQ.
func UserEmail(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserEmail, v))
}

// UserPassword applies equality check predicate on the "user_password" field. It's identical to UserPasswordEQ.
func UserPassword(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserPassword, v))
}

// UserStatus applies equality check predicate on the "user_status" field. It's identical to UserStatusEQ.
func UserStatus(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserStatus, v))
}

// UserType applies equality check predicate on the "user_type" field. It's identical to UserTypeEQ.
func UserType(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserType, v))
}

// FailCount applies equality check predicate on the "fail_count" field. It's identical to FailCountEQ.
func FailCount(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFailCount, v))
}

// LastSignedIn applies equality check predicate on the "last_signed_in" field. It's identical to LastSignedInEQ.
func LastSignedIn(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastSignedIn, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserEmailEQ applies the EQ predicate on the "user_email" field.
func UserEmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserEmail, v))
}

// UserEmailNEQ applies the NEQ predicate on the "user_email" field.
func UserEmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUserEmail, v))
}

// UserEmailIn applies the In predicate on the "user_email" field.
func UserEmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUserEmail, vs...))
}

// UserEmailNotIn applies the NotIn predicate on the "user_email" field.
func UserEmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUserEmail, vs...))
}

// UserEmailGT applies the GT predicate on the "user_email" field.
func UserEmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUserEmail, v))
}

// UserEmailGTE applies the GTE predicate on the "user_email" field.
func UserEmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUserEmail, v))
}

// UserEmailLT applies the LT predicate on the "user_email" field.
func UserEmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUserEmail, v))
}

// UserEmailLTE applies the LTE predicate on the "user_email" field.
func UserEmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUserEmail, v))
}

// UserEmailContains applies the Contains predicate on the "user_email" field.
func UserEmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUserEmail, v))
}

// UserEmailHasPrefix applies the HasPrefix predicate on the "user_email" field.
func UserEmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUserEmail, v))
}

// UserEmailHasSuffix applies the HasSuffix predicate on the "user_email" field.
func UserEmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUserEmail, v))
}

// UserEmailEqualFold applies the EqualFold predicate on the "user_email" field.
func UserEmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUserEmail, v))
}

// UserEmailContainsFold applies the ContainsFold predicate on the "user_email" field.
func UserEmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUserEmail, v))
}

// UserPasswordEQ applies the EQ predicate on the "user_password" field.
func UserPasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserPassword, v))
}

// UserPasswordNEQ applies the NEQ predicate on the "user_password" field.
func UserPasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUserPassword, v))
}

// UserPasswordIn applies the In predicate on the "user_password" field.
func UserPasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUserPassword, vs...))
}

// UserPasswordNotIn applies the NotIn predicate on the "user_password" field.
func UserPasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUserPassword, vs...))
}

// UserPasswordGT applies the GT predicate on the "user_password" field.
func UserPasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUserPassword, v))
}

// UserPasswordGTE applies the GTE predicate on the "user_password" field.
func UserPasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUserPassword, v))
}

// UserPasswordLT applies the LT predicate on the "user_password" field.
func UserPasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUserPassword, v))
}

// UserPasswordLTE applies the LTE predicate on the "user_password" field.
func UserPasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUserPassword, v))
}

// UserPasswordContains applies the Contains predicate on the "user_password" field.
func UserPasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUserPassword, v))
}

// UserPasswordHasPrefix applies the HasPrefix predicate on the "user_password" field.
func UserPasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUserPassword, v))
}

// UserPasswordHasSuffix applies the HasSuffix predicate on the "user_password" field.
func UserPasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUserPassword, v))
}

// UserPasswordEqualFold applies the EqualFold predicate on the "user_password" field.
func UserPasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUserPassword, v))
}

// UserPasswordContainsFold applies the ContainsFold predicate on the "user_password" field.
func UserPasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUserPassword, v))
}

// UserStatusEQ applies the EQ predicate on the "user_status" field.
func UserStatusEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserStatus, v))
}

// UserStatusNEQ applies the NEQ predicate on the "user_status" field.
func UserStatusNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUserStatus, v))
}

// UserStatusIn applies the In predicate on the "user_status" field.
func UserStatusIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUserStatus, vs...))
}

// UserStatusNotIn applies the NotIn predicate on the "user_status" field.
func UserStatusNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUserStatus, vs...))
}

// UserStatusGT applies the GT predicate on the "user_status" field.
func UserStatusGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUserStatus, v))
}

// UserStatusGTE applies the GTE predicate on the "user_status" field.
func UserStatusGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUserStatus, v))
}

// UserStatusLT applies the LT predicate on the "user_status" field.
func UserStatusLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUserStatus, v))
}

// UserStatusLTE applies the LTE predicate on the "user_status" field.
func UserStatusLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUserStatus, v))
}

// UserStatusContains applies the Contains predicate on the "user_status" field.
func UserStatusContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUserStatus, v))
}

// UserStatusHasPrefix applies the HasPrefix predicate on the "user_status" field.
func UserStatusHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUserStatus, v))
}

// UserStatusHasSuffix applies the HasSuffix predicate on the "user_status" field.
func UserStatusHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUserStatus, v))
}

// UserStatusEqualFold applies the EqualFold predicate on the "user_status" field.
func UserStatusEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUserStatus, v))
}

// UserStatusContainsFold applies the ContainsFold predicate on the "user_status" field.
func UserStatusContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUserStatus, v))
}

// UserTypeEQ applies the EQ predicate on the "user_type" field.
func UserTypeEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserType, v))
}

// UserTypeNEQ applies the NEQ predicate on the "user_type" field.
func UserTypeNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUserType, v))
}

// UserTypeIn applies the In predicate on the "user_type" field.
func UserTypeIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUserType, vs...))
}

// UserTypeNotIn applies the NotIn predicate on the "user_type" field.
func UserTypeNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUserType, vs...))
}

// UserTypeGT applies the GT predicate on the "user_type" field.
func UserTypeGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUserType, v))
}

// UserTypeGTE applies the GTE predicate on the "user_type" field.
func UserTypeGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUserType, v))
}

// UserTypeLT applies the LT predicate on the "user_type" field.
func UserTypeLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUserType, v))
}

// UserTypeLTE applies the LTE predicate on the "user_type" field.
func UserTypeLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUserType, v))
}

// UserTypeContains applies the Contains predicate on the "user_type" field.
func UserTypeContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUserType, v))
}

// UserTypeHasPrefix applies the HasPrefix predicate on the "user_type" field.
func UserTypeHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUserType, v))
}

// UserTypeHasSuffix applies the HasSuffix predicate on the "user_type" field.
func UserTypeHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUserType, v))
}

// UserTypeEqualFold applies the EqualFold predicate on the "user_type" field.
func UserTypeEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUserType, v))
}

// UserTypeContainsFold applies the ContainsFold predicate on the "user_type" field.
func UserTypeContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUserType, v))
}

// FailCountEQ applies the EQ predicate on the "fail_count" field.
func FailCountEQ(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFailCount, v))
}

// FailCountNEQ applies the NEQ predicate on the "fail_count" field.
func FailCountNEQ(v int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldFailCount, v))
}

// FailCountIn applies the In predicate on the "fail_count" field.
func FailCountIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldFailCount, vs...))
}

// FailCountNotIn applies the NotIn predicate on the "fail_count" field.
func FailCountNotIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldFailCount, vs...))
}

// FailCountGT applies the GT predicate on the "fail_count" field.
func FailCountGT(v int) predicate.User {
	return predicate.User(sql.FieldGT(FieldFailCount, v))
}

// FailCountGTE applies the GTE predicate on the "fail_count" field.
func FailCountGTE(v int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldFailCount, v))
}

// FailCountLT applies the LT predicate on the "fail_count" field.
func FailCountLT(v int) predicate.User {
	return predicate.User(sql.FieldLT(FieldFailCount, v))
}

// FailCountLTE applies the LTE predicate on the "fail_count" field.
func FailCountLTE(v int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldFailCount, v))
}

// LastSignedInEQ applies the EQ predicate on the "last_signed_in" field.
func LastSignedInEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastSignedIn, v))
}

// LastSignedInNEQ applies the NEQ predicate on the "last_signed_in" field.
func LastSignedInNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldLastSignedIn, v))
}

// LastSignedInIn applies the In predicate on the "last_signed_in" field.
func LastSignedInIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldLastSignedIn, vs...))
}

// LastSignedInNotIn applies the NotIn predicate on the "last_signed_in" field.
func LastSignedInNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldLastSignedIn, vs...))
}

// LastSignedInGT applies the GT predicate on the "last_signed_in" field.
func LastSignedInGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldLastSignedIn, v))
}

// LastSignedInGTE applies the GTE predicate on the "last_signed_in" field.
func LastSignedInGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldLastSignedIn, v))
}

// LastSignedInLT applies the LT predicate on the "last_signed_in" field.
func LastSignedInLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldLastSignedIn, v))
}

// LastSignedInLTE applies the LTE predicate on the "last_signed_in" field.
func LastSignedInLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldLastSignedIn, v))
}

// LastSignedInIsNil applies the IsNil predicate on the "last_signed_in" field.
func LastSignedInIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldLastSignedIn))
}

// LastSignedInNotNil applies the NotNil predicate on the "last_signed_in" field.
func LastSignedInNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldLastSignedIn))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldUpdatedAt))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
