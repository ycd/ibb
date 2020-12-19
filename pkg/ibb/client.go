package ibb

import (
	"context"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"

	"github.com/pkg/errors"
)

var (
	logger, _ = zap.NewProduction()
	sugar     = logger.Sugar()
)

const (
	headerKeyAccept = "Accept"
	contentTypeJSON = "application/json"
)

// Client is an API client object for accessing the IBB Open Dataset
type Client struct {
	c     *http.Client
	debug bool
}

// NewClient creates a new Client for accessing
// IBB's open dataset through the API.
func NewClient() *Client {
	return &Client{
		c:     http.DefaultClient,
		debug: false,
	}
}

func (c *Client) get(ctx context.Context, resource string) ([]byte, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, resource, nil)
	if err != nil {
		return nil, errors.Wrap(err, errCreateRequest)
	}

	r.Header.Add(headerKeyAccept, contentTypeJSON)

	if c.debug {
		sugar.Infof("[DEBUG] sending a request to: %s", resource)
	}

	resp, err := c.c.Do(r)
	if err != nil {
		return nil, errors.Wrap(err, errDoRequest)
	}
	defer resp.Body.Close()

	if c.debug {
		sugar.Infof("[DEBUG] Response: %s", resp)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(ErrRequestFailed, errReqDataF, resp.StatusCode, resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}
