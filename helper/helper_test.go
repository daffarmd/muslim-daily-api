package helper

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"api-go-test/model/domain"
)

type trackingReadCloser struct {
	reader io.Reader
	closed bool
}

func (t *trackingReadCloser) Read(p []byte) (int, error) {
	return t.reader.Read(p)
}

func (t *trackingReadCloser) Close() error {
	t.closed = true
	return nil
}

func TestReadFromReqBody(t *testing.T) {
	body := &trackingReadCloser{reader: strings.NewReader(`{"name":"calendar"}`)}
	req := httptest.NewRequest(http.MethodPost, "/api/data", body)

	var payload struct {
		Name string `json:"name"`
	}

	ReadFromReqBody(req, &payload)

	if payload.Name != "calendar" {
		t.Fatalf("expected decoded body name to be %q, got %q", "calendar", payload.Name)
	}

	if !body.closed {
		t.Fatal("expected request body to be closed")
	}
}

func TestWriteResponseBody(t *testing.T) {
	recorder := httptest.NewRecorder()

	WriteResponseBody(recorder, map[string]string{"status": "ok"})

	if got := recorder.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("expected content type application/json, got %q", got)
	}

	var response map[string]string
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if response["status"] != "ok" {
		t.Fatalf("expected status %q, got %q", "ok", response["status"])
	}
}

func TestToDataResponses(t *testing.T) {
	datas := []domain.Data{
		{Id: 1, Name: "Muharram", Status: 1},
		{Id: 2, Name: "Safar", Status: 1},
	}

	responses := ToDataResponses(datas)

	if len(responses) != 2 {
		t.Fatalf("expected 2 mapped responses, got %d", len(responses))
	}

	if responses[0].Name != "Muharram" || responses[1].Name != "Safar" {
		t.Fatalf("unexpected mapped responses: %#v", responses)
	}
}
