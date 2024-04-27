package simulation

import (
	"testing"

	"github.com/jpshrader/financial-simulation/internal/assets"
	"github.com/stretchr/testify/assert"
)

func Test_SimulationYears_Cash(t *testing.T) {
	cash := assets.CashEquivalent{
		CostBasis:           float64(100_000),
		InterestIncome:      float64(0),
		DefinedContribution: float64(10_000),
		RateOfReturn:        float64(0.05),
	}
	years := 10
	simYears := SimulateYears(years, cash)
	firstCash := simYears[0].(asset)
	lastCash := simYears[years].(asset)

	assert.Len(t, simYears, years+1)
	assert.Equal(t, float64(100_000), firstCash.GetCostBasis())
	assert.Equal(t, float64(0), firstCash.GetInterestIncome())
	assert.Equal(t, float64(100_000), firstCash.GetValue())

	assert.Equal(t, float64(200_000), lastCash.GetCostBasis())
	assert.Equal(t, float64(88_668.39), lastCash.GetInterestIncome())
	assert.Equal(t, float64(288_668.39), lastCash.GetValue())
}
