package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"frontend-app/models"
)

func parseJSONBody(w http.ResponseWriter, r *http.Request, target interface{}) error {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "invalid content type", http.StatusBadRequest)
		return errors.New("invalid content type")
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

func parseQueryParam(r *http.Request, key string) (string, error) {
	value := r.URL.Query().Get(key)
	if value == "" {
		return "", fmt.Errorf("query parameter %s is missing", key)
	}
	return value, nil
}

func parseQueryIntParam(r *http.Request, key string) (int, error) {
	value, err := parseQueryParam(r, key)
	if err != nil {
		return 0, err
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid query parameter %s: %w", key, err)
	}
	return intValue, nil
}

func logError(err error) {
	log.Printf("error: %v\n", err)
}

func writeJSONResponse(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		logError(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func writeJSONError(w http.ResponseWriter, err error) {
	writeJSONResponse(w, map[string]string{"error": err.Error()})
}

func parseUserAgent(r *http.Request) string {
	return r.Header.Get("User-Agent")
}

func parseAcceptLanguage(r *http.Request) string {
	return r.Header.Get("Accept-Language")
}