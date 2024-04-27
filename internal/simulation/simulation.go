package simulation

import "github.com/jpshrader/financial-simulation/internal/assets"

func Simulate(as []assets.Asset) []assets.Asset {
	newAssets := make([]assets.Asset, len(as))
	for i, a := range as {
		newAssets[i] = a.Compound()
	}
	return newAssets
}

func SimulateYears(years int, as []assets.Asset) [][]assets.Asset {
	simYears := make([][]assets.Asset, years+1)
	simYears[0] = as
	for i := range years {
		simYears[i+1] = Simulate(simYears[i])
	}
	return simYears
}
