package handler

import "encoding/json"

func JsonError(msg string) []byte {
	errorMsg := struct {
		Message string `json:"message"`
	}{
		msg,
	}

	r, err := json.Marshal(errorMsg)
	if err != nil {
		return []byte(err.Error())
	}
	return r
}
