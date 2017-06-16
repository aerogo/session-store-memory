package benchmark

import (
	"testing"

	"github.com/aerogo/session"
	"github.com/aerogo/session-store-memory"
)

func BenchmarkMemoryStore(t *testing.B) {
	store := memstore.New()

	t.ReportAllocs()
	t.ResetTimer()

	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			store.Get("sid")
			store.Set("sid", session.New("sid", nil))
			store.Get("sid")
		}
	})
}

func BenchmarkMemoryStoreRWLock(t *testing.B) {
	store := NewMemoryStoreRWLock()

	t.ReportAllocs()
	t.ResetTimer()

	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			store.Get("sid")
			store.Set("sid", session.New("sid", nil))
			store.Get("sid")
		}
	})
}
