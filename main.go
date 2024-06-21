package main

import (
	"fmt"
	"github.com/jay13jay/metaPractice/combosum"
)

func main() {
	ret := combosum.ComboSum([]int{2, 3, 6, 7}, 7)

	fmt.Printf("Combinations: %v\n", ret)
}