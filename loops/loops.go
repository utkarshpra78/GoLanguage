package main

import "fmt"

func main() {
	for i:=0;i<10;i++{
		fmt.Println(i)
	}

	//while loop is for loop
    i:=5;
	for i>0{
		fmt.Println("less")
		i--;
	}

	//continue and break key word exista in go
	fmt.Println("loops added")
}