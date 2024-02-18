package list

import (
	"github.com/google/uuid"
	"testing"
)

func TestNewList(t *testing.T) {
	ids := NewList[int]()
	t.Logf("size : %d", ids.Size())
}

func TestList_Add(t *testing.T) {
	datas := NewList[uuid.UUID]()
	datas.Add(uuid.New())

	t.Logf("size: %d", datas.Size())
}
