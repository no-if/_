package work

import (
	"time"
)

func Once(task_time time.Time, task_func func()) bool {
	var sub = task_time.Sub(time.Now())
	if sub > 0 {
		<-time.After(sub)
		go task_func()
		return true
	}
	return false
}

func Do(hms string, second int, begin_exec bool, task_func func()) {
	if begin_exec {
		go task_func()
	}

	var (
		now       time.Time
		task_time time.Time
		err       error
	)
	if second == 0 {
		second = 60 * 60 * 24
	}

	now = time.Now()
	if task_time, err = time.Parse("15:04:05", hms); err != nil {
		task_time = now
	} else {
		task_time = time.Date(now.Year(), now.Month(), now.Day(), task_time.Hour(), task_time.Minute(), task_time.Second(), 0, time.Local)
	}

	go func() {
		for {
			if Once(task_time, task_func) == false {
				task_time = task_time.Add(time.Duration(1e9 * second))
			}
		}
	}()

}
