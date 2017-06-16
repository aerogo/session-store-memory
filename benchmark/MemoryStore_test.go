package benchmark

import (
	"testing"

	"github.com/aerogo/aero"
	"github.com/aerogo/session-store-memory"
)

func BenchmarkMemoryStore(t *testing.B) {
	store := memstore.NewMemoryStore()

	t.ReportAllocs()
	t.ResetTimer()

	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			store.Get("sid")
			store.Set("sid", aero.NewSession("sid", nil))
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
			store.Set("sid", aero.NewSession("sid", nil))
			store.Get("sid")
		}
	})
}
