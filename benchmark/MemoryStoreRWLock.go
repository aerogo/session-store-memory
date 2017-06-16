package benchmark

import (
	"sync"

	"github.com/aerogo/session"
)

// MemoryStoreRWLock is the default session store.
// You should use it for prototyping, not for production.
type MemoryStoreRWLock struct {
	sessions map[string]*session.Session
	lock     sync.RWMutex
}

// NewMemoryStoreRWLock creates a session memory store.
func NewMemoryStoreRWLock() *MemoryStoreRWLock {
	return &MemoryStoreRWLock{
		sessions: make(map[string]*session.Session),
	}
}

// Get a session by its ID.
func (store *MemoryStoreRWLock) Get(id string) *session.Session {
	store.lock.RLock()
	session := store.sessions[id]
	store.lock.RUnlock()
	return session
}

// Set saves a session so it can be retrieved by its ID.
func (store *MemoryStoreRWLock) Set(id string, session *session.Session) {
	store.lock.Lock()
	store.sessions[id] = session
	store.lock.Unlock()
}
