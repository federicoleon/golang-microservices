package polo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	polo = "polo"
)

func Marco(c *gin.Context) {
	c.String(http.StatusOK, polo)
}
