package main

type asset interface {
	Compound() asset
	GetValue() float64
	GetCostBasis() float64
	GetInterestIncome() float64
}

func Simulate(asset asset) asset {
	return asset.Compound()
}

func SimulateYears(years int, a asset) []asset {
	simYears := make([]asset, years+1)
	simYears[0] = a
	for i := range years {
		simYears[i+1] = Simulate(simYears[i])
	}
	return simYears
}
