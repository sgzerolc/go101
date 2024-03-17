package main

import (
	"fmt"
	"strings"
)

type Vtx struct {
	X, Y float64
}

// key, value
// zero value of map is nil. A nil is empty and cannot be added keys
func genMap() {
	var m map[string]Vtx
	m = make(map[string]Vtx)
	m["sun"] = Vtx{ // like python dictionary {}
		0.1, 0.1,
	}
	fmt.Println(m["sun"])

	var n = map[string]Vtx{
		// a map literal is: 1) must contain key
		"water": Vtx{0.2, 0.2},
		// can omit type when it is same as the top level of map
		"mouse": {0.3, 0.3},
	}
	fmt.Println(n["mouse"])

	// insertion
	n["book"] = Vtx{0.4, 0.4}
	// retrieve an element
	kBook := n["book"]
	fmt.Println(kBook)
	delete(n, "water")
	// return 0 if is_elem is not there; ok is boolean
	isElem, ok := m["no"]
	fmt.Println("The value:", isElem, "Present?", ok)
}

func myAdd(x float64, y float64) float64 {
	return x + y*2
}
func myMultipy(x float64, y float64) float64 {
	return x * y * 2
}

// function values: think about the function pointer in C. It's a bit odd
// that function is the argument of input
func cal(fn func(float64, float64) float64) float64 {
	return fn(4, 5) // the input is fixed for the function that is passed to cal
}

// function closures
func adder() func(int) int {
	i := 0
	return func(x int) int { // imply a new argument of input (closure: referenced var)
		i += x
		return i
	}
}

func WordCount(s string) map[string]int {
	var dict map[string]int
	dict = make(map[string]int)
	words := strings.Fields(s)

	for _, w := range words {
		dict[w] = len(w)
	}
	return dict
}

func main() {
	genMap()
	m := WordCount("apple ambition coke koala")
	for k, v := range m {
		fmt.Printf("%s -> %d\n", k, v)
	}

	fmt.Println(cal(myAdd)) // function values
	fmt.Println(cal(myMultipy))

	a, b := adder(), adder()
	for i := 0; i < 5; i++ {
		fmt.Println(a(i), b(2*i))
	}
}
