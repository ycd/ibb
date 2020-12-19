package ibb

// RailSystems is a holds the data about rail systems in Istanbul
// The routes the year and the period etc.
type RailSystems struct {
	railSystemsResult `json:"result"`
}

type railSystemsResult struct {
	Records []*RailSystemTimelineRecords `json:"records"`
}

type RailSystemTimelineRecords struct {
	Year   int    `json:"Yil"`
	Period string `json:"Periyot"`
	RailSystemRoutes
}

type RailSystemRoutes struct {
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
