package main

import "fmt"

const boilingPointOfWater = 373.00

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
