package assets

import (
	"github.com/jpshrader/financial-simulation/internal/common"
)

type CashEquivalent struct {
	CostBasis      float64
	InterestIncome float64
	Contribution   float64
	RateOfReturn   float64
}

func (c CashEquivalent) Compound() Asset {
	c.InterestIncome += c.GetGrossValue() * c.RateOfReturn
	c.CostBasis += c.Contribution
	return c
}

func (c CashEquivalent) GetCostBasis() float64 {
	return common.To2f(c.CostBasis)
}

func (c CashEquivalent) GetGrossValue() float64 {
	return common.To2f(c.CostBasis + c.InterestIncome)
}

func (c CashEquivalent) GetNetValue(capitalGainsRate float64) float64 {
	return common.To2f(c.GetCostBasis() + c.GetNetIncome(capitalGainsRate))
}

func (c CashEquivalent) GetGrossIncome() float64 {
	return common.To2f(c.InterestIncome)
}

func (c CashEquivalent) GetNetIncome(capitalGainsRate float64) float64 {
	return common.To2f(c.InterestIncome * (1 - capitalGainsRate))
}
