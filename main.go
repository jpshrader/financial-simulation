package main

import (
	"fmt"

	"github.com/jpshrader/financial-simulation/internal/assets"
	"github.com/jpshrader/financial-simulation/internal/simulation"
)

func main() {
	cash := assets.CashEquivalent{
		CostBasis:           float64(100_000),
		InterestIncome:      float64(0),
		DefinedContribution: float64(10_000),
		RateOfReturn:        float64(0.05),
		Value:               float64(100_000),
	}
	simYears := simulation.SimulateYears(10, cash)
	for year, simYear := range simYears {
		logCash(year, simYear.(assets.CashEquivalent))
	}
}

func logCash(year int, cash assets.CashEquivalent) {
	fmt.Printf("year %d:\n", year)
	fmt.Printf("basis: $%.2f\n", assets.To2f(cash.CostBasis))
	fmt.Printf("interest income: $%.2f\n", assets.To2f(cash.InterestIncome))
	fmt.Printf("defined contribution: $%.2f\n", assets.To2f(cash.DefinedContribution))
	fmt.Printf("rate of return: %.2f\n", assets.To2f(cash.RateOfReturn))
	fmt.Printf("value: $%.2f\n", assets.To2f(cash.Value))
}
