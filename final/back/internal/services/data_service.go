package services

import (
	"sync"
)

func FetchUserData() string {
	// Simulate fetching user data
	return "User data"
}

func FetchOrderData() string {
	// Simulate fetching order data
	return "Order data"
}

func FetchDataConcurrently() (string, string) {
	var wg sync.WaitGroup
	var userData, orderData string

	wg.Add(2)

	go func() {
		defer wg.Done()
		userData = FetchUserData()
	}()

	go func() {
		defer wg.Done()
		orderData = FetchOrderData()
	}()

	wg.Wait()
	return userData, orderData
}
