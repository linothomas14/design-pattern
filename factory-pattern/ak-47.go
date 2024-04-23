package main

type AK47 struct {
	Gun
}

func newAK47() IGun {
	return &AK47{
		Gun: Gun{
			Name:   "AK47",
			Damage: 40,
		},
	}
}
