package quit

import (
	"os"
	"os/signal"
)

func Run(exit_func func()) {
	var quit_chan = make(chan os.Signal)
	signal.Notify(quit_chan)
	<-quit_chan
	exit_func()
}
