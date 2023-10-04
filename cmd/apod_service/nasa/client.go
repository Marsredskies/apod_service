package nasa

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/Marsredskies/apod_service/cmd/apod_service"
	"github.com/Marsredskies/apod_service/cmd/apod_service/database"
	"github.com/Marsredskies/apod_service/envconfig"
)

type NasaClient struct {
	BaseURL   string
	ApiKey    string
	requester interface {
		Do(*http.Request) (*http.Response, error)
	}
	fetchIntervalHours int
	db                 *database.DB
}

type Response struct {
	Message string   `json:"message"`
	Urls    []string `json:"urls"`
}

func MustInitClient(config envconfig.Apod, db *database.DB) *NasaClient {
	client, err := InitClient(config, db)
	if err != nil {
		panic(fmt.Errorf("failed to init nasa api client: %v", err))
	}
	return client
}

func InitClient(config envconfig.Apod, db *database.DB) (*NasaClient, error) {
	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	return &NasaClient{
		BaseURL:            config.BaseURL,
		ApiKey:             config.ApiKey,
		requester:          client,
		fetchIntervalHours: config.IntervalHours,
		db:                 db,
	}, nil
}

func (n *NasaClient) FetchAndSaveAPOD(ctx context.Context, date time.Time) error {
	imageData, err := n.GetAPOD(ctx, date)
	if err != nil {
		return err
	}

	return n.db.Save(ctx, *imageData)
}
func (n *NasaClient) GetAPOD(ctx context.Context, date time.Time) (*apod.ImageData, error) {
	req, err := n.doRequest(date)
	if err != nil {
		return nil, err
	}

	metadata, err := n.getMetadata(ctx, req)
	if err != nil {
		return nil, err
	}

	url := n.selectURL(metadata)
	if url == "" {
		return nil, errors.New("couldn't get image url")
	}

	raw, err := n.downloadImage(ctx, url)
	if err != nil {
		return nil, err
	}

	metadata.RAW = raw

	return metadata, nil
}

func (n *NasaClient) doRequest(date time.Time) (string, error) {
	url, err := url.Parse(n.BaseURL)
	if err != nil {
		return "", err
	}

	q := url.Query()
	q.Set("api_key", n.ApiKey)
	q.Set("date", date.Format(dateFormat))
	q.Set("thumbs", "true")
	q.Set("hd", "true")

	url.RawQuery = q.Encode()

	return url.String(), nil
}

func (n *NasaClient) getMetadata(ctx context.Context, u string) (*apod.ImageData, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := n.requester.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var metadata apod.ImageData
	if err := json.Unmarshal(body, &metadata); err != nil {
		return nil, err
	}

	return &metadata, nil
}

func (n *NasaClient) downloadImage(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := n.requester.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (n *NasaClient) selectURL(img *apod.ImageData) string {
	switch {
	case img.MediaType == typeImage && img.HDURL != "":
		return img.HDURL
	case img.MediaType == typeImage && img.URL != "":
		return img.URL
	case img.MediaType == typeVideo && img.ThumbURL != "":
		return img.ThumbURL
	default:
		return ""
	}
}
