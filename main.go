package main

import (
	"fmt"
	"os"
	"time"

	"workspace/container"
	"workspace/model/domain"
)

// ==================================================================================================
// MAIN
// ==================================================================================================
func main() {

	scheduler, _ := container.NewISchedulerService()
	scheduler.ScheduleTask(task, time.Duration(10*time.Second))
	select {}
}

func task() error {
	// -----------------
	// 0. Setup
	// -----------------
	myService, err := container.NewIWeatherReportService()
	if err != nil {
		fmt.Println(err)
	}

	// -----------------
	// 1. Read dates from file
	// -----------------
	datesFilePath := os.Getenv("DATES_FILE_PATH")
	if datesFilePath == "" {
		datesFilePath = "datesToObserve.csv"
	}
	dates := myService.ReadDatesFrom(datesFilePath).([]string)
	fmt.Println("Dates read from file:")
	fmt.Println(dates)

	// -----------------
	// 2. Get weather report from API for each date
	// -----------------
	var reportsFromAPI []domain.Daily

	fmt.Println("Weather reports from API:")
	for _, date := range dates {

		report, err := myService.WeatherReportByDate(date)
		if err != nil {
			fmt.Println(err)
		}

		/* Extracts relevant data from report */
		relevantDailyData := report.(domain.WeatherReport).Daily
		fmt.Println(relevantDailyData)

		/* Stores for further use */
		reportsFromAPI = append(reportsFromAPI, relevantDailyData)
	}

	// -----------------
	// 3. Check if weather report exists in DB for each date (if does, update, if not, create)
	// -----------------
	fmt.Println("Checking and storing weather report...")
	for _, report := range reportsFromAPI {

		err := myService.CheckAndStoreWeatherReport(report)
		if err != nil {
			fmt.Println(err)
		}

	}
	fmt.Println("Success!")

	// -----------------
	// 4. Release resources
	// -----------------
	myService = nil

	return nil
}
