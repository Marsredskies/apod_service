package api

const (
	tableTemplateStart = `
<table border="3">
	<tr>
		<th>Date</th>
		<th>Title</th>
		<th>URL</th>
		<th>HDURL</th>
		<th>ThumbURL</th>
		<th>MediaType</th>
		<th>Copyright</th>
		<th>Explanation</th>
	</tr>
`
	imageDataTemplate = `
	<tr>
		<td>%s</td>
		<td>%s</td>
		<td><a href="%s" target="_blank">%s</a></td>
		<td><a href="%s" target="_blank">%s</a></td>
		<td><a href="%s" target="_blank">%s</a></td>
		<td>%s</td>
		<td>%s</td>
		<td>%s</td>
	</tr>
	`

	tableTemplateEnd = `
</table>
`
)
