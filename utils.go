package main

import (
	"fmt"
	"runtime"
)

func SqRoot(x float64) float64 {
	z := 1.0
	k := 0.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if k == z {
			break
		}
		k = z
		fmt.Printf("SqRoot i: %v z: %v\n", i, z)
	}
	return z
}

func currentOS() {
	fmt.Print("Go runs on ")
	// 'switch' statement evaluates cases from top to bottom
	// _only_ runs the matched case
	switch os := runtime.GOOS; os { // a short statement is available
	// switch cases (os) need not be constants or integers
	case "darwin":
		fmt.Println("OS X.")
		// cases do not need a 'break' statement
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	/*
		// switch with no condition is like a cleaner if-else statement
			t := time.Now()
			switch {
			case t.Hour() < 12:
				fmt.Println("Good morning!")
			case t.Hour() < 17:
				fmt.Println("Good afternoon.")
			default:
				fmt.Println("Good evening.")
			}
	*/
}

func deferredPrint(deferredString string) { // A defer statement defers the execution of a function.


	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	defer fmt.Println(deferredString) // until the surrounding function returns
	fmt.Print("your string is: ")
}
