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
// func init_status(){
//   nodes := []node{
//
//   }
//   return nodes;
// }

func print_table(tables [][]bool){
  for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(" ", tables[i][j])
		}
		fmt.Print("\n")
	}
}

// checks if at least one node is alive
func alive(tables [][]bool) bool{
  for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if tables[i][j]{
        return true
      }
		}
	}
  return false
}

func main(){
  // initialize global table with every node's status
	tables := make([][]bool, 8)
	for i := 0; i < 8; i++ {
		tables[i] = make([]bool, 8)
    for j := 0; j < 8; j++ {
      tables[i][j] = true
    }
	}
  // print status of heartbeats
  print_table(tables)
  for alive(tables){
    
  }
}
