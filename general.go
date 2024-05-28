package utils

func Max[T int | int32 | int64 | float32 | float64 | string](in ...T) T {
	var out T

	if len(in) == 0 {
		return out
	}

	out = in[0]
	for i := 1; i < len(in); i++ {
		if in[i] > out {
			out = in[i]
		}
	}

	return out
}

func Min[T int | int32 | int64 | float32 | float64 | string](in ...T) T {
	var out T

	if len(in) == 0 {
		return out
	}

	out = in[0]
	for i := 1; i < len(in); i++ {
		if in[i] < out {
			out = in[i]
		}
	}

	return out
}

