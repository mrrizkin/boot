package app

import "github.com/mrrizkin/boot/third-party/scheduler"

// Schedule sets up scheduled tasks for the application.
// It takes a scheduler.Scheduler as an argument to define and manage recurring tasks.
func (a *App) Schedule(schedule scheduler.Scheduler) {
	// Example usage:
	// schedule.Add("@every 1m", func() {
	//     a.Log("info", "Scheduled task ran at %s", time.Now().Format(time.RFC3339))
	// })
	//
	// Add your scheduled tasks here. Use cron syntax for defining intervals.
	// Refer to the scheduler documentation for more advanced usage.
}
