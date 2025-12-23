package Gym

import "fmt"

type Gym struct {
	Members map[int]Member
}

func (g *Gym) AddMember(id int, member Member) {
	g.Members[id] = member
}

func (g *Gym) ListMembers() {
	for id, member := range g.Members {
		fmt.Println("ID:", id, "-", member.GetDetails())
	}
}
