package assets

import (
	"github.com/jpshrader/financial-simulation/internal/common"
)

type RothIra struct {
	CostBasis    float64
	CapitalGains float64
	Contribution float64
	RateOfReturn float64
}

func (r RothIra) Compound() Asset {
	r.CapitalGains += r.GetGrossValue() * r.RateOfReturn
	r.CostBasis += r.Contribution
	return r
}

func (r RothIra) GetCostBasis() float64 {
	return common.To2f(r.CostBasis)
}

func (r RothIra) GetGrossValue() float64 {
	return common.To2f(r.CostBasis + r.CapitalGains)
}

func (r RothIra) GetNetValue(capitalGainsRate float64) float64 {
	return common.To2f(r.GetCostBasis() + r.GetNetIncome(capitalGainsRate))
}

func (r RothIra) GetGrossIncome() float64 {
	return common.To2f(r.CapitalGains)
}

func (r RothIra) GetNetIncome(capitalGainsRate float64) float64 {
	return common.To2f(r.CapitalGains)
}
