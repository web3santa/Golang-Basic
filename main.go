package main

import "main.go/helpers"

func main() {
	println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "Some name"
	myVar.TypeNumber = 42

	println("TypeName:", myVar.TypeName)
	println("TypeNumber:", myVar.TypeNumber)
}
