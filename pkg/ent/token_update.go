// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"nursing_api/pkg/ent/predicate"
	"nursing_api/pkg/ent/token"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TokenUpdate is the builder for updating Token entities.
type TokenUpdate struct {
	config
	hooks    []Hook
	mutation *TokenMutation
}

// Where appends a list predicates to the TokenUpdate builder.
func (tu *TokenUpdate) Where(ps ...predicate.Token) *TokenUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUserID sets the "user_id" field.
func (tu *TokenUpdate) SetUserID(u uuid.UUID) *TokenUpdate {
	tu.mutation.SetUserID(u)
	return tu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableUserID(u *uuid.UUID) *TokenUpdate {
	if u != nil {
		tu.SetUserID(*u)
	}
	return tu
}

// SetAccessToken sets the "access_token" field.
func (tu *TokenUpdate) SetAccessToken(s string) *TokenUpdate {
	tu.mutation.SetAccessToken(s)
	return tu
}

// SetNillableAccessToken sets the "access_token" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableAccessToken(s *string) *TokenUpdate {
	if s != nil {
		tu.SetAccessToken(*s)
	}
	return tu
}

// SetRefreshToken sets the "refresh_token" field.
func (tu *TokenUpdate) SetRefreshToken(s string) *TokenUpdate {
	tu.mutation.SetRefreshToken(s)
	return tu
}

// SetNillableRefreshToken sets the "refresh_token" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableRefreshToken(s *string) *TokenUpdate {
	if s != nil {
		tu.SetRefreshToken(*s)
	}
	return tu
}

// SetAccessTokenExpires sets the "access_token_expires" field.
func (tu *TokenUpdate) SetAccessTokenExpires(t time.Time) *TokenUpdate {
	tu.mutation.SetAccessTokenExpires(t)
	return tu
}

// SetNillableAccessTokenExpires sets the "access_token_expires" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableAccessTokenExpires(t *time.Time) *TokenUpdate {
	if t != nil {
		tu.SetAccessTokenExpires(*t)
	}
	return tu
}

// SetRefreshTokenExpires sets the "refresh_token_expires" field.
func (tu *TokenUpdate) SetRefreshTokenExpires(t time.Time) *TokenUpdate {
	tu.mutation.SetRefreshTokenExpires(t)
	return tu
}

// SetNillableRefreshTokenExpires sets the "refresh_token_expires" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableRefreshTokenExpires(t *time.Time) *TokenUpdate {
	if t != nil {
		tu.SetRefreshTokenExpires(*t)
	}
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TokenUpdate) SetCreatedAt(t time.Time) *TokenUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableCreatedAt(t *time.Time) *TokenUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TokenUpdate) SetUpdatedAt(t time.Time) *TokenUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableUpdatedAt(t *time.Time) *TokenUpdate {
	if t != nil {
		tu.SetUpdatedAt(*t)
	}
	return tu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tu *TokenUpdate) ClearUpdatedAt() *TokenUpdate {
	tu.mutation.ClearUpdatedAt()
	return tu
}

