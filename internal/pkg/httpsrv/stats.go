package httpsrv

import "log"

type sessionStats struct {
	id   string
	sent int
}

func (w *sessionStats) print() {
	log.Printf("session %s has received %d messages\n", w.id, w.sent)
}

func (w *sessionStats) inc() {
	w.sent++
}

func (s *Server) incStats(id string) {
	// Find and increment.
	for i := range s.sessionStats {
		if s.sessionStats[i].id == id {
			s.sessionStats[i].inc()
			return
		}
		/* if ws.id == id {
			ws.inc()
			return
		} */
	}
	// Not found, add new.
	s.sessionStats = append(s.sessionStats, sessionStats{id: id, sent: 1})
}
