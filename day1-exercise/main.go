package main

import "fmt"

func classify(n int) string {
	switch {
	case n < 0:
		return "negative"
	case n == 0:
		return "zero"
	default:
		return "positive"
	}
}

func sumTo(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("sumTo: n must be non-negative, got %d", n)
	}
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum, nil
}

func main() {
	for _, n := range []int{5, 0, -3, 10} {
		class := classify(n)
		total, err := sumTo(n)
		if err != nil {
			fmt.Println(n, class, "error:", err)
			continue
		}
		fmt.Println(n, class, total)
	}
}
