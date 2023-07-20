package service

import "time"

//==================================================================================================
// IWeatherReportService interface
//==================================================================================================

// -----------------
// methods to implement
// -----------------
type ISchedulerService interface {
	ScheduleTask(p_task func() error, p_interval time.Duration) error
}
