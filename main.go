package main

import (
	"fmt"
	"math"
)

func main() {
	cash := CashEquivalent{
		CostBasis:           float64(100_000),
		InterestIncome:      float64(0),
		DefinedContribution: float64(10_000),
		RateOfReturn:        float64(0.05),
		Value:               float64(100_000),
	}
	simYears := SimulateYears(10, cash)
	for year, simYear := range simYears {
		logCash(year, simYear.(CashEquivalent))
	}
}

func logCash(year int, cash CashEquivalent) {
	fmt.Printf("year %d:\n", year)
	fmt.Printf("basis: $%.2f\n", to2f(cash.CostBasis))
	fmt.Printf("interest income: $%.2f\n", to2f(cash.InterestIncome))
	fmt.Printf("defined contribution: $%.2f\n", to2f(cash.DefinedContribution))
	fmt.Printf("rate of return: %.2f\n", to2f(cash.RateOfReturn))
	fmt.Printf("value: $%.2f\n", to2f(cash.Value))
}

func to2f(f float64) float64 {
	return float64(math.Round(f*100)) / 100
}
