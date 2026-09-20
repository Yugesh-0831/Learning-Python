package main

import (
	"fmt"
	"time"
)

// // interface definition -> a user defined data type which implements all the methods of a interfce are said to satisfy/implement that interface
// type Abser interface {
// 	Abs() float64
// }

// // any is alias for empty interface
// type any interface{}

// // empty interface
// var i interface{}

// func do(i interface{}) {

// 	// l = len(i) --> can't do since no value is assigned here

// 	switch i.(type) {
// 	case int:
// 		fmt.Println("int came")
// 	case string:
// 		fmt.Println("string came")
// 	default:
// 		fmt.Println("Don't know type")
// 	}
// }

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func Interface() {

	if err := run(); err != nil {
		fmt.Println(err)
	}

	// var a Abser
	// f := MyFloat(-math.Sqrt2)
	// v := Vertex{3, 4}
	// a = f
	// fmt.Println(a.Abs())

	// // you must assign as a pointer or val depending on how method signature is for that data type
	// a = &v

	// // a = v -----> problem
	// fmt.Println(a.Abs())

	// can assign anythign to an empty interface
	// i = 42

	// fmt.Println(i)

	// var i any
	// i = "hello"
	// fmt.Println(i)

	// do(42)
	// do("hello")
	// do(true)
}

// type MyFloat float64

// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }
