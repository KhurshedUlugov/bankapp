package main

import(
	"fmt"
	"bank/pkg/bank/types"
   "bank/pkg/bank/card"
   "bank/pkg/bank/payment"
)

func main(){
   pan := card.IssueCard("TJS", "green", "ASU AS")
   fmt.Println(pan)

   result := card.Withdraw(types.Card{Balance:20_000_00, Active: true}, 10_000_00)
   fmt.Println("Pos: ",result.Balance);

   result = card.Withdraw(types.Card{Balance:10_000_00, Active: true}, 20_000_00)
   fmt.Println("no Money: ",result.Balance);

   result = card.Withdraw(types.Card{Balance:20_000_00, Active: false}, 10_000_00)
   fmt.Println("inactive: ", result.Balance);

   result = card.Withdraw(types.Card{Balance:20_000_00, Active: true}, 21_000_00)
   fmt.Println("lilit: ",result.Balance);

   cards := []types.Card{
   {   
      Balance: 10000,
      Active: true,
   },
   {
      Balance: 15000,
      Active: true,
   },
   {
      Balance: 12000,
      Active: true,
   },  
   }

   total := card.Total(cards)
   fmt.Println("Total: ",total);

   payments := []types.Payment{
      {
         ID: 1,
         Amount: 100,
      },
      {
         ID: 2,
         Amount: 101,
      },
      {
         ID: 3,
         Amount: 102,
      },
   }
   max := payment.Max(payments)
   fmt.Println("Max: ", max)
}