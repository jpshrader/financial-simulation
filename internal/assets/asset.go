package assets

import "github.com/jpshrader/financial-simulation/internal/common"

type Asset interface {
	GetName() string
	Compound(schedule common.CompoundingSchedule, isNewYear bool) Asset
	GetCostBasis() float64
	GetGrossValue() float64
	GetNetValue(capitalGainsRate float64) float64
	GetGrossIncome() float64
	GetNetIncome(capitalGainsRate float64) float64
}
