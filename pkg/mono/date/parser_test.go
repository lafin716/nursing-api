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
