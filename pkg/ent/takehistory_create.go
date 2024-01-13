// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"nursing_api/pkg/ent/takehistory"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TakeHistoryCreate is the builder for creating a TakeHistory entity.
type TakeHistoryCreate struct {
	config
	mutation *TakeHistoryMutation
	hooks    []Hook
}

// Mutation returns the TakeHistoryMutation object of the builder.
func (thc *TakeHistoryCreate) Mutation() *TakeHistoryMutation {
	return thc.mutation
}

// Save creates the TakeHistory in the database.
func (thc *TakeHistoryCreate) Save(ctx context.Context) (*TakeHistory, error) {
	return withHooks(ctx, thc.sqlSave, thc.mutation, thc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (thc *TakeHistoryCreate) SaveX(ctx context.Context) *TakeHistory {
	v, err := thc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (thc *TakeHistoryCreate) Exec(ctx context.Context) error {
	_, err := thc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (thc *TakeHistoryCreate) ExecX(ctx context.Context) {
	if err := thc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (thc *TakeHistoryCreate) check() error {
	return nil
}

func (thc *TakeHistoryCreate) sqlSave(ctx context.Context) (*TakeHistory, error) {
	if err := thc.check(); err != nil {
		return nil, err
	}
	_node, _spec := thc.createSpec()
	if err := sqlgraph.CreateNode(ctx, thc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	thc.mutation.id = &_node.ID
	thc.mutation.done = true
	return _node, nil
}

func (thc *TakeHistoryCreate) createSpec() (*TakeHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &TakeHistory{config: thc.config}
		_spec = sqlgraph.NewCreateSpec(takehistory.Table, sqlgraph.NewFieldSpec(takehistory.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// TakeHistoryCreateBulk is the builder for creating many TakeHistory entities in bulk.
type TakeHistoryCreateBulk struct {
	config
	err      error
	builders []*TakeHistoryCreate
}

// Save creates the TakeHistory entities in the database.
func (thcb *TakeHistoryCreateBulk) Save(ctx context.Context) ([]*TakeHistory, error) {
	if thcb.err != nil {
		return nil, thcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(thcb.builders))
	nodes := make([]*TakeHistory, len(thcb.builders))
	mutators := make([]Mutator, len(thcb.builders))
	for i := range thcb.builders {
		func(i int, root context.Context) {
			builder := thcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TakeHistoryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, thcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, thcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, thcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (thcb *TakeHistoryCreateBulk) SaveX(ctx context.Context) []*TakeHistory {
	v, err := thcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (thcb *TakeHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := thcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (thcb *TakeHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := thcb.Exec(ctx); err != nil {
		panic(err)
	}
}