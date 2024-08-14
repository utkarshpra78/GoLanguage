package main

import "fmt"

func sub(x int, y int) int {
	return x - y
}

func concat(s1 string, s2 string) string {
	return s1 + s2
}

//multiple parameter of same data types
func add(x, y, z int) int {
	return x + y + z
}

//function as a parameter
func f(a int, b int) int {
	return a * b
}

func f1(f2 func(int, int) int, x int) int {
	return f2(x, x)
}

//pass by reference
func inc(x *int) {
	*x++
}

//multiple return data types

func getPoint(x int, y int) (int, int) {
	return x, y
}

func main() {
	x := 5
	fmt.Println("function")
	fmt.Println(sub(6, 2))
	fmt.Println(concat("hello ", "world"))
	fmt.Println(add(2, 3, 4))
	fmt.Println(f1(f, 5))
	inc(&x)
	fmt.Println(x)
	y, _ := getPoint(3, 4)
	fmt.Println(y)
}
