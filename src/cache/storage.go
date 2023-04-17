package cache

import "sync"

var (
	storage map[string]interface{}
	mutex   sync.Mutex
)

func init() {
	storage = make(map[string]interface{})
}

// Puts key-value pair into in-memory storage
// If the exists, then the function overwrites it.
func Put(key string, value interface{}) {

	done := make(chan bool)
	go addKeyValue(key, value, done)

	<-done
}

// Gets value from storage by key
// If key does not exists, returns nil
func GetValue(key string) interface{} {
	result := make(chan interface{})
	go getByKey(key, result)

	return <-result
}

func addKeyValue(key string, value interface{}, done chan bool) {
	mutex.Lock()
	storage[key] = value
	done <- true
	mutex.Unlock()
}

func getByKey(key string, result chan interface{}) {
	mutex.Lock()
	result <- storage[key]
	mutex.Unlock()
}
