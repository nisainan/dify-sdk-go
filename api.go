package dify

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

const (
	Chat    = "Chat"
	Dataset = "Dataset"
)

type API struct {
	c          *Client
	chatSecret string
	dataSecret string
}

func (api *API) WithChatSecret(secret string) *API {
	api.chatSecret = secret
	return api
}

func (api *API) getChatSecret() string {
	if api.chatSecret != "" {
		return api.chatSecret
	}
	return api.c.getChatAPISecret()
}

func (api *API) WithDatasetSecret(secret string) *API {
	api.dataSecret = secret
	return api
}

func (api *API) getDatasetSecret() string {
	if api.dataSecret != "" {
		return api.dataSecret
	}
	return api.c.getDatasetAPISecret()
}

func (api *API) createBaseRequest(ctx context.Context, method, apiUrl string, body interface{}, apiType string) (*http.Request, error) {
	var b io.Reader
	if body != nil {
		reqBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		b = bytes.NewBuffer(reqBytes)
	} else {
		b = http.NoBody
	}
	req, err := http.NewRequestWithContext(ctx, method, api.c.getHost()+apiUrl, b)
	if err != nil {
		return nil, err
	}
	var token string
	switch apiType {
	case Chat:
		token = api.getChatSecret()
	case Dataset:
		token = api.getDatasetSecret()
	default:
		token = api.getChatSecret()
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	return req, nil
}
