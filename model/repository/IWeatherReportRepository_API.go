package repository

// ==================================================================================================
// IWeatherReportRepository_API interface
// ==================================================================================================

// -----------------
// methods to implement
// -----------------
type IWeatherReportRepository_API interface {
	ReadBy(p_attribute string) (any, error)
}
