package main

import "fmt"

func main() {
	fmt.Println("helloo Golang") //sequential
	foo()

	fmt.Printf("exited foo")
}

func foo() {
	for i := 0; i < 20; i++ { //loop
		if i%2 == 0 { //conditional
			fmt.Println(i)
		}
	}
}
