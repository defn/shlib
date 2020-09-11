package main

import (
	"fmt"

	"github.com/defn/shlib/v2"
	"github.com/dgwyn/gorillama"
)

func main() {
	fmt.Println(shlib.Hello())
	a, b := 1, 2
	fmt.Printf("%d + %d = %d\n", a, b, gorillama.Sum(a, b))
}
