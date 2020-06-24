package cachesynchronizer

import (
	"sync"
	"time"

	"github.com/espal-digital-development/espal-core/logger"
	"github.com/espal-digital-development/espal-core/stores/cachenotify"
	"github.com/juju/errors"
)

var _ Synchronizer = &CacheSynchronizer{}

// TODO :: 777 Can cleanup older entries every X-period

// Synchronizer represents an object that is responsible
// for registering cache updates so other cluster instances
// can be notified of changes being done in different
// geo-located clusters.
// The most obvious use-cases will be simple cache-
// invalidations, but more complex instructions
// could be passed through the key-information too.
type Synchronizer interface {
	Notify(target uint, key string) error
	Subscribe(target uint) (Subscription, error)
	Unsubscribe(subscription Subscription)
}

// CacheSynchronizer service object.
type CacheSynchronizer struct {
	loggerService    logger.Loggable
	cacheNotifyStore cachenotify.Store
	subscribers      map[uint][]*subscription
	cyclers          map[uint]bool
	subscribersMutex *sync.RWMutex
	cyclersMutex     *sync.RWMutex
	checkInterval    time.Duration
	idCounter        uint
}

// Subscription represents listening binding to
// a synchronizer.
type Subscription interface {
	ID() uint
	Target() uint
	Items() chan string
}

type subscription struct {
	id     uint
	target uint
	items  chan string
}

// ID returns the unique subscription ID.
func (subscription *subscription) ID() uint {
	return subscription.id
}

// Target returns the subscription target descriptor.
func (subscription *subscription) Target() uint {
	return subscription.target
}

// Items returns the subscription items channel to list to.
func (subscription *subscription) Items() chan string {
	return subscription.items
}

// Notify broadcasts a key changed for the given target.
func (cacheSynchronizer *CacheSynchronizer) Notify(target uint, key string) error {
	return cacheSynchronizer.cacheNotifyStore.Save(target, key)
}

// Subscribe returns a channel to listen to changes for the given target.
func (cacheSynchronizer *CacheSynchronizer) Subscribe(target uint) (Subscription, error) {
	cacheSynchronizer.subscribersMutex.Lock()
	defer cacheSynchronizer.subscribersMutex.Unlock()
	cacheSynchronizer.idCounter++
	sub := &subscription{
		id:     cacheSynchronizer.idCounter,
		items:  make(chan string),
		target: target,
	}
	_, ok := cacheSynchronizer.subscribers[target]
	if !ok {
		cacheSynchronizer.subscribers[target] = make([]*subscription, 0)
	}
	cacheSynchronizer.subscribers[target] = append(cacheSynchronizer.subscribers[target], sub)
	if !ok {
		cacheSynchronizer.startCycler(target)
	}
	return sub, nil
}

// Unsubscribe revokes a given subscription.
func (cacheSynchronizer *CacheSynchronizer) Unsubscribe(subscription Subscription) {
	target := subscription.Target()
	cacheSynchronizer.subscribersMutex.Lock()
	defer cacheSynchronizer.subscribersMutex.Unlock()
	var subscriptionsRemaining uint
	for k := range cacheSynchronizer.subscribers[target] {
		if cacheSynchronizer.subscribers[target][k] != nil {
			subscriptionsRemaining++
		}
	}
	if subscriptionsRemaining <= 1 {
		cacheSynchronizer.cyclersMutex.Lock()
		defer cacheSynchronizer.cyclersMutex.Unlock()
		delete(cacheSynchronizer.cyclers, target)
		delete(cacheSynchronizer.subscribers, target)
	}
}

func (cacheSynchronizer *CacheSynchronizer) startCycler(target uint) {
	cacheSynchronizer.cyclersMutex.Lock()
	if _, ok := cacheSynchronizer.cyclers[target]; !ok {
		cacheSynchronizer.cyclers[target] = true
		defer cacheSynchronizer.cyclersMutex.Unlock()
		go cacheSynchronizer.cycler(target)
	}
}

func (cacheSynchronizer *CacheSynchronizer) cycler(target uint) {
	for {
		if !cacheSynchronizer.cycle(target) {
			break
		}
		time.Sleep(cacheSynchronizer.checkInterval)
	}
}

func (cacheSynchronizer *CacheSynchronizer) cycle(target uint) bool {
	// TODO :: 77 This package started off well, but is a bit of a mess now. Needs a revision in the future
	// cacheSynchronizer.subscribersMutex.Lock()
	// defer cacheSynchronizer.subscribersMutex.Unlock()

	if _, ok := cacheSynchronizer.subscribers[target]; !ok {
		return false
	}
	cacheNotifications, ok, err := cacheSynchronizer.cacheNotifyStore.GetLatest(cacheSynchronizer.checkInterval)
	if err != nil {
		cacheSynchronizer.loggerService.Error(errors.ErrorStack(err))
		return false
	}
	if ok {
		for i := range cacheNotifications {
			for k2 := range cacheSynchronizer.subscribers[target] {
				// if cacheSynchronizer.subscribers[target][k2] == nil {
				// 	return false
				// }
				cacheSynchronizer.subscribers[target][k2].Items() <- cacheNotifications[i].Key()
			}
		}
	}
	return true
}

// New returns a new instance of CacheSynchronizer.
func New(loggerService logger.Loggable, cacheNotifyStore cachenotify.Store, checkInterval time.Duration) *CacheSynchronizer {
	return &CacheSynchronizer{
		loggerService:    loggerService,
		cacheNotifyStore: cacheNotifyStore,
		subscribers:      map[uint][]*subscription{},
		cyclers:          map[uint]bool{},
		subscribersMutex: &sync.RWMutex{},
		cyclersMutex:     &sync.RWMutex{},
		checkInterval:    checkInterval,
	}
}
