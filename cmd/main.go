package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"

	"github.com/oklog/oklog/pkg/group"

	"github.com/hecomp/session-management/internal/util"
	. "github.com/hecomp/session-management/pkg/in_memory"
	. "github.com/hecomp/session-management/pkg/repository"
	"github.com/hecomp/session-management/pkg/session_management"
)

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
		inMemStore = NewInMemStore(logger)
	)

	var (
		sessionMgmntRepo = NewSessionMgmntRepository(inMemStore, logger)
	)

	var sessionMgmnt session_management.SessionMgmntService
	{
		sessionMgmnt = session_management.NewService(sessionMgmntRepo, logger)
		sessionMgmnt = session_management.NewLoggingService(log.With(logger, "component", "sessionMgmnt"), sessionMgmnt)
	}

	httpLogger := log.With(logger, "component", "http")

	var (
		httpHandler = session_management.MakeHandler(sessionMgmnt, sessionMgmntRepo, httpLogger)
	)

	// Now we're to the part of the func main where we want to start actually
	// running things, like servers bound to listeners to receive connections.
	//
	// The method is the same for each component: add a new actor to the group
	// struct, which is a combination of 2 anonymous functions: the first
	// function actually runs the component, and the second function should
	// interrupt the first function and cause it to return. It's in these
	// functions that we actually bind the Go kit server/handler structs to the
	// concrete transports and run them.
	//
	// Putting each component into its own block is mostly for aesthetics: it
	// clearly demarcates the scope in which each listener/socket may be used.
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


