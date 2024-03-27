package main

type CookiePepero struct {
	Pepero
}

func newCookiePepero() IPepero {
	return &CookiePepero{
		Pepero: Pepero{
			flavor: "cookie",
		},
	}
}
