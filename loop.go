package main

import (
	"fmt"
	"math"
	"runtime"
	"strings"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func Sqrt(x float64) float64 {
	z, iter := 1.0, 1
	var next float64
	limit := 0.000001
	next = z - (z*z-x)/(2*z) // initial guess
	fmt.Println(next)
	for math.Abs(next-z) > limit {
		z = next
		next -= (z*z - x) / (2 * z)
		iter++
	}
	fmt.Printf("guess %d times \n", iter)
	//    for ;iter < 10; iter++ {
	//        z -= (z*z - x) / (2*z)
	//        fmt.Printf("guess[%d] is %g\n", iter, z)
	//    }
	return z
}

func power(x, n, r float64) float64 {
	// if with short declaration; the short statement and variable inside if are shared with else blocks.
	if v := math.Pow(x, n); v < r {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, r)
	}
	return r
}

func goswitch() {
	fmt.Print("Go is switching ")
	switch os := runtime.GOOS; os { // only support constants
	//evaluation order: top->bottom until one case goes through
	case "linux": // break is automatically provided in go. Only runs the selected case once.
		fmt.Println("Linux.")
	case "darwin":
		fmt.Println("OS x.")
	default:
		fmt.Printf("%s don't care\n", os)
	}

	// switch without condition
	t := time.Now()
	switch { // same as switch true
	case t.Hour() < 12:
		fmt.Println("honestly don't care")
	case t.Hour() < 17:
		fmt.Println("fff")
	default:
		fmt.Println("bye")
	}
}

func godefer() {
	defer goswitch() // defers returning a call until the surrounding function returns but evaluates arguments imediately
	fmt.Println("deferring")
}

// defer, panic, recover
func stackdefer() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // last in, first out
	}

	fmt.Println("done")
}

func gopointer() { // null ptr: nil
	// unlike C, go has no pointer arithmetic
	i := 4
	p := &i
	fmt.Printf("type of p is %T %d\n", p, *p)
}

type Vertex struct {
	X int
	Y int
}

var ( // struct literals
	v1 = Vertex{2, 4}
	v2 = Vertex{X: 1} // field name: value, meaning that let x of v2 be 1
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), s)
}

func goarray() {
	var a [2]string // [number]Type - array
	a[0] = "Hi"     // accessed by a[0], pretty straight forward
	a[1] = "Bye"
	primes := [6]int{2, 3, 5, 7, 11, 13} //declare first, right part is an array literal
	fmt.Println(a, primes)

	// slices: a[low:high], default bounds for slices are zero. It is like python list slices[start, step, stop] without step
	var s []int = primes[1:4] //reassign.
	fmt.Println(s)

	// It's references to arrays, which does not store actual data.
	s[0] = 100
	fmt.Println(primes, s)

	// slice literals: array literal without the length
	q := []int{24, 4, 4, 4, 4, 4, 4}
	fmt.Println(q[:2]) // index, no negative value
	q = q[:0]
	printSlice(q)
	q = q[:4]
	printSlice(q)
	q = q[2:] // Capacity counts from the first element. Here it dropped two elements
	printSlice(q)
	var si []int
	printSlice(si)
	if si == nil {
		fmt.Println("hey nil!")
	}

	// allocate a dynamically-sized arrays. In C, must alloacte/free manually. In python, it doesn't care
	ai := make([]int, 5) //len
	printSlice(ai)
	b := make([]int, 0, 5) // type, len, cap
	printSlice(b)
	b = b[:4]
	printSlice(b)
	b = b[1:]
	printSlice(b)

	// slices of slices: like double arr
	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"=", "-", "="},
	}

	// change value of index (0, 1)
	board[0][1] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], ""))
	}

	// go append: it appends new element (same type) to the original slice. If the capacity of slice is large enough, it just
	// places the new value. Otherwise, it allocates a new arr to hold all the given values
	var cont []int
	printSlice(cont)
	cont = append(cont, 0)
	printSlice(cont)
	cont = append(s, 2, 3, 4)
	printSlice(cont)

	// ranging over a slice, it returns (index, copy of arr[index]) in each iteration.
	var pow = []int{1, 2, 4, 8}
	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}

	// short for:
	// for i, _ := range pow
	// for _, value := range pow
	// for i := range pow  --> only index
}

func Pic(dx, dy int) [][]uint8 {
	pixels := make([][]uint8, dy)
	fmt.Println(len(pixels), cap(pixels), pixels)
	for i := 0; i < dy; i++ {
		pixels[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			pixels[i][j] = uint8((j + i) / 2)
		}
		fmt.Println(pixels[i])
	}
	return pixels
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ { //no parentheses
		sum += i
	}
	fmt.Println(sum)

	for sum < 100 { // omit-expression in for is just like C
		sum += sum
	}
	fmt.Println(sum)
	for sum < 1000 { // While in C
		sum += sum
	}
	fmt.Println(sum)

	// will run forever
	// for {
	// }

	fmt.Println(
		sqrt(3), sqrt(-3),
		power(4, 2, 13),
	)

	//    for i := 1; i < 10; i++ {
	//        fmt.Printf("Sqrt (%d) is %g\n", i, Sqrt(float64(i)))
	//    }

	//    goswitch()
	//    godefer()
	//    stackdefer()
	//    gopointer()

	//    v := Vertex{1, 2} // initialize a struct
	//    fmt.Println(v.X)
	//    v.X = 4
	//    fmt.Println(v.X)
	//
	//    p := &v
	//    p.X = 1e9 // go permits this notation. Should be (*p).X
	//    fmt.Println(v)
	fmt.Println(v1, p, v2, v3)
	goarray()
	Pic(3, 4)
}
