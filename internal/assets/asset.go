package assets

type Asset interface {
	GetName() string
	Compound() Asset
	GetCostBasis() float64
	GetGrossValue() float64
	GetNetValue(capitalGainsRate float64) float64
	GetGrossIncome() float64
	GetNetIncome(capitalGainsRate float64) float64
}
