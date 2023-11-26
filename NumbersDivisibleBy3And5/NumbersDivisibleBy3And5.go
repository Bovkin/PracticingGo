package main

import "fmt"

const boilingPointOfWater = 373.00

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Pin-Pan")
		} else if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}
