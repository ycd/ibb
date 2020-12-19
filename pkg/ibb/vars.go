package ibb

// RailSystems is a holds the data about rail systems in Istanbul
// The routes the year and the period etc.
type RailSystems struct {
	railSystemsRecords `json:"result"`
}

type railSystemsRecords struct {
	Records []*railSystemTimelineRecords `json:"records"`
}

type railSystemTimelineRecords struct {
	Yil     int    `json:"Yil"`
	Periyot string `json:"Periyot"`
	Hatlar
}

// Hatlar contains information about the railway routes
type Hatlar struct {
	M1Hatti  int `json:"M1 Hatti"`
	M2Hatti  int `json:"M2 Hatti"`
	M3Hatti  int `json:"M3 Hatti"`
	M4Hatti  int `json:"M4 Hatti"`
	M5Hatti  int `json:"M5 Hatti"`
	M6Hatti  int `json:"M6 Hatti"`
	T1Hatti  int `json:"T1 Hatti"`
	T3Hatti  int `json:"T3 Hatti"`
	T4Hatti  int `json:"T4 Hatti"`
	F1Hatti  int `json:"F1 Hatti"`
	TF1Hatti int `json:"TF1 Hatti"`
	TF2Hatti int `json:"TF2 Hatti"`
}

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

// MunicipalityPopulation2019 contains the data about the population
// based on municipality's on Istanbul in 2019
// https://data.ibb.gov.tr/en/dataset/belediye-nufuslari-veri-seti
type MunicipalityPopulation2019 struct {
	municipalityPopulation2019Records `json:"result"`
}

type municipalityPopulation2019Records struct {
	Records []*struct {
		Belediye string `json:"Belediyeler"`
		Nufus    int    `json:"2019 yili nufuslari"`
	} `json:"records"`
}

// WasteAmount ..
type WasteAmount struct {
	wasteAmountRecords `json:"result"`
}

type wasteAmountRecords struct {
	Records []*struct {
		Ilce     string `json:"Ilce"`
		VeriTuru string `json:"Veri Turu"`
		years
	} `json:"records"`
}

type years struct {
	Y2004 int `json:"2004,omitempty"`
	Y2005 int `json:"2005,omitempty"`
	Y2006 int `json:"2006,omitempty"`
	Y2007 int `json:"2007,omitempty"`
	Y2008 int `json:"2008,omitempty"`
	Y2009 int `json:"2009,omitempty"`
	Y2010 int `json:"2010,omitempty"`
	Y2011 int `json:"2011,omitempty"`
	Y2012 int `json:"2012,omitempty"`
	Y2013 int `json:"2013,omitempty"`
	Y2014 int `json:"2014,omitempty"`
	Y2015 int `json:"2015,omitempty"`
	Y2016 int `json:"2016,omitempty"`
	Y2017 int `json:"2017,omitempty"`
	Y2018 int `json:"2018,omitempty"`
	Y2019 int `json:"2019,omitempty"`
	Y2020 int `json:"2020,omitempty"`
	Y2021 int `json:"2021,omitempty"`
}
