package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/h4shu/shiritori-go/model"
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

	res := &model.GetWordchainResponse{
		Wordchain: wc.ToStrSlice(),
		Len:       wc.Len(),
	}
	return c.JSON(http.StatusOK, res)
}

func (h *WordchainHandler) AddWordchain(c echo.Context) error {
	var req model.AddWordchainRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	w := model.NewWord(req.Word)
	err = h.svc.TryAddWord(c.Request().Context(), &w)
	if err != nil {
		switch err.(type) {
		case *model.ErrWordInvalid:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
