package main

import (
	"fmt"

	module "github.com/yolandesmita/module"
)

func main() {

	fmt.Println(module.SayHello())
	fmt.Println(module.NamaPengunjung())

	A := module.Diskon{"Agus"}
	fmt.Println(A.DiskonMakan())

	B := module.Diskon{"Agus"}
	fmt.Println(B.DiscoutUsed())

	C := module.Customer{"Agus", 29, "Kasryawan Swasta"}
	fmt.Println(C)

	D := module.Nilai{2, 5, 10}
	fmt.Println(D.Perkalian())
}
