package main

import (
 "fmt"
 "time"
)

func main(){
  fmt.Println("Hello World")
  newtime := time.NewTimer(5*time.Second)
  <-newtime.C
  fmt.Println("After 5 Seconds")
 }
