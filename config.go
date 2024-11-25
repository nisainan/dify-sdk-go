package dify

import (
	"net/http"
	"time"
)

type ClientConfig struct {
	Host             string
	ChatAPISecret    string
	DatasetAPISecret string
	Timeout          time.Duration
	Transport        *http.Transport
}
