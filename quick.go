package main

import (
  "fmt"
  "math/rand"
)


/* TODO: find biggest / smallest element of shuffled list */

func quick_sort(list []int) {
  length, left, right := len(list), 0, len(list) - 1

  /* list is empty or one elem (sorted!) */
  if length < 2 { return }

  /* choose random pivot -- less screwed with sorted arrays */
  pivot := list[rand.Intn(length)]

  for left <= right {

    /* this element is fine */
    for list[left] < pivot {
      left++
    }

    /* this element is fine */
    for list[right] > pivot {
      right--
    }

    /* if we've not crossed over, these elements need swappin'! */
    if left <= right {
      list[left], list[right] = list[right], list[left]
      left++
      right--
    }
  }

  /* recurse */
  quick_sort(list[:right + 1])
  quick_sort(list[left:])
}

func main() {
  var list = []int{1,6,2,9,3,4,5,0}
  quick_sort(list)
  fmt.Println(list)

}
