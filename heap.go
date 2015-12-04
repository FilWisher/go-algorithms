package main

import (
  "fmt"
)

type Tree struct {
  left *Tree
  value int
  right *Tree
}

/*

  INVARIANTS: 
    - Greatest value stays at top of tree.
 
  OPERATIONS
    - Insert
    - Remove
    - Balance(?)

*/

func main() {

  fmt.Printf("Hello\n")

}
