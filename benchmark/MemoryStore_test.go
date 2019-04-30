package benchmark

import (
	"testing"

	"github.com/aerogo/session"
	memstore "github.com/aerogo/session-store-memory"
)

func BenchmarkMemoryStore(b *testing.B) {
	store := memstore.New()

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := store.Get("sid")

			if err != nil {
				b.Fail()
			}

			err = store.Set("sid", session.New("sid", nil))

			if err != nil {
				b.Fail()
			}

			_, err = store.Get("sid")

			if err != nil {
				b.Fail()
			}
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
