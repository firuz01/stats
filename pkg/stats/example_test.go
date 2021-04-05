package stats

import (
	"github.com/firuz01/bank/v2/pkg/types"
	"fmt"
	
)

func ExampleAvg()  {
	payments := []types.Payment {
		{
			ID: 1,
			Amount: 60,
			Category: "shop",
			Status: "FAIL",
		},{
			ID: 2,
			Amount: 60,
			Category: "shop",
			Status: "INPROGRESS",
		},{
			ID: 3,
			Amount: 60,
			Category: "shop",
			Status: "FAIL",
		},{
			ID: 3,
			Amount: 60,
			Category: "shop",
			Status: "OK",
		},{
			ID: 3,
			Amount: 60,
			Category: "shop",
			Status: "OK",
		},
	}
	fmt.Println(Avg(payments))
	// Output: 60
}
func ExampleTotalInCategory()  {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10_000_00,
			Category: "auto",
			Status: "FAIL",
		},
		{
			ID:       2,
			Amount:   20_000_00,
			Category: "auto",
			Status: "OK",
		},
		{
			ID:       3,
			Amount:   30_000_00,
			Category: "auto",
			Status: "INPROGRESS",
		},
		{
			ID:       3,
			Amount:   50_000_00,
			Category: "shop",
			Status: "FAIL",
		},
		{
			ID:       3,
			Amount:   10_000_00,
			Category: "shop",
			Status: "OK",
		},
	}

	inCategory := types.Category("auto")
	totalInCategory := TotalInCategory(payments, inCategory)
fmt.Println(totalInCategory)
// Output: 5000000
}