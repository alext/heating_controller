package scheduler

import (
	"fmt"
	"time"

	"github.com/alext/heating-controller/logger"
	"github.com/alext/heating-controller/output"
)

type Event struct {
	Hour   int
	Min    int
	Action Action
}

func (e Event) NextOccurance() time.Time {
	return e.nextOccuranceAfter(time_Now().Local())
}

func (e Event) nextOccuranceAfter(current time.Time) time.Time {
	next := time.Date(current.Year(), current.Month(), current.Day(), e.Hour, e.Min, 0, 0, time.Local)
	if next.Before(current) {
		current = current.AddDate(0, 0, 1)
		next = time.Date(current.Year(), current.Month(), current.Day(), e.Hour, e.Min, 0, 0, time.Local)
	}
	return next
}

func (e Event) after(hour, min int) bool {
	return e.Hour > hour || (e.Hour == hour && e.Min > min)
}

func (e Event) do(out output.Output) {
	var err error
	if e.Action == TurnOn {
		logger.Infof("[Scheduler:%s] Activating output", out.Id())
		err = out.Activate()
	} else {
		logger.Infof("[Scheduler:%s] Deactivating output", out.Id())
		err = out.Deactivate()
	}
	if err != nil {
		logger.Warnf("[Scheduler:%s] Output error: %v", out.Id(), err)
	}
}

func (e Event) String() string {
	return fmt.Sprintf("%d:%02d %s", e.Hour, e.Min, e.Action)
}

type eventList []*Event

func (el eventList) Len() int      { return len(el) }
func (el eventList) Swap(i, j int) { el[i], el[j] = el[j], el[i] }
func (el eventList) Less(i, j int) bool {
	a, b := el[i], el[j]
	return a.Hour < b.Hour || (a.Hour == b.Hour && a.Min < b.Min)
}
