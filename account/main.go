package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/iradukunda1/account/handlers"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

// errOnShutdownCode is the exit code if there is an error shutting down
const errOnShutdownCode = 1

// harTimeout is the time to wait before forcing a shutdown
const hardTimeout = 5 * time.Second

func main() {

	ctx := context.Background()

	log.Printf("starting account service")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r := http.NewServeMux()

	r.HandleFunc("/create", handlers.Create)
	r.HandleFunc("/deposit", handlers.Deposit)
	r.HandleFunc("/withdraw", handlers.WithDraw)
	r.HandleFunc("/balance", handlers.Balance)

	g := errgroup.Group{}

	g.Go(func() error {
		logrus.WithFields(logrus.Fields{
			"port": port,
		}).Infof("starting the application http server")
		return ListenAndServe(ctx, r, port)
	})

	if err := g.Wait(); err != nil {
		logrus.Fatal("main: program terminated")
	}

}

func ListenAndServe(ctx context.Context, handler http.Handler, addr string) error {
	var g errgroup.Group

	srv := &http.Server{
		Addr:    ":" + addr,
		Handler: handler,
	}

	// Setup our Ctrl+C handler
	g.Go(func() error {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

		killSignal := <-interrupt
		switch killSignal {
		case os.Interrupt:
			logrus.Print("got SIGINT...")
		case syscall.SIGTERM:
			logrus.Print("got SIGTERM...")
		}

		ctx, cancel := context.WithTimeout(ctx, hardTimeout)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			logrus.Errorf("shutting down with error: %v", err)
			os.Exit(errOnShutdownCode)
		} else {
			os.Exit(0)
		}
		return nil
	})

	g.Go(func() error {
		<-ctx.Done()
		return srv.Shutdown(ctx)
	})

	g.Go(srv.ListenAndServe)

	return g.Wait()
}
