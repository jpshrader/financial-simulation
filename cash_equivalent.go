package main

type CashEquivalent struct {
	CostBasis           float64
	InterestIncome      float64
	Value               float64
	DefinedContribution float64
	RateOfReturn        float64
}

func (c CashEquivalent) Compound() asset {
	c.InterestIncome += c.Value * c.RateOfReturn
	c.CostBasis += c.DefinedContribution
	c.Value = c.CostBasis + c.InterestIncome
	return c
}

func (c CashEquivalent) GetValue() float64 {
	return to2f(c.Value)
}

func (c CashEquivalent) GetCostBasis() float64 {
	return to2f(c.CostBasis)
}

func (c CashEquivalent) GetInterestIncome() float64 {
	return to2f(c.InterestIncome)
}
