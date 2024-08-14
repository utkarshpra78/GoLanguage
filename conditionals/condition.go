package main

import "fmt"

func getLength(email string)int{
	return len(email);
}

func main() {
	email:="utkarsh_108"
	//one way to declare the condition
	length := getLength(email)
	if length < 1 {
		fmt.Println("Email is invalid")
	}

	//another way to declare the condition
	if length := getLength(email); length < 1 {
		fmt.Println("Email is invalid")
	}
	fmt.Println("this is condition")
}