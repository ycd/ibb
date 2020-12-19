package ibb

import (
	"context"
	"encoding/json"

	"github.com/ycd/ibb/pkg/resources"
)

// IsparkParkingLots information about ISPARK Parking Lots
// (Park ID, Park Name, Location2 ID, Location Code, Location Name, Park Type ID, Park Type, Park Capacity,
// Hours of Operation, Region ID, Region, Sub Area ID, Sub Area, District ID, District, Address,
// Latitude / Longitude, Polygon Data, Latitude, Longitude, Monthly Subscription Fee, Free Parking Time,
// Timetable, Park Continue Point).
// https://data.ibb.gov.tr/en/dataset/ispark-o
func (c *Client) IsparkParkingLots(ctx context.Context) (*ParkingLots, error) {
	var pl ParkingLots

	url := resources.BaseURL + resources.IsparkParkingLots

	resp, err := c.get(ctx, url)
	if err != nil {
		return &ParkingLots{}, err
	}

	json.Unmarshal(resp, &pl)

	return &pl, nil
}
