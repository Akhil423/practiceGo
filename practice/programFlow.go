package practice

import (
	"fmt"
	"math/rand"
	"time"
)

func ConditionalStatements() {

	password := "Conditional"

	if len(password) == 0 {
		fmt.Println("enter password")
	} else if len(password) < 5 {
		fmt.Println("bad password")
	} else if len(password) > 8 {
		fmt.Println("good password")
	} else {
		fmt.Println("Un execpected")
	}

	// condition with initialization

	if password = "Changed"; len(password) == 0 {
		fmt.Println("enter password")

	} else if len(password) < 5 {
		fmt.Println("bad password")
	} else if len(password) > 8 {
		fmt.Println("good password")
	} else {
		fmt.Println("Un execpected")
	}
}

func SwitchTest() {

	rand.Seed(time.Now().Unix())
	var result string
	switch dayVal := rand.Intn(7) + 1; dayVal {
	case 1:
		result = "This is sunday!"
	case 2:
		result = "This is monday!"
	default:
		result = "the other day"

	}
	fmt.Println(result)

	// fallthrough in any case can be used to fall to next case if parent is fail

}

func ForLoops() {

	colors := []string{"Green", "White", "Blue"}

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	for _, color := range colors {
		fmt.Println(color)
	}

	// while loop like
	x := 5
	for x > 0 {
		fmt.Println("Value: ", x)
		x--
	}

	x = 2

	for x > 20 {
		fmt.Println("Value:", x)
		x++

		if x == 15 {
			goto itsFifteen
		}
	}

itsFifteen:
	fmt.Println("it's fifteenth position")
}
