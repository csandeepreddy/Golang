package main

import (

  "fmt"

)
//  func main() {



  // Go program to illustrate
  // how to create a channel

  	// Creating a channel
  	// Using var keyword
  /*	var mychannel chan int
  	fmt.Println("Value of the channel: ", mychannel)
  	fmt.Printf("Type of the channel: %T ", mychannel)

  	// Creating a channel using make() function
  	mychannel1 := make(chan int)
  	fmt.Println("\nValue of the channel1: ", mychannel1)
  	fmt.Printf("Type of the channel1: %T ", mychannel1)
}  */

func myfunc(ch chan int) {

	fmt.Println(234 + <-ch)
}
func main() {
	fmt.Println("start Main method")
	// Creating a channel
	ch := make(chan int)

	go myfunc(ch)
	ch <- 567    //receive the data 
	fmt.Println("End Main method")
}
