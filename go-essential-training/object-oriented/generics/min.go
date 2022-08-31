package main

import "fmt"

type Ordered interface {
	int | float64 | string
}

func min[T Ordered](values []T) (T, error) {
	if len(values) == 0 {
		var zero T
		return zero, fmt.Errorf("empty slice")
	}

	m := values[0]
	for _, v := range values {
		if v < m {
			m = v
		}
	}
	return m, nil
}

func main() {
	fmt.Println(min([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(min([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}))
	fmt.Println(min([]string{"joe", "billy", "perry", "biden", "lilly"}))
}
