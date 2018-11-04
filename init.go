package main

import (
	"fmt"
)

type Node struct {
	prev_index int
	next_index int
	index int
	status bool
}

func main(){
	tables := make([][]bool, 8)
	for i := 0; i < 8; i++ {
		tables[i] = make([]bool, 8)
	}
	fmt.Println("starts");
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(tables[i][j])
		}
		fmt.Print("\n")
	}
}