// Mutation returns the TokenMutation object of the builder.
func (tu *TokenUpdate) Mutation() *TokenMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TokenUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TokenUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TokenUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TokenUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(token.Table, token.Columns, sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UserID(); ok {
		_spec.SetField(token.FieldUserID, field.TypeUUID, value)
	}
	if value, ok := tu.mutation.AccessToken(); ok {
		_spec.SetField(token.FieldAccessToken, field.TypeString, value)
	}
	if value, ok := tu.mutation.RefreshToken(); ok {
		_spec.SetField(token.FieldRefreshToken, field.TypeString, value)
	}
	if value, ok := tu.mutation.AccessTokenExpires(); ok {
		_spec.SetField(token.FieldAccessTokenExpires, field.TypeTime, value)
	}
	if value, ok := tu.mutation.RefreshTokenExpires(); ok {
		_spec.SetField(token.FieldRefreshTokenExpires, field.TypeTime, value)
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.SetField(token.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(token.FieldUpdatedAt, field.TypeTime, value)
	}
	if tu.mutation.UpdatedAtCleared() {
		_spec.ClearField(token.FieldUpdatedAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{token.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TokenUpdateOne is the builder for updating a single Token entity.
type TokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TokenMutation
}

// SetUserID sets the "user_id" field.
func (tuo *TokenUpdateOne) SetUserID(u uuid.UUID) *TokenUpdateOne {
	tuo.mutation.SetUserID(u)
	return tuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableUserID(u *uuid.UUID) *TokenUpdateOne {
	if u != nil {
		tuo.SetUserID(*u)
	}
	return tuo
}

// SetAccessToken sets the "access_token" field.
func (tuo *TokenUpdateOne) SetAccessToken(s string) *TokenUpdateOne {
	tuo.mutation.SetAccessToken(s)
	return tuo
}

// SetNillableAccessToken sets the "access_token" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableAccessToken(s *string) *TokenUpdateOne {
	if s != nil {
		tuo.SetAccessToken(*s)
	}
	return tuo
}

// SetRefreshToken sets the "refresh_token" field.
func (tuo *TokenUpdateOne) SetRefreshToken(s string) *TokenUpdateOne {
	tuo.mutation.SetRefreshToken(s)
	return tuo
}

// SetNillableRefreshToken sets the "refresh_token" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableRefreshToken(s *string) *TokenUpdateOne {
	if s != nil {
		tuo.SetRefreshToken(*s)
	}
	return tuo
}

// SetAccessTokenExpires sets the "access_token_expires" field.
func (tuo *TokenUpdateOne) SetAccessTokenExpires(t time.Time) *TokenUpdateOne {
	tuo.mutation.SetAccessTokenExpires(t)
	return tuo
}

// SetNillableAccessTokenExpires sets the "access_token_expires" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableAccessTokenExpires(t *time.Time) *TokenUpdateOne {
	if t != nil {
		tuo.SetAccessTokenExpires(*t)
	}
	return tuo
}

// SetRefreshTokenExpires sets the "refresh_token_expires" field.
func (tuo *TokenUpdateOne) SetRefreshTokenExpires(t time.Time) *TokenUpdateOne {
	tuo.mutation.SetRefreshTokenExpires(t)
	return tuo
}

// SetNillableRefreshTokenExpires sets the "refresh_token_expires" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableRefreshTokenExpires(t *time.Time) *TokenUpdateOne {
	if t != nil {
		tuo.SetRefreshTokenExpires(*t)
	}
	return tuo
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TokenUpdateOne) SetCreatedAt(t time.Time) *TokenUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableCreatedAt(t *time.Time) *TokenUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TokenUpdateOne) SetUpdatedAt(t time.Time) *TokenUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableUpdatedAt(t *time.Time) *TokenUpdateOne {
	if t != nil {
		tuo.SetUpdatedAt(*t)
	}
	return tuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tuo *TokenUpdateOne) ClearUpdatedAt() *TokenUpdateOne {
	tuo.mutation.ClearUpdatedAt()
	return tuo
}

// Mutation returns the TokenMutation object of the builder.
func (tuo *TokenUpdateOne) Mutation() *TokenMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TokenUpdate builder.
func (tuo *TokenUpdateOne) Where(ps ...predicate.Token) *TokenUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TokenUpdateOne) Select(field string, fields ...string) *TokenUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Token entity.
func (tuo *TokenUpdateOne) Save(ctx context.Context) (*Token, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TokenUpdateOne) SaveX(ctx context.Context) *Token {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TokenUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TokenUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TokenUpdateOne) sqlSave(ctx context.Context) (_node *Token, err error) {
	_spec := sqlgraph.NewUpdateSpec(token.Table, token.Columns, sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Token.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, token.FieldID)
		for _, f := range fields {
			if !token.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != token.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UserID(); ok {
		_spec.SetField(token.FieldUserID, field.TypeUUID, value)
	}
	if value, ok := tuo.mutation.AccessToken(); ok {
		_spec.SetField(token.FieldAccessToken, field.TypeString, value)
	}
	if value, ok := tuo.mutation.RefreshToken(); ok {
		_spec.SetField(token.FieldRefreshToken, field.TypeString, value)
	}
	if value, ok := tuo.mutation.AccessTokenExpires(); ok {
		_spec.SetField(token.FieldAccessTokenExpires, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.RefreshTokenExpires(); ok {
		_spec.SetField(token.FieldRefreshTokenExpires, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.SetField(token.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(token.FieldUpdatedAt, field.TypeTime, value)
	}
	if tuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(token.FieldUpdatedAt, field.TypeTime)
	}
	_node = &Token{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{token.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
