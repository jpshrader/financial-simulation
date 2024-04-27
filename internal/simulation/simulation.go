package simulation

type asset interface {
	Compound() any
	GetValue() float64
	GetCostBasis() float64
	GetInterestIncome() float64
}

func Simulate(asset asset) any {
	return asset.Compound()
}

func SimulateYears(years int, a asset) []any {
	simYears := make([]any, years+1)
	simYears[0] = a
	for i := range years {
		simYears[i+1] = Simulate(simYears[i].(asset))
	}
	return simYears
}
