package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Структура для JSON-ответа
type Response struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Разрешаем CORS
	w.Header().Set("Access-Control-Allow-Origin", "*") // Здесь укажите ваш фронтенд URL вместо "*", если необходимо
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		// Обрабатываем preflight запрос
		return
	}

	// Создаем ответ в формате JSON
	response := Response{
		Message: "Привет, от Голанга на Апи Шлюзе!",
		Status:  true,
	}

	// Устанавливаем заголовок Content-Type для JSON
	w.Header().Set("Content-Type", "application/json")

	// Кодируем ответ в JSON и отправляем его
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/api/hello-world", helloHandler)
	fmt.Println("Сервер запущен на порту 5000...")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Println(err)
	}
}
