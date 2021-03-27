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