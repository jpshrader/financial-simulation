package main

import (
	"fmt"
	"strings"

	"github.com/jpshrader/financial-simulation/internal/assets"
	"github.com/jpshrader/financial-simulation/internal/compound"
	"github.com/jpshrader/financial-simulation/internal/format"
	"github.com/jpshrader/financial-simulation/internal/simulation"
)

func main() {
	instructions := simulation.Instructions{
		Years:               10,
		CompoundingSchedule: compound.Annually,
		Snapshot: simulation.Snapshot{
			Year: 2024,
			Assets: []assets.Asset{
				assets.CashEquivalent{
					Name:         "cash",
					CostBasis:    float64(100_000),
					Contribution: float64(10_000),
					RateOfReturn: float64(0.05),
				},
				assets.RothIra{
					Name:         "roth",
					CostBasis:    float64(70_000),
					Contribution: float64(7_000),
					RateOfReturn: float64(0.07),
				},
				assets.Hsa{
					Name:         "hsa",
					CostBasis:    float64(25_000),
					Contribution: float64(3_500),
					RateOfReturn: float64(0.05),
				},
			},
		},
	}
	sim := simulation.Simulate(instructions)
	for _, snap := range sim {
		fmt.Printf("====== YEAR %d ======\n", snap.Year)
		for _, a := range snap.Assets {
			logAsset(a)
		}
	}
}

func logAsset(a assets.Asset) {
	fmt.Printf("%s\n", strings.ToUpper(a.GetName()))
	fmt.Printf("cost basis: $%.2f\n", format.To2f(a.GetCostBasis()))
	fmt.Printf("capital gains: $%.2f\n", format.To2f(a.GetGrossIncome()))
	fmt.Printf("value: $%.2f\n", format.To2f(a.GetGrossValue()))
}
