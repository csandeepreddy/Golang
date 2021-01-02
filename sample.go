package main

import "fmt"

func main() {
	var i int
	var j int
	i=45
	j=600
	fmt.Println(i * j)

type	biodata struct{   //struct
reg int
name string
age int
marks []int

}                         //struct



type father struct{
        name  string
				land  int
}
type son struct{
        father
				car   string
				bike  string
				speed  float32
			}


// without var and declaring variables

/*name :="sandeep"
fmt.Println("my name is " + name)

//without var and int

name="suresh"
k :=24
fmt.Printf("%v, %T", k,name)

//percentage of printf value = %t and %var name type

var (
name string="sandeep"
c int=999
)

fmt.Println(name , c)

var m int=65
var sname string="sudheer"
sname=string(m)
fmt.Println("my value is ",m,"and my name is "+sname)


//

var m int=67
var sname string
 sname=strconv.Itoa(m)
 fmt.Println(sname)

//captial i=integer and a standsfor= ascii */

var z bool=true
n := 1 ==1

fmt.Printf("%T , %v\n" ,z , z )
fmt.Printf("%T , %v\n" ,n , n )

s:="sandeep"
fmt.Printf("%v,%T" ,string(s[2]) ,s[2])

alpha:="sandeep"
show:=[]byte(alpha)

fmt.Printf("%v,",show)

fmt.Println("/n/n")
grades:= [7]int {78,64,90,53,65,87,43}
fmt.Println(grades[0])

fmt.Println("/n/n")
fmt.Println("/n/n")
fmt.Println("/n/n")

//population
statepopulation:= map[string]int{
  "india": 634234678683,
	"usa": 2384783748284,
	"mp":6374632874672834627,
}
fmt.Println(statepopulation["usa"])
statepopulation["japan"]=78234683248 // adding statepopulation
fmt.Println(statepopulation["japan"])
delete(statepopulation,"india") //deleting statepopulation
fmt.Println(statepopulation)
fmt.Println(len(statepopulation))
sp:=statepopulation
delete(sp, "usa")
fmt.Println(statepopulation)

//struct

details:=biodata{
	reg:76,
	name:"sandeep",
	age: 38,
	marks:[]int{
		23,23,23,23,23,
	},

}

fmt.Println(details.age)




adoctor:=struct {name string}{name:"sandeep reddy"}
anotherdoctor:=&adoctor              //we give "&" it will print same name two times
anotherdoctor.name="suresh reddy"

fmt.Println(adoctor.name)
fmt.Println(anotherdoctor.name)



				b :=  son{}
				b.name= "srinivasulu"
				b.land= 8
				b.car= "lamborghni"
				b.bike= "ktm"
				b.speed= 80
	fmt.Println(b.name )
	fmt.Println(b.land )
	fmt.Println(b.car )
	fmt.Println(b.bike )















}
