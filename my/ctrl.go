package my

import (
	"syscall"
	"os"
	"os/signal"
)

//ctrl+cをハンドル
func HandleCtrlc(fn func()){
	HandleSig(fn, syscall.SIGINT)
}

func HandleSig(fn func(), sigtype syscall.Signal){
	signal_chan := make(chan os.Signal)
	signal.Notify(signal_chan,
		sigtype,
	)
	for{
		sig := <-signal_chan
		switch sig {
		case sigtype:
			fn()
		}
	}
}