package assets

type Asset interface {
	Compound() Asset
	GetGrossValue() float64
	GetCostBasis() float64
	GetGrossIncome() float64
	GetNetIncome(capitalGainsRate float64) float64
}
