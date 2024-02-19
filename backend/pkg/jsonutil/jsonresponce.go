package jsonutil

import (
	"encoding/json"
	"net/http"
)

func sendJSON(jsonData map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	// Преобразуем map в JSON
	responseData, err := json.Marshal(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок Content-Type как application/json
	w.Header().Set("Content-Type", "application/json")

	// Отправляем JSON данные на клиентскую сторону
	w.Write(responseData)
}
