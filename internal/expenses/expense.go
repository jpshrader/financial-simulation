package expenses

type Expense interface{
	GetName() string
	GetAmount() string
	GetInflationRate() float64
}