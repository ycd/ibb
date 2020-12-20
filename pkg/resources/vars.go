// Package resources contains data about the resources
// based on IBB's open dataset platform.
// https://data.ibb.gov.tr/en/dataset
package resources

const (
	// Limit100000 is a query parameter for getting
	// all data through the pagination.
	Limit100000 = "&limit=100000"

	// BaseURL is a base url for accessing IBB data through the API
	BaseURL = "https://data.ibb.gov.tr/en/api/3/action/datastore_search?resource_id="

	// DailyMaximumJourneyByRailSystem contains the maximum number of daily trips using rail system lines in Istanbul.
	// MARMARAY maximum line-based daily trips are included.
	// https://data.ibb.gov.tr/en/dataset/rayli-sistem-gunluk-maksimum-yolculuk-sayilari
	DailyMaximumJourneyByRailSystem = "38f170b4-5746-4672-81b1-a75516712c2a"

	// RailSystemsTimeline contains data on the number of voyages of the
	// rail system lines in Istanbul. The data are in daily, monthly and yearly formats.
	// https://data.ibb.gov.tr/en/dataset/rayli-sistemler-gunluk-aylik-yillik-hat-bazli-sefer-sayilari
	RailSystemsTimeline = "e20b32f7-c346-4803-8c0e-837f149c7fa2"

	// IsparkParkingLots contains information about ISPARK Parking Lots
	// (Park ID, Park Name, Location2 ID, Location Code, Location Name, Park Type ID, Park Type, Park Capacity,
	// Hours of Operation, Region ID, Region, Sub Area ID, Sub Area, District ID, District, Address,
	// Latitude / Longitude, Polygon Data, Latitude, Longitude, Monthly Subscription Fee, Free Parking Time,
	// Timetable, Park Continue Point).
	// https://data.ibb.gov.tr/en/dataset/ispark-otoparklarina-ait-bilgiler
	IsparkParkingLots = "c3eb0d72-1ce4-4983-a3a8-6b0b4b19fcb9"

	// IGDASSubscribers contains IGDAS (gas company) subscription numbers
	// in the 39 districts of Istanbul.
	// https://data.ibb.gov.tr/en/dataset/ilcelere-gore-abone-sayilari
	IGDASSubscribers = "b6a8daa2-074d-4536-ae68-a6fe15b965c9"

	// WasteAmounts contains information on the amount of waste by district, year and type of waste
	// (solid waste collection activities, Medical waste collection activities,
	// mechanical sweeping activities to prevent pollution in the main arteries and squares)
	// https://data.ibb.gov.tr/en/dataset/ilce-yil-ve-atik-turu-bazinda-atik-miktari
	WasteAmounts = "50036dfd-aea5-4f06-832f-f7020fdaaa5a"

	// MunicipalityPopulation2019 contains the data about the population
	// based on municipality's on Istanbul in 2019
	// https://data.ibb.gov.tr/en/dataset/belediye-nufuslari-veri-seti
	MunicipalityPopulation2019 = "c6c9b289-2824-41b3-ab3d-4fd655ed4e24"

	// IETTTicketPrices2020 Contains the price information of the
	// public transportation tickets used in Istanbul.
	IETTTicketPrices2020 = "31c7b38f-c952-43cf-98dd-b69c44079764"

	// IETTTicketPrices2019 Contains the price information of the
	// public transportation tickets used in Istanbul.
	IETTTicketPrices2019 = "8e132527-2eb7-4d68-a549-d224d233ab16"

	// DamOccupancyRates contains information on the daily and annual
	// changes of the occupancy rates of dams in Istanbul.
	// https://data.ibb.gov.tr/en/dataset/istanbul-baraj-doluluk-oranlari-verisi
	DamOccupancyRates = "b68cbdb0-9bf5-474c-91c4-9256c07c4bdf&limit"
)
