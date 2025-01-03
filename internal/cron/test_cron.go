package cron

import (
	"fmt"
	"github.com/zngue/zng_app/pkg/cron"
)

type TestCron struct {
	cron.IAppServer
}

func NewTestCron() *TestCron {
	return &TestCron{}
}
func (t *TestCron) Run() {
	fmt.Println("TestCron->run")
}
func (t *TestCron) Stop() {
	fmt.Println("TestCron->stop")
}
