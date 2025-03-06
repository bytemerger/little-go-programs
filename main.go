package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	mu := sync.RWMutex{}
	memory := make(map[string]string)

	router := http.NewServeMux()

	router.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("processing a get request")
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "Please pass a valid key", http.StatusBadRequest)
			return
		}

		mu.RLock()
		val, ok := memory[key]
		mu.RUnlock()

		if ok {
			fmt.Fprintf(w, "%s", val)
			return
		}

		http.Error(w, "Key not found", http.StatusNotFound)
	})
	router.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("processing a set request")

		queryParams := r.URL.Query()
		// only sets the first query param
		var key, value string
		for queryKey, queryValue := range queryParams {
			key, value = queryKey, queryValue[0]
			break
		}
		if len(key) < 0 || len(value) < 0 {
			http.Error(w, "Malformed request", http.StatusBadRequest)
			return
		}

		mu.Lock()
		memory[key] = value
		mu.Unlock()

		fmt.Fprintf(w, "The key %s has been set in the memory\n", key)
	})

	fmt.Println("Server starting at port 4000")
	err := http.ListenAndServe(":4000", router)
	if err != nil {
		fmt.Println("Error starting server ", err)
	}
}
