package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type KVStore struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewKVStore() *KVStore {
	return &KVStore{
		data: make(map[string]string),
	}
}

type Item struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (s *KVStore) handleSet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "methond not allowed. use POST", http.StatusMethodNotAllowed)
		return
	}

	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	s.data[item.Key] = item.Value
	s.mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "key '%s' saved successfully \n", item.Key)
}

func main() {

	store := NewKVStore()

	fmt.Println("Store initialized:", store)

	http.HandleFunc("/set", store.handleSet)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error  Starting Server", err)
	}

}
