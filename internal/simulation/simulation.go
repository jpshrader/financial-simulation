package simulation

import (
	"github.com/jpshrader/financial-simulation/internal/assets"
	"github.com/jpshrader/financial-simulation/internal/common"
	"github.com/jpshrader/financial-simulation/internal/expenses"
)

type Instructions struct {
	Iterations          int
	CompoundingSchedule common.CompoundingSchedule
	Snapshot            Snapshot
}

type Snapshot struct {
	Year     int
	Assets   []assets.Asset
	Expenses []expenses.Expense
}

func Simulate(instructions Instructions) []Snapshot {
	snapshots := make([]Snapshot, (instructions.Iterations*int(instructions.CompoundingSchedule))+1)
	snapshots[0] = instructions.Snapshot
	for i := range instructions.Iterations * int(instructions.CompoundingSchedule) {
		snapshots[i+1] = simulate(snapshots[i], i, instructions.CompoundingSchedule)
	}
	return snapshots
}

func simulate(data Snapshot, iteration int, schedule common.CompoundingSchedule) Snapshot {
	newAssets := make([]assets.Asset, len(data.Assets))
	for i, a := range data.Assets {
		newAssets[i] = a.Compound(schedule, iteration%int(schedule) == 0)
	}
	return Snapshot{
		Year:   data.Year + 1,
		Assets: newAssets,
	}
}
