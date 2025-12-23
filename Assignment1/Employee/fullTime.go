package Employee

type FullTime struct {
	MonthlySalary float64
	BonusRate     float64
}

func (f FullTime) CalculateSalary() float64 {
	return f.MonthlySalary + f.CalculateBonus()
}

func (f FullTime) CalculateBonus() float64 {
	return f.MonthlySalary * f.BonusRate
}

func (f FullTime) GetWorkHours() int {
	return 160
}
