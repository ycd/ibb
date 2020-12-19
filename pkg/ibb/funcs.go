package ibb

import (
	"context"
	"encoding/json"

	"github.com/ycd/ibb/pkg/resources"
)

func (c *Client) RealSystemsTimeline(ctx context.Context) (*RailSystems, error) {
	var rs RailSystems

	resp, err := c.get(ctx, resources.BaseURL+resources.RailSystemsTimeline)
	if err != nil {
		return &RailSystems{}, err
	}

	json.Unmarshal(resp, &rs)

	return &rs, nil
}
