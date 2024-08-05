package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/h4shu/shiritori-go/adapters/controllers"
	"github.com/h4shu/shiritori-go/domain/entities"
	"github.com/h4shu/shiritori-go/infrastructure/web/models"
)

type WordchainHandler struct {
	c *controllers.WordchainController
}

func NewWordchainHandler(c *controllers.WordchainController) *WordchainHandler {
	return &WordchainHandler{
		c: c,
	}
}

func (h *WordchainHandler) GetLast(c echo.Context) error {
	m, err := h.c.GetLast(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	res := &models.WordchainGetLastResponse{
		Word: m.GetWord(),
	}
	return c.JSON(http.StatusOK, res)
}

func (h *WordchainHandler) List(c echo.Context) error {
	var req models.WordchainListRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	m, err := h.c.List(c.Request().Context(), req.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	res := &models.WordchainListResponse{
		Wordchain: m.GetWordchain(),
		Len:       m.GetLen(),
	}
	return c.JSON(http.StatusOK, res)
}

func (h *WordchainHandler) Append(c echo.Context) error {
	var req models.WordchainAppendRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err = h.c.Append(c.Request().Context(), req.Word)
	if err != nil {
		switch err.(type) {
		case *entities.ErrWordInvalid:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
