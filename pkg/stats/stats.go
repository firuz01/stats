package stats

import (
	"github.com/firuz01/bank/v2/pkg/types"
	

)

func Avg(payments []types.Payment) types.Money {
	numPayments := types.Money(0)
	sumPayments := types.Money(0)
	for _, payment := range payments {
		if payment.Status == types.StatusFail {	
			continue	
		} 
		sumPayments += payment.Amount
		numPayments ++
	}
	return sumPayments/numPayments
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {

		sumPayments := types.Money(0)
		for _, payment := range payments {
			if payment.Category != category || payment.Status == types.StatusFail {
				continue
			}
			sumPayments += payment.Amount
		}
		return sumPayments
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	totalinCategory := map[types.Category]types.Money{}
	numPayinCategory := map[types.Category]types.Money{}
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue		
		}
		totalinCategory[payment.Category] += payment.Amount
		numPayinCategory[payment.Category] ++
	}
	for money := range totalinCategory {
		totalinCategory[money] /= numPayinCategory[money]		
	}
	return totalinCategory
}