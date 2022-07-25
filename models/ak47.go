package models

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			Name:  "AK47 Gun",
			Power: 4,
		},
	}
}
