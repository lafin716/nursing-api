package database

import (
	"context"
	"errors"
	"nursing_api/pkg/ent"
)

type TransactionManager interface {
	Register(tx Transactional) error
	IsRunningTransaction() bool
	GetTx() *ent.Tx
	BeginTx() error
	CommitTx() error
	RollbackTx() error
	EndTx()
}

type Transactional interface {
	BeginTx() error
	CommitTx() error
	RollbackTx() error
}

type transactionManager struct {
	client *ent.Client
	ctx    context.Context
	tx     *ent.Tx
	txList []Transactional
}

func (t *transactionManager) Register(tx Transactional) error {
	t.txList = append(t.txList, tx)
	return nil
}

func (t *transactionManager) IsRunningTransaction() bool {
	return t.tx != nil
}

func (t *transactionManager) GetTx() *ent.Tx {
	return t.tx
}

func (t *transactionManager) BeginTx() error {
	if t.tx != nil {
		return errors.New("이미 실행중인 트랜잭션이 있습니다")
	}
	tx, err := t.client.Tx(t.ctx)
	if err != nil {
		return err
	}
	t.tx = tx
	return nil
}

func (t *transactionManager) CommitTx() error {
	defer t.EndTx()
	if t.tx == nil {
		return errors.New("트랜잭션이 시작되지 않았습니다")
	}
	err := t.tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (t *transactionManager) RollbackTx() error {
	defer t.EndTx()
	if t.tx == nil {
		return errors.New("트랜잭션이 시작되지 않았습니다")
	}
	err := t.tx.Rollback()
	if err != nil {
		return err
	}
	return nil
}

func (t *transactionManager) EndTx() {
	if t.tx != nil {
		t.tx = nil
	}
}

func NewTransactionManager(
	client *ent.Client,
	ctx context.Context,
) TransactionManager {
	return &transactionManager{
		client: client,
		ctx:    ctx,
		tx:     nil,
		txList: make([]Transactional, 0),
	}
}
