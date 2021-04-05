package stats

import (
	"github.com/firuz01/bank/pkg/types"
	"fmt"
	
)

func ExampleAvg()  {
	payments := []types.Payment {
		{
			ID: 1,
			Amount: 60,
			Category: "shop",
		},{
			ID: 2,
			Amount: 80,
			Category: "shop",

		},
	}
	fmt.Println(Avg(payments))
	// Output: 70
}
func ExampleTotalInCategory()  {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10_000_00,
			Category: "auto",
		},
		{
			ID:       2,
			Amount:   20_000_00,
			Category: "auto",
		},
		{
			ID:       3,
			Amount:   30_000_00,
			Category: "auto",
		},
		{
			ID:       3,
			Amount:   50_000_00,
			Category: "shop",
		},
		{
			ID:       3,
			Amount:   10_000_00,
			Category: "shop",
		},
	}

	inCategory := types.Category("auto")
	totalInCategory := TotalInCategory(payments, inCategory)
fmt.Println(totalInCategory)
// Output: 6000000
}