package card

import (
	"bank/pkg/bank/types"
)

func IssueCard(currency types.Currency, color string, name string) types.Card {
	newCard := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Active:   true,
		Color:    color,
		Name:     name,
	}
	return newCard
}

func Withdraw(card types.Card, amount types.Money) types.Card {
	limit := types.Money(20_000_00)
	if card.Active && card.Balance >= amount && amount > 0 && amount <= limit {
		card.Balance = card.Balance - amount
	}
	return card
}

func Deposit(card *types.Card, amount types.Money) {
	depositLimit := 50_000_00
	if !card.Active {
		return
	}
	if int(amount) > depositLimit {
		return
	}
	if int(amount) < 0 {
		return
	}
	(*card).Balance += types.Money(amount)
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	minBalance := 10_000
	maxBonus := 5_000

	if !card.Active {
		return
	}

	if int(card.Balance) < minBalance {
		return
	}

	bonus := int(int(card.Balance) * percent / 100 * daysInMonth / daysInYear)
	if bonus > maxBonus {
		bonus = maxBonus
	}

	(*card).Balance += types.Money(bonus)
}

func Total(cards []types.Card) types.Money {
	var totalBalance types.Money = 0
	for _, card := range cards {
		if card.Balance > 0 && card.Active {
			totalBalance += card.Balance
		}
	}
	return totalBalance
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var result []types.PaymentSource
	cardType := "card"
	for _, card := range cards {
		if !card.Active || card.Balance < 1 {
			continue;
		}
		pmtSource := types.PaymentSource{
			Type: cardType,
			Number: string(card.PAN),
			Balance: card.Balance,
		}
		result = append(result, pmtSource)
	}
	return result
}