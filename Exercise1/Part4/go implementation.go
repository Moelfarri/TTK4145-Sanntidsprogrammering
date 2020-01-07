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
    
    channel_1 := make(chan bool)
    channel_2 := make(chan bool)
    
    go increasing(channel_1)
    go decreasing(channel_2)
    
    
    <-channel_1
    <-channel_2
	                              
    Println(i)
}
