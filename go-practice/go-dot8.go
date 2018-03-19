package main

import "fmt"

func hi(name string) (msg string){
  // fmt.Println("hi!" + name)
  // msg := "hi!" + name
  // return msg

  msg = "hi" + name
  return
}

func main() {
  // hi("taguchi")
  fmt.Println(hi("taguchi"))
}
