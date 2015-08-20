package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func OK(c *gin.Context) {
	result := Result{true, "ok", nil}
	c.JSON(http.StatusOK, result)
}

func ERROR(msg string) Result {
	return Result{false, msg, nil}
}

func renderOk(c *gin.Context, data interface{}) {
	result := Result{true, "ok", data}
	c.JSON(http.StatusOK, result)
}

func renderErrorMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, ERROR(msg))
}

func renderError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, ERROR(err.Error()))
}
