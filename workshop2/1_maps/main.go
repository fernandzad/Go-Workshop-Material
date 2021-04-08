package main

import (
	"fmt"
	"sort"
)

//https://blog.golang.org/maps

func main() {
	// DECLARATION
	var m map[string]int // This only allocates memory but doesn't initialize our map
	// m["my-key"]=1 // This will panic because we are trying to do an assignment to a 'nil' map
	fmt.Println("My map: ", m)
	fmt.Println()

	// DECLARATION & INITIALIZATION
	m = make(map[string]int) // This allocates memory AND initializes our map
	m["my-key-1"] = 100
	m["my-key-2"] = 200
	m["my-key-3"] = 300
	fmt.Println("My map with values assigned: ", m)
	fmt.Println()

	// VALUE RECOVERY
	// Retrieves the value stored under the key
	v := m["my-key-2"]
	fmt.Println("The value of 'my-key-2' in my map is: ", v)
	// If the requested key doesn't exist, we get the value type's zero value. In this case the value type is int, so the zero value is 0
	fmt.Println("The value of 'my-key-4' in my map is: ", m["my-key-4"])
	fmt.Println()

	// BUILT-IN FUNCTIONS
	// The built in len function returns on the number of items in a map:
	fmt.Println("The length of my map is: ", len(m))
	// The built in delete function removes an entry from the map:
	delete(m, "my-key-2")
	// The delete function doesn't return anything, and will do nothing if the specified key doesn't exist.
	fmt.Println("My map with 'my-key-2' deleted:", m)
	fmt.Println()

	// VALIDATIONS
	// A two-value assignment tests for the existence of a key:
	v, ok := m["my-key-4"]
	// In this statement, the first value (v) is assigned the value stored under the key "my-key-4".
	// If that key doesn't exist, v is the value type's zero value (0).
	// The second value (ok) is a bool that is true if the key exists in the map, and false if not.
	fmt.Println("Checking existence and value of key 'my-key-4': ", v, ok)
	// To test for a key without retrieving the value, use an underscore in place of the first value:
	_, ok = m["route"]
	fmt.Println("Checking existence of key 'my-key-4': ", ok)
	fmt.Println()

	// ITERATION
	// When iterating over a map with a range loop, the iteration order is not specified
	// and is not guaranteed to be the same from one iteration to the next.
	// If you require a stable iteration order you must maintain a separate data structure that specifies that order.
	// This example uses a separate sorted slice of keys to print a map[int]string in key order:
	scores := make(map[string]int)
	scores["A"]=11
	scores["B"]=22
	scores["C"]=33
	scores["D"]=44
	scores["E"]=55
	scores["F"]=66
	scores["G"]=77
	for key, value := range scores {
		fmt.Println("Key:", key, "Value:", value)
	}
	// We create a new slice that stores our keys
	var keys []string
	for k := range scores {
		keys = append(keys, k)
	}
	fmt.Println("My Keys: ",keys)
	// Sorting our keys
	sort.Strings(keys)
	fmt.Println("My Keys (Ordered): ",keys)
	// We iterate over our map sorted by key
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", scores[k])
	}
}
