package cachenotifymock

//go:generate moq -pkg cachenotifymock -out cachenotify_synthesized.go .. CacheNotifyEntity
//go:generate moq -pkg cachenotifymock -out store_synthesized.go .. Store
