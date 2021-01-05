package main

import (
  "fmt"
)

func main(){

  fmt.Println("hello golang")

/*
  a := 42
  b := a
fmt.Println(a, b)
a = 54
fmt.Println(a,b)   */


/* var  a int  =67
var  b *int =&a
fmt.Println(a, *b)
a = 87
fmt.Println(a, *b)
    *b = 45
fmt.Println(a, *b)  */


/* a := [3] int{1,2,3}
 b := &a[0]
 c := &a[1]
 d := &a[2]
 fmt.Println(" %v ,%p ,%p\n" , a, b, c ,d)   */


  /*  var ms *myStruct
 ms = new(myStruct)
 ms.foo = 42
 fmt.Println(ms.foo)

}
type myStruct struct {
   foo int                  */


 a :=[3]int{1,2,3}
 b := a
 fmt.Println(a,b)
 a [1] =42
 fmt.Println(a , b)




}
