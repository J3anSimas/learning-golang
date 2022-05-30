package main

import "fmt"

func generateMultTable() {
	for i := 1; i < 11; i++ {
		for j := 1; j < 11; j++ {
			fmt.Println(fmt.Sprintf("%v x %v = %v", i, j, i*j))
		}
		fmt.Println()
	}
}

func main() {
	generateMultTable()
}
