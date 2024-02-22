package crontab

import "github.com/robfig/cron"

func Cron() {
	c := cron.New()
	//c.AddFunc("0 0 * * * *", CronCalculationGameSum)	// 想起来应该交给监控系统
	c.Start()
}
