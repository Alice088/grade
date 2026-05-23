package booking

import (
	httpx "alice088/booking_sql/internal/http"

	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	Service BookingService
}

func NewHandler(service BookingService) BookingHandler {
	return BookingHandler{
		Service: service,
	}
}

func (h *BookingHandler) Book(ctx *gin.Context) {
	var req httpx.BookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ok":    false,
			"error": err.Error(),
		})
	}
}
