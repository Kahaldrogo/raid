package main

import (
	"fmt"
)

func test( n,x int) {
	fmt.Print("o")
	for i:= 0; i<n-2 ;i++ {
	
	      fmt.Print("-")
	  
	
	}
	if n >=2 {
	fmt.Print("o")
	}
	
	fmt.Print("\n")
	for j := 0; j<n-2 ;j++{
	fmt.Print("|")
	for i:= 0; i<n-2 ;i++ {
	
	      fmt.Print(" ")
	  
	
	}
	if n >=2 {
	fmt.Print("|")
	}
	}
	fmt.Print("\n")
	
	fmt.Print("o")
	for i:= 0; i<n-2 ;i++ {
	
	      fmt.Print("-")
	  
	
	}
	if n >=2 {
	fmt.Print("o")
	}
	
	
}
func main (){
   test(5,2)

}
