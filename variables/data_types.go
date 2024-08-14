package main

import "fmt"

func main() {
	//int  int8  int16  int32  int64
	//positive whole number is unsigned int
	//uint uint8 uint16 uint32 uint64 uintptr
	//float32 float64
	var data1 string = "name1"

	//use assignment operator always
	data2 := "name2"

	//multiple variable in same line
	average,message :=24.5,"message here!"

	//constants
	const basicPlan = "basic plan"

	fmt.Println("happy nation")
	fmt.Println(data1)
	fmt.Println(data2)
	fmt.Println(message)
	fmt.Println(average)
	fmt.Println(basicPlan)
}
