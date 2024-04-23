package main

import "fmt"

func main()  {
	//array 
	numbers := [5]int{1,34,55,77,90}
	for index,item := range numbers{
		//fmt.Println("Index and value is :",index,item)
		fmt.Println("Index is :",index,"Value is :",item)
	}
	rangeWithString()
	rangeWithMap()
	accesKeyUsingRang()
	sliceWithRange()
	sliceRange()
}

func rangeWithString(){
	strings := "Golang"
	fmt.Println("Index : Character")

	 // i access index of each character
     // item access each character

	 for i,item := range strings{
		fmt.Printf("%d = %c \n",i,item)
	 }
}

func rangeWithMap(){
	//create map
	subjectsMark := map[string]float32{"Java" : 80,"Golang" : 90, "PHP" : 90, "Android" : 95}
	fmt.Println("Marks Obtained :")
	for sub,marks := range subjectsMark{
		fmt.Println(sub,"=>",marks)
	}
}

func accesKeyUsingRang()  {
	//create map
	subjectsMark := map[string]float32{"Java" : 80,"Golang" : 90, "PHP" : 90, "Android" : 95}
	fmt.Println("Keys :")
	for sub := range subjectsMark{
		fmt.Println(sub)
	}
}

func sliceWithRange(){
	my_slice := []int{11,22,33,44,55,66,77,88}
	my_new_slice := my_slice[1:7]
	fmt.Println("My New Slice is :",my_new_slice)
	for i,value := range my_new_slice{
		fmt.Println(i,":",value)
	}
}

func sliceRange(){
	my_slice := []int{11,22,33,44,55,66,77,88}
	my_new_slice := my_slice[1:7]
	fmt.Println("My New Slice is :",my_new_slice)
	for i,value := range my_new_slice{
		fmt.Println(i,":",value)
	}
}