package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jay13jay/metaPractice/bossfight"
)
var size int32 = 1000
var source rand.Source = rand.NewSource(time.Now().UnixNano())
var rng *rand.Rand = rand.New(source)

func main() {
	for i := 0; i < 10; i ++ {
		// Health
		var H []int32 = createSlice(size, 10)
		// var H []int32 = []int32{1, 1, 2, 100}
		N := size

		// Fighter Damage / sec
		// var D []int32 = []int32{1, 2, 1, 3}
		D := createSlice(size, 10)
		// Boss Damage / sec
		var B int32 = int32(rng.Intn(20))

		start := time.Now()
		ret := bossfight.GetDamage(N, H, D, B)
		end := time.Now()
		elapsed := end.Sub(start)
		fmt.Printf("Fighters: %d \t Damage: %.2f\tTime: %v\n", N, ret, elapsed)
		
		size += 1000
	}
}

func createSlice(size int32, card int) []int32 {
	arr := make([]int32, size)
	for i := 0; i < int(size); i++ {
		arr[i] = int32(rng.Intn(card))
	}
	return arr
}
