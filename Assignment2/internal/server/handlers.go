package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) PostData(w http.ResponseWriter, r *http.Request) {
	var body map[string]string
	if json.NewDecoder(r.Body).Decode(&body) != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	s.requests++
	for k, v := range body {
		s.data[k] = v
	}
	s.mu.Unlock()

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) GetData(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	s.requests++
	dataCopy := make(map[string]string)
	for k, v := range s.data {
		dataCopy[k] = v
	}
	s.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dataCopy)
}

func (s *Server) DeleteData(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	s.mu.Lock()
	s.requests++
	delete(s.data, key)
	s.mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) Stats(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	s.requests++
	stats := map[string]int{
		"total_requests": s.requests,
		"db_size":        len(s.data),
	}
	s.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
