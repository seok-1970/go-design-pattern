package main

type ChocoPepero struct {
	Pepero
}

func newChocoPepero() IPepero {
	return &ChocoPepero{
		Pepero: Pepero{
			flavor: "choco",
		},
	}
}
