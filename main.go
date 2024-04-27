package main

import (
	"fmt"

	"github.com/jpshrader/financial-simulation/internal/assets"
	"github.com/jpshrader/financial-simulation/internal/common"
	"github.com/jpshrader/financial-simulation/internal/simulation"
)

func main() {
	cash := assets.CashEquivalent{
		CostBasis:    float64(100_000),
		Contribution: float64(10_000),
		RateOfReturn: float64(0.05),
	}
	rothIra := assets.RothIra{
		CostBasis:    float64(70_000),
		Contribution: float64(7_000),
		RateOfReturn: float64(.07),
	}

	simYears := simulation.SimulateYears(10, []assets.Asset{cash, rothIra})
	for year, simYear := range simYears {
		fmt.Printf("year %d:\n", year)
		c := simYear[0].(assets.CashEquivalent)
		logCash(c)
		roth := simYear[1].(assets.RothIra)
		logRothIra(roth)
	}
}

func logCash(cash assets.CashEquivalent) {
	fmt.Print("====== CASH ======")
	fmt.Printf("cost basis: $%.2f\n", common.To2f(cash.CostBasis))
	fmt.Printf("interest income: $%.2f\n", common.To2f(cash.InterestIncome))
	fmt.Printf("contribution: $%.2f\n", common.To2f(cash.Contribution))
	fmt.Printf("rate of return: %.2f\n", common.To2f(cash.RateOfReturn))
	fmt.Printf("value: $%.2f\n", common.To2f(cash.GetGrossValue()))
}

func logRothIra(ira assets.RothIra) {
	fmt.Println("====== ROTH IRA ======")
	fmt.Printf("cost basis: $%.2f\n", common.To2f(ira.CostBasis))
	fmt.Printf("capital gains: $%.2f\n", common.To2f(ira.CapitalGains))
	fmt.Printf("contribution: $%.2f\n", common.To2f(ira.Contribution))
	fmt.Printf("rate of return: %.2f\n", common.To2f(ira.RateOfReturn))
	fmt.Printf("value: $%.2f\n", common.To2f(ira.GetGrossValue()))
}
