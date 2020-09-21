package card

import (
		"fmt"
		"bank/pkg/bank/types"
		)

func ExampleWithdraw_positive(){
	result := Withdraw(types.Card{Balance:20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance);
    
    // Output: 1000000
}

func ExampleWithdraw_noMoney(){
	result := Withdraw(types.Card{Balance:10_000_00, Active: true}, 20_000_00)
	fmt.Println(result.Balance);
    
    // Output: 1000000
}

func ExampleWithdraw_inactive(){
	result := Withdraw(types.Card{Balance:20_000_00, Active: false}, 10_000_00)
	fmt.Println(result.Balance);
    
    // Output: 2000000
}

func ExampleWithdraw_limit(){
	result := Withdraw(types.Card{Balance:20_000_00, Active: true}, 21_000_00)
	fmt.Println(result.Balance);
    
    // Output: 2000000
}

func ExampleDeposit_positive(){
	card := types.Card{Balance:20_000_00, Active: true}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance);
    
    // Output: 3000000
}

func ExampleDeposit_inactive(){
	card := types.Card{Balance:20_000_00, Active: false}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance);
    
    // Output: 2000000
}

func ExampleDeposit_limit(){
	card := types.Card{Balance:20_000_00, Active: true}
	Deposit(&card, 51_000_00)
	fmt.Println(card.Balance);
    
    // Output: 2000000
}

func ExamplePaymentSources() {
	cards := []types.Card{
			{
					Balance: 100,
					Active:  true,
					PAN:     "5555",
			},
			{
					Balance: 1001,
					Active:  false,
					PAN:     "777777",
			},
	}
	res := PaymentSources(cards)

	for _, c := range res {
			fmt.Println(c.Number)
	}

	// Output: 5555
}
