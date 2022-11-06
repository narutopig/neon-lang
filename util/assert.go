package util

// Assert checks if the given value is true, and throws an error if otherwise
func Assert(value bool, message string) {
	if !value {
		panic(message)
	}
}
