package notifier

import (
	"sync"
	"time"

	"github.com/espal-digital-development/espal-core/logger"
	"github.com/espal-digital-development/espal-core/stores/notification"
	"github.com/juju/errors"
)

var _ Hub = &Notifier{}

const (
	defaultTargetChannelSize = 25
	defaultCycleTime         = time.Second * 6
)

// TODO :: 777777 Can cleanup older entries every X-period
// TODO :: 777777 Rate limit target additions to prevent exploits for Denial of Service weakness
// TODO :: 777777
// - Cluster awareness via database:
// - s.Notify() to specific cluster entry (in case of rolling upgrade, go offline, back online)
// - Upgrade, but if DB changed, any unupdated node would be instantly invalid?
// - Options to interact with Network-layer banlists (Cloudflare etc.)

// Hub represents an object that is responsible for registering updates so other cluster instances can be notified of
// changes being done in different geo-located clusters.
// The most obvious use-cases will be simple cache-invalidations, but more complex instructions could be passed through
// the key-information too.
type Hub interface {
	Notify(target string, key string) error
	NotifyWithValue(target string, key string, value string) error
	ListenToTarget(target string) (string, string)
}

type data struct {
	key   string
	value string
}

// Notifier service object.
type Notifier struct {
	loggerService     logger.Loggable
	notificationStore notification.Store

	targets map[string]chan *data
	mutex   *sync.RWMutex

	selfProducedIDs      map[string]bool
	selfProducedIDsMutex *sync.RWMutex
}

// Notify pushes the key onto the target's notification list.
func (n *Notifier) Notify(target string, key string) error {
	return errors.Trace(n.NotifyWithValue(target, key, ""))
}

// NotifyWithValue pushes the key/value combination onto the target's notification list.
func (n *Notifier) NotifyWithValue(target string, key string, value string) error {
	id, err := n.notificationStore.Save(target, key, value)
	if err != nil {
		return errors.Trace(err)
	}
	n.selfProducedIDsMutex.Lock()
	n.selfProducedIDs[id] = true
	n.selfProducedIDsMutex.Unlock()
	return nil
}

// ListenToTarget will listen blocking until the target feed returns a new value.
func (n *Notifier) ListenToTarget(target string) (string, string) {
	var channel chan *data
	n.mutex.RLock()
	if _, ok := n.targets[target]; !ok {
		n.mutex.RUnlock()
		n.mutex.Lock()
		channel = make(chan *data, defaultTargetChannelSize)
		n.targets[target] = channel
		n.mutex.Unlock()
	} else {
		channel = n.targets[target]
		n.mutex.RUnlock()
	}
	data := <-channel
	return data.key, data.value
}

func (n *Notifier) listen() {
	// TODO :: 777777 Implement cycler and pinger
	// n.mutex.Lock()
	// defer n.mutex.Unlock()
	// if _, ok := n.targets[target]; !ok {
	// 	n.targets[target] = make(chan *data, defaultTargetChannelSize)
	// }
	// n.targets[target] <- &data{
	// 	key: key,
	// }
	for {
		time.Sleep(defaultCycleTime)

	}
}

// New returns a new instance of s.
func New(loggerService logger.Loggable, notificationStore notification.Store) *Notifier {
	n := &Notifier{
		loggerService:     loggerService,
		notificationStore: notificationStore,

		targets:              map[string]chan *data{},
		selfProducedIDs:      map[string]bool{},
		mutex:                &sync.RWMutex{},
		selfProducedIDsMutex: &sync.RWMutex{},
	}

	n.listen()

	return n
}
