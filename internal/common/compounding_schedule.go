package common

type CompoundingSchedule int

const (
	Annually  CompoundingSchedule = 1
	Quarterly CompoundingSchedule = 4
	Monthly   CompoundingSchedule = 12
)

func (c CompoundingSchedule) IsValid() bool {
	switch c {
	case Annually, Quarterly, Monthly:
		return true
	default:
		return false
	}
}
