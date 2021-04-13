package main

import {
  "fmt"
  "time"
}

func main(){
  fmt.Println("Hello World")
  newtime := time.newtime(5*time.second)
  <-newtime.C
  fmt.Println("After 5 Seconds")
 }
