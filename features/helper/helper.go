package helper

func ResponseFailed(message string) map[string]interface{} {
	return map[string]interface{}{
		"status" : "Error",
		"message": message,
	}
}

func ResponseSuccessNoData(message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "Success",
		"message": message,
	}
}

//	Nggak dipakai
func ResponseSuccessWithData(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "Success",
		"message": message,
		"data":    data,
	}
}