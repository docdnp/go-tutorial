// Example: Invalid duplicate declaration
package main

func add(a, b int) int {
	return a + b
}

func add(a, b float64) float64 {
	return a + b
}

func main() {
	println("Sum:", add(5, 5))
}
