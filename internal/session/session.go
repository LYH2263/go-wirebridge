package session

import (
	"sync"
	"time"
)

// Session 简易会话（管理试帧关联）。
type Session struct {
	ID        string
	CreatedAt time.Time
	LastOp    uint16
	Count     int
}

// Store 会话表。
type Store struct {
	mu sync.Mutex
	m  map[string]*Session
}

// NewStore 构造。
func NewStore() *Store {
	return &Store{m: make(map[string]*Session)}
}

// Touch 更新或创建。
func (s *Store) Touch(id string, op uint16) *Session {
	s.mu.Lock()
	defer s.mu.Unlock()
	cur, ok := s.m[id]
	if !ok {
		cur = &Session{ID: id, CreatedAt: time.Now().UTC()}
		s.m[id] = cur
	}
	cur.LastOp = op
	cur.Count++
	return &Session{ID: cur.ID, CreatedAt: cur.CreatedAt, LastOp: cur.LastOp, Count: cur.Count}
}

// Get 拷贝。
func (s *Store) Get(id string) (Session, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	cur, ok := s.m[id]
	if !ok {
		return Session{}, false
	}
	return *cur, true
}

// Delete 删除。
func (s *Store) Delete(id string) {
	s.mu.Lock()
	delete(s.m, id)
	s.mu.Unlock()
}

// Len 数量。
func (s *Store) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.m)
}
