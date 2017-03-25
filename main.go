package main

import (
       "fmt"
       "runtime"
       "reflect"
)

var (
	dogName string 
	dogWeight float64
	dogBreed = "Lurcher"
)
//Following Go Fundamentals course by Nigel Poulson on Pluralsight

func main () {
	sayHello()
    getDogInformation()
    doSimpleCalculation()
    haveFunWithPointers()
}

func sayHello () {
	fmt.Println("Hello world")
    fmt.Println("Hello", runtime.GOOS)
}

func getDogInformation () {
	fmt.Println("Dog's name is", dogName)
	fmt.Println("Dog's weight is", dogWeight)
	fmt.Println("Dog's breed is", dogBreed)
	
	fmt.Println("dogBreed is", reflect.TypeOf(dogBreed))
}

func doSimpleCalculation () {
	a := 3.000000 //shorthand declare and initialize only inside functions
	b := 6.999

	c := a + b

	fmt.Println("Calculation result", c)
}

func haveFunWithPointers () {
	fmt.Println("Memory address of dogBreed is", &dogBreed)
}