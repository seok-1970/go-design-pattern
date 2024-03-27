package main

type Pepero struct {
	flavor string
}

func (p *Pepero) setFlavor(flavor string) {
	p.flavor = flavor
}

func (p *Pepero) getFlavor() string {
	return p.flavor
}
