package main

import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	newCard := card.IssueCard(types.TJS, "white", "Abdu")
	fmt.Printf("%+v", newCard)
}