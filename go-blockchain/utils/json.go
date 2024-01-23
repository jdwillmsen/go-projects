package utils

import "fmt"

func JsonStatus(message string) []byte {
	return []byte(fmt.Sprintf("{\"message\":\"%v\"}", message))
}
