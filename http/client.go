package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type client struct {
	c *http.Client
}

func NewClient() *client {
	return &client{
		c: &http.Client{},
	}
}

func (c *client) PostJSON(ctx context.Context, url string, payload interface{}, target interface{}) error {
	contentType := "application/json"
	encodedPayload, err := json.Marshal(payload)
	if err != nil {
		return errors.Wrapf(err, "http.client.PostJSON: encode payload")
	}

	buffer := bytes.NewBuffer(encodedPayload)
	if err := c.post(ctx, url, contentType, buffer, target); err != nil {
		return errors.Wrapf(err, "http.client.PostJSON: send request")
	}

	return nil
}

func (c *client) PostBytes(ctx context.Context, url string, payload []byte, target interface{}) error {
	contentType := "application/octet-stream"
	buffer := bytes.NewBuffer(payload)

	if err := c.post(ctx, url, contentType, buffer, target); err != nil {
		return errors.Wrapf(err, "http.client.PostBytes")
	}

	return nil
}

func (c *client) post(ctx context.Context, url string, contentType string, payload io.Reader, target interface{}) error {
	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", contentType)

	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var bodyText string
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			bodyText = "failed to read response body"
		} else {
			bodyText = string(body)
		}

		return errors.Errorf("http.client.Post: %d - %s", resp.StatusCode, bodyText)
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return errors.Wrapf(err, "decode response")
	}

	return nil
}
