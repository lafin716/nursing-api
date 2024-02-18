package list

import (
	"errors"
	"strings"
)

type list[T any] struct {
	Data []T
}

type List[T any] interface {
	Add(data T) error
	AddForce(data T)
	Pop() (T, error)
	Push(data T) error
	PushForce(data T)
	Insert(index uint, data T) error
	InsertForce(index uint, data T)
	Convert(converter func(data any) []T)
	Serialize() string
	Size() uint
	ISize() uint
}

func NewList[T any]() List[T] {
	return &list[T]{
		Data: make([]T, 0),
	}
}

func (l list[T]) Serialize() string {
	serials := []string{}
	return strings.Join(serials, ",")
}

func (l list[T]) Add(data T) error {
	l.Data = append(l.Data, data)
	return nil
}

func (l list[T]) AddForce(data T) {
	l.Data = append(l.Data, data)
}

func (l list[T]) Pop() (T, error) {
	return l.Data[0], nil
}

func (l list[T]) Push(data T) error {
	l.Data = append([]T{data}, l.Data...)
	return nil
}

func (l list[T]) PushForce(data T) {
	err := l.Push(data)
	if err != nil {
		return
	}
}

func (l list[T]) Insert(index uint, data T) error {
	if index > l.Size() {
		return errors.New("index cannot be bigger than size")
	}
	l.Data = append(l.Data[:index], append([]T{data}, l.Data[index:]...)...)
	return nil
}

func (l list[T]) InsertForce(index uint, data T) {
	err := l.Insert(index, data)
	if err != nil {
		return
	}
}

func (l list[T]) Convert(converter func(data any) []T) {
	//TODO implement me
	panic("implement me")
}

func (l list[T]) Size() uint {
	return uint(len(l.Data))
}

func (l list[T]) ISize() uint {
	return uint(len(l.Data) - 1)
}
