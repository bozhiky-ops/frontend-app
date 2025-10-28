package helpers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func generateUUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid, nil
}

func hashString(s string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(s))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func isEmailValid(email string) bool {
	if strings.Contains(email, "@") {
		parts := strings.Split(email, "@")
		if len(parts) != 2 {
			return false
		}
		domainParts := strings.Split(parts[1], ".")
		if len(domainParts) < 2 {
			return false
		}
		return true
	}
	return false
}

func getCurrentTime() string {
	return time.Now().Format(time.RFC3339)
}

func getHTTPStatusCode(res *http.Response) int {
	return res.StatusCode
}

func handleHTTPError(res *http.Response) error {
	if res.StatusCode >= 400 {
		return errors.New(res.Status)
	}
	return nil
}

func logError(err error) {
	log.Printf("Error: %v", err)
}