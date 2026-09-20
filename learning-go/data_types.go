package main

import (
	"fmt"
	"math"
)

// struct
type Vertex struct {
	X, Y float64
}

// var m map[string]Vertex

// // we just need to give a name and signature of func
// func compute(fn func(float64, float64) float64) float64 {
// 	return fn(3, 4)
// }

// pass by value
func (v *Vertex) Abs() float64 {
	v.X = 2
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// pass by reference
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func DataTypes() {
	// ------------ methods ------------

	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(v)

	v.Scale(10)
	fmt.Println(v)

	// ------------- func -------------

	//functions in golang are first class citiens (can be assigned to a variable and can be returned and passed in reference to another function)
	// hypot := func(x, y float64) float64 {
	// 	return math.Sqrt(x*x + y*y)
	// }
	// fmt.Println(hypot(3, 4))
	// fmt.Println(compute(hypot))
	// fmt.Println(compute(math.Pow))

	// golang also supports closure -> to study the concept use case!!

	// ------------- maps -------------
	// m = make(map[string]Vertex)

	// m["Bell Labs"] = Vertex{1, 2}

	// fmt.Println(m["Bell Labs"])

	// var m map[string]int
	// error! never initialise map with this syntax either -> m := map[string]int{} or see below

	// m := make(map[string]int)
	// m["a"] = 1
	// fmt.Println(m["a"])
	// m["a"] = 2
	// fmt.Println(m["a"])

	// delete(m, "a")
	// fmt.Println(m["a"])

	// syntax to check if value is present or not
	// v, ok := m["a"]
	// fmt.Println("value: ", v, "Present: ", ok)

	// for key, value := range m {
	// 	fmt.Println(key, value)
	// }

	// ---------- struct, arrays, slices -----------

	// v := Vertex{1, 2}
	// v.Y = 4
	// p := &v

	// // p.X is same as  -> (*p).X
	// p.X = 2
	// fmt.Println(v)

	// type V struct {
	// 	X int
	// }
	// v1 := V{1}
	// v2 := v1
	// v2.X = 2

	// p1 := &v1
	// fmt.Println(v1, v2, *p1)

	// arrays
	// var a [2]int
	// a[0] = 1

	// fmt.Println(a)

	// primes := [6]int{2, 3, 5, 7, 11, 13}
	// // fmt.Println(primes)

	// // slices
	// var s []int = primes[1:4]
	// fmt.Println(s, len(s))

	// names := [4]string{
	// 	"John",
	// 	"Paul",
	// 	"George",
	// 	"RIngo",
	// }

	// a := names[0:2]
	// b := names[1:3]

	// fmt.Println(a, b)

	// b[0] = "XXX"

	// // Every value gets changed since slice are reference types
	// fmt.Println(a, b, names)

	// s := []int{2, 3, 5, 7, 11, 13}

	// // slices capacity changes based on starting index/ pointer of the slice to the last element of the backing array
	// // but length is simply based on no of elements in the slice
	// s = s[1:4]
	// fmt.Println(s, len(s), cap(s))

	// s = s[:5]
	// fmt.Println(s, len(s), cap(s))

	// s = s[:2]
	// fmt.Println(s, len(s), cap(s))

	// s = s[1:]
	// fmt.Println(s, len(s), cap(s))

	// var s []int
	// if s == nil {
	// 	fmt.Println("Starting value of reference type is Nil!")
	// }

	// board := [][]string{
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// }

	// for i := 0; i < len(board); i++ {
	// 	fmt.Printf("%s \n", strings.Join(board[i], " "))
	// }

	// var s []int
	// fmt.Printf("Value: %v \tPointer: %p \tType: %T\n", s, s, s)

	// // always need to reassign as golang creates new backing array
	// s = append(s, 1)
	// fmt.Printf("Value: %v \tPointer: %p \tType: %T\n", s, s, s)

	// s = append(s, 2, 3, 4)
	// fmt.Printf("Value: %v \tPointer: %p \tType: %T\n", s, s, s)

	// var s []int
	// for i := 0; i < 20; i++ {
	// 	s = append(s, i)
	// 	// Capacity doubles when backing array is full
	// 	fmt.Printf("Length: %d, Capacity: %d, Pointer: %p \n", len(s), cap(s), s)
	// }

	// better syntax to loop -> but v is a copy not the reference here!!
	// for i, v := range s {
	// 	fmt.Println(i, v)
	// }
}
