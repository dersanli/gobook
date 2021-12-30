package main

import "fmt"

var c, python, java bool // variable declaration, package or function level

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

func main() {
	var i, j int = 1, 2 // variable declaration with initializers (type (int) can be omitted)
	fmt.Println(i, j, c, python, java)

	fmt.Println(add(661651651, 6548498))

	a, b := swap("swap", "us") // handling multiple return values
	fmt.Println(a, b)

	fmt.Println(split(256))
}
