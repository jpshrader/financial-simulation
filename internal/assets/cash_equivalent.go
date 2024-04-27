package assets

import (
	"math"
)

type CashEquivalent struct {
	CostBasis           float64
	InterestIncome      float64
	DefinedContribution float64
	RateOfReturn        float64
}

func (c CashEquivalent) Compound() Asset {
	c.InterestIncome += c.GetValue() * c.RateOfReturn
	c.CostBasis += c.DefinedContribution
	return c
}

func (c CashEquivalent) GetValue() float64 {
	return To2f(c.CostBasis + c.InterestIncome)
}

func (c CashEquivalent) GetCostBasis() float64 {
	return To2f(c.CostBasis)
}

func (c CashEquivalent) GetInterestIncome() float64 {
	return To2f(c.InterestIncome)
}

func To2f(f float64) float64 {
	return float64(math.Round(f*100)) / 100
}
