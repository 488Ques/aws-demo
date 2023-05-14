package controllers

func buildErrorJSON(err error) map[string]interface{} {
	return map[string]interface{}{
		"message": err.Error(),
	}
}

func buildResponseJSON(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"data": data,
	}
}
