package benchmark

import (
	"sync"

	"github.com/aerogo/aero"
)

// MemoryStoreRWLock is the default session store.
// You should use it for prototyping, not for production.
type MemoryStoreRWLock struct {
	sessions map[string]*aero.Session
	lock     sync.RWMutex
}

// NewMemoryStoreRWLock creates a session memory store.
func NewMemoryStoreRWLock() *MemoryStoreRWLock {
	return &MemoryStoreRWLock{
		sessions: make(map[string]*aero.Session),
	}
}

// Get a session by its ID.
func (store *MemoryStoreRWLock) Get(id string) *aero.Session {
	store.lock.RLock()
	session := store.sessions[id]
	store.lock.RUnlock()
	return session
}

// Set saves a session so it can be retrieved by its ID.
func (store *MemoryStoreRWLock) Set(id string, session *aero.Session) {
	store.lock.Lock()
	store.sessions[id] = session
	store.lock.Unlock()
}
