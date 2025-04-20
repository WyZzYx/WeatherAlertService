package tests

import (
	"WeatherApp/services"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSubscribeEndpoint(t *testing.T) {

	r := services.SetupRouter()
	body := `{"email":"tests@example.com","city":"Kyiv","condition":"temperature < 20"}`
	req, _ := http.NewRequest("POST", "/subscriptions", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}
}
