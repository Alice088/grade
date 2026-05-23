package handlers

import (
	"alice088/umbrella/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Time struct {
	Service service.Time
}

func NewTime(service service.Time) Time {
	return Time{
		Service: service,
	}
}

func (t *Time) Since() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, t.Service.Since())
	}
}
