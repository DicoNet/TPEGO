package main

import (
	"fmt"

	"go2021.com/model"
)

func main() {
	fmt.Println("hola xd")
	Result, err := model.VerificarCaracteres("TX04ABCE")

	fmt.Println(Result, err)
}
