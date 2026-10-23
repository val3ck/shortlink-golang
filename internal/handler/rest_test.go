package rest

import (
	encodeS "golang-api/internal/service"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Hardcoded for now
func TestCreateLink(t *testing.T) {
	_ = os.Chdir("../../")
	req := httptest.NewRequest("POST", "/li", nil)
	recording := httptest.NewRecorder()
	createLink(recording, req)
	if recording.Code != http.StatusUnauthorized {
		t.Errorf("Expected status %d, got %d", http.StatusUnauthorized, recording.Code)
	}

}

func TestEncodedURL(t *testing.T) {
	encoded := encodeS.EncodingURL("https://google.com", 12)
	t.Log(encoded)
	t.Error()
}
