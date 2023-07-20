package service

//==================================================================================================
// IWeatherReportService interface
//==================================================================================================

// -----------------
// methods to implement
// -----------------
type IWeatherReportService interface {
	WeatherReportByDate(p_date any) (any, error)
	ReadDatesFrom(p_path string) any
	CheckAndStoreWeatherReport(p_report any) error
}
