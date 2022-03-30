package routes

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestNoRoute(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	noRoute(c)
	if w.Code != 404 {
		t.Error("Code different then 404")
	}
}
