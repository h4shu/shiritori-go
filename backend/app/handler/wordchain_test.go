package handler_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/h4shu/shiritori-go/handler"
	"github.com/h4shu/shiritori-go/service"
)

func TestGetWordchain(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	svc := service.NewWordchainService(rdb)
	h := handler.NewWordchainHandler(svc)
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/wc", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	val := []string{}
	mock.ExpectLRange(service.WordchainKey, 0, service.WordchainLimit-1).SetVal(val)
	err := h.GetWordchain(c)
	if assert.NoErrorf(t, err, "unexpected error: %v", err) {
		assert.Equalf(t, rec.Code, http.StatusOK, "got %d; want %d", rec.Code, http.StatusOK)
		wantJson := "{\"wordchain\":null,\"len\":\"0\"}\n"
		assert.Equalf(t, rec.Body.String(), wantJson, "got '%s'; want '%s'", rec.Body.String(), wantJson)
	}
}

func TestAddWordchain(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	svc := service.NewWordchainService(rdb)
	h := handler.NewWordchainHandler(svc)
	e := echo.New()

	var httpErr *echo.HTTPError
	str := ""
	json := fmt.Sprintf(`{"word": "%s"}`, str)
	req := httptest.NewRequest(http.MethodPost, "/wc", strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	val := []string{}
	mock.ExpectLRange(service.WordchainKey, 0, 0).SetVal(val)
	err := h.AddWordchain(c)
	if assert.ErrorAsf(t, err, &httpErr, "got wrong error: %v", err) {
		assert.Equalf(t, err.(*echo.HTTPError).Code, http.StatusBadRequest, "got %d; want %d", err.(*echo.HTTPError).Code, http.StatusBadRequest)
	}

	mock.ClearExpect()
	str = "しりとり"
	json = fmt.Sprintf(`{"word": "%s"}`, str)
	req = httptest.NewRequest(http.MethodPost, "/wc", strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	val = []string{"しりとり"}
	mock.ExpectLRange(service.WordchainKey, 0, 0).SetVal(val)
	err = h.AddWordchain(c)
	if assert.ErrorAsf(t, err, &httpErr, "got wrong error: %v", err) {
		assert.Equalf(t, err.(*echo.HTTPError).Code, http.StatusBadRequest, "got %d; want %d", err.(*echo.HTTPError).Code, http.StatusBadRequest)
	}

	mock.ClearExpect()
	str = "りんご"
	json = fmt.Sprintf(`{"word": "%s"}`, str)
	req = httptest.NewRequest(http.MethodPost, "/wc", strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	val = []string{"しりとり"}
	mock.ExpectLRange(service.WordchainKey, 0, 0).SetVal(val)
	mock.Regexp().ExpectLPush(service.WordchainKey, str).SetErr(errors.New("FAIL"))
	err = h.AddWordchain(c)
	if assert.ErrorAsf(t, err, &httpErr, "got wrong error: %v", err) {
		assert.Equalf(t, err.(*echo.HTTPError).Code, http.StatusInternalServerError, "got %d; want %d", err.(*echo.HTTPError).Code, http.StatusInternalServerError)
	}

	mock.ClearExpect()
	str = "りんご"
	json = fmt.Sprintf(`{"word": "%s"}`, str)
	req = httptest.NewRequest(http.MethodPost, "/wc", strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	val = []string{"しりとり"}
	mock.ExpectLRange(service.WordchainKey, 0, 0).SetVal(val)
	mock.Regexp().ExpectLPush(service.WordchainKey, str).SetVal(0)
	mock.Regexp().ExpectLPush(service.WordchainKey, str).SetErr(nil)
	err = h.AddWordchain(c)
	assert.NoErrorf(t, err, "unexpected error: %v", err)
}
