package temp

const Boiling = 100.0

const offset = 32.0

func ToF(x float64) float64 {
	return convert(x)
}

func convert(x float64) float64 {
	return x*9/5 + offset
}
