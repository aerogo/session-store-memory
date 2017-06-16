package memstore

import (
	"github.com/aerogo/aero"
	"golang.org/x/sync/syncmap"
)

// MemoryStore is the default session store.
// You should use it for prototyping, not for production.
type MemoryStore struct {
	sessions syncmap.Map
}

// NewMemoryStore creates a session memory store.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

// Get a session by its ID.
func (store *MemoryStore) Get(id string) *aero.Session {
	session, ok := store.sessions.Load(id)

	if !ok {
		return nil
	}

	return session.(*aero.Session)
}

// Set saves a session so it can be retrieved by its ID.
func (store *MemoryStore) Set(id string, session *aero.Session) {
	store.sessions.Store(id, session)
}
