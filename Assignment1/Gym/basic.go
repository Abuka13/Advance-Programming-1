package Gym

type BasicMember struct {
	Name string
	Age  int
}

func (b BasicMember) GetDetails() string {
	return "Basic Member: " + b.Name
}
