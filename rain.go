package main

import (
    "fmt"
)

func main() {
    var name string
    var age int
    var birthYear int

    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)

    fmt.Print("Enter your age: ")
    fmt.Scanln(&age)

    fmt.Print("Enter your birth year: ")
    fmt.Scanln(&birthYear)

    fmt.Println("\n--- Your Information ---")
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Birth Year:", birthYear)
}