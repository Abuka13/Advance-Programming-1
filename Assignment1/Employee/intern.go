package Employee

type Intern struct {
	DailyRate  float64
	DaysWorked int
}

func (i Intern) CalculateSalary() float64 {
	return float64(i.DaysWorked) * i.DailyRate
}

func (i Intern) CalculateBonus() float64 {
	return 0
}

func (i Intern) GetWorkHours() int {
	return i.DaysWorked * 8
}
