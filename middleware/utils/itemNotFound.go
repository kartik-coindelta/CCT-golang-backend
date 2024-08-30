package utils

import (
	"net/http"
)

// ItemNotFound checks if an item is found and returns an error if not.
func ItemNotFound(err error, item interface{}, message string) map[string]interface{} {
	if err != nil {
		return BuildErrObject(http.StatusUnprocessableEntity, err.Error())
	}
	if item == nil {
		return BuildErrObject(http.StatusNotFound, message)
	}
	return nil
}
