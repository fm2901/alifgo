package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 1000000
}

func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 0, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 0
}

func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000
}

func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 30_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000
}

func ExampleDeposit_positive() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 30_000_00)
	fmt.Println(card.Balance)
	// Output: 5000000
}

func ExampleDeposit_negativeBalance() {
	card := types.Card{Balance: -20_000_00, Active: true}
	Deposit(&card, 30_000_00)
	fmt.Println(card.Balance)
	// Output: 1000000
}

func ExampleDeposit_negative() {
	card := types.Card{Balance: -20_000_00, Active: false}
	Deposit(&card, 30_000_00)
	fmt.Println(card.Balance)
	// Output: -2000000
}

func ExampleDeposit_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Deposit(&card, 30_000_00)
	fmt.Println(card.Balance)
	// Output: 2000000
}
func ExampleDeposit_limitEqual() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 50_000_00)
	fmt.Println(card.Balance)
	// Output: 7000000
}

func ExampleDeposit_limitAbove() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 60_000_00)
	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleAddBonus_positive() {
	card := types.Card{Balance: 10_000, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 10024
}

func ExampleAddBonus_notActive() {
	card := types.Card{Balance: 10_000, Active: false}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 10000
}

func ExampleAddBonus_negativeBalance() {
	card := types.Card{Balance: -10_000, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: -10000
}

func ExampleAddBonus_limit() {
	card := types.Card{Balance: 2100000, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 2105000
}

func ExampleTotal() {
	cards := []types.Card{
		{Balance: 210, Active: true},
		{Balance: 120, Active: true},
		{Balance: 2000, Active: false},
	}
	total := Total(cards)
	fmt.Println(total)
	// Output: 330
}