package main

type Musket struct {
	Gun
}

func NewMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
