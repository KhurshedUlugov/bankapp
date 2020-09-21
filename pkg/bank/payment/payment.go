package payment

import(
	//"fmt"
	"bank/pkg/bank/types"
)

func Max(payments []types.Payment) types.Payment{
	var	max types.Payment
	max.Amount = 0
	for _, payment := range payments{
		if max.Amount < payment.Amount{
			max = payment
		}
	}
	return max
}
