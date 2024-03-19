package main

// notes taken from A Tour of Go: https://go.dev/tour/

// import packages: factored import statement is a good style

// single package:
// import "fmt"
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// variable declaration: var name type = expression
var foo, bar int
var pig, apple = "xxx", "yyy" // omit type when an initializer is present

// functions: func name(variable type) returned type
func add(x int, y int) int {
	return x + y
}

// function parameters share a type: omit the type from all but the last
func addi(x, y int) int {
	return x + y
}

// return any number of results (a bit like python tuples)
func swap(x, y string) (string, string) {
	return y, x
}

//named return values
func split(sum int) (x, y int) {
	x = sum + 1
	y = sum - 1
	return //naked return for named return values; suitable only in short functions (readability)
}

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("Hello, olc")
	fmt.Println("The time is", time.Now())
	fmt.Println("number is", rand.Intn(100))
	fmt.Printf("..... %g ....\n", math.Sqrt(7))

	// exported names start with capitals; unexported names of a package cannot be accessed from outside
	fmt.Println(math.Pi)
	fmt.Println(add(1, 3))

	// short variable declaration: declare at least one variable
	var c rune = 'a'  // double quotes represent a string composed of chars; single quotes represent a char
	c, d := 'q', "ss" // assign c a new value and declare d as a new var
	fmt.Printf("%q\n", c)
	fmt.Println(c, d)

	a, b := swap("crazy", "rabbit") // := is a declaration; = is an assignment
	fmt.Println(b, a)
	fmt.Println(split(4))
	fmt.Println(foo, bar) // unassigned, 0; false; "" (like python)
	fmt.Println(pig, apple)

	// type conversions (like python)
	i := 42
	f := float64(i)
	u := uint(f)
	fmt.Println(i, f, u)

	//type inference: number type depends on the precision of right-hand number if it is untyped
	v := 42.1
	j := 3 + 0.5i // i is an imaginary number...
	fmt.Printf("type(v) %T, type(j) %T\n", v, j)

	//constants: cannot be declared using :=
	const world = "green"
	fmt.Println("leaves are", world)

	const truth = true
	fmt.Println("2>1?", truth)

	fmt.Println(needInt(Small))
	//    fmt.Println(needInt(Big)) overflow
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
