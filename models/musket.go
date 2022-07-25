package models

type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{
		Gun: Gun{
			Name:  "Musket Gun",
			Power: 1,
		},
	}
}
