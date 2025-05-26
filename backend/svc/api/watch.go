package api

import "github.com/oxidnova/novadm/backend/internal/config"

type watcher struct {
	s *Server
}

func (w *watcher) Action(data interface{}) {
	s := w.s
	conf, ok := data.(*config.Config)
	if !ok {
		s.d.Logger().Warn("reloading invalid config")
		return
	}

	s.restart(s.d.Logger(), conf)
}

func (w *watcher) Error(err error) {
	w.s.d.Logger().With("err", err).Fatal("reloading config")
}
