package main

import (
	"fmt"

	"github.com/WissameMekhilef/euler-project/ex_1/helpers"
)

func main() {
	x := helpers.GetMultiplesOfUnder(3, 5, 10000)
	fmt.Println(x)
}
