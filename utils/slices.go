package utils

// Filter will filter a slice of string using given condition
func Filter(a []string, condition func(string) bool) []string {
	n := 0
	for _, x := range a {
		if condition(x) {
			a[n] = x
			n++
		}
	}
	a = a[:n]

	return a
}
