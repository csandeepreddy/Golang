package main

import "fmt"

func main()  {


fmt.Println("hello world")

//switch statement
i := 10
switch {
case i <=10:
 fmt.Println("it is correct")
   fallthrough                  // This stmt means it will print both case  case i
case i <=20:
 fmt.Println("it is not corect")
default:
fmt.Println("both statements correct")

}
}
