package service

import (
	"fmt"
	"time"
)

//==================================================================================================
// WeatherReportService_Default struct
//==================================================================================================

// -----------------
// properties, fields
// -----------------
type SchedulerService_Default struct {
}

// -----------------
// constructors
// -----------------
func NewSchedulerService_Default() (*SchedulerService_Default, error) {
	return &SchedulerService_Default{}, nil
}

// -----------------
// methods
// -----------------
func (s *SchedulerService_Default) ScheduleTask(p_task func() error, p_interval time.Duration) error {
	/* Starts to run the task every interval duration */
	go func() {
		for {
			fmt.Println("Starting task...")
			err := p_task()
			if err != nil {
				fmt.Println("Task failed: ", err)
			}
			fmt.Println("Task completed. Now sleeping for ", p_interval)
			time.Sleep(p_interval)
		}
	}()

	return nil
}
