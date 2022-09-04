package generics

import "fmt"

type Ordered interface {
	int | float64 | string
}

func MinArray[T Ordered](values []T) (T, error) {
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

func MinArgs[T Ordered](values ...T) (T, error) {
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
