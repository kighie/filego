package common

import (
	"time"
)

type StopWatch struct {
	start 	time.Time
	mark	time.Time
	stop	time.Time
}

func (sw *StopWatch) Start() time.Time{
	sw.start = time.Now()
	sw.mark = sw.start
	return sw.start
}

// mark now. return FromLastMark() duration
func (sw *StopWatch) Mark() time.Duration {
	d := time.Since(sw.mark)
	sw.mark = time.Now()
	return d
}

// return ellapsed time from last mark
func (sw *StopWatch) FromLastMark() time.Duration{
	return time.Since(sw.mark)
}

func (sw *StopWatch) FromStart() time.Duration{
	return time.Since(sw.start)
}

func (sw *StopWatch) Stop() time.Duration{
	sw.stop = time.Now()
	return time.Since(sw.start)
}


func (sw *StopWatch) GetStartTime() time.Time{
	return sw.start
}

func (sw *StopWatch) GetEndTime() time.Time{
	return sw.stop
}



