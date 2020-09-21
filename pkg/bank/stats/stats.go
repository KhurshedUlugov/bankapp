package stats

import (
	"github.com/shohinsherov/bank/pkg/bank/types"
)
// Avg рассчитает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	n := 0
	for _, pay := range payments {
		n++;
		sum += pay.Amount;
	}
	return sum/types.Money(n)
}