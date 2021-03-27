package payment

import (
	"fmt"
	"bank/pkg/bank/types"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 10,
		},
		{
			ID: 99,
			Amount: 15,
		},
		{
			ID: 3,
			Amount: 11,
		},
	}

	max := Max(payments)
	fmt.Println(max.ID)
	// Output: 99
}

func ExamplePaymentSources() {
	var cards []types.Card
	cards = append(cards, 
		types.Card {
			PAN:		"1234 5678 9123 4567 8901",
			Balance: 	100.0,
			Active:		false,
		},
		types.Card {
			PAN:		"1234 5678 9123 4567 8902",
			Balance: 	200.0,
			Active:		true,
		},
		types.Card {
			PAN:		"1234 5678 9123 4567 8903",
			Balance: 	0,
			Active:		true,
		},
	)

	result := PaymentSources(cards)
	for _,ps := range result {
		fmt.Println(ps.Number)
	}
	// Output: 1234 5678 9123 4567 8902
}