package main

import (
	"fmt"
)

func main() {
	first := 10
	second := 90
	fmt.Println(first)
	fmt.Println(second)

	change(&first, &second)

	fmt.Println(first)
	fmt.Println(second)

}

func change(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
