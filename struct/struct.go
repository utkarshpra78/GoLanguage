package main

import "fmt"

//simple struct
type car struct {
	name   string
	model  string
	height int
	width  int
}

type messageToSend struct {
	phoneNUmber int
	message     string
}

//nested struct in GO
type cars struct {
	name       string
	model      string
	doors      int
	mileage    int
	frontWheel wheel
	backWheel  wheel
}

type wheel struct {
	radius   int
	material string
}

//anynomus struct in go

type car1 struct {
	name   string
	model  string
	doors  int
	wheels struct {
		radius   int
		material string
	}
}

//embedded structs

type cars2 struct {
	name  string
	model string
}

type truck struct {
	cars2
	bedSize int
}

//struct method in GO
type rect struct {
	width  int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}

func main() {
	//making instane of struct
	myCar := cars{}
	myCar.frontWheel.radius = 5

	//anynomus struct in go
	myCar1 := struct {
		name  string
		doors int
	}{
		name:  "tesla",
		doors: 4,
	}
	myCar1.doors = 5
	myCar1.name = "TATA"

	//embedded VS nested
	lanesTruck := truck{
		bedSize: 10,
		cars2: cars2{
			name:  "tesla",
			model: "suv",
		},
	}
	fmt.Println(lanesTruck.bedSize)

	//struct method in go
	var r = rect{
		width:  5,
		height: 5,
	}
	fmt.Println(r.area())
	fmt.Println("hello there")
}
