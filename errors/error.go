package main

import (
	"errors"
	"fmt"
	"strconv"
)

//error is a interface
type error interface {
	Error() string
}

//custom error interface

type userError struct{
	name string
}

func (e userError) Error() string{
	return fmt.Sprintf("name is %v",e.name)
}

//formatting strings
func formatString(name string ,age float64){
	s:=fmt.Sprintf("%v is %v years old",name,age);
	fmt.Println(s) 
	//prints upto 2 decimal places 
	fmt.Printf("I am %.2f years old\n", 10.523)
}

func main() {
	atoiValue, err := strconv.Atoi("97b")
	if err != nil {
		//custon error call
		fmt.Println(userError{name:"utkarsh prakash"})
		return
	}
	fmt.Println(atoiValue);

	formatString("utkarsh",23.456)


	//error package
	var _ error = errors.New("something went wrong")
	
	fmt.Println("error handling")
}
