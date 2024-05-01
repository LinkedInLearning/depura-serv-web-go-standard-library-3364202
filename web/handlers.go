package web

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/linkedinlearning/depura-go/webservice/user"
)

type UserHandler struct {
	DB *sql.DB
}

func (h *UserHandler) Info(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("uid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse userID"})
		return
	}

	ui, err := user.GetInfo(h.DB, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"User": ui,
	})
}
