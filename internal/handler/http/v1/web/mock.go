package webv1

import "encoding/json"

func CreateExpectedResponse(data interface{}) string {
	response := AppResponse{
		Data:    data,
		Success: true,
	}
	jsonResponse, _ := json.Marshal(response)
	return string(jsonResponse) + "\n"
}
