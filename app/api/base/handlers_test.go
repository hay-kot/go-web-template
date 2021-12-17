package base

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hay-kot/git-web-template/internal/tests"
)

func GetTestHandler(t *testing.T) *Handlersv1 {
	return NewHandlerV1(tests.GetConsoleLogger(t))
}

func TestHandlersv1_HandleBase(t *testing.T) {
	// Setup
	hdlrFunc := GetTestHandler(t).HandleBase("v1")

	// Call Handler Func
	rr := httptest.NewRecorder()
	hdlrFunc(rr, nil)

	// Validate Status Code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code to be %d, got %d", http.StatusOK, rr.Code)
	}

	// Validate Json Payload
	expected := `{"health":true,"versions":["v1"],"title":"Go API Template","message":"Welcome to the Go API Template Application!"}`

	if rr.Body.String() != expected {
		t.Errorf("Expected json to be %s, got %s", expected, rr.Body.String())
	}

}
