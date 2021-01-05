package main

import (
  "fmt"
)

func main() {

 // fmt.Println("hello my golang")
 /* for i := 0 ; i <5 ; i ++ {
 sayMessage("hello world",i)
}
}
 func sayMessage(msg string, idx int) {
   fmt.Println(msg)
   fmt.Println("the value is ", idx)
 }   */






/*  sayMessage("i love u","sweety")
}
func sayMessage(greeting string,name string) {
  fmt.Println(greeting, name)
  name = "ted"
  fmt.Println(name)
}                */






 s := sum(1,2,3,4,5)
fmt.Println("The sum is " , s)
}

func sum(values ...int) (result int) {
              fmt.Println(values)
for _ , v := range values {
  result +=  v
         }
return
}
