package practice

import (
	"fmt"
	"math"
)

func MathOps() {

	f1, f2, f3 := 2.3, 5.4, 5.6
	sumOfFloats := f1 + f2 + f3
	fmt.Println("sum of floats:", sumOfFloats)

	percisionVal := math.Round(sumOfFloats) * 100 / 100
	fmt.Println("sum with precision of 2:", percisionVal)
}
