package darndefer

import (
	"sync"
	"testing"
)

func withDefer() {
	m := &sync.Mutex{}
	m.Lock()
	defer m.Unlock()
}

func without() {
	m := &sync.Mutex{}
	m.Lock()
	m.Unlock()
}

func withSyncFunc() {
	m := &sync.Mutex{}
	syncFunc(m, func() {})
}
func BenchmarkWithDefer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withDefer()
	}
}

func BenchmarkWithout(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		without()
	}
}

func BenchmarkSyncFunc(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withSyncFunc()
	}
}

func syncFunc(l sync.Locker, block func()) {
	l.Lock()
	block()
	l.Unlock()
}
