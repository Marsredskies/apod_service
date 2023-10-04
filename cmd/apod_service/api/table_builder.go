package api

import (
	"fmt"

	"github.com/Marsredskies/apod_service/cmd/apod_service"
)

func generateHTMLTable(data []apod.ImageData) string {
	startHtml := tableTemplateStart
	for _, d := range data {
		startHtml += fmt.Sprintf(imageDataTemplate, d.Date, d.Title, d.URL, d.URL, d.HDURL, d.HDURL, d.ThumbURL, d.ThumbURL, d.MediaType, d.Copyright, d.Explanation)
	}
	return startHtml + tableTemplateEnd

}
