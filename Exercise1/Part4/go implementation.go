package main

import (
    . "fmt"
    "runtime"
    "time"
)

var i = 0

func increasing(c chan<- bool) {
   for j := 0; j < 1000000; j++ {
		i++
	}
}

func decreasing(c chan<- bool) {
  for j := 0; j < 1000000; j++ {
		i--
	}
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())   
    
    thread_1 := make(chan bool)
    thread_2 := make(chan bool)
    
    go increasing(thread_1)
    go decreasing(thread_2)
    
    
    <-thread_1
    <-thread_2
	                              
    Println(i)
}
