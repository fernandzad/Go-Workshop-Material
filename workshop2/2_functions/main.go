package main

import (
	"fmt"
	"strings"
)

func main()  {
	numbers := []int{1,2,3,4,5,6,7,8,9,0}
	// Case 1 -- No returns
	printNumbers(numbers)

	// Case 2 -- No params
	sayHello()

	// Case 3 -- Variadic args
	fmt.Println(sumNumbers(numbers...))

	// Case 4 -- Multiple Returns
	sum,count := sumAndCountNumbers(numbers...)
	fmt.Println(sum,count)

	// Case 5 -- Named Returns
	name,lastname := changeValues("Gustavo Aguilar")
	fmt.Println("Name: ",name,"Last Name: ",lastname)

	// Case 6 - Anonymous function
	af := func(name,lastname string){
		fmt.Println("Welcome, ",name,lastname)
	}
	af(name,lastname)
	af2 := func(){
		fmt.Println("To this Go Bootcamp!")
	}
	af2()
	func(a,b int){
		fmt.Println(a,"*",b," is:",a*b)
	}(2,10)
}

func printNumbers(numbers []int){
	fmt.Println("Your numbers ->",numbers)
}

func sayHello(){
	fmt.Println("Hello there!")
}

func sumNumbers(numbers ...int) int{
	sum := 0
	for _,v := range numbers{
		sum+=v
	}
	return sum
}

func sumAndCountNumbers(numbers ...int) (int,int){
	count:=0
	for i:=0;i<len(numbers);i++{
		count++
	}
	return sumNumbers(numbers...),count
}

func changeValues(s string) (name,lastname string){
	split := strings.Split(s," ")
	name ,lastname = split[0],split[1]
	return
}