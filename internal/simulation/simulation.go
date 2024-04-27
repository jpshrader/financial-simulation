package simulation

import "github.com/jpshrader/financial-simulation/internal/assets"

func Simulate(asset assets.Asset) assets.Asset {
	return asset.Compound()
}

func SimulateYears(years int, a assets.Asset) []assets.Asset {
	simYears := make([]assets.Asset, years+1)
	simYears[0] = a
	for i := range years {
		simYears[i+1] = Simulate(simYears[i])
	}
	return simYears
}
