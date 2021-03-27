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