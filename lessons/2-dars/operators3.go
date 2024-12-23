// mantiqiy operatorlar

package main

import "fmt"

func operators3() {
	a := 14
	b := 30

	// &&(va)
	if a != b && a <= b {
		fmt.Println("true")
	}

	// ||(yoki)
	if a != b || a <= b {
		fmt.Println("true")
	}

	// !(yo'q)
	if !(a == b) {
		fmt.Println("true")
	}

}
