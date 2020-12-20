package ibb

import (
	"context"
	"encoding/json"

	"github.com/ycd/ibb/pkg/resources"
)

// RailSystemsTimeline contains data on the number of voyages of the
// rail system lines in Istanbul. The data are in daily, monthly and yearly formats.
// https://data.ibb.gov.tr/en/dataset/rayli-sistemler-gunluk-aylik-yillik-hat-bazli-sefer-sayilari
func (c *Client) RailSystemsTimeline(ctx context.Context) (*RailSystemsTimeline, error) {
	var rs RailSystemsTimeline

	url := resources.BaseURL + resources.RailSystemsTimeline

	resp, err := c.get(ctx, url)
	if err != nil {
		return &RailSystemsTimeline{}, err
	}

	json.Unmarshal(resp, &rs)

	return &rs, nil
}

// DailyMaximumJourneyByRailSystem contains the maximum number of daily trips using rail system lines in Istanbul.
// MARMARAY maximum line-based daily trips are included.
// https://data.ibb.gov.tr/en/dataset/rayli-sistem-gunluk-maksimum-yolculuk-sayilari
func (c *Client) DailyMaximumJourneyByRailSystem(ctx context.Context) (*RailSystemsDailyMaximumJourneys, error) {
	var rs RailSystemsDailyMaximumJourneys

	url := resources.BaseURL + resources.DailyMaximumJourneyByRailSystem

	resp, err := c.get(ctx, url)
	if err != nil {
		return &RailSystemsDailyMaximumJourneys{}, err
	}

	json.Unmarshal(resp, &rs)

	return &rs, nil
}
