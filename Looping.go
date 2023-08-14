/*
package main

import "fmt"

	func main() {
		for i := 1; i <= 100; i++ {
			fmt.Println("Angka", i)
		}
	}
*/
package main

import "fmt"

func loopEdu() {
	for i := 1; i <= 33; i++ {
		fmt.Println("Edu - Angka", i)
	}
}

func loopWork() {
	for i := 34; i <= 66; i++ {
		fmt.Println("Work - Angka", i)
	}
}

func loopEduWork() {
	for i := 67; i <= 100; i++ {
		fmt.Println("EduWork - Angka", i)
	}
}

func main() {
	loopEdu()
	loopWork()
	loopEduWork()
}
