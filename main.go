package main

import (
	"fmt"
)

type StringShower interface {
	fmt.Stringer
	Show()
}

type stringShower struct {
	fmt.Stringer
}

func (s *stringShower) Show() {
	fmt.Println(s.String())
}

func NewStringShower(s fmt.Stringer) StringShower {
	return &stringShower{s}
}

type People struct {
	Name string
}

func (p *People) String() string {
	return p.Name
}

func main() {

	var p fmt.Stringer = &People{Name: "Guillermo"}
	shower := NewStringShower(p)
	shower.Show()

}
