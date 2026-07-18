package web

import (
	"io/fs"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSetupRouterRendersHome(t *testing.T) {
	if _, err := fs.Stat(Dist, "dist/server/server.js"); err != nil {
		t.Skip("build web/dist before running the SSR integration test")
	}

	gin.SetMode(gin.TestMode)
	router := gin.New()
	if err := SetupRouter(router); err != nil {
		t.Fatal(err)
	}

	request := httptest.NewRequest(http.MethodGet, "http://example.test/", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("unexpected status %d: %s", response.Code, response.Body.String())
	}
	for _, expected := range []string{"gossr is ready", "xtpl SSR is ready", "window.__SSR_DATA__"} {
		if !strings.Contains(response.Body.String(), expected) {
			t.Errorf("response does not contain %q: %s", expected, response.Body.String())
		}
	}
}
