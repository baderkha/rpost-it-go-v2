package http

import (
	"strconv"
	"strings"
)

const (
	DefaultHttpCode = 500
)

func StatusFromError(err error) int {
	codeStr := strings.Split(err.Error(), ",")[0]
	status, _ := strconv.Atoi(codeStr)
	if status == 0 {
		return DefaultHttpCode
	}
	return status
}
