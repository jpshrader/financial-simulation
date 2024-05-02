package common

type CompoundingSchedule int

const (
	Annually CompoundingSchedule = 1
	Monthly  CompoundingSchedule = 12
	Weekly   CompoundingSchedule = 52
)

func (c CompoundingSchedule) IsValid() bool {
	switch c {
	case Annually, Monthly, Weekly:
		return true
	default:
		return false
	}
}
