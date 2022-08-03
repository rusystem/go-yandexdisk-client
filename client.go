package yandexdisk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rusystem/go-yandexdisk-client/responses"
	"net/http"
	"os"
	"time"
)

const (
	BaseURLV1 = "https://cloud-api.yandex.net/v1/disk"
)

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

func NewClient(apiKey string, timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can`t be zero")
	}

	return &Client{
		BaseURL: BaseURLV1,
		apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c *Client) sendRequest(req *http.Request, data interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("OAuth %s", c.apiKey))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		var errResp responses.ErrorResponse
		if err = json.NewDecoder(resp.Body).Decode(&errResp); err == nil {
			return errors.New(errResp.Info())
		}

		return fmt.Errorf("unknown error, status code %d\n", resp.StatusCode)
	}

	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil
	}

	return nil
}

func (c *Client) GetDiskInfo(ctx context.Context) (*responses.Disk, error) {
	req, err := http.NewRequest("GET", c.BaseURL, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	d := responses.Disk{}
	if err = c.sendRequest(req, &d); err != nil {
		return nil, err
	}

	return &d, nil
}

func (c *Client) GetFiles(ctx context.Context) ([]responses.Item, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/resources/files", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	f := responses.FilesResourceList{}
	if err = c.sendRequest(req, &f); err != nil {
		return nil, err
	}

	return f.Items, nil
}

func (c *Client) UploadFile(ctx context.Context, fileName string, url string) error {
	if fileName == "" && url == "" {
		return errors.New("fileName or url can`t be null")
	}

	path := fmt.Sprintf("disk:/%s", fileName)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/resources/upload?path=%s&url=%s", c.BaseURL, path, url), nil)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)

	var r responses.SuccessResponse
	if err = c.sendRequest(req, &r); err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteResource(ctx context.Context, resourceName string) error {
	if resourceName == "" {
		return errors.New("resource name can`t be null")
	}

	path := fmt.Sprintf("disk:/%s", resourceName)

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/resources?path=%s", c.BaseURL, path), nil)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)

	var r responses.SuccessResponse
	if err = c.sendRequest(req, &r); err != nil {
		return err
	}

	return nil
}
