package main

import (
	"fmt"
	"io"
	"strings"
)

// Errors: the `error` type is an interface as fmt.Stringer
// A nil `error` means success; a non-nill `error` means failure
// type error interface{
// 	Error() string
// }

type newError struct {
	s string
}

func (e *newError) Error() string {
	return fmt.Sprintf("new error is %s", e.s)
}

func anytime() error {
	return &newError{"chk"}
}

// Readers: io.Reader(), the read end of a stream of data
// func (T) Read(b []byte) (n int, err error) -> Returns io.EOF error when the read stream ends

// Generics
// Type parameters: appears between brackets & before the arguments.
// type T fulfills the built-in constraint `comparable`, making it possible to use `==`,`!=` on
// values of the type.
// func Index[T comparable](s []T, x T) int

// Generic types: nothing special. Just like oop of java.
// keyword `any`

type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	if err := anytime(); err != nil {
		fmt.Println(err)
	}

	r := strings.NewReader("Harry potter")
	b := make([]byte, 5)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:%v] = %q\n", n, b[:n])
		if err == io.EOF {
			break
		}
	}
}
