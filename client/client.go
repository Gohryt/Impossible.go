package client

import (
	"net/http"
	"time"
)

var (
	Client = http.Client{Timeout: time.Second * 10}
)
