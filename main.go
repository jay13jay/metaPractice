package main

import (
	"fmt"

	"github.com/jay13jay/metaPractice/bossfight"
)

func main() {
	var N int32 = 3
	var H []int32 = []int32{1, 1, 2, 100}
	var D []int32 = []int32{1, 2, 1, 3}
	var B int32 = 8
	ret := bossfight.GetDamage(N, H, D, B)
	fmt.Printf("Damage: %.2f", ret)
}