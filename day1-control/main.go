package main

import "fmt"

func dayType(day string) string {
	switch day {
	case "Saturday", "Sunday":
		return "weekend"
	default:
		return "weekday"
	}
}

func power(base, exp int) (int, error) {
	if exp < 0 {
		return 0, fmt.Errorf("negative exponent %d not supported", exp)
	}
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result, nil
}

func main() {
	fmt.Println(dayType("Sunday"))

	for _, e := range []int{0, -1} {
		p, err := power(3, e)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		fmt.Println("2 to the", e, "=", p)
	}
}
