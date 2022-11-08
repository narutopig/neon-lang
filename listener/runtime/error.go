package runtime

// Error is used for capturing errors without using the default Go error (allows for a list of errors)
type Error struct {
	Message string
}

// E returns a new Error
func E(message string) Error {
	return Error{Message: message}
}
