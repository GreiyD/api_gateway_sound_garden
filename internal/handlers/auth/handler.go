package auth

import (
	"api_gateway/internal/config"
	httpHandler "api_gateway/internal/transport/http"
	"fmt"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request, conf *config.AuthService) {
	requestUrl := conf.Url + "/v1/users/register"

	resp, err := httpHandler.ProxyRequest(requestUrl, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при отправке запроса: %v", err), http.StatusInternalServerError)
		return
	}

	body, statusCode, err := httpHandler.ProcessServiceResponse(resp)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to process service response: %v", err), http.StatusInternalServerError)
		return
	}

	httpHandler.RespondToClient(w, body, statusCode)
}

func LoginUser(w http.ResponseWriter, r *http.Request, conf *config.AuthService) {
	requestUrl := conf.Url + "/v1/users/login"

	resp, err := httpHandler.ProxyRequest(requestUrl, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при отправке запроса: %v", err), http.StatusInternalServerError)
		return
	}

	body, statusCode, err := httpHandler.ProcessServiceResponse(resp)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to process service response: %v", err), http.StatusInternalServerError)
		return
	}

	httpHandler.RespondToClient(w, body, statusCode)
}
