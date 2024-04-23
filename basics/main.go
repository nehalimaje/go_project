package main

import "fmt"

type person struct{
	fname string
	lname string
}

type secreteAgent struct{
	person
	licenseToKill bool
}

//SOC INTERFACE EXAMPLE BLOCK
type human interface{
	speak()
}

func interfaceDemo(h human){
	fmt.Println("Using interfaceDemo")
	h.speak()
	
}
//EOC INTERFACE BLOCK

func(p person)speak(){
	fmt.Println(p.fname,p.lname,`says,"Hello this is my first Go program."`)
}

func(sa secreteAgent) speak(){
	fmt.Println(sa.fname,sa.lname,`says,"Shaken, not stirred."`)
}

func main(){
	xi := []int{2,3,4,5,}
	fmt.Println(xi)

	m := map[string]int{
		"Rober" : 45,
		"James" : 50,
	}
	fmt.Println(m)

	p1 := person{
		"Neha",
		"Limaje",
	}
	fmt.Println(p1)
	p1.speak()

	sag := secreteAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	sag.speak()

	//SOC INTERFACE BLOCK
	interfaceDemo(p1)
	interfaceDemo(sag)
	//EOC INTERFACE BLOCK

	fmt.Println("......From Slice file......")
	sliceDemo()
}