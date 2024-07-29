package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/h4shu/shiritori-go/entity"
	"github.com/h4shu/shiritori-go/service"
)

type WordchainHandler struct {
	svc *service.WordchainService
}

func NewWordchainHandler(svc *service.WordchainService) *WordchainHandler {
	return &WordchainHandler{
		svc: svc,
	}
}

func (h *WordchainHandler) GetWordchain(c echo.Context) error {
	wc, err := h.svc.GetWordchain(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := &entity.GetWordchainResponse{
		Wordchain: wc.ToStrSlice(),
		Len:       wc.Len(),
	}
	return c.JSON(http.StatusOK, res)
}

func (h *WordchainHandler) AddWordchain(c echo.Context) error {
	var req entity.AddWordchainRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	w := entity.NewWord(req.Word)
	err = h.svc.TryAddWord(c.Request().Context(), &w)
	if err != nil {
		switch err.(type) {
		case *entity.ErrWordInvalid:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
