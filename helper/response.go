package helper

import "net/http"

func SuccessResponses(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": message,
		"data":    data,
	}
}

func SuccessWithoutDataResponses(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"status":  "success",
		"message": message,
	}
}

func FailedResponses(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"status":  "failed",
		"message": message,
	}
}

func UnauthorizedResponses(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusUnauthorized,
		"status":  "failed",
		"message": message,
	}
}
