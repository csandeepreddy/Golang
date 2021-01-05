package main

 import  "fmt"
//func main(){

                           // defer function

  /*  fmt.Println("sandeep reddy")
    defer fmt.Println("bangalore")
    fmt.Println("jain university") */


                               // panic

/*    a , b := 1,0
    ans := a / b
    fmt.Println("ans")    */


/*  fmt.Println("start")
    defer fmt.Println("more more")
    panic("i have more doughts")
    fmt.Println("end")  /*



                     // Example for panic function

    // Simple Go program which illustrates
// the concept of panic

// Main function

	// Creating an array of string type
	// Using var keyword
	var myarr [3]string

	// Elements are assigned
	// using an index
	myarr[0] = "GFG"
	myarr[1] = "GeeksforGeeks"
	myarr[2] = "Geek"

	// Accessing the elements
	// of the array
	// Using index value
	fmt.Println("Elements of Array:")
	 defer fmt.Println("Element 1: ", myarr[0])

	// Program panics because
	// the size of the array is
	// 3 and we try to access
	// index 5 which is not
	// available in the current array,
	// So, it gives an runtime error
	fmt.Println("Element 2: ", myarr[1])          */

                          // recover

  func handlepanic() {

      if a := recover(); a != nil {
          fmt.Println("RECOVER", a)
      }
  }

  // Function
  func entry(lang *string, aname *string) {

      // Normal function
      handlepanic()

      // When the value of lang
      // is nil it will panic
      if lang == nil {
          panic("Error: Language cannot be nil")
      }

      // When the value of aname
      // is nil it will panic
      if aname == nil {
          panic("Error: Author name cannot be nil")
      }

      fmt.Printf("Author Language: %s \n Author Name: %s\n", *lang, *aname)
      fmt.Printf("Return successfully from the entry function")
  }

  // Main function
  func main() {

      A_lang := "GO Language"
      entry(&A_lang, nil)
      fmt.Printf("Return successfully from the main function")







}
