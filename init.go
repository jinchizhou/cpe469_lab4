package main

import (
	"fmt"
  "time"
)

type Node struct {
	prev_index int
	next_index int
	index int
	status bool
}

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

// prints out the data from Nodes
func print_Node(nodes []Node){
  for i := 0; i < 8; i++ {
    fmt.Println("Candidate # %s, prev_index: %s, next_index: %s, curr_index: %s, status: ", i, nodes[i].prev_index, nodes[i].next_index, nodes[i].index, nodes[i].status)
  }
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
  // init a list of node's global status
  candidates := []Node{
      {
        prev_index: 7,
      	next_index: 1,
      	index: 0,
      	status: true,
      },
      {
        prev_index: 0,
      	next_index: 2,
      	index: 1,
      	status: true,
      },
      {
        prev_index: 1,
      	next_index: 3,
      	index: 2,
      	status: true,
      },
      {
        prev_index: 2,
      	next_index: 4,
      	index: 3,
      	status: true,
      },
      {
        prev_index: 3,
      	next_index: 5,
      	index: 4,
      	status: true,
      },
      {
        prev_index: 4,
      	next_index: 6,
      	index: 5,
      	status: true,
      },
      {
        prev_index: 5,
      	next_index: 7,
      	index: 6,
      	status: true,
      },
      {
        prev_index: 6,
      	next_index: 0,
      	index: 7,
      	status: true,
      },
  }
  // print data from each Node
  print_Node(candidates)
  // print status of heartbeats
  print_table(tables)
  // every 0.1 sec, exchange table
  time_ex_table := time.Now()
  // every 5 sec, one node dies
  time_node := time.Now()
  for alive(tables){
    // exchange tables
    if (time.Since(time_ex_table).Seconds() > 0.2){

      fmt.Println(time.Since(time_ex_table).Seconds())
      time_ex_table = time.Now()
    }
    // one node dies
    if (time.Since(time_node).Seconds() > 5){

      fmt.Println("One dies ", time.Since(time_node).Seconds())
      time_node = time.Now()
    }
  }

}
