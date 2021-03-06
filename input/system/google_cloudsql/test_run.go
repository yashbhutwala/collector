package google_cloudsql

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/pganalyze/collector/logs"
	"github.com/pganalyze/collector/state"
	"github.com/pganalyze/collector/util"
)

func LogTestRun(server *state.Server, globalCollectionOpts state.CollectionOpts, logger *util.Logger) error {
	cctx, cancel := context.WithCancel(context.Background())

	// We're testing one server at a time during the test run for now
	servers := []*state.Server{server}

	logTestSucceeded := make(chan bool, 1)
	gcpLogStream := make(chan LogStreamItem, 500)
	wg := sync.WaitGroup{}
	err := SetupLogSubscriber(cctx, &wg, globalCollectionOpts, logger, servers, gcpLogStream)
	if err != nil {
		cancel()
		return err
	}
	logReceiver(cctx, servers, gcpLogStream, globalCollectionOpts, logger, logTestSucceeded)

	logs.EmitTestLogMsg(server, globalCollectionOpts, logger)

	select {
	case <-logTestSucceeded:
		cancel()
		return nil
	case <-time.After(10 * time.Second):
		cancel()
		return fmt.Errorf("Timeout")
	}
}
