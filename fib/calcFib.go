package fib

import (
	"math/big"
)

var doneTest map[[2]int]bool
var doneRes map[[2]int]*big.Int

func init() {
    // Initialize the maps
    doneTest = make(map[[2]int]bool)
    doneRes = make(map[[2]int]*big.Int)
}

// CalcFib calculates the fibonacci number for n
func CalcFib(n int) *big.Int {
    if n <= 1 {
        return big.NewInt(int64(n))
    }
    // Check if the result is already computed
    if doneTest[[2]int{n-1, n-2}] {
        return doneRes[[2]int{n-1, n-2}]
    }
    res := new(big.Int).Add(CalcFib(n-1), CalcFib(n-2))
    doneTest[[2]int{n-1, n-2}] = true
    doneRes[[2]int{n-1, n-2}] = res
    return res
}

func FormatScientific(n *big.Int) string {
    f := new(big.Float).SetInt(n)
    return f.Text('e', 10) // 'e' for scientific notation with 10 decimal places
}