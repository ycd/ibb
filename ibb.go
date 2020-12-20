package ibb

import (
	"github.com/ycd/ibb/pkg/ibb"
)

// NewClient creates a new Client for accessing
// IBB's open dataset through the API.
func NewClient() *ibb.Client {
	return ibb.NewClient()
}
