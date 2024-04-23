package main

type Gun struct {
	Name   string
	Damage int
}

func (g *Gun) setName(name string) {
	g.Name = name
}

func (g *Gun) getName() string {
	return g.Name
}

func (g *Gun) setPower(power int) {
	g.Damage = power
}

func (g *Gun) getPower() int {
	return g.Damage
}
