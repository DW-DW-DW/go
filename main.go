package main

import (
       "fmt"
       "runtime"
       "reflect"
       "os"
)

var (
    dogName string 
    dogWeight float64
    dogBreed = "Lurcher" //type inference
)
//Following Go Fundamentals course by Nigel Poulson on Pluralsight
//Using Sublime Text and Terminal
//Run this with "go run main.go"

func main () {
    sayHello()
    getDogInformation()
    doSimpleCalculation()
    haveFunWithPointers()
    testRandomFunctionality()
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
    breedPointer := &dogBreed
    fmt.Println("Memory address of dogBreed is", &breedPointer)
    //& references a pointer
    fmt.Println("The value of memory address", breedPointer, "is", *breedPointer)
    //* dereferences a pointer

    //You want to pass by reference
    //Otherwise Go will just create a copy of the variable
    changeBreedName(&dogBreed)
    fmt.Println("The value of memory address still is", breedPointer, 
        "\nBut value has changed to", *breedPointer)
}

func changeBreedName(dogBreed *string) string {
 
    *dogBreed = "Japanese Akita" 

    return *dogBreed
}

func testRandomFunctionality () {
    const daysOfTheWeek = 7 //can't use short hand with consts
    //const not used but Go doesn't seem to freak out like I thought it would

    for _, env := range os.Environ() {
        fmt.Println(env)
    }

}