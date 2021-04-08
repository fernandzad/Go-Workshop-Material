package main

import "fmt"

func main() {
	// Array declarations
	// Long way
	var myArray [5]int
	myArray[0]=2
	myArray[1]=4
	myArray[2]=6
	myArray[3]=8
	myArray[4]=10
	// Short way
	myArray2 := [10]int{}
	// Short way with composite literal
	myArray3 := [15]int{2,3,9,12,15}
	// As a pointer
	myArray4 := new([20]int)
	fmt.Println("Array 1 ->", myArray)
	fmt.Println("Array 2 ->", myArray2)
	fmt.Println("Array 3 ->", myArray3)
	fmt.Println("Array 4 ->", myArray4)

	// Slice declarations
	// Long way
	var mySlice []int
	//mySlice[0]=2 // THIS WILL PANIC (The zero value of a slice is nil)
	// Short way
	mySlice2 := []int{}
	// Short way with composite literal
	mySlice3 := []int{2,3,9,12,15}
	// As a pointer
	mySlice4 := new([]int)
	// Using Make
	mySlice5 := make([]int,5)
	fmt.Println("Slice 1 ->", mySlice)
	fmt.Println("Slice 2 ->", mySlice2)
	fmt.Println("Slice 3 ->", mySlice3)
	fmt.Println("Slice 4 ->", mySlice4)
	fmt.Println("Slice 5 ->", mySlice5)

	// CAP & LEN
	fmt.Println("LEN -> ", len(mySlice5)," CAP-> ",cap(mySlice5))
	mySlice6 := make([]int,5,10)
	fmt.Println("Slice 6 ->", mySlice6)
	fmt.Println("LEN -> ", len(mySlice6)," CAP-> ",cap(mySlice6))
	mySlice6 = append(mySlice6,6)
	fmt.Printf("Address-> %p\n",mySlice6)
	mySlice6 = []int{0,0,0,0,0,6} // This changes the underlying array
	fmt.Printf("Address-> %p\n",mySlice6)
	fmt.Println("Slice 6 ->", mySlice6)
	fmt.Println("LEN -> ", len(mySlice6)," CAP-> ",cap(mySlice6))
	mySlice6 = append(mySlice6,7,8,9,10,11) // We are exceeding original capacity at this point
	fmt.Println("Slice 6 ->", mySlice6)
	fmt.Println("LEN -> ", len(mySlice6)," CAP-> ",cap(mySlice6))
	fmt.Println()

	// Basic Operations
	// Append
	myColors := []string{"Red", "Green","Blue"}
	fmt.Println("My colors ->", myColors)
	myColors = append(myColors,"White")
	fmt.Println("My colors ->", myColors)

	// Slicing -> T[from : up to but not included]
	// Slicing - Last element
	fmt.Println("Slicing last element", myColors[:len(myColors)-1])
	// Slicing - First element
	fmt.Println("Slicing first element", myColors[1:])
	// Slicing - N element
	fmt.Println("Slicing third element", myColors[:2], myColors[3:])

	// Variadic function
	myNewColors := myColors[:2]
	myNewColors = append(myNewColors, myColors[3:]...)
	fmt.Println("My new colors ->", myNewColors)
}
