package test_utils

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"net/http"
)

func GetMockedContext(request *http.Request, response *httptest.ResponseRecorder) *gin.Context {
	c, _ := gin.CreateTestContext(response)
	c.Request = request
	return c
}
