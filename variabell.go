package main

import "fmt"

func main() {
	var lastName string = " My"
	const firstName string = " Skill"
	var bilanganBulat uint8 = 20
	var bilanganDesimal = 2.5

	var numberA int = 4
	var numberB *int = &numberA
	*numberB = 8
	fmt.Print("Halo", firstName, lastName, bilanganBulat, bilanganDesimal, numberA, "!\n")
}
