package input

import (
	"time"

	flag "github.com/spf13/pflag"
)

var (
	Visualize bool
	Delay     time.Duration
)

func ParseFlags() {
	flag.BoolVar(&Visualize, "v", false, "show visualization of solve")
	d := flag.Int64("d", 300, "visualization delay")
	flag.Parse()
	Delay = time.Duration(*d)
}
