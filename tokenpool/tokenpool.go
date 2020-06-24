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

func (tokenPool *TokenPool) set(key int) {
	tokenPool.mutex.Lock()
	defer tokenPool.mutex.Unlock()
	tokenPool.entries[key] = time.Now()
}

// Validate check the given token key and returns whether it is valid
// and will consume it if so, expiring it instantly.
func (tokenPool *TokenPool) Validate(key int) bool {
	tokenPool.mutex.RLock()
	createdAt, ok := tokenPool.entries[key]
	if !ok {
		tokenPool.mutex.RUnlock()
		return false
	}
	valid := time.Since(createdAt) <= tokenPool.expiration
	tokenPool.mutex.RUnlock()
	tokenPool.Expire(key)
	return valid
}

// Expire ensures the token is removed so it can't be hijacked afterwards.
func (tokenPool *TokenPool) Expire(key int) {
	tokenPool.mutex.Lock()
	defer tokenPool.mutex.Unlock()
	delete(tokenPool.entries, key)
}

// RequestToken requests a new token from the stack.
func (tokenPool *TokenPool) RequestToken() (int, error) {
	key := int(tokenPool.randSource.Int63())
	tokenPool.set(key)
	return key, nil
}

// Count returns the current pool entries size.
func (tokenPool *TokenPool) Count() uint {
	tokenPool.mutex.RLock()
	defer tokenPool.mutex.RUnlock()
	return uint(len(tokenPool.entries))
}

// New returns a new instance of Pool.
// TODO :: 7 The cleanup cycler will still be too slow for high visit counts. What to do?
func New(expiration time.Duration, cleanupInterval time.Duration) *TokenPool {
	tokenPool := &TokenPool{
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
			tokenPool.mutex.Unlock()
		}
	}(tokenPool)

	return tokenPool
}
