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

func breakLoopWithConditional() {
	count := 1

	for {
		fmt.Println(count)
		if count == 100 {
			break
		}
		count += 1
	}
}
func main() {
	breakLoopWithConditional()
}
