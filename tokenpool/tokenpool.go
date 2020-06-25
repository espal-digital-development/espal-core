package tokenpool

import (
	"math/rand"
	"sync"
	"time"
)

var _ Pool = &TokenPool{}

// Pool represents an object that manages
// token-style instances.
type Pool interface {
	Validate(key int) bool
	Expire(key int)
	RequestToken() (int, error)
	Count() uint
}

// Pool holds all tokens and handles cycling.
type TokenPool struct {
	entries         map[int]time.Time
	expiration      time.Duration
	cleanupInterval time.Duration
	mutex           *sync.RWMutex
	randSource      rand.Source
}

func (p *TokenPool) set(key int) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.entries[key] = time.Now()
}

// Validate check the given token key and returns whether it is valid
// and will consume it if so, expiring it instantly.
func (p *TokenPool) Validate(key int) bool {
	p.mutex.RLock()
	createdAt, ok := p.entries[key]
	if !ok {
		p.mutex.RUnlock()
		return false
	}
	valid := time.Since(createdAt) <= p.expiration
	p.mutex.RUnlock()
	p.Expire(key)
	return valid
}

// Expire ensures the token is removed so it can't be hijacked afterwards.
func (p *TokenPool) Expire(key int) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	delete(p.entries, key)
}

// RequestToken requests a new token from the stack.
func (p *TokenPool) RequestToken() (int, error) {
	key := int(p.randSource.Int63())
	p.set(key)
	return key, nil
}

// Count returns the current pool entries size.
func (p *TokenPool) Count() uint {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return uint(len(p.entries))
}

// New returns a new instance of Pool.
// TODO :: 7 The cleanup cycler will still be too slow for high visit counts. What to do?
func New(expiration time.Duration, cleanupInterval time.Duration) *TokenPool {
	p := &TokenPool{
		entries:         map[int]time.Time{},
		mutex:           &sync.RWMutex{},
		expiration:      expiration,
		cleanupInterval: cleanupInterval,
		randSource:      rand.NewSource(time.Now().UnixNano()),
	}

	go func(pool *TokenPool) {
		for pool != nil {
			time.Sleep(pool.cleanupInterval)
			pool.mutex.Lock()
			for k := range pool.entries {
				if time.Since(pool.entries[k]) > pool.expiration {
					delete(pool.entries, k)
				}
			}
			p.mutex.Unlock()
		}
	}(p)

	return p
}
