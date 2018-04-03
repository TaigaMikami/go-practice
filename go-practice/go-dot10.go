// 配列

package main

import "fmt"

func main() {
  // var a [5]int // a[0] to [5]
  // a[2] = 3
  // a[4] = 10
  // fmt.Println(a)
  // b := [3]int{1, 3, 5}
  b := [...]int{1,2}
  fmt.Println(len(b))
}
