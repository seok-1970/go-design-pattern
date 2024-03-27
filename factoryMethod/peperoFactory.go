package main

import "fmt"

func GetPepero(flavor string) (IPepero, error) {
	if flavor == "choco" {
		return newChocoPepero(), nil
	} else if flavor == "cookie" {
		return newCookiePepero(), nil
	}
	return nil, fmt.Errorf("Invalid flavor")
}
