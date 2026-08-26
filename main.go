package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Signal struct {
	Device string `json:"device"`
	Lane   string `json:"lane"`
	Value  string `json:"value"`
}

type Router struct {
	mu      sync.Mutex
	lanes   map[string][]Signal
	workers int
}

func NewRouter(workers int) *Router {
	return &Router{lanes: make(map[string][]Signal), workers: workers}
}

func (r *Router) Enqueue(s Signal) error {
	if s.Device == "" || s.Lane == "" {
		return errors.New("device and lane are required")
	}
	r.mu.Lock()
	r.lanes[s.Lane] = append(r.lanes[s.Lane], s)
	r.mu.Unlock()
	return nil
}

func (r *Router) Drain(lane string) []Signal {
	r.mu.Lock()
	defer r.mu.Unlock()
	items := append([]Signal(nil), r.lanes[lane]...)
	delete(r.lanes, lane)
	return items
}

func (r *Router) Count(lane string) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return len(r.lanes[lane])
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost || req.URL.Path != "/signals" {
		http.NotFound(w, req)
		return
	}
	var signal Signal
	if err := json.NewDecoder(req.Body).Decode(&signal); err != nil {
		http.Error(w, "invalid signal", http.StatusBadRequest)
		return
	}
	if err := r.Enqueue(signal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"accepted": true, "lane": signal.Lane})
}

func (r *Router) WaitFor(lane string, n int, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if r.Count(lane) >= n {
			return nil
		}
		time.Sleep(time.Millisecond)
	}
	return fmt.Errorf("lane %s did not receive %d signals", lane, n)
}

func main() {
	_ = http.ListenAndServe(":8080", NewRouter(2))
}
