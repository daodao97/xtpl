package job

import (
	"fmt"
	"time"

	"github.com/daodao97/xgo/xapp"
	"github.com/daodao97/xgo/xcron"
	"github.com/daodao97/xgo/xredis"
)

func NewCronServer() xapp.NewServer {
	return func() xapp.Server {
		return xcron.New2(
			xcron.WithJobs(
				xcron.Job{
					Name: "CronJobName",
					Spec: "@every 10s",
					Func: func() {
						fmt.Println("CronJobName test", time.Now())
					},
					Immediate: true,
				},
			),
			xcron.WithRdb(xredis.Get()),
		)
	}
}
