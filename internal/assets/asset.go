package assets

type Asset interface {
	Compound() Asset
	GetValue() float64
	GetCostBasis() float64
	GetInterestIncome() float64
}
