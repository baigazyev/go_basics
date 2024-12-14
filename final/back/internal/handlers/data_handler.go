package handlers

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint"},
	)
)

func init() {
	// Register custom metrics
	prometheus.MustRegister(httpRequestsTotal)
}

func fetchUserData(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done() // Mark this Goroutine as done
	time.Sleep(2 * time.Second)
	ch <- "User data fetched"
}

func fetchOrderData(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done() // Mark this Goroutine as done
	time.Sleep(3 * time.Second)
	ch <- "Order data fetched"
}

func FetchDataHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	httpRequestsTotal.WithLabelValues("GET", "/fetch-data").Inc()

	var wg sync.WaitGroup
	wg.Add(2)

	dataChannel := make(chan string, 2)

	go fetchUserData(&wg, dataChannel)
	go fetchOrderData(&wg, dataChannel)

	wg.Wait()
	close(dataChannel)

	response := []string{}
	for data := range dataChannel {
		response = append(response, data)
	}

	duration := time.Since(start)
	log.Printf("INFO: /fetch-data completed in %v", duration)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	for _, msg := range response {
		fmt.Fprintf(w, "%s\n", msg)
	}
}
