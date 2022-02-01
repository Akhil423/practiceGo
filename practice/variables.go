package practice

import (
	"fmt"
)

func Variables() {
	var aString string = "This is my go practice"

	fmt.Println(aString)
	fmt.Printf("the variable type is %T\n", aString)
	var anInteger int = 24
	fmt.Println(anInteger)

	var defaultVal int

	fmt.Println(defaultVal)

	var newString = "definition without type"
	fmt.Println(newString)
	fmt.Printf("the variable type is %T\n", newString)

	newString2 := "colon declarion" //only accessible inside the class
	fmt.Println(newString2)
	fmt.Printf("the variable type is %T\n", newString2)

}
