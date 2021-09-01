package shutdown

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

// SignalContext returns a context which will be canceled when either the SIGINT or
// SIGTERM signal is caught. It also returns a function that can be used to
// programmatically cancel the same context at any time. If either signal is
// caught a second time, the program is terminated immediately with exit code 1.
func SignalContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigCh
		logrus.Infof("Received signal %s, beginning shutdown", sig.String())
		cancel()
		sig = <-sigCh
		logrus.Fatalf(
			`Received signal "%s" during shutdown; exiting immediately`,
			sig,
		)
	}()
	return ctx
}
