package assets

import (
	"github.com/jpshrader/financial-simulation/internal/compound"
	"github.com/jpshrader/financial-simulation/internal/format"
)

type Hsa struct {
	Name             string
	CostBasis        float64
	CapitalGains     float64
	Contribution     float64
	RateOfReturn     float64
	YearToDateGrowth float64
}

func (h Hsa) GetName() string {
	return h.Name
}

func (h Hsa) Compound(schedule compound.CompoundingSchedule, isNewYear bool) Asset {
	if isNewYear {
		h.YearToDateGrowth = 0
	}
	gains := (h.GetGrossValue() - h.YearToDateGrowth) * (h.RateOfReturn / float64(schedule))
	h.CapitalGains += gains
	contribution := h.Contribution / float64(schedule)
	h.CostBasis += contribution
	h.YearToDateGrowth += contribution + gains
	return h
}

func (h Hsa) GetCostBasis() float64 {
	return format.To2f(h.CostBasis)
}

func (h Hsa) GetGrossValue() float64 {
	return format.To2f(h.CostBasis + h.CapitalGains)
}

func (h Hsa) GetNetValue(capitalGainsRate float64) float64 {
	return format.To2f(h.GetCostBasis() + h.GetNetIncome(capitalGainsRate))
}

func (h Hsa) GetGrossIncome() float64 {
	return format.To2f(h.CapitalGains)
}

func (h Hsa) GetNetIncome(capitalGainsRate float64) float64 {
	return format.To2f(h.CapitalGains)
}
