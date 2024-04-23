package main

import "fmt"


func sliceDemo(){
	var my_slice = []int{1,2,3,4,5,6,7,8,9} 
	fmt.Println("Slice",my_slice)
	fmt.Printf("Length of Slice is = %d and capacity is = %d",len(my_slice),cap(my_slice)) //
	fmt.Println()
	

	var my_slice_add = append(my_slice,10)
	fmt.Println("New Slice",my_slice_add)
	fmt.Printf("Length of Slice is = %d and capacity is = %d",len(my_slice_add),cap(my_slice_add)) //
	fmt.Println()
	

	var my_slice2 = my_slice_add[1:5]
	fmt.Println("Slice 2 ",my_slice2)
	fmt.Printf("Length of Slice 2 is = %d and capacity is = %d",len(my_slice2),cap(my_slice2))
	fmt.Println()
	

	var my_slice3 = my_slice[2:6]
	fmt.Println("Slice 3 ",my_slice3)
	fmt.Printf("Length of Slice 3 is = %d and capacity is = %d",len(my_slice3),cap(my_slice3))
	fmt.Println()
	

	var my_slice4 = my_slice[3:6]
	fmt.Println("Slice 4 ",my_slice4)
	fmt.Printf("Length of Slice 4 is = %d and capacity is = %d",len(my_slice4),cap(my_slice4))
	fmt.Println()

		
	// fmt.Println("Original Slice ",my_slice)
	// fmt.Printf("Length of Original Slice is = %d and capacity is = %d",len(my_slice),cap(my_slice))
	// var elem = my_slice[3]
	// var deletedIndex = append(my_slice[:3], my_slice[3+1:]...)
	// fmt.Println()
	// fmt.Printf("The element %d was deleted.\n", elem)
    // fmt.Println("Slice after deleting elements:", my_slice)
	// fmt.Printf("Length of Original Slice is = %d and capacity is = %d",len(deletedIndex),cap(deletedIndex))
	// fmt.Println()

	// fmt.Println("Original Slice ",my_slice)
	// fmt.Printf("Length of Original Slice is = %d and capacity is = %d",len(my_slice),cap(my_slice))
	// var elem = my_slice[3]
	// var deletedIndex = append(my_slice[:3], my_slice[3+1:]...)
	// fmt.Println()
	// fmt.Printf("The element %d was deleted.\n", elem)
    // fmt.Println("Slice after deleting elements:", my_slice)
	// fmt.Printf("Length of Original Slice is = %d and capacity is = %d",len(deletedIndex),cap(deletedIndex))
	// fmt.Println()

	var index int = 3
     
    // Get the element at a given index in the slice
    var elems = my_slice[index]
     
    // Call the function delete_at_index which 
    // returns a slice of integers
    var numbers = delete_at_index(my_slice, index)
 
    fmt.Printf("The element %d was deleted.\n", elems)
    fmt.Println("Slice after deleting element:", numbers)
	fmt.Printf("Length of Original Slice is = %d and capacity is = %d",len(numbers),cap(numbers))
	fmt.Println()
}

func delete_at_index(slice []int, index int) []int {
     
    // Append function used to append elements to a slice
    // first parameter as the slice to which the elements 
    // are to be added/appended second parameter is the 
    // element(s) to be appended into the slice
    // return value as a slice
    return append(slice[:index], slice[index+1:]...)
}