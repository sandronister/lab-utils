package main

import (
	"fmt"

	"github.com/sandronister/lab-utils/pkg/utils/array"
)

func main() {
	list := []string{"Asmodeus", "Balrog", "Paimon"}
	value, err := array.SearchKey[[]string](2, list)
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
}
