package main
import "fmt"

func main() {
	fmt.Println(sum(10, 10))
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func times(a int, b int) int {
	return a * b
}

func sumA(a int, b int) int {
	return a + b + a
}

func sumB(a int, b int) int {
	return a + b + b
}
