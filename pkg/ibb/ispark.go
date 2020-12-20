package ibb

import (
	"context"
	"encoding/json"

	"github.com/ycd/ibb/pkg/resources"
)

// ParkingLots information about ISPARK Parking Lots
// (Park ID, Park Name, Location ID, Location Code, Location Name, Park Type ID, Park Type, Park Capacity,
// Hours of Operation, Region ID, Region, Sub Area ID, Sub Area, District ID, District, Address,
// Latitude / Longitude, Polygon Data, Latitude, Longitude, Monthly Subscription Fee, Free Parking Time,
// Timetable, Park Continue Point).
type ParkingLots struct {
	parkingRecords `json:"result"`
}

type parkingRecords struct {
	Records []*struct {
		ParkID               int     `json:"Park ID"`
		ParkADI              string  `json:"Park Adi"`
		LokasyonID           int     `json:"Lokasyon ID"`
		LokasyonKodu         int     `json:"LokasyonKodu"`
		LokasyonAdi          string  `json:"Lokasyon Adi"`
		ParkTipiID           int     `json:"Park Tipi ID"`
		ParkTipi             string  `json:"YOL ÜSTÜ"`
		ParkKapasitesi       int     `json:"Park Kapasitesi"`
		CalismaSaatleri      string  `json:"Calisma Saatleri"`
		BolgeID              int     `json:"Bolge ID"`
		Bolge                string  `json:"Avrupa"`
		AltBolgeID           int     `json:"Alt Bolge ID"`
		AltBolge             string  `json:"Alt Bolge"`
		IlceID               int     `json:"Ilce ID"`
		Ilce                 string  `json:"Ilce"`
		Adres                string  `json:"Adres"`
		EnlemBoylam          string  `json:"Enlem/Boylam"`
		PolygonVerisi        string  `json:"Polygon Verisi"`
		Boylam               float64 `json:"Boylam"`
		Enlem                float64 `json:"Enlem"`
		AylikAbonelikUcreti  int     `json:"Aylik Abonelik Ucreti"`
		UcretsizParkSuresi   int     `json:"Ucretsiz Parklanma Suresi (dakika)"`
		Tarifesi             string  `json:"Tarifesi"`
		ParkEtDevamEtNoktasi int     `json:"Park Et Devam Et Noktasi"`
	} `json:"records"`
}

// IsparkParkingLots contains information about
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
