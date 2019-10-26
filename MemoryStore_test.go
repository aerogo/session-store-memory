package memstore_test

import (
	"testing"

	"github.com/aerogo/session"
	memstore "github.com/aerogo/session-store-memory"
	"github.com/akyoto/assert"
)

func TestStore(t *testing.T) {
	store := memstore.New()
	session := session.New("1", nil)
	err := store.Set(session.ID(), session)
	assert.Nil(t, err)
	session2, err := store.Get(session.ID())
	assert.Nil(t, err)
	assert.Equal(t, session, session2)
	store.Delete(session.ID())
	session2, err = store.Get(session.ID())
	assert.NotNil(t, err)
	assert.Nil(t, session2)
}
