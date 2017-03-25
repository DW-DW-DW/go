package main

import (
       "fmt"
       "runtime"
)

var (
	dogName string 
	dogWeight float64
)
//Following Go Fundamentals course by Nigel Poulson on Pluralsight

func main () {
	fmt.Println("Hello world")
    fmt.Println("Hello", runtime.GOOS)

    getDogInformation()
}

func getDogInformation () {
	fmt.Println("Dog's name is ", dogName)
	fmt.Println("Dog's weight is ", dogWeight)
}