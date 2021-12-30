package main

import "fmt"

func add(x, y int) int { //shortened parameter type declaration
	return x + y
}

func swap(x, y string) (string, string) { // multiple return from a function
	return y, x
}

func split(sum int) (x, y int) { // named return values (x, y)
	x = sum * 4 / 9
	y = sum - x
	return // naked return
}

func main() {
	fmt.Println(add(661651651, 6548498))
	a, b := swap("swap", "us") // handling multiple return values
	fmt.Println(a, b)
	fmt.Println(split(256))
}
