package repository

// ==================================================================================================
// IWeatherReportRepository_DataStore interface
// ==================================================================================================

// -----------------
// methods to implement
// -----------------
type IWeatherReportRepository_DataStore interface {
	Create(p_weatherReport any) error
	Update(p_weatherReport any) error
}
