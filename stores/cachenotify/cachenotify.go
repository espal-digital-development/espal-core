package cachenotify

import (
	"time"
)

const (
	CacheNotifyKeyDomain uint = iota + 1
	CacheNotifyKeySite
)

// TODO :: 77777 ObjectEntity interface name, check all used instances and rebuild Mocks that need to

// CacheNotify database object.
// @synthesize
type CacheNotify struct {
	id                 string
	createdByID        *string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	target             uint
	key                string
}
