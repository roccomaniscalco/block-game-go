package main

import "fmt"

func main() {

	arr := [5]int{1,2,3,4,5}

	for _, num := range arr {
		fmt.Println(num)
	}
}
