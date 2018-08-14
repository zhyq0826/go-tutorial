package main

import (
	"fmt"
)

func main() {
	months := [...]string{1: "January", 2: "February", 3: "Marcy", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	Q2 := months[4:7]
	fmt.Println(Q2)
	fmt.Println("Q2 len", len(Q2))
	fmt.Println("Q2 cap", cap(Q2))

	summer := months[6:9]
	fmt.Println(summer)
	fmt.Println("summer len", len(summer))
	fmt.Println("summer cap", cap(summer))

	afterSummer := summer[:4]
	fmt.Println(afterSummer)
	fmt.Println("afterSummer len", len(afterSummer))
	fmt.Println("afterSummer cap", cap(afterSummer))

	// afterSummer = summer[:8]
	// afterSummer = summer[4:]
}
