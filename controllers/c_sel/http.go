package csel

import (
	"mini_project/config"
	msel "mini_project/models/m_sel"
	msipir "mini_project/models/m_sipir"
	ssel "mini_project/service/s_sel"
	"net/http"

	"github.com/labstack/echo/v4"
)

var selController ssel.SelService = ssel.NewSel()

func GetAll(c echo.Context) error {
	result := selController.GetAll()

	return c.JSON(http.StatusOK, result)
}

func GetByID(c echo.Context) error {
	id := c.Param("id")
	result := selController.GetByID(id)
	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"messege": "data not found",
		})
	}
	return c.JSON(http.StatusOK, result)

}

func Create(c echo.Context) error {
	input := new(msel.Input)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"messege": "failed created",
		})
	}
	var sipir msipir.Sipir
	config.DB.Preload("Sel").First(&sipir, "id=?", input.SipirID)

	if len(sipir.Sel) >= 3 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"messege": "sipir penuh",
		})
	}
	err := input.Validator()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Validation Failed",
		})
	}
	result := selController.Create(*input)

	return c.JSON(http.StatusAccepted, result)
}

func Update(c echo.Context) error {
	id := c.Param("id")
	validateId := selController.GetByID(id)

	if validateId.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"messege": "data not found",
		})
	}

	input := new(msel.Input)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"messege": "failed created",
		})
	}

	var sipir msipir.Sipir
	config.DB.Preload("Sel").First(&sipir, "id=?", input.SipirID)

	if len(sipir.Sel) >= 3 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"messege": "sipir penuh",
		})
	}
	err := input.Validator()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Validation Failed",
		})
	}
	result := selController.Update(id, *input)

	return c.JSON(http.StatusAccepted, result)
}

func Delete(c echo.Context) error {
	id := c.Param("id")
	validateId := selController.GetByID(id)

	if validateId.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"messege": "data not found",
		})
	}

	result := selController.Delete(id)

	if !result {
		return c.JSON(http.StatusNotFound, map[string]string{
			"messege": "failed deleted",
		})
	}

	// use http.StatusOK
	return c.JSON(http.StatusOK, map[string]string{
		"messege": "success deleted",
	})
}
