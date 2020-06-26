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
func (s *subscription) ID() uint {
	return s.id
}

// Target returns the subscription target descriptor.
func (s *subscription) Target() uint {
	return s.target
}

// Items returns the subscription items channel to list to.
func (s *subscription) Items() chan string {
	return s.items
}

// Notify broadcasts a key changed for the given target.
func (s *CacheSynchronizer) Notify(target uint, key string) error {
	return s.cacheNotifyStore.Save(target, key)
}

// Subscribe returns a channel to listen to changes for the given target.
func (s *CacheSynchronizer) Subscribe(target uint) (Subscription, error) {
	s.subscribersMutex.Lock()
	defer s.subscribersMutex.Unlock()
	s.idCounter++
	sub := &subscription{
		id:     s.idCounter,
		items:  make(chan string),
		target: target,
	}
	_, ok := s.subscribers[target]
	if !ok {
		s.subscribers[target] = make([]*subscription, 0)
	}
	s.subscribers[target] = append(s.subscribers[target], sub)
	if !ok {
		s.startCycler(target)
	}
	return sub, nil
}

// Unsubscribe revokes a given subscription.
func (s *CacheSynchronizer) Unsubscribe(subscription Subscription) {
	target := subscription.Target()
	s.subscribersMutex.Lock()
	defer s.subscribersMutex.Unlock()
	var subscriptionsRemaining uint
	for k := range s.subscribers[target] {
		if s.subscribers[target][k] != nil {
			subscriptionsRemaining++
		}
	}
	if subscriptionsRemaining <= 1 {
		s.cyclersMutex.Lock()
		defer s.cyclersMutex.Unlock()
		delete(s.cyclers, target)
		delete(s.subscribers, target)
	}
}

func (s *CacheSynchronizer) startCycler(target uint) {
	s.cyclersMutex.Lock()
	if _, ok := s.cyclers[target]; !ok {
		s.cyclers[target] = true
		defer s.cyclersMutex.Unlock()
		go s.cycler(target)
	}
}

func (s *CacheSynchronizer) cycler(target uint) {
	for {
		if !s.cycle(target) {
			break
		}
		time.Sleep(s.checkInterval)
	}
}

func (s *CacheSynchronizer) cycle(target uint) bool {
	// TODO :: 77 This package started off well, but is a bit of a mess now. Needs a revision in the future
	// s.subscribersMutex.Lock()
	// defer s.subscribersMutex.Unlock()

	if _, ok := s.subscribers[target]; !ok {
		return false
	}
	cacheNotifications, ok, err := s.cacheNotifyStore.GetLatest(s.checkInterval)
	if err != nil {
		s.loggerService.Error(errors.ErrorStack(err))
		return false
	}
	if ok {
		for i := range cacheNotifications {
			for k2 := range s.subscribers[target] {
				// if s.subscribers[target][k2] == nil {
				// 	return false
				// }
				s.subscribers[target][k2].Items() <- cacheNotifications[i].Key()
			}
		}
	}
	return true
}

// New returns a new instance of s.
func New(loggerService logger.Loggable, cacheNotifyStore cachenotify.Store,
	checkInterval time.Duration) *CacheSynchronizer {
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
