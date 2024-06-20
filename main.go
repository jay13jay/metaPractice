package main

import (
	"fmt"

	"github.com/jay13jay/metaPractice/uniformintegers"
)

func main() {
	var a int64 = 234
	var b int64 = 10556
	ret := uniformintegers.GetUniforms(a, b)
	fmt.Println(ret)

}