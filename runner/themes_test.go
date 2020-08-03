package runner_test

import (
	"sync"
	"testing"

	"github.com/espal-digital-development/espal-core/text"
)

// var data = []string{
// 	"asdasdas",
// 	"a0s89hdaisd",
// 	"oasdbk",
// 	"noausidboa0shudn",
// 	"a0us9dbi",
// 	"9yas8dyashd",
// 	"aoshidgasud",
// 	"ybat7sdcyasvjdb",
// }

var defaultValue = "asdsadsa"

func BenchmarkMutexLock(b *testing.B) {
	list := map[string]string{}
	m := &sync.RWMutex{}

	for i := 0; i < b.N; i++ {
		m.Lock()
		if _, ok := list[defaultValue]; ok {
			m.Unlock()
			continue
		}
		list[defaultValue] = defaultValue
		m.Unlock()
	}
}

func BenchmarkMutexLockBigList(b *testing.B) {
	list := map[string]string{}
	for i := 0; i < 1e5; i++ {
		v := text.RandomString(72)
		list[v] = v
	}
	m := &sync.RWMutex{}

	for i := 0; i < b.N; i++ {
		m.Lock()
		if _, ok := list[defaultValue]; ok {
			m.Unlock()
			continue
		}
		list[defaultValue] = defaultValue
		m.Unlock()
	}
}

func BenchmarkMutexRLockLock(b *testing.B) {
	list := map[string]string{}
	m := &sync.RWMutex{}

	for i := 0; i < b.N; i++ {
		m.RLock()
		if _, ok := list[defaultValue]; ok {
			m.RUnlock()
			continue
		}
		m.RUnlock()
		m.Lock()
		list[defaultValue] = defaultValue
		m.Unlock()
	}
}

func BenchmarkMutexRLockLockBigList(b *testing.B) {
	list := map[string]string{}
	for i := 0; i < 1e5; i++ {
		v := text.RandomString(72)
		list[v] = v
	}
	m := &sync.RWMutex{}

	for i := 0; i < b.N; i++ {
		m.RLock()
		if _, ok := list[defaultValue]; ok {
			m.RUnlock()
			continue
		}
		m.RUnlock()
		m.Lock()
		list[defaultValue] = defaultValue
		m.Unlock()
	}
}
