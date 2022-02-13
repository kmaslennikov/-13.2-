package main

import (
	"fmt"
)

func main() {
  fmt.Println("Введите два числовых значения:")
	var first,	second int
  fmt.Scan(&first)
  fmt.Scan(&second)
  
  fmt.Println("Значения до:")
	fmt.Println(first)
	fmt.Println(second)

	change(&first, &second)

  fmt.Println("Значения после:")
	fmt.Println(first)
	fmt.Println(second)

}

func change(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
