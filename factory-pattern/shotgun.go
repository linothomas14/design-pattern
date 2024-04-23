package main

type Shotgun struct {
	Gun
}

func newShotgun() IGun {
	return &Shotgun{
		Gun: Gun{
			Name:   "Shotgun",
			Damage: 67,
		},
	}
}
