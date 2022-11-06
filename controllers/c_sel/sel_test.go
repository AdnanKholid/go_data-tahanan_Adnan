package csel

import (
	"bytes"
	"encoding/json"
	"mini_project/config"
	msel "mini_project/models/m_sel"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	config.InitTestDB()
	e := echo.New()

	return e
}

func TestGetAll_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/sel/get_all",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	},
	}

	e := InitEcho()

	req := httptest.NewRequest(http.MethodGet, "/sel/get_all", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetAll(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestCreate_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/sel/create",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	},
	}

	e := InitEcho()

	input := msel.Input{
		NoSel:   1,
		SipirID: 1,
	}

	jsonBody, _ := json.Marshal(&input)
	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPost, "/sel/create", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")

	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, Create(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestGetByID_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/sel/get_by_id/:id",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	},
	}

	e := InitEcho()

	test := config.SeedSel()
	id := strconv.Itoa(int(test.ID))

	req := httptest.NewRequest(http.MethodGet, "/sel/get_by_id/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(id)

		if assert.NoError(t, GetByID(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TesUpdate_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/sel/update/:id",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	},
	}

	e := InitEcho()

	test := config.SeedSel()

	input := msel.Input{
		NoSel:   2,
		SipirID: 2,
	}

	jsonBody, _ := json.Marshal(&input)
	bodyReader := bytes.NewReader(jsonBody)

	id := strconv.Itoa(int(test.ID))

	req := httptest.NewRequest(http.MethodPut, "/sel/update/:id", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")

	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(id)

		if assert.NoError(t, Update(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestDelete_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/sel/delete/:id",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	},
	}

	e := InitEcho()

	tes := config.SeedSel()
	id := strconv.Itoa(int(tes.ID))

	req := httptest.NewRequest(http.MethodDelete, "/sel/delete/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(id)

		if assert.NoError(t, Delete(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}

}
