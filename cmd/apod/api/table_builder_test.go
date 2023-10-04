package api

import (
	"testing"

	"github.com/Marsredskies/apod_service/cmd/apod"
	"github.com/stretchr/testify/require"
)

func TestTableBuilder(t *testing.T) {
	images := []apod.ImageData{
		{Date: "2023-10-04", Title: "Example Image 1", URL: "http://example.com/image", HDURL: "", ThumbURL: "", MediaType: "image/jpeg", Copyright: "Copyright Example", Explanation: "test"},
		{Date: "2023-10-03", Title: "Example Image 2", URL: "http://example.com/image", HDURL: "", ThumbURL: "", MediaType: "image/jpeg", Copyright: "Copyright Example", Explanation: "test"},
		{Date: "2023-10-02", Title: "Example Image 3", URL: "http://example.com/image", HDURL: "", ThumbURL: "", MediaType: "image/jpeg", Copyright: "Copyright Example", Explanation: "test"},
	}

	html := generateHTMLTable(images)
	require.NotEmpty(t, html)
}
