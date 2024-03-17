package main

import (
	"fmt"
	"math"
)

type thirdVtx struct {
	X, Y float64
}

// Methods: add a new method to a variable
// a method is a function + receiver argument
// Requirement: the type of receiver and method must be defined in the same package

func (v thirdVtx) Abs() float64 { // func receiver method FuncName(Input) returned type {}
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type ThisFloat float64 // float64 is built-in type which is not in the same package as the method
func (f ThisFloat) Abs() float64 {
	return float64(f * f)
}

// same as:
// func Abs(v thirdVtx) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }
// Abs(v)

// Pointer receivers Vs. Value receiver
// concept: if the type of receiver has a pointer, then it is a pointer receiver. Otherwise, it is
// a value receiver.
// - Methods with pointer receivers can take a value. Go interprets the value because the pointer
// receiver is there. Methods with value receivers can take a pointer, too.
// - Function with pointer arguments must take a pointer.

func (v *thirdVtx) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

// Interfaces: a set of methods (like java, python). Struct level.

type Abser interface {
	Abs() float64 // funcName | returned type
}

// Implement interfaces implicitly

type I interface {
	thisOne() // Our identifier
}

type Y struct {
	N string
}

func (y Y) thisOne() { // No explicit declaration of implementing interface
	fmt.Println(y.N, "implicit")
}

// Check nil
// if the concrete value inside the interface is nil, the method will be called with a nil receiver.
// A nil receiver is like a nullptr
func (v *thirdVtx) thisOne() {
	if v == nil {
		fmt.Println("<nil>")
		return
	}
	return
}

// Empty interface: an interface with no methods -> to handle unknown type
// It holds values of any type

type emptyI interface{}

// Type switches: like ordinary switch, but only switch type not value.
func swtch(i interface{}) {
	switch s := i.(type) { // like a type assertion, but use keyword `type`
	case int:
		fmt.Println("int", s)
	case float64:
		fmt.Println("float", s)
	case string:
		fmt.Println("string", s)
	default:
		fmt.Println("unknown", s)
	}
}

// Stringers
// In fmt package: type Stringer interface{String() string}
func (v thirdVtx) String() string {
	return fmt.Sprintf("ordinate (%d, %d)", int(v.X), int(v.Y)) // formatted string
}

func main() {
	v := thirdVtx{3, 4}
	v.Scale(0.1)
	fmt.Println(v.Abs())

	var abs Abser
	fi := 0.25
	q := thirdVtx{6, 8}
	abs = ThisFloat(fi) // An interface value holds any value that implements its methods
	fmt.Println(abs.Abs())
	abs = q
	fmt.Println(abs.Abs())

	// Interface values are like a tuple of (value, type). The value passed in to the interfaces
	// gets interpreted based on its type.
	var i I = Y{"interface"}
	i.thisOne()
	var nullVtx *thirdVtx
	i = nullVtx
	i.thisOne()

	// Nil interface values: a nil interface value holds nether value nor concrete type
	// Calling a method on a nil interface is a run-time err
	// var nilI I
	// nilI.thisOne()

	var nullI emptyI
	nullI = 3
	fmt.Println(nullI)
	nullI = "hello"
	fmt.Println(nullI)

	// Type assertions: t := i.(T), assign t the value of type T inside the interface i.
	var randi interface{} = "rand"
	s := randi.(string)
	fmt.Println(s)

	// a safeguard: If the string is not the type of any methods of the interface, the statement
	// returns false on the second argument.
	s, ok := randi.(string)
	fmt.Println(s, ok)

	f, ok := randi.(float64)
	fmt.Println(f, ok)

	// panic here
	// f = randi.(float64)
	// fmt.Println(f)

	swtch(23)
	swtch(7.2)
	swtch("hello")

	var er thirdVtx = thirdVtx{1, 4}
	fmt.Println(er)

}
