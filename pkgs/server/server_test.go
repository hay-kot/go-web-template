package server

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func testServer(t *testing.T, r http.Handler) *Server {
	svr := NewServer("127.0.0.1", "19245")

	go func() {
		err := svr.Start(r)
		assert.NoError(t, err) // Just in case?
	}()

	// Hack to get wait for the server to start
	time.Sleep(time.Second)

	return svr
}

func Test_ServerStarts(t *testing.T) {
	svr := testServer(t, nil)
	err := svr.Shutdown("test")
	assert.NoError(t, err)
}

func Test_GracefulServerShutdownWithWorkers(t *testing.T) {
	isFinished := false

	svr := testServer(t, nil)

	svr.Background(func() {
		time.Sleep(time.Second * 4)
		isFinished = true
	})

	err := svr.Shutdown("test")

	assert.NoError(t, err)
	assert.True(t, isFinished)

}

func Test_GracefulServerShutdownWithRequests(t *testing.T) {
	isFinished := false

	router := http.NewServeMux()

	router.HandleFunc("/test", func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 4)
		isFinished = true
	})

	svr := testServer(t, router)

	// Make request to "/test"
	go func() {
		_, err := http.Get("http://localhost:19245/test") // This is probably bad?
		assert.NoError(t, err)
	}()

	time.Sleep(time.Second) // Hack to wait for the request to be made

	err := svr.Shutdown("test")
	assert.NoError(t, err)

	assert.True(t, isFinished)
}
