package main

import  (

  "fmt"
   "time"

)


/*  var msg= "hello"
   go func() {

    fmt.Println(msg)
   }

}               */



//In the above example, we simply create a display() function and then call this function\new in two different ways first one is a Goroutine, i.e. go display(“Welcome”) and another one is a normal function, i.e. display(“GeeksforGeeks”). But there is a problem, it only displays the result of the normal function that does not display the result of Goroutine because when a new Goroutine executed, the Goroutine call return immediately. The control does not wait for Goroutine to complete their execution just like normal function they always move forward to the next line after the Goroutine call and ignores the value returned by the Goroutine.




// Go program to illustrate
// the concept of Goroutine
/*  func display(str string) {
	for w := 0; w < 6; w++ {
		fmt.Println(str)
	}
}

func main() {

	// Calling Goroutine
	go display("Welcome")

	// Calling normal function
	display("to golang land")
} */






// Go program to illustrate the concept of Goroutine
// Go program to illustrate how
// to create an anonymous Goroutine
// Main function
func main() {

	fmt.Println("Welcome!! to Main function")

	// Creating Anonymous Goroutine
	 func() {

		fmt.Println("Welcome!! to goland")
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("GoodBye!! to Main function")
}
