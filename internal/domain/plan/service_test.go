package plan

import (
	"fmt"
	"nursing_api/pkg/mono"
	"testing"
)

func TestPlanService_Add(t *testing.T) {
	utl := mono.NewMono()
	str, err := utl.Date.Parse("Y-m-d", "2024-07-17")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	takeDays := 10
	for i := 0; i <= takeDays; i++ {
		nw := utl.Date.TruncateToDateAddDay(str, i)
		fmt.Println(nw)
	}
}
