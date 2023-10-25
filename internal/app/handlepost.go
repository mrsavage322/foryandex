package app

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	link := strings.TrimSpace(string(bodyBytes))
	if link == "" {
		http.Error(w, "Empty link", http.StatusBadRequest)
		return
	}

	id := GenerateRandomID(5)
	shortURL := fmt.Sprintf("%s/%s", BaseURL, id)
	URLMap.Set(id, link)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortURL))
	//sugar.Infof("Status: %s", http.StatusCreated)
	//sugar.Infof("Response sent", zap.Int("StatusCode", http.StatusCreated), zap.Int("ContentLength", len(shortURL)))

}

func GenerateRandomID(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
