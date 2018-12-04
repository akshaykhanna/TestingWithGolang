package TestingWithGolang

import "fmt"

func add(numbers ...int) interface{} {
	sum := 0

	// missing in code coverage
	if numbers == nil {
		fmt.Println("No arguments passed")
		return 0
	}
	for _, n := range numbers {
		sum += n
	}
	return sum
}
