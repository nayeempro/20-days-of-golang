// ================== Array =====================
// Arrays in Go are fixed-size collections of elements, all of the same data type.
// They provide a way to store multiple values in a single variable,
// enabling efficient access and manipulation of data.

// ## How Arrays Work in Go
/*
1. An array is a sequence of elements, each indexed by a position, starting from 0.
2. In Go, arrays have a fixed size, meaning once an array is created, its length cannot change.
3. Arrays are strongly typed; an array defined for integers cannot hold string or float values.
4. Arrays in Go are value types, meaning when you assign one array to another, a copy is created, not a reference.
*/

// ## Declaring and Initializing Arrays

var arr [5]int  // Declares an array of 5 integers, all initialized to 0 by default.

arr := [5]int{1, 2, 3, 4, 5}  // Infers the data type as [5]int based on the provided values.

var arr = [5]int{1, 2}  // Remaining elements will be initialized to 0, so arr = [1, 2, 0, 0, 0].

arr := [...]int{1, 2, 3, 4, 5}  // Go determines the array length automatically based on the elements.



var intArray [5]int               // Integer array
var stringArray [3]string         // String array
var floatArray = [4]float64{2.5, 3.8, 4.9, 1.2}  // Array with float values

package main

import "fmt"

func main() {
	fmt.Println("Array concept")
}

//=================================Slice=====================================

// Slices in Go
// Slices in Go are more flexible than arrays, allowing dynamic resizing and more
//  versatile data handling. Slices build on arrays but have additional functionality,
//   making them ideal for many programming tasks.


// How Slices Differ from Arrays
// Dynamic Size: Unlike arrays, slices do not have a fixed size; they can grow and shrink as needed.
// References to Arrays: A slice is a reference to a segment of an array. Modifying the slice affects the underlying array.
// Built-in Functions: Functions like append and copy make slices highly flexible and efficient for data manipulation.

// Declare
1. From an array
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4]  // Creates a slice with elements from index 1 to 3, so slice = [2, 3, 4]

2. make function
slice := make([]int, 5)  // Creates a slice of integers with length 5, all elements initialized to 0

3. slice Literal
slice := []int{1, 2, 3, 4, 5}  // Creates and initializes a slice with specific values

// some manipulation
1.Appending elements

slice := []int{1, 2, 3}
slice = append(slice, 4, 5)  // slice now contains [1, 2, 3, 4, 5]

2. Copying

src := []int{1, 2, 3}
dest := make([]int, len(src))
copy(dest, src)  // Copies elements of src into dest

3. slicing and reSlicing
data := []int{1, 2, 3, 4, 5, 6}
subData := data[2:5]          // subData = [3, 4, 5]
resliced := subData[1:]       // resliced = [4, 5]


// Understanding Maps in Go
// Maps in Go are collections that store data as key-value pairs. They are highly 
// useful for cases where you need to associate unique identifiers (keys)
// with specific data (values).

// Declare (map[keyType]valueType)
var ageMap map[string]int                // Declares a map (initially nil)
ageMap = make(map[string]int)            // Initializes the map to allow adding entries
idMap := map[int]string{1: "Alice", 2: "Bob"} // Declares and initializes a map with two entries


// Map with custom type
userMap := map[string][]string{"Alice": {"Math", "Physics"}} // Map with a slice of strings as values


// Adding ,updating, deleting 
ageMap["John"] = 25    // Adds a new entry with "John" as the key and 25 as the value
ageMap["Doe"] = 30     // Adds another entry
ageMap["John"] = 26    // Updates the value associated with "John" to 26
delete(ageMap, "Doe")  // Removes the entry with key "Doe"

// loop in go
for key, value := range ageMap {
    fmt.Printf("Key: %s, Value: %d\n", key, value)
}
for key := range ageMap {
    fmt.Println("Key:", key)
}






