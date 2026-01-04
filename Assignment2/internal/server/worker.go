package server

import (
	"fmt"
	"time"
)

func (s *Server) StartWorker() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.mu.Lock()
			fmt.Printf("Status: requests=%d, db_size=%d\n", s.requests, len(s.data))
			s.mu.Unlock()

		case <-s.shutdownCh:
			fmt.Println("Worker stopped")
			return
		}
	}
}
