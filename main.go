package main

import "fmt"

var c, python, java = true, false, "string!" // variable declaration, package or function level
const (
	Pi = 3.14 // constant
	k  = 1 << 100
)

func add(x, y int) int { // shortened parameter type declaration
	return x + y
}

func swap(x, y string) (string, string) { // multiple return from a function
	return y, x
}

func split(sum int) (x, y int) { // named return values (x, y)
	x = sum * 4 / 9
	y = sum - x
	return // naked return (can harm readability in longer functions)
}

func intConvert(x int) float64 {
	return float64(x)
}

func forLoop(max int) (int, int) {
	sum := 0
	for i := 0; i < max; i++ {
		sum += i
	}

	/*
		// init and post statements are optional
			sum := 1
			for ; sum < max; {
				sum += sum
			}
	*/

	/*
		// 'while' is 'for' as well
			sum := 1
			for sum < max {
				sum += sum
			}
	*/

	/*
		// infinite loop
			for {
			}
	*/
	return max, sum
}

func main() {
	var i, j int = 1, 2 // variable declaration with initializers (type (int) can be omitted)
	k := 5              // short assigment statement (only used inside a function)
	fmt.Println(i, j, k, c, python, java)

	fmt.Println(add(661651651, 6548498))

	a, b := swap("swap", "us") // handling multiple return values
	fmt.Println(a, b)

	fmt.Println(split(256))

	t := intConvert(256) // type inference (it understands t is float64)
	fmt.Printf("Type: %T Value: %v\n\n", t, t)

	max, sum := forLoop(100)
	fmt.Printf("max: %v sum: %v\n\n", max, sum)

}
