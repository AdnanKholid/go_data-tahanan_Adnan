package ctahanan

import (
	"mini_project/config"
	msel "mini_project/models/m_sel"
	mtahanan "mini_project/models/m_tahanan"
	stahanan "mini_project/service/s_tahanan"
	"net/http"

	"github.com/labstack/echo/v4"
)

var selController stahanan.TahananService = stahanan.NewTahanan()

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
	input := new(mtahanan.Input)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"messege": "failed created",
		})
	}
	var sel msel.Sel
	config.DB.Preload("Tahanan").First(&sel, "id=?", input.SelID)

	if len(sel.Tahanan) >= 8 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"messege": "sel sudah penuh",
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

	input := new(mtahanan.Input)
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
	var sel msel.Sel
	config.DB.Preload("Tahanan").First(&sel, "id=?", input.SelID)

	if len(sel.Tahanan) >= 8 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"messege": "sel sudah penuh",
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
	return c.JSON(http.StatusNotFound, map[string]string{
		"messege": "success deleted",
	})
}
