package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// TODO: #3 Implement graceful shutdown
// TODO: #2 Implement Go routine pool/job queue

type Server struct {
	Host string
	Port string
	Wg   *sync.WaitGroup
}

func NewServer(host, port string) *Server {
	return &Server{
		Host: host,
		Port: port,
		Wg:   &sync.WaitGroup{},
	}
}

func (s *Server) Start(router http.Handler) error {
	httpServer := &http.Server{
		Addr:         s.Host + ":" + s.Port,
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		// Create a quit channel which carries os.Signal values.
		quit := make(chan os.Signal, 1)

		// Use signal.Notify() to listen for incoming SIGINT and SIGTERM signals and
		// relay them to the quit channel.
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// Read the signal from the quit channel. block until received
		s := <-quit
		fmt.Println("received termination request -> ", s.String())

		// Create a context with a 5-second timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		shutdownError <- httpServer.Shutdown(ctx)

		// Exit the application with a 0 (success) status code.
		os.Exit(0)
	}()

	err := httpServer.ListenAndServe()

	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	fmt.Println("Server shutdown successfully")

	return nil
}

// Background starts a go routine that runs on the servers pool. In the event of a shutdown
// request, the server will wait until all open goroutines have finished before shutting down.
func (s *Server) Background(task func()) {
	s.Wg.Add(1)
	go func() {
		defer s.Wg.Done()
		task()
	}()
}
