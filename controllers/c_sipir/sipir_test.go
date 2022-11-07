package csipir

import (
	"bytes"
	"encoding/json"
	"mini_project/config"
	msel "mini_project/models/m_sel"
	msipir "mini_project/models/m_sipir"
	ssel "mini_project/service/s_sel"
	ssipir "mini_project/service/s_sipir"
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
		expectedBodyStartsWith: "[{\"id\":",
	},
	}

	e := InitEcho()

	// create sipir data
	sipirService := ssipir.NewSipir()
	sipir := sipirService.Create(msipir.Input{
		Nama:    "test",
		Jabatan: "test",
	})

	// create sel data
	selService := ssel.NewSel()
	selService.Create(msel.Input{
		NoSel:   12,
		SipirID: sipir.ID,
	})

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
		expectedBodyStartsWith: "{\"id\":",
	},
	}

	e := InitEcho()

	// create sipir data
	sipirService := ssipir.NewSipir()
	sipir := sipirService.Create(msipir.Input{
		Nama:    "test",
		Jabatan: "test",
	})

	input := msel.Input{
		NoSel:   1,
		SipirID: sipir.ID,
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
			assert.Equal(t, http.StatusAccepted, rec.Code)
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
		expectedBodyStartsWith: "{\"id\":",
	},
	}

	e := InitEcho()

	sipirService := ssipir.NewSipir()
	sipir := sipirService.Create(msipir.Input{
		Nama:    "test",
		Jabatan: "test",
	})

	// create sel data
	selService := ssel.NewSel()
	sel := selService.Create(msel.Input{
		NoSel:   12,
		SipirID: sipir.ID,
	})

	id := strconv.Itoa(int(sel.ID))

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

	sipirService := ssipir.NewSipir()
	sipir := sipirService.Create(msipir.Input{
		Nama:    "test",
		Jabatan: "test",
	})

	// create sel data
	selService := ssel.NewSel()
	sel := selService.Create(msel.Input{
		NoSel:   12,
		SipirID: sipir.ID,
	})

	input := msel.Input{
		NoSel:   2,
		SipirID: sipir.ID,
	}

	jsonBody, _ := json.Marshal(&input)
	bodyReader := bytes.NewReader(jsonBody)

	id := strconv.Itoa(int(sel.ID))

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
		expectedBodyStartsWith: "{\"messege\":",
	},
	}

	e := InitEcho()

	sipirService := ssipir.NewSipir()
	sipir := sipirService.Create(msipir.Input{
		Nama:    "test",
		Jabatan: "test",
	})

	// create sel data
	selService := ssel.NewSel()
	sel := selService.Create(msel.Input{
		NoSel:   12,
		SipirID: sipir.ID,
	})

	id := strconv.Itoa(int(sel.ID))

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
