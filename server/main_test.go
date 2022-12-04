package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flutter_go/app/routers"
	"gorm.io/gorm/utils/tests"
)

func TestPingRoute(t *testing.T) {
	router := routers.InitRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	tests.AssertEqual(t, 200, w.Code)
	tests.AssertEqual(t, "pong", w.Body.String())

}
