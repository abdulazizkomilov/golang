package main

// func main() {
// 	fmt.Println("Hello World!")
// }

// func greeting(name string) string {
// 	return "Hello " + name
// }

// func main() {
// 	var name string
// 	fmt.Print("Enter your name: ")
// 	fmt.Scanln(&name)
// 	fmt.Println(greeting(name))
// }

// func sum(a int, b int) int {
// 	return a + b
// }

// func main() {
// 	var a, b int
// 	fmt.Print("Enter a: ")
// 	fmt.Scanln(&a)
// 	fmt.Print("Enter b: ")
// 	fmt.Scanln(&b)
// 	fmt.Println(sum(a, b))
// }

// func main() {
// 	var x int = 10
// 	y := 20
// 	c := x + y
// 	fmt.Println(c)
// }

// func main() {
// 	num := 10
// 	if num >= 5 {
// 		fmt.Println("num is greater than 5")
// 	} else {
// 		fmt.Println("num is less than 5")
// 	}
// }

// func main() {
// 	num := 1
// 	for num <= 10 {
// 		fmt.Println(num)
// 		num++
// 	}
// }

// func main() {
// 	day := "Friday"
// 	switch day {
// 	case "Monday":
// 		fmt.Println("Start of the work week!")
// 	case "Friday":
// 		fmt.Println("Almost the weekend!")
// 	default:
// 		fmt.Println("Middle of the week!")
// 	}
// }

// func main() {
// 	var arr [3]int = [3]int{1, 2, 3}
// 	fmt.Println(arr)
// }

// func main() {
// 	slice := []int{1, 2, 3, 4}
// 	slice = append(slice, 5)
// 	fmt.Println(slice)
// }

// func main() {
// 	var x = 5
// 	var p *int = &x
// 	fmt.Println(*p)
// }

// type Person struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	person := Person{"Alice", 30}
// 	fmt.Println(person.Name, person.Age)
// }

// func sum(a int, b int, result chan int) {
// 	result <- a + b // Send sum to channel
// }

// func main() {
// 	result := make(chan int)
// 	go sum(3, 4, result)
// 	fmt.Println(<-result) // Receive result from channel
// }
