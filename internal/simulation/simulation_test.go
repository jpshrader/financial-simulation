package simulation

import (
	"testing"

	"github.com/jpshrader/financial-simulation/internal/assets"
	"github.com/stretchr/testify/assert"
)

func Test_SimulateYears_Cash_And_Roth(t *testing.T) {
	instructions := Instructions{
		Iterations: 10,
		Snapshot: Snapshot{
			Assets: []assets.Asset{
				assets.CashEquivalent{
					CostBasis:    float64(100_000),
					Contribution: float64(10_000),
					RateOfReturn: float64(0.05),
				},
				assets.RothIra{
					CostBasis:    float64(70_000),
					Contribution: float64(7_000),
					RateOfReturn: float64(0.07),
				},
			},
		},
	}
	sim := Simulate(instructions)

	assert.Len(t, sim, 11)
	firstSim := sim[0]
	firstCash := firstSim.Assets[0]
	assert.Equal(t, float64(100_000), firstCash.GetCostBasis())
	assert.Equal(t, float64(0), firstCash.GetGrossIncome())
	assert.Equal(t, float64(100_000), firstCash.GetGrossValue())
	firstRoth := firstSim.Assets[1]
	assert.Equal(t, float64(70_000), firstRoth.GetCostBasis())
	assert.Equal(t, float64(0), firstRoth.GetGrossIncome())
	assert.Equal(t, float64(70_000), firstRoth.GetGrossValue())

	lastSim := sim[10]
	lastCash := lastSim.Assets[0]
	assert.Equal(t, float64(200_000), lastCash.GetCostBasis())
	assert.Equal(t, float64(88_668.39), lastCash.GetGrossIncome())
	assert.Equal(t, float64(288_668.39), lastCash.GetGrossValue())
	lastRoth := lastSim.Assets[1]
	assert.Equal(t, float64(140_000), lastRoth.GetCostBasis())
	assert.Equal(t, float64(94_415.73), lastRoth.GetGrossIncome())
	assert.Equal(t, float64(234_415.73), lastRoth.GetGrossValue())
}
