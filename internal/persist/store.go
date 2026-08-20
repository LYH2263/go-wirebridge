package persist

import (
	"encoding/json"
	"os"
	"path/filepath"

	"example.com/wirebridge/internal/route"
)

type fileDoc struct {
	Version int          `json:"version"`
	Routes  []route.Meta `json:"routes"`
}

// Save 原子写路由快照。
func Save(path string, rows []route.Meta) error {
	if path == "" {
		return nil
	}
	doc := fileDoc{Version: 1, Routes: make([]route.Meta, 0, len(rows))}
	for _, r := range rows {
		doc.Routes = append(doc.Routes, route.CloneMeta(r))
	}
	raw, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return err
	}
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, raw, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

// Load 读取快照。
func Load(path string) ([]route.Meta, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var doc fileDoc
	if err := json.Unmarshal(raw, &doc); err != nil {
		return nil, err
	}
	out := make([]route.Meta, 0, len(doc.Routes))
	for _, r := range doc.Routes {
		out = append(out, route.CloneMeta(r))
	}
	return out, nil
}
