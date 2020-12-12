package input

import (
	"time"

	flag "github.com/spf13/pflag"
)

var (
	Visualize bool
	Delay     time.Duration
	Timing    bool
	Day       int
)

func ParseFlags() {
	flag.BoolVar(&Timing, "time", false, "Times the solution, do not show it")
	flag.IntVar(&Day, "day", time.Now().Day(), "What day to solve [doesn't work for timing, which is applied to all days and then shows the results separated]")
	flag.BoolVar(&Visualize, "visualize", false, "show visualization of solve")
	d := flag.Int64("delay", 300, "visualization delay")
	flag.Parse()
	if !Visualize {
		Delay = 0
	} else {
		Delay = time.Duration(*d)
	}
}
