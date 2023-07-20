package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"workspace/model/domain"
)

// ==================================================================================================
// WeatherReportRepository_meteoAPI struct
// ==================================================================================================

// -----------------
// properties, fields
// -----------------
type WeatherReportRepository_meteoAPI struct {
}

// -----------------
// constructors
// -----------------
func NewWeatherReportRepository_meteoAPI() (*WeatherReportRepository_meteoAPI, error) {
	return &WeatherReportRepository_meteoAPI{}, nil
}

// -----------------
// methods
// -----------------
func (r *WeatherReportRepository_meteoAPI) ReadBy(p_date string) (any, error) {
	/* Generates the URL for the OpenMeteo API request */
	getRequest := generateAPIURL(
		47.50,
		19.04,
		[]string{"temperature_2m_max", "temperature_2m_min", "rain_sum", "windspeed_10m_max", "winddirection_10m_dominant"},
		"Europe/Berlin",
		p_date,
		p_date,
	)

	/* Makes a GET request to the OpenMeteo API and returns the response as a WeatherInfo struct */
	res, err := http.Get(getRequest)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var weatherReport domain.WeatherReport
	err = json.Unmarshal(body, &weatherReport)
	if err != nil {
		return nil, err
	}

	return weatherReport, nil
}

func generateAPIURL(latitude float64, longitude float64, daily []string, timezone string, startDate string, endDate string) string {
	baseURL := "https://api.open-meteo.com/v1/forecast?"
	dailyParams := strings.Join(daily, ",")
	return fmt.Sprintf("%slatitude=%f&longitude=%f&daily=%s&timezone=%s&start_date=%s&end_date=%s", baseURL, latitude, longitude, dailyParams, timezone, startDate, endDate)
}
