package main

import (
	"fmt"
	"reflect"
)

func main() {
	chocoPepero, _ := GetPepero("choco")
	cookiePepero, _ := GetPepero("cookie")
	fmt.Println(reflect.TypeOf(chocoPepero))
	fmt.Println(reflect.TypeOf(cookiePepero))

}
