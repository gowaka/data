package data

type SummaryData struct {
	GrandTotal struct {
		Digital      string `json:"digital"`
		Hours        int    `json:"hours"`
		Minutes      int    `json:"minutes"`
		Text         string `json:"text"`
		TotalSeconds int    `json:"total_seconds"`
	} `json:"grand_total"`

	Projects []struct {
		Name         string `json:"name"`
		TotalSeconds int    `json:"total_seconds"`
		Digital      string `json:"digital"`
		Text         string `json:"text"`
		Hours        int    `json:"hours"`
		Minutes      int    `json:"minutes"`
		Seconds      int    `json:"seconds"`
	} `json:"projects"`

	Languages []struct {
		Name    string  `json:"name"`
		Percent float64 `json:"percent"`
	} `json:"languages"`

	Range struct {
		Start      int    `json:"start"`
		StartDate  string `json:"start_date"`
		StartHuman string `json:"start_human"`
		End        int    `json:"end"`
		EndDate    string `json:"end_date"`
		EndHuman   string `json:"end_human"`
		Text       string `json:"text"`
		Timezone   string `json:"timezone"`
	} `json:"range"`
}

// Summary maps to the data returned by resource `/api/v1/summary/daily`
type Summary struct {
	Data  []SummaryData `json:"data"`
	Start int           `json:"start"`
	End   int           `json:"end"`
}
