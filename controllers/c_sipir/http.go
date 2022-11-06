package csipir

import (
	msipir "mini_project/models/m_sipir"
	ssipir "mini_project/service/s_sipir"
	"net/http"

	"github.com/labstack/echo/v4"
)

var sipirController ssipir.SipirService = ssipir.NewSipir()

func GetAll(c echo.Context) error {
	result := sipirController.GetAll()

	return c.JSON(http.StatusOK, result)
}

func GetByID(c echo.Context) error {
	id := c.Param("id")
	result := sipirController.GetByID(id)
	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"messege": "data not found",
		})
	}

	return c.JSON(http.StatusOK, result)

}

func Create(c echo.Context) error {
	input := new(msipir.Input)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"messege": "failed created",
		})
	}
	err := input.Validator()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Validation Failed",
		})
	}
	result := sipirController.Create(*input)

	return c.JSON(http.StatusAccepted, result)
}

func Update(c echo.Context) error {
	id := c.Param("id")
	validateId := sipirController.GetByID(id)

	if validateId.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"messege": "data not found",
		})
	}

	input := new(msipir.Input)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"messege": "failed created",
		})
	}
	err := input.Validator()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Validation Failed",
		})
	}
	result := sipirController.Update(id, *input)

	return c.JSON(http.StatusAccepted, result)
}

func Delete(c echo.Context) error {
	id := c.Param("id")
	validateId := sipirController.GetByID(id)

	if validateId.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"messege": "data not found",
		})
	}

	result := sipirController.Delete(id)

	if !result {
		return c.JSON(http.StatusNotFound, map[string]string{
			"messege": "failed deleted",
		})
	}
	return c.JSON(http.StatusNotFound, map[string]string{
		"messege": "success deleted",
	})
}
