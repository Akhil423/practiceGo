package practice

import (
	"fmt"
)

func Pointers() {

	anInt := 45
	pointer := &anInt
	fmt.Println("value by pointer:", *pointer)

	value1 := 43.4
	pointer1 := &value1
	fmt.Println("value by pointer", *pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("value with pointer:", *pointer1)

}
