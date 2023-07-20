package container

import (
	"workspace/model/repository"
	"workspace/service"
)

// ==================================================================================================
// DI Container functions
// ==================================================================================================

func NewIWeatherReportRepository_DataStore() (repository.IWeatherReportRepository_DataStore, error) {
	repoDS, err := repository.NewWeatherReportRepository_MongoDB()
	return repoDS, err
}

func NewIWeatherReportRepository_API() (repository.IWeatherReportRepository_API, error) {
	repoAPI, err := repository.NewWeatherReportRepository_meteoAPI()
	return repoAPI, err
}

func NewIWeatherReportService() (service.IWeatherReportService, error) {
	repoDS, err := NewIWeatherReportRepository_DataStore()
	if err != nil {
		return nil, err
	}
	repoAPI, err := NewIWeatherReportRepository_API()
	if err != nil {
		return nil, err
	}

	service, err := service.NewWeatherReportService_Default(repoDS, repoAPI)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func NewISchedulerService() (service.ISchedulerService, error) {
	service, err := service.NewSchedulerService_Default()
	if err != nil {
		return nil, err
	}

	return service, nil
}
