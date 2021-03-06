package number

import "math"

func ContinuedFraction(f float64, eps ...float64) ([]int, int, int) {
	e := 1e-3
	if len(eps) > 0 {
		e = eps[0]
	}

	list := make([]int, 0)

	r := f
	for {
		t := math.Trunc(r)
		list = append(list, int(t))

		diff := r - t
		if diff < e {
			break
		}

		r = 1.0 / diff
	}

	if len(list) == 1 {
		return list, 1, list[0]
	}

	n, d := 1, list[len(list)-1]
	for i := 2; i < len(list); i++ {
		n, d = d, list[len(list)-i]*d+n
	}

	return list, n, d
}

func IsOdd(v int) bool {
	return !IsEven(v)
}

func IsEven(v int) bool {
	if v%2 == 0 {
		return true
	}

	return false
}
