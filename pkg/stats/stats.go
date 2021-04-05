package stats

import (
	"github.com/firuz01/bank/v2/pkg/types"
	

)

func Avg(payments []types.Payment) types.Money {
	numPayments := types.Money(len(payments))
	sumPayments := types.Money(0)
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		} 
		moneyAmount := payment.Amount
		sumPayments += moneyAmount
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