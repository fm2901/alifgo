package payment

import "bank/pkg/bank/types"

func Max(payments []types.Payment) types.Payment {
	maxPaymentId := 0
	for id, payment := range payments {
		if payment.Amount > payments[maxPaymentId].Amount {
			maxPaymentId = id
		}
	}
	return payments[maxPaymentId]
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