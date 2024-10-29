package httpsrv

import (
	"goapp/internal/pkg/watcher"
)

func deleteStat(s []sessionStats, index int) []sessionStats {
	ret := make([]sessionStats, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func (s *Server) addWatcher(w *watcher.Watcher) {
	s.watchersLock.Lock()
	defer s.watchersLock.Unlock()
	s.watchers[w.GetWatcherId()] = w
}

func (s *Server) removeWatcher(w *watcher.Watcher) {
	s.watchersLock.Lock()
	defer s.watchersLock.Unlock()
	// Print satistics before removing watcher.
	for i := range s.sessionStats {
		if s.sessionStats[i].id == w.GetWatcherId() {
			s.sessionStats[i].print()
		}
	}

	delete(s.watchers, w.GetWatcherId())
}

func (s *Server) notifyWatchers(str string) {
	s.watchersLock.RLock()
	defer s.watchersLock.RUnlock()

	// Send message to all watchers and increment stats.
	for id := range s.watchers {
		s.watchers[id].Send(str)
		s.incStats(id)
	}
}
