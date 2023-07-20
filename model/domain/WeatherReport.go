package domain

// ==================================================================================================
// WeatherReport struct
// ==================================================================================================

// -----------------
// properties,fields
// -----------------
/* time, space data */
type WeatherReport struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`

	DailyUnits DailyUnits `json:"daily_units"`
	Daily      Daily      `json:"daily"`
}

/* formatting purpose */
type DailyUnits struct {
	Time                     string `json:"time"`
	Temperature2MMax         string `json:"temperature_2m_max"`
	Temperature2MMin         string `json:"temperature_2m_min"`
	RainSum                  string `json:"rain_sum"`
	Windspeed10MMax          string `json:"windspeed_10m_max"`
	Winddirection10MDominant string `json:"winddirection_10m_dominant"`
}

/* relevant data */
type Daily struct {
	Time                     []string  `json:"time"`
	Temperature2MMax         []float64 `json:"temperature_2m_max"`
	Temperature2MMin         []float64 `json:"temperature_2m_min"`
	RainSum                  []float64 `json:"rain_sum"`
	Windspeed10MMax          []float64 `json:"windspeed_10m_max"`
	Winddirection10MDominant []int     `json:"winddirection_10m_dominant"`
}
