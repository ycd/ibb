package ibb

import (
	"context"
	"encoding/json"

	"github.com/ycd/ibb/pkg/resources"
)

// RailSystemsTimeline contains data on the number of voyages of the
// rail system lines in Istanbul. The data are in daily, monthly and yearly formats.
// https://data.ibb.gov.tr/en/dataset/rayli-sistemler-gunluk-aylik-yillik-hat-bazli-sefer-sayilari
func (c *Client) RailSystemsTimeline(ctx context.Context) (*RailSystems, error) {
	var rs RailSystems

	url := resources.BaseURL + resources.RailSystemsTimeline

	resp, err := c.get(ctx, url)
	if err != nil {
		return &RailSystems{}, err
	}

	json.Unmarshal(resp, &rs)

	return &rs, nil
}
