package practice

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputConsole() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter some test: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Entered test is: ", input)

	fmt.Println("Enter number: ")
	numinput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numinput), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Entered test is: ", aFloat)
	}
}
