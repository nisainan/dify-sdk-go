package dify

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
)

type DatasetDocumentsRequest struct {
	DatasetID string `json:"dataset_id"`
	Keyword   string `json:"keyword,omitempty"`
	Page      int    `json:"page,omitempty"`
	Limit     int    `json:"limit,omitempty"`
}

type DatasetDocumentsDataResponse struct {
	ID                   string `json:"id"`
	Position             int    `json:"position"`
	DataSourceType       string `json:"data_source_type"`
	DataSourceInfo       any    `json:"data_source_info"`
	DatasetProcessRuleId any    `json:"dataset_process_rule_id"`
	Name                 string `json:"name"`
	CreatedFrom          string `json:"created_from"`
	CreatedBy            string `json:"created_by"`
	CreatedAt            int    `json:"created_at"`
	Tokens               int    `json:"tokens"`
	IndexingStatus       string `json:"indexing_status"`
	Error                any    `json:"error"`
	Enabled              bool   `json:"enabled"`
	DisabledAt           any    `json:"disabled_at"`
	DisabledBy           any    `json:"disabled_by"`
	Archived             bool   `json:"archived"`
}

type DatasetDocumentsResponse struct {
	Limit   int                            `json:"limit"`
	HasMore bool                           `json:"has_more"`
	Total   int                            `json:"total"`
	Page    int                            `json:"page"`
	Data    []DatasetDocumentsDataResponse `json:"data"`
}

func (api *API) DatasetDocuments(ctx context.Context, req *DatasetDocumentsRequest) (resp *DatasetDocumentsResponse, err error) {
	httpReq, err := api.createBaseRequest(ctx, http.MethodGet, fmt.Sprintf("/v1/datasets/%s/documents", req.DatasetID), nil, Dataset)
	if err != nil {
		return
	}
	query := httpReq.URL.Query()
	query.Set("keyword", req.Keyword)
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

// ------------------------------

type DatasetDocumentCreatByFileRequest struct {
	DatasetID string         `json:"dataset_id"`
	File      multipart.File `json:"file"`
	FileName  string         `yaml:"file_name"`
	Data      struct {
		OriginalDocumentID string `json:"original_document_id"`
		IndexingTechnique  string `json:"indexing_technique"`
		ProcessRule        struct {
			Mode  string `json:"mode"`
			Rules []any  `json:"rules"`
		} `json:"process_rule"`
	} `json:"data"`
}

type DatasetDocumentCreatByFileResponse struct {
	Document struct {
		ID             string `json:"id"`
		Position       int    `json:"position"`
		DataSourceType string `json:"data_source_type"`
		DataSourceInfo struct {
			UploadFileId string `json:"upload_file_id"`
		} `json:"data_source_info"`
		DatasetProcessRuleId string `json:"dataset_process_rule_id"`
		Name                 string `json:"name"`
		CreatedFrom          string `json:"created_from"`
		CreatedBy            string `json:"created_by"`
		CreatedAt            int    `json:"created_at"`
		Tokens               int    `json:"tokens"`
		IndexingStatus       string `json:"indexing_status"`
		Error                any    `json:"error"`
		Enabled              bool   `json:"enabled"`
		DisabledAt           any    `json:"disabled_at"`
		DisabledBy           any    `json:"disabled_by"`
		Archived             bool   `json:"archived"`
		DisplayStatus        string `json:"display_status"`
		WordCount            int    `json:"word_count"`
		HitCount             int    `json:"hit_count"`
		DocForm              string `json:"doc_form"`
	} `json:"document"`
	Batch string `json:"batch"`
}

func (api *API) DatasetDocumentCreatByFile(ctx context.Context, req *DatasetDocumentCreatByFileRequest) (resp *DatasetDocumentCreatByFileResponse, err error) {
	reqData, err := json.Marshal(req.Data)
	if err != nil {
		return nil, fmt.Errorf("error json.Marshal form data: %v", err)
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", req.FileName)
	if err != nil {
		return nil, fmt.Errorf("error creating form file: %v", err)
	}
	_, err = io.Copy(part, req.File)
	if err != nil {
		return nil, fmt.Errorf("error copying file: %v", err)
	}
	err = writer.WriteField("data", string(reqData))
	if err != nil {
		return nil, fmt.Errorf("error writer.WriteField: %v", err)
	}
	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("error closing writer: %v", err)
	}
	httpReq, err := api.createBaseRequest(ctx, http.MethodGet, fmt.Sprintf("/v1/datasets/%s/documents/create_by_file", req.DatasetID), body, Dataset)
	if err != nil {
		return
	}
	httpReq.Header.Set("Content-Type", writer.FormDataContentType())
	err = api.c.sendJSONRequest(httpReq, &resp)
	return
}
