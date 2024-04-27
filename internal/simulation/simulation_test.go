package simulation

import (
	"testing"

	"github.com/jpshrader/financial-simulation/internal/assets"
	"github.com/stretchr/testify/assert"
)

func Test_SimulationYears_Cash_And_Roth(t *testing.T) {
	cash := assets.CashEquivalent{
		CostBasis:    float64(100_000),
		Contribution: float64(10_000),
		RateOfReturn: float64(0.05),
	}
	rothIra := assets.RothIra{
		CostBasis:    float64(70_000),
		Contribution: float64(7_000),
		RateOfReturn: float64(0.07),
	}
	years := 10
	as := []assets.Asset{cash, rothIra}
	simYears := SimulateYears(years, as)

	assert.Len(t, simYears, years+1)
	firstSim := simYears[0]
	firstCash := firstSim[0]
	assert.Equal(t, float64(100_000), firstCash.GetCostBasis())
	assert.Equal(t, float64(0), firstCash.GetGrossIncome())
	assert.Equal(t, float64(100_000), firstCash.GetGrossValue())
	firstRoth := firstSim[1]
	assert.Equal(t, float64(70_000), firstRoth.GetCostBasis())
	assert.Equal(t, float64(0), firstRoth.GetGrossIncome())
	assert.Equal(t, float64(70_000), firstRoth.GetGrossValue())

	lastSim := simYears[years]
	lastCash := lastSim[0]
	assert.Equal(t, float64(200_000), lastCash.GetCostBasis())
	assert.Equal(t, float64(88_668.39), lastCash.GetGrossIncome())
	assert.Equal(t, float64(288_668.39), lastCash.GetGrossValue())
	lastRoth := lastSim[1]
	assert.Equal(t, float64(140_000), lastRoth.GetCostBasis())
	assert.Equal(t, float64(94_415.73), lastRoth.GetGrossIncome())
	assert.Equal(t, float64(234_415.73), lastRoth.GetGrossValue())
}
