package main

import "fmt"

func main(){
	a,b,c := 3,4,5
	
	fmt.Println(isTriangle(a,b,c))

}

func isTriangle(a,b,c int) bool {
	if c*c == a*a + b*b {
		return true
	}
	return false
}
