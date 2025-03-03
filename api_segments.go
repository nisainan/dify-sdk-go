package dify

import (
	"context"
	"fmt"
	"net/http"
)

type SegmentsRequest struct {
	DatasetID  string `json:"dataset_id"`
	DocumentID string `json:"document_id"`
	Keyword    string `json:"keyword,omitempty"`
	Status     string `json:"status,omitempty"`
}

type SegmentsDataResponse struct {
	Id            string   `json:"id"`
	Position      int      `json:"position"`
	DocumentId    string   `json:"document_id"`
	Content       string   `json:"content"`
	Answer        string   `json:"answer"`
	WordCount     int      `json:"word_count"`
	Tokens        int      `json:"tokens"`
	Keywords      []string `json:"keywords"`
	IndexNodeId   string   `json:"index_node_id"`
	IndexNodeHash string   `json:"index_node_hash"`
	HitCount      int      `json:"hit_count"`
	Enabled       bool     `json:"enabled"`
	DisabledAt    any      `json:"disabled_at"`
	DisabledBy    any      `json:"disabled_by"`
	Status        string   `json:"status"`
	CreatedBy     string   `json:"created_by"`
	CreatedAt     int      `json:"created_at"`
	IndexingAt    int      `json:"indexing_at"`
	CompletedAt   int      `json:"completed_at"`
	Error         any      `json:"error"`
	StoppedAt     any      `json:"stopped_at"`
}

type SegmentsResponse struct {
	Data    []SegmentsDataResponse `json:"data"`
	DocForm string                 `json:"doc_form"`
}

func (api *API) Segments(ctx context.Context, req *SegmentsRequest) (resp *SegmentsResponse, err error) {
	httpReq, err := api.createBaseRequest(ctx, http.MethodGet, fmt.Sprintf("/v1/datasets/%s/documents/%s/segments", req.DatasetID, req.DocumentID), nil, Dataset)
	if err != nil {
		return
	}
	query := httpReq.URL.Query()
	if len(req.Keyword) > 0 {
		query.Set("keyword", req.Keyword)
	}
	if len(req.Status) > 0 {
		query.Set("status", req.Status)
	}
	httpReq.URL.RawQuery = query.Encode()
	err = api.c.sendJSONRequest(httpReq, &resp)
	return
}

// ------------------------------

type SegmentDeleteRequest struct {
	DatasetID  string `json:"dataset_id"`
	DocumentID string `json:"document_id"`
	SegmentID  string `json:"segment_id"`
}

type SegmentDeleteResponse struct {
	Result string `json:"result"`
}

func (api *API) SegmentDelete(ctx context.Context, req *SegmentDeleteRequest) (resp *SegmentDeleteResponse, err error) {
	httpReq, err := api.createBaseRequest(ctx, http.MethodDelete, fmt.Sprintf("/v1/datasets/%s/documents/%s/segments/%s", req.DatasetID, req.DocumentID, req.SegmentID), nil, Dataset)
	if err != nil {
		return
	}
	err = api.c.sendJSONRequest(httpReq, &resp)
	return
}

// ------------------------------

type Segment struct {
	Content  string   `json:"content"`
	Answer   string   `json:"answer"`
	Keywords []string `json:"keywords,omitempty"`
}

type SegmentCreateRequest struct {
	DatasetID  string    `json:"dataset_id"`
	DocumentID string    `json:"document_id"`
	Segments   []Segment `json:"segments"`
}

func (api *API) SegmentCreate(ctx context.Context, req *SegmentCreateRequest) (resp *SegmentsResponse, err error) {
	httpReq, err := api.createBaseRequest(ctx, http.MethodPost, fmt.Sprintf("/v1/datasets/%s/documents/%s/segments", req.DatasetID, req.DocumentID), req, Dataset)
	if err != nil {
		return
	}
	err = api.c.sendJSONRequest(httpReq, &resp)
	return
}

// ------------------------------

type SegmentUpdateRequest struct {
	DatasetID  string  `json:"dataset_id"`
	DocumentID string  `json:"document_id"`
	SegmentID  string  `json:"segment_id"`
	Segment    Segment `json:"segment"`
}

type SegmentUpdateResponse struct {
	Data    SegmentsDataResponse `json:"data"`
	DocForm string               `json:"doc_form"`
}

func (api *API) SegmentUpdate(ctx context.Context, req *SegmentUpdateRequest) (resp *SegmentUpdateResponse, err error) {
	httpReq, err := api.createBaseRequest(ctx, http.MethodPost, fmt.Sprintf("/v1/datasets/%s/documents/%s/segments/%s", req.DatasetID, req.DocumentID, req.SegmentID), req, Dataset)
	if err != nil {
		return
	}
	err = api.c.sendJSONRequest(httpReq, &resp)
	return
}
