package main

import (
    "fmt"
                    //"math" it is a type of package
)

func main()  {


fmt.Println("hello world")

//switch statement
 /*  i := 10
switch {
case i <=10:
 fmt.Println("it is correct")
   fallthrough                  // This stmt means it will print both case  case i
case i <=20:
 fmt.Println("it is not a corect")
default:
fmt.Println("both statements correct")
} */


      // switch stmt model 2
      // switch types   switch 1:, switch i:= 10 -5;i , case i<=10
      /* i := 10
      switch  {
      case i <=10:
        fmt.Println("ohhh!  u choose me")
        fallthrough                    // if we this it will print both statement while the 2 cases is or not it will print both cases
      case i >=20:
        fmt.Println("ok ur going with another one")
      default:
        fmt.Println("bye bye")
      }    */




      //switch type 3
var i interface{} = 8
switch i.(type) {
    case int:
        fmt.Println(" it is integer type")
        break                         // if we use break hear it break and print another or same line
        fmt.Println("only in number")
      case float32:
          fmt.Println("it is float32 type")
      case string:
        fmt.Println("it is string type")
      case [2]int:
          fmt.Println("it is array type")
      default:
        fmt.Println("bye bye")
      }







}
