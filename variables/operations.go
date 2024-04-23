package main

import "fmt"

func add(a,b int) int{
	return a + b
}

func boolean() bool{
	var a bool
	a = false
	return a
}

func pattern1(num int){
	for i := 1; i<= num; i++ {
		for j := 1; j<= i; j++ {
			fmt.Print(" *")
		}
		fmt.Println()
	}
}

func pattern2(num int){
	for i := 1; i <= num; i++ {
		for j:= 1; j<= i; j++ {
			fmt.Print(" *")
		}
		fmt.Println()
	}
	for a := num-1; a >= 1; a-- {
		for b := 1 ; b <= a ; b++ {
			fmt.Print(" *")
		}
		fmt.Println()
	}
}

func matrix(){
	var a [2][3]int
	n := 1
	for i := 0; i < 2; i++ {
		for j := 0; j < 3 ; j++ {
			a[i][j] = n;
			fmt.Print(a[i][j])
			n++;
		}
		fmt.Println();
	}
}

func imatrix(){
	var a[3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {			
			if(i == j){
				a[i][j] = 1				
			}else{
				a[i][j] = 0
			}	
			fmt.Print(a[i][j])		
		}
		fmt.Println();
	}
}

func pointer1(){
	var x,y *int //DECLARING AS A POINTER VARIABLE
	b := 6
	c := 10
	x = &c //STORING DETAILS TO ANOTHER VARIABLE
	y = &b
	fmt.Println("Value of x is ",*x)
	fmt.Println("Value of y is ",*y)
}

func recursion(n int){
	z := fact(n)
	fmt.Printf("Factorial : %d\n" , z)
}
func fact(n int)int {
	if n == 1 {
		return 1
	} else  {
		return n * fact(n-1)
	}
}



func recursionPattern(n int){
	if n == 1 {
		fmt.Println(" *")
	}else{
		for i := 1 ; i<= n; i++ {
			fmt.Print(" *")
		}
		fmt.Println()
		recursionPattern(n-1)
	}
}