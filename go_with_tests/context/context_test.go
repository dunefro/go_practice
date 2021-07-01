package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("tells store to cancel the work if request is cancelled", func(t *testing.T) {
		data := "Hello world"
		store := &SpyStore{
			response: data,
		}
		svr := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)
		store.assertWasNotCancelled(t)

	})
	t.Run("No cancellation of request", func(t *testing.T) {
		data := "Hello world"
		store := &SpyStore{
			response: data,
		}
		svr := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)
		if response.Body.String() != data {
			t.Errorf("Expected %s , Recieved %s", data, response.Body.String())
		}
		store.assertWasCancelled(t)
	})

}

func (store *SpyStore) assertWasCancelled(t *testing.T) {
	t.Helper()
	if store.cancelled {
		t.Errorf("Expected the store to not get cancelled")
	}
}

func (store *SpyStore) assertWasNotCancelled(t *testing.T) {
	t.Helper()
	if !store.cancelled {
		t.Error("Expected the store to get cancelled")
	}
}
