package memstore

import (
	"errors"

	"github.com/aerogo/session"
	"golang.org/x/sync/syncmap"
)

// MemoryStore is the default session store.
// You should use it for prototyping, not for production.
type MemoryStore struct {
	sessions syncmap.Map
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
