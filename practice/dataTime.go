package practice

import (
	"fmt"
	"time"
)

func DateTimeGo() {

	presentTime := time.Now()
	fmt.Println(presentTime)

	dateAt := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go launched on", dateAt)
	fmt.Println(dateAt.Format(time.ANSIC))

}
