package Employee

type PartTime struct {
	HourlyRate  float64
	HoursWorked int
}

func (p PartTime) CalculateSalary() float64 {
	return float64(p.HoursWorked) * p.HourlyRate
}

func (p PartTime) CalculateBonus() float64 {
	return 0
}

func (p PartTime) GetWorkHours() int {
	return p.HoursWorked
}
