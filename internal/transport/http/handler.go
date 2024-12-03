package http

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func DecodeJSON(w http.ResponseWriter, r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)
	if err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return fmt.Errorf("error decoding JSON: %w", err)
	}

	return nil
}

func ProxyRequest(url string, r *http.Request) (*http.Response, error) {
	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		return nil, err
	}

	req.Header = r.Header

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func decodeResponseBody(resp *http.Response) (io.ReadCloser, error) {
	if resp.Header.Get("Content-Encoding") == "gzip" {
		return gzip.NewReader(resp.Body)
	}
	return resp.Body, nil
}

func ProcessServiceResponse(resp *http.Response) ([]byte, int, error) {
	reader, err := decodeResponseBody(resp)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	defer reader.Close()
	defer resp.Body.Close()

	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return body, resp.StatusCode, nil
}

func RespondToClient(w http.ResponseWriter, body []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(body)
}
