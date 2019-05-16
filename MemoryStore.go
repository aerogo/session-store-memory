package memstore

import (
	"errors"
	"sync"

	"github.com/aerogo/session"
)

// MemoryStore is the default session store.
// You should use it for prototyping, not for production.
type MemoryStore struct {
	sessions sync.Map
}

// New creates a session memory store.
func New() *MemoryStore {
	return &MemoryStore{}
}

// Get a session by its ID.
func (store *MemoryStore) Get(id string) (*session.Session, error) {
	value, ok := store.sessions.Load(id)

	if !ok {
		return nil, errors.New("Session not found")
	}

	return value.(*session.Session), nil
}

// Set saves a session so it can be retrieved by its ID.
func (store *MemoryStore) Set(id string, session *session.Session) error {
	store.sessions.Store(id, session)
	return nil
}

// Delete deletes the session with the given ID.
func (store *MemoryStore) Delete(id string) {
	store.sessions.Delete(id)
}
