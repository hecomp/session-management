package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/oklog/oklog/pkg/group"

	"github.com/hecomp/session-management/internal/util"
	. "github.com/hecomp/session-management/pkg/in_memory"
	. "github.com/hecomp/session-management/pkg/repository"
	"github.com/hecomp/session-management/pkg/session_management"
)

// SessionInterval parameter controls how frequently expired session data is removed by the
// background cleanup goroutine
const SessionInterval = 2 * time.Minute

func main() {

	// Define our flags. Your service probably won't need to bind listeners for
	// *all* supported transports, or support both Zipkin and LightStep, and so
	// on, but we do it here for demonstration purposes.
	fs := flag.NewFlagSet("sessionManagementSvc", flag.ExitOnError)

	var (
		httpAddr = fs.String("http_response-addr", ":8081", "HTTP listen address")
	)

	fs.Usage = util.UsageFor(fs, os.Args[0]+" [flags]")
	fs.Parse(os.Args[1:])

	// Create a single logger, which we'll use and give to other components.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var (
		inMemStore = NewInMemStore(SessionInterval, logger)
	)

	var (
		sessionMgmntRepo = NewSessionMgmntRepository(inMemStore, logger)
	)

	var sessionMgmnt session_management.SessionMgmntService
	{
		sessionMgmnt = session_management.NewService(sessionMgmntRepo, logger)
		sessionMgmnt = session_management.NewLoggingService(log.With(logger, "component", "sessionMgmnt"), sessionMgmnt)
	}

	var (
		httpHandler = session_management.MakeHandler(sessionMgmnt)
	)

	var g group.Group
	{
		// The HTTP listener mounts the Go kit HTTP handler we created.
		httpListener, err := net.Listen("tcp", *httpAddr)
		if err != nil {
			logger.Log("transport", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "HTTP", "addr", *httpAddr)
			return http.Serve(httpListener, httpHandler)
		}, func(error) {
			httpListener.Close()
		})
	}
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	logger.Log("exit", g.Run())
}


