package main

import (
  "fmt"
)

type Heap struct {
  left *Heap
  value int
  right *Heap
}

/*

  INVARIANTS: 
    - Greatest value stays at top of heap.
 
  OPERATIONS
    - Insert
    - Remove
    - Balance(?)

*/

/* TODO: shouldn't only build up on the left hand side */
func insert(heap *Heap, val int) *Heap {
  if heap == nil {
    heap = &Heap{ nil, val, nil }
  } else if val > heap.value {
    heap.left = insert(heap.left, val)
  } else if val < heap.value {
    heap.left = insert(heap.left, heap.value)
    heap.value = val
  }
  return heap
}


func print_heap(heap *Heap) {
  if heap == nil {
    return
  }
  print_heap(heap.left)
  fmt.Printf("%d \n", heap.value)
  print_heap(heap.right)
}

func main() {

  var heap *Heap

  /* TODO: get values from somewhere more interesting (array?) */
  heap = insert(heap, 1)
  heap = insert(heap, 3)
  heap = insert(heap, 2)
  print_heap(heap)

}
