package assets

import (
	"github.com/jpshrader/financial-simulation/internal/compound"
	"github.com/jpshrader/financial-simulation/internal/format"
)

type RothIra struct {
	Name             string
	CostBasis        float64
	CapitalGains     float64
	Contribution     float64
	RateOfReturn     float64
	YearToDateGrowth float64
}

func (r RothIra) GetName() string {
	return r.Name
}

func (r RothIra) Compound(schedule compound.CompoundingSchedule, isNewYear bool) Asset {
	if isNewYear {
		r.YearToDateGrowth = 0
	}
	gains := ((r.GetGrossValue() - r.YearToDateGrowth) * r.RateOfReturn) / float64(schedule)
	contribution := r.Contribution / float64(schedule)

	r.CapitalGains += gains
	r.CostBasis += contribution
	r.YearToDateGrowth += gains + contribution
	return r
}

func (r RothIra) GetCostBasis() float64 {
	return format.To2f(r.CostBasis)
}

func (r RothIra) GetGrossValue() float64 {
	return format.To2f(r.CostBasis + r.CapitalGains)
}

func (r RothIra) GetNetValue(capitalGainsRate float64) float64 {
	return format.To2f(r.GetCostBasis() + r.GetNetIncome(capitalGainsRate))
}

func (r RothIra) GetGrossIncome() float64 {
	return format.To2f(r.CapitalGains)
}

func (r RothIra) GetNetIncome(capitalGainsRate float64) float64 {
	return format.To2f(r.CapitalGains)
}
