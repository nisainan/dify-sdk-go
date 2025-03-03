package dify

import (
	"context"
	"net/http"
	"strconv"
)

type DatasetsRequest struct {
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
}

type DatasetsDataResponse struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	Permission        string `json:"permission"`
	DataSourceType    string `json:"data_source_type"`
	IndexingTechnique string `json:"indexing_technique"`
	AppCount          int    `json:"app_count"`
	DocumentCount     int    `json:"document_count"`
	WordCount         int    `json:"word_count"`
	CreatedBy         string `json:"created_by"`
	CreatedAt         int    `json:"created_at"`
	UpdatedBy         string `json:"updated_by"`
	UpdatedAt         int    `json:"updated_at"`
}

type DatasetsResponse struct {
	Limit   int                    `json:"limit"`
	HasMore bool                   `json:"has_more"`
	Total   int                    `json:"total"`
	Page    int                    `json:"page"`
	Data    []DatasetsDataResponse `json:"data"`
}

func (api *API) Datasets(ctx context.Context, req *DatasetsRequest) (resp *DatasetsResponse, err error) {
	httpReq, err := api.createBaseRequest(ctx, http.MethodGet, "/v1/datasets", nil, Dataset)
	if err != nil {
		return
	}
	query := httpReq.URL.Query()
	if req.Page > 0 {
		query.Set("page", strconv.FormatInt(int64(req.Page), 10))
	}
	if req.Limit > 0 {
		query.Set("limit", strconv.FormatInt(int64(req.Limit), 10))
	}
	httpReq.URL.RawQuery = query.Encode()

	err = api.c.sendJSONRequest(httpReq, &resp)
	return
}
