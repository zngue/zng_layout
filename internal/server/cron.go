package server

import (
	c "github.com/zngue/zng_app/pkg/cron"
	"github.com/zngue/zng_layout/internal/cron"
)

func NewCron(testCron *cron.TestCron) (items []c.ICron, err error) {
	items = []c.ICron{
		testCron,
	}
	return
}
