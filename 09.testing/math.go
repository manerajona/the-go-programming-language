package math

// Adder accepts diferent int and return the sum of them
func Adder(xs ...int) (result int) {
	for _, v := range xs {
		result += v
	}
	return
}
