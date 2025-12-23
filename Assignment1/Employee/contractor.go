package Employee

type Contractor struct {
	ProjectRate       float64
	ProjectsCompleted int
}

func (c Contractor) CalculateSalary() float64 {
	return float64(c.ProjectsCompleted) * c.ProjectRate
}

func (c Contractor) CalculateBonus() float64 {
	return 0
}

func (c Contractor) GetWorkHours() int {
	return 0
}
