package api

import (
	"net/http"
	"time"

	"example.com/wirebridge"
)

// Options HTTP 服务选项。
type Options struct {
	WebDir    string
	AllowCORS bool
}

// Server 管理 API。
type Server struct {
	bridge *wirebridge.Bridge
	opts   Options
	mux    *http.ServeMux
}

// New 构造。
func New(b *wirebridge.Bridge, opts Options) *Server {
	s := &Server{bridge: b, opts: opts, mux: http.NewServeMux()}
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.opts.AllowCORS {
		withCORS(w)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	s.mux.ServeHTTP(w, r)
}

func (s *Server) routes() {
	s.mux.HandleFunc("/api/health", s.handleHealth)
	s.mux.HandleFunc("/api/routes", s.handleRoutes)
	s.mux.HandleFunc("/api/serve", s.handleServe)
	s.mux.HandleFunc("/api/stats", s.handleStats)
	s.mux.HandleFunc("/api/encode", s.handleEncode)
	if s.opts.WebDir != "" {
		s.mux.Handle("/", staticHandler(s.opts.WebDir))
	}
}

// ListenAndServe 便捷启动。
func ListenAndServe(addr string, b *wirebridge.Bridge, opts Options) error {
	srv := &http.Server{
		Addr:              addr,
		Handler:           New(b, opts),
		ReadHeaderTimeout: 5 * time.Second,
	}
	return srv.ListenAndServe()
}
