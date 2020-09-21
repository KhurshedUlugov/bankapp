package card

import(
	//"fmt"
	"bank/pkg/bank/types"
)

func IssueCard(currency types.Currency, color string, name string) types.Card {
	return types.Card {
	                ID:       1000,
	                PAN:      "5058 xxxx xxxx 0001",
                    Balance:  0,
	                Currency: currency,
	                Color:    color,
	                Name:     name,
	                Active:   true,
	}
}

func Withdraw(card types.Card, amount types.Money) types.Card {
	if (!card.Active){
		//fmt.Println("inactive. Balance: ", card.Balance)
	    return card
	}
	if (amount > card.Balance){
		//fmt.Println("no money. Balance: ", card.Balance)
	    return card
	}
	if (amount < 0){
		//fmt.Println("less zero. Balance: ", card.Balance)
	    return card
	}
	if (amount > 20_000_00){
		//fmt.Println("over limit. Balance: ", card.Balance)
	    return card
	}
	card.Balance = card.Balance - amount
	//fmt.Println("Bal: ",card.Balance)
    return card
}

func Deposit(card *types.Card, amount types.Money){
	if (!card.Active){
		//fmt.Println("inactive. Balance: ", card.Balance)
	    return
	}
	if (amount > 50_000_00){
		//fmt.Println("over limit. Balance: ", card.Balance)
	    return
	}
	if (amount < 0){
		//fmt.Println("over limit. Balance: ", card.Balance)
	    return
	}
	card.Balance = card.Balance + amount
	return
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int){
	if (!card.Active){
		//fmt.Println("inactive. Balance: ", card.Balance)
	    return
	}
	if (card.Balance <= 0){
		//fmt.Println("over limit. Balance: ", card.Balance)
	    return
	}
	bonus := int(card.MinBalance) * percent/100 * daysInMonth/daysInYear
	if (bonus > 5_000_00){
		return
	}

	card.Balance = card.Balance + types.Money(bonus)
}

func Total(cards []types.Card) types.Money{
	total := int64(0)
	for _, card := range cards{
		if (card.Balance > 0 && card.Active){
	        total += int64(card.Balance)
		}
	}
	return types.Money(total)
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var PaymentCards []types.PaymentSource
	for _, card := range cards {
			if card.Balance > 0 && card.Active {
					PaymentCards = append(PaymentCards, types.PaymentSource{
							Type:    `cart`,
							Number:  string(card.PAN),
							Balance: card.Balance,
					})
			}
	}
	return PaymentCards
}
