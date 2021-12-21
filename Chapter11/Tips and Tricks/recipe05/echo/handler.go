package echo

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GET /
func rootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Gopher!")
}

// GET /item/:id
func getItemHandler(c echo.Context) error {
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	switch itemId {
	case 1:
		return c.JSON(http.StatusOK, &item{ID: 1, Name: "first item"})
	case 2, 3:
		return c.JSON(http.StatusOK, &item{ID: itemId, Name: "other item"})
	default:
		return c.JSON(http.StatusNotFound, "")
	}
}

// POST /item/:id
func postItemHandler(c echo.Context) error {
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	if itemId < 0 || itemId > 3 {
		return c.JSON(http.StatusInternalServerError, "")
	}
	var item item
	err = c.Bind(&item)
	if err != nil {
		return err
	}
	if itemId != item.ID {
		return c.JSON(http.StatusBadRequest, "")
	}
	return c.JSON(http.StatusOK, item)
}
