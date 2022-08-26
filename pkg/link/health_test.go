package link

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteChecker(t *testing.T) {
	t.Run("returns no error if link is not broken", func(t *testing.T) {
		serverA := makeHttpServer(http.StatusOK)
		defer serverA.Close()

		err := CheckHealth(context.Background(), serverA.URL)
		assertNoError(t, err)
	})

	t.Run("returns error, if link is broken", func(t *testing.T) {
		serverA := makeHttpServer(http.StatusNotFound)
		defer serverA.Close()

		err := CheckHealth(context.Background(), serverA.URL)
		assertError(t, err)
	})

	t.Run("returns error on timeout", func(t *testing.T) {
		serverA := makeDelayedHttpServer(100 * time.Millisecond)
		defer serverA.Close()

		cancellingCtx, cancel := context.WithCancel(context.Background())
		time.AfterFunc(10*time.Millisecond, cancel)

		err := CheckHealth(cancellingCtx, serverA.URL)
		assertError(t, err)
	})
}

func assertError(t testing.TB, err error) {
	t.Helper()

	if err == nil {
		t.Error("wanted error, got nil")
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Error("got error, wanted nil")
	}
}

func makeHttpServer(statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
	}))
}

func makeDelayedHttpServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
