package service

import (
	"bufio"
	"log"
	"os"
	"strings"
	"workspace/model/repository"
)

//==================================================================================================
// WeatherReportService_Default struct
//==================================================================================================

// -----------------
// properties, fields
// -----------------
type WeatherReportService_Default struct {
	RepoDS  repository.IWeatherReportRepository_DataStore
	RepoAPI repository.IWeatherReportRepository_API
}

// -----------------
// constructors
// -----------------
func NewWeatherReportService_Default(p_repoDS repository.IWeatherReportRepository_DataStore, p_repoAPI repository.IWeatherReportRepository_API) (*WeatherReportService_Default, error) {
	return &WeatherReportService_Default{RepoDS: p_repoDS, RepoAPI: p_repoAPI}, nil
}

// -----------------
// methods
// -----------------
func (s *WeatherReportService_Default) WeatherReportByDate(p_date any) (any, error) {
	weatherReport, err := s.RepoAPI.ReadBy(p_date.(string))
	return weatherReport, err
}

func (s *WeatherReportService_Default) CheckAndStoreWeatherReport(p_report any) error {
	err := s.RepoDS.Update(p_report)
	if err != nil {
		err = s.RepoDS.Create(p_report)
	}
	return err
}

func (s *WeatherReportService_Default) ReadDatesFrom(p_path string) any {

	file, err := os.Open(p_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	/* Reads the file into string */
	scanner := bufio.NewScanner(file)
	// // var dates []time.Time
	var dates []string
	for scanner.Scan() {
		/* Splits the string into a slice of strings by comma then trims spaces and parses to time */
		dateStrings := strings.Split(scanner.Text(), ",")
		for i := 0; i < len(dateStrings); i++ {
			// // date, err := time.Parse("2006-01-02", dateStrings[i])
			// // if err != nil {
			// // 	log.Fatal(err)
			// // }
			dates = append(dates, dateStrings[i])
		}
	}
	return dates
}
