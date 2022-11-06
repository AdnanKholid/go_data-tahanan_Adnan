package ctahanan

import (
	"bytes"
	"encoding/json"
	"mini_project/config"
	mtahanan "mini_project/models/m_tahanan"
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
		path:                   "/tahanan/get_all",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	},
	}

	e := InitEcho()

	req := httptest.NewRequest(http.MethodGet, "/tahanan/get_all", nil)
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
		path:                   "/tahanan/create",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	},
	}

	e := InitEcho()

	input := mtahanan.Input{
		SelID:         1,
		Nama:          "tes",
		Usia:          12,
		MasaTahanan:   "tes",
		Pelanggaran:   "tes",
		TanggalMasuk:  "tes",
		TanggalKeluar: "tes",
	}

	jsonBody, _ := json.Marshal(&input)
	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPost, "/tahanan/create", bodyReader)
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
		path:                   "/tahanan/get_by_id/:id",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	},
	}

	e := InitEcho()

	test := config.SeedTahanan()
	id := strconv.Itoa(int(test.ID))

	req := httptest.NewRequest(http.MethodGet, "/tahanan/get_by_id/:id", nil)
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
		path:                   "/tahanan/update/:id",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	},
	}

	e := InitEcho()

	test := config.SeedTahanan()

	input := mtahanan.Input{
		SelID:         1,
		Nama:          "tes",
		Usia:          12,
		MasaTahanan:   "tes",
		Pelanggaran:   "tes",
		TanggalMasuk:  "tes",
		TanggalKeluar: "tes",
	}

	jsonBody, _ := json.Marshal(&input)
	bodyReader := bytes.NewReader(jsonBody)

	id := strconv.Itoa(int(test.ID))

	req := httptest.NewRequest(http.MethodPut, "/tahanan/update/:id", bodyReader)
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
		path:                   "/tahanan/delete/:id",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	},
	}

	e := InitEcho()

	tes := config.SeedTahanan()
	id := strconv.Itoa(int(tes.ID))

	req := httptest.NewRequest(http.MethodDelete, "/tahanan/delete/:id", nil)
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
