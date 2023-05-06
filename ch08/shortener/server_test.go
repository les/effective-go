package shortener

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/inancgumus/effective-go/ch08/short"
	"github.com/inancgumus/effective-go/ch08/sqlx/sqlxtest"
)

func TestHandleShorten(t *testing.T) {
	t.Parallel()

	body, err := json.Marshal(map[string]any{
		"key": "go",
		"url": "https://go.dev",
	})
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, shorteningRoute, bytes.NewReader(body))

	svc := &Service{
		LinkStore: &short.LinkStore{
			DB: sqlxtest.Dial(t),
		},
	}
	handler := handleShorten(svc)
	handler.ServeHTTP(w, r)

	if w.Code != http.StatusCreated {
		t.Errorf("got status code = %d, want %d", w.Code, http.StatusCreated)
	}
	if !strings.Contains(w.Body.String(), `"go"`) {
		t.Errorf("got body = %s\twant contains %s", w.Body.String(), `"go"`)
	}
}
