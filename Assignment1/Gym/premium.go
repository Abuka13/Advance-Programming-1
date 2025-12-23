package Gym

type PremiumMember struct {
	Name       string
	Age        int
	HasTrainer bool
}

func (p PremiumMember) GetDetails() string {
	return "Premium Member: " + p.Name
}
