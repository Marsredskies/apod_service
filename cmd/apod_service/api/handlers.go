package api

import (
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/labstack/echo"
)

func (a *API) getAllImagesInfo(c echo.Context) error {
	imagesData, err := a.db.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			fmt.Sprintf("internal error: %s", err.Error()))
	}

	if imagesData == nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "no images found"})
	}

	return c.HTML(http.StatusOK, generateHTMLTable(imagesData))
}

func (a *API) getImageByDate(c echo.Context) error {
	date := c.FormValue("date")
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)

	if !re.MatchString(date) {
		return c.JSON(http.StatusBadRequest,
			echo.Map{"error": "invalid data format, use YYYY-MM-YY"})
	}

	imageData, err := a.db.GetByDate(c.Request().Context(), date)
	if err != nil {
		return jsonIntenalError(c, err)
	}

	if imageData == nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "no images found"})
	}

	fileName := fmt.Sprintf("%s.%s", imageData.Title, imageData.Extension)

	tmpfile, err := os.CreateTemp("", fileName)
	if err != nil {
		return jsonIntenalError(c, err)
	}
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.Write(imageData.RAW)
	if err != nil {
		return jsonIntenalError(c, err)
	}

	c.Response().Header().Set(echo.HeaderContentType, http.DetectContentType(imageData.RAW))
	c.Response().Header().Set(echo.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%s", fileName))

	return c.File(tmpfile.Name())
}

func jsonIntenalError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError,
		fmt.Sprintf("internal error: %s", err.Error()))
}
