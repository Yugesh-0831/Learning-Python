package main

// func add(x int, y int) int {
// 	return x + y
// }

// // add in brackets if multiple values are returned
// func swap(x, y string) (string, string) {
// 	return y, x
// }

// var c, python, java bool

func main() {
	// a, b := swap("Hello", "World")
	// fmt.Println(a, b)

	//NOTE: Walrus operator only accessible inside a function
	// i := 0
	// fmt.Println(i, c, python, java)

	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("OS X.")
	// case "linux":
	// 	fmt.Println("Linux.")
	// default:
	// 	fmt.Printf("%s. \n", os)
	// }

	// today := time.Now().Weekday()

	// // You can evaluate a condition in switch statemnts
	// switch time.Monday {
	// case today + 0:
	// 	fmt.Println("Today.")
	// case today + 1:
	// 	fmt.Println("Tomorrow.")
	// case today + 2:
	// 	fmt.Println("In two days.")
	// default:
	// 	fmt.Println("Too far away.")
	// }

	//NOTE: switch can be used for multiple if else if situation/

	// for i := 0; i < 10; i++ {
	// 	// defer pushes data in a lifo stack
	// 	defer fmt.Println(i)
	// }

	//pointers ->
	// i, j := 42, 2701

	// p := &i
	// fmt.Println(*p)
	// *p = 21
	// fmt.Println(i)

	// p = &j
	// *p = *p / 37
	// fmt.Println(j)

	// DataTypes()
	// Interface()
	Concurrency()
}
