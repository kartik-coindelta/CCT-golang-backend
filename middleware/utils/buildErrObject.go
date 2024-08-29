package utils

// BuildErrObject creates a structured error object.
func BuildErrObject(statusCode int, message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  statusCode,
		"message": message,
	}
}
