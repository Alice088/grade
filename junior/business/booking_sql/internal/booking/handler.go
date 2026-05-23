package booking

import (
	httpx "alice088/booking_sql/internal/http"
	"errors"

	errorsx "alice088/booking_sql/pkg/errors"

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
		ctx.JSON(http.StatusBadRequest, httpx.Err(err.Error()))
		return
	}

	bookID, err := h.Service.Repo.Book(ctx.Request.Context(), Booking{
		UserName: req.UserName,
		RoomName: req.RoomName,
		StartAt:  req.StartAt,
		EndAt:    req.EndAt,
	})

	if err != nil {
		if fErr, ok := errors.AsType[errorsx.HTTPFriendly](err); !ok {
			ctx.JSON(http.StatusInternalServerError, httpx.Err("failed to book"))
		} else {
			ctx.JSON(fErr.Status(), httpx.Err(fErr.FMsg()))
		}
		return
	}

	ctx.JSON(http.StatusOK, httpx.OK(
		httpx.Book{
			ID: bookID,
		},
	))
}
