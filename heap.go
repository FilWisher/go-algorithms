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

func insert(heap *Heap, val int) *Heap {
  insertion := val
  if heap == nil {
    heap = &Heap{ nil, val, nil }
    return heap
  } else if val > heap.value {
    insertion = heap.value
    heap.value = val
  }
  // not sure about the costs of these 'assignments' when already assigned
  if heap.left != nil {
    heap.right = insert(heap.right, insertion)
  } else {
    heap.left = insert(heap.left, insertion)
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
