package practice

import (
	"fmt"
	"sort"
)

func ArraysCheck() {
	var components [3]string
	components[0] = "cpu"
	components[1] = "monitor"
	components[2] = "logical"
	fmt.Println(components)
	fmt.Println("value at 2nd position", components[1])

	var even = [5]int{2, 4, 6, 8, 10}
	fmt.Println(even)
	fmt.Println("value at 3rd position", even[2])

	///total number of components
	fmt.Println(components)

	//length of components array
	fmt.Println("length of components", len(components))
}

func SlicesCheck() {
	var components = []string{"cpu", "monitor", "logical"}
	fmt.Println(components)
	components = append(components, "mouse")
	fmt.Println(components)

	components = append(components[1:])
	fmt.Println(components)
	components = append(components[:len(components)-1])

	///total number of components
	fmt.Println(components)

	//memory allocation with make

	numbers := make([]int, 5)
	numbers[0] = 2
	numbers[1] = 8
	numbers[2] = 6
	numbers[3] = 4
	numbers[4] = 10

	fmt.Println(numbers)

	numbers = append(numbers, 456)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)

}

func Mapss() {
	states := make(map[string]string)
	fmt.Println(states)
	states["TS"] = "telangana"
	states["KA"] = "karnataka"
	states["MR"] = "maharastra"
	fmt.Println("print map val:", states)

	fmt.Println("print a val", states["TS"])

	delete(states, "MR")
	states["TN"] = "Tamil Nadu"
	fmt.Println(states)

	for key, val := range states {

		fmt.Printf("%v: %v ", key, val)
	}

	//sort order map

	keys := make([]string, len(states))
	i := 0

	for k := range states {
		keys[i] = k
		i++
	}

	fmt.Println("keys print", keys)
	sort.Strings(keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}

}
