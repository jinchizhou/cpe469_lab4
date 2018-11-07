package main

import (
  "fmt"
  "time"
  "sync"
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
func alive(cand []Node) bool{
  for i := 0; i < 8; i++ {
    if (cand[i].status){
      return true
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
// kills one Node and updates status
func killOneNode(live_table [][]bool, candidates []Node){
  for i:=0; i < 8; i++{
    if (candidates[i].status){
      // kill it, update statuses
      live_table[i][i] = false
      candidates[i].status = false
      return
    }
  }
}
func exchangeTables(i int, live_table [][]bool, candidates []Node, connector []int, wg *sync.WaitGroup, mutex *sync.RWMutex){
  defer wg.Done()
  mutex.Lock()
  // assume if status of candidate is false, it is dead, so return
  if (!candidates[i].status){
    mutex.Unlock()
    return
  }
  p_index := candidates[i].prev_index
  n_index := candidates[i].next_index
  // if previous is dead
  if (!candidates[p_index].status){
    connector[0] = i
  }
  // if next is dead
  if (!candidates[n_index].status){
    connector[1] = i;
  }
  // loop through prev_index and update on live_table
  for j:=0; j < 8; j++{
    if(!live_table[p_index][j]){
      live_table[i][j] = false
    }
  }
  mutex.Unlock()
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
  var wg sync.WaitGroup
  mutex := sync.RWMutex{}
  for alive(candidates){
    // one node dies
    if (time.Since(time_node).Seconds() > 2){
      killOneNode(tables, candidates)
      fmt.Println("One dies ", time.Since(time_node).Seconds())
      time_node = time.Now()
    }
    // exchange tables
    if (time.Since(time_ex_table).Seconds() > 0.2){
      // connect two if one died
      connector := make([]int, 2)
      connector[0] = -1
      connector[1] = -1
      // exchange tables
      for i:= 0; i < 8; i++{
        wg.Add(1)
        go exchangeTables(i, tables, candidates, connector, &wg, &mutex)
      }
      wg.Wait()
      if (connector[0] != -1){
        // one node died, so connect connector[0] to connector[1]
        candidates[connector[0]].prev_index = connector[1]
        candidates[connector[1]].next_index = connector[0]
      }
      print_table(tables)
      fmt.Println(time.Since(time_ex_table).Seconds())
      time_ex_table = time.Now()
    }
  }
  // done
}
