package date

import (
	"fmt"
	"testing"
)

func TestReplaceLayout(t *testing.T) {
	//layout := "Y-m-d"
	layout := MakeLayout("-", YEAR, MONTH, DATE)
	expected := fmt.Sprintf("%s-%s-%s", YEAR, MONTH, DATE)
	newLayout := ReplaceLayout(layout)

	t.Logf("layout : %s", layout)
	t.Logf("expected : %s", expected)
	t.Logf("newLayout : %s", newLayout)

	if newLayout != expected {
		t.Fatal("날짜 변환 실패")
	}
}

func TestValidate(t *testing.T) {
	layout := MakeLayout("-", YEAR, MONTH, DATE)
	newLayout := ReplaceLayout(layout)

	t.Logf("layout : %s", layout)
	t.Logf("newLayout : %s", newLayout)

	parser := NewParser()
	result, err := parser.Parse(newLayout, "")
	if err == nil {
		t.Fatal("have to be an error ")
	}

	t.Logf("result : %s", result)
}

func TestParse(t *testing.T) {
	layout := MakeLayout("-", YEAR, MONTH, DATE)
	newLayout := ReplaceLayout(layout)

	t.Logf("layout : %s", layout)
	t.Logf("newLayout : %s", newLayout)

	parser := NewParser()
	result, err := parser.Parse(newLayout, "2024-01-01")
	if err != nil {
		t.Fatalf("parse error : %v", err)
	}

	t.Logf("result : %s", result)
}
