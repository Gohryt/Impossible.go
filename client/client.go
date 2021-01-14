package client

import (
	"net/http"
	"time"
)

var (
	//Default is predefined http client
	Default = http.Client{Timeout: time.Second * 10}
)
