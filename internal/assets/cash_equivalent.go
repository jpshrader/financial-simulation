package assets

import (
	"github.com/jpshrader/financial-simulation/internal/compound"
	"github.com/jpshrader/financial-simulation/internal/format"
)

type CashEquivalent struct {
	Name             string
	CostBasis        float64
	InterestIncome   float64
	Contribution     float64
	RateOfReturn     float64
	YearToDateGrowth float64
}

func (c CashEquivalent) GetName() string {
	return c.Name
}

func (c CashEquivalent) Compound(schedule compound.CompoundingSchedule, isNewYear bool) Asset {
	if isNewYear {
		c.YearToDateGrowth = 0
	}
	income := ((c.GetGrossValue() - c.YearToDateGrowth) * c.RateOfReturn) / float64(schedule)
	c.InterestIncome += income
	contribution := c.Contribution / float64(schedule)
	c.CostBasis += contribution
	c.YearToDateGrowth += income + contribution
	return c
}

func (c CashEquivalent) GetCostBasis() float64 {
	return format.To2f(c.CostBasis)
}

func (c CashEquivalent) GetGrossValue() float64 {
	return format.To2f(c.CostBasis + c.InterestIncome)
}

func (c CashEquivalent) GetNetValue(capitalGainsRate float64) float64 {
	return format.To2f(c.GetCostBasis() + c.GetNetIncome(capitalGainsRate))
}

func (c CashEquivalent) GetGrossIncome() float64 {
	return format.To2f(c.InterestIncome)
}

func (c CashEquivalent) GetNetIncome(capitalGainsRate float64) float64 {
	return format.To2f(c.InterestIncome * (1 - capitalGainsRate))
}
