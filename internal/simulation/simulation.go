package simulation

import (
	"github.com/jpshrader/financial-simulation/internal/assets"
	"github.com/jpshrader/financial-simulation/internal/expenses"
)

type Instructions struct {
	Iterations int
	Snapshot   Snapshot
}

type Snapshot struct {
	Year     int
	Assets   []assets.Asset
	Expenses []expenses.Expense
}

func Simulate(instructions Instructions) []Snapshot {
	snapshots := make([]Snapshot, instructions.Iterations+1)
	snapshots[0] = instructions.Snapshot
	for i := range instructions.Iterations {
		snapshots[i+1] = simulate(snapshots[i])
	}
	return snapshots
}

func simulate(data Snapshot) Snapshot {
	newAssets := make([]assets.Asset, len(data.Assets))
	for i, a := range data.Assets {
		newAssets[i] = a.Compound()
	}
	return Snapshot{
		Year:   data.Year + 1,
		Assets: newAssets,
	}
}
