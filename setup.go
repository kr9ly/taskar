package tasker

import (
	"github.com/robfig/cron"
)

func Setup(conf *Config, c *cron.Cron) {
	for _, t := range conf.Triggers {
		c.AddFunc(t.Time, func() {
			for _, a := range t.Actions {
				ExecAction(a, conf)
			}
		})
	}
}
