package main

import "fmt"

func main() {
  a := 5
  var pa *int
  pa = &a // aのアドレス
  //pa の領域にあるデータの値 = *pa
  fmt.Println(pa)
  fmt.Println(*pa)
}
