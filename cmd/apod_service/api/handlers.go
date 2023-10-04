package api

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"text/template"

	"github.com/labstack/echo"
)

func (a *API) getAllImagesInfo(c echo.Context) error {
	imagesData, err := a.db.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("internal error: %v", err))
	}

	if imagesData == nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": errors.New("no images found")})
	}

	return c.HTML(http.StatusOK, generateHTMLTable(imagesData))
}

var tmpl *template.Template

func (a *API) getImageByDate(c echo.Context) error {
	date := c.FormValue("date")
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)

	if !re.MatchString(date) {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid data format, use YYYY-MM-YY"})
	}

	imageData, err := a.db.GetByDate(c.Request().Context(), date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Errorf("internal error: %v", err)})
	}

	if imageData == nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "no images found"})
	}

	return c.Blob(http.StatusOK, "multipart/form-data", imageData.RAW)
}
