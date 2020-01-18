package utils

// CheckError util
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
