package api

import (
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"example.com/wirebridge"
)

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]any{
		"ok":   !s.bridge.IsClosed(),
		"name": s.bridge.Name(),
		"max":  s.bridge.MaxFrame(),
	})
}

func (s *Server) handleRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rows, err := s.bridge.ListRoutes()
		if err != nil {
			writeErr(w, err, 500)
			return
		}
		writeJSON(w, rows)
	case http.MethodPut:
		var rows []wirebridge.RouteMeta
		if err := json.NewDecoder(r.Body).Decode(&rows); err != nil {
			writeErr(w, err, 400)
			return
		}
		if err := s.bridge.ApplyRoutes(rows); err != nil {
			writeErr(w, err, 500)
			return
		}
		writeJSON(w, map[string]string{"status": "ok"})
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleServe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	raw, err := io.ReadAll(io.LimitReader(r.Body, int64(s.bridge.MaxFrame())+64))
	if err != nil {
		writeErr(w, err, 400)
		return
	}
	// 允许 hex 文本
	if ct := r.Header.Get("Content-Type"); ct == "text/plain" || r.URL.Query().Get("hex") == "1" {
		dec, err := hex.DecodeString(trimSpace(string(raw)))
		if err != nil {
			writeErr(w, err, 400)
			return
		}
		raw = dec
	}
	out, err := s.bridge.ServeFrame(raw)
	if err != nil {
		writeErr(w, err, 400)
		return
	}
	writeJSON(w, map[string]any{
		"opcode":  out.Opcode,
		"flags":   out.Flags,
		"payload": hex.EncodeToString(out.Payload),
	})
}

func (s *Server) handleStats(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, s.bridge.SnapshotStats())
}

func (s *Server) handleEncode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Opcode  uint16 `json:"opcode"`
		Flags   uint8  `json:"flags"`
		Payload string `json:"payload"` // hex
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErr(w, err, 400)
		return
	}
	pl, err := hex.DecodeString(req.Payload)
	if err != nil {
		writeErr(w, err, 400)
		return
	}
	raw, err := s.bridge.Encode(wirebridge.Frame{
		Opcode:  wirebridge.Opcode(req.Opcode),
		Flags:   wirebridge.Flags(req.Flags),
		Payload: pl,
	})
	if err != nil {
		writeErr(w, err, 400)
		return
	}
	writeJSON(w, map[string]string{
		"hex":  hex.EncodeToString(raw),
		"size": strconv.Itoa(len(raw)),
	})
}

func trimSpace(s string) string {
	for len(s) > 0 && (s[0] == ' ' || s[0] == '\n' || s[0] == '\r' || s[0] == '\t') {
		s = s[1:]
	}
	for len(s) > 0 {
		c := s[len(s)-1]
		if c != ' ' && c != '\n' && c != '\r' && c != '\t' {
			break
		}
		s = s[:len(s)-1]
	}
	return s
}
