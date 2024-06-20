package uniformintegers

import (
	"fmt"
	"math"
)

func getUniformIntegerCountInInterval(A int64, B int64) int32 {
  // Write your code here

	// get # of digits in A and B
	digitsA := int64(math.Log10(float64(A))) + 1
	digitsB := int64(math.Log10(float64(B))) + 1
	fmt.Printf("digitsA: %d\tdigitsB: %d\n", digitsA, digitsB)

	difference := digitsB - digitsA
	fmt.Println("difference: ", difference)

	// get 10^digitsA and 10^digitsB"
	capA := int64(math.Pow(10, float64(digitsA)))
	capB := int64(math.Pow(10, float64(digitsB)))
	fmt.Printf("capA: %d\tcapB: %d\n", capA, capB)

	// 
	floorA := int64((1.0 / 9.0) * float64(capA))
	floorB := int64((1.0 / 9.0) * float64(capB))
	fmt.Printf("floorA: %d\t floorB: %d\n", floorA, floorB)


	var uni int64
	// if # of digits difference is more than 0, 
	// subtract floor from cap. then subtract
	// else subtract floor from original value
	if difference > 0 {
		uniA := int64(float64(capA-1)/float64(floorA)) - int64(float64(A)/float64(floorA))
		fmt.Println("uniA: ", uniA)
		uniB := int64(float64(B) / float64(floorB))
		fmt.Println("uniB: ", uniB)
		uni = uniA + uniB + (difference-1)*9
	} else {
		uniA := int64(float64(A) / float64(floorA))
		uniB := int64(float64(B) / float64(floorB))
		uni = uniB - uniA
	}

	if A%floorA == 0 {
		uni++
	}

	return int32(uni)
}

func GetUniforms(a, b int64) int32 {
	return getUniformIntegerCountInInterval(a,b)
}