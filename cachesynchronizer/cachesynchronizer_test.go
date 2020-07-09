package cachesynchronizer_test

// TODO :: 777 Should test multi-broadcasts too to be actually sure that the functionality works solidly

// var (
// 	defaultNotifyValue    = "4"
// 	checkInterval         = time.Millisecond * 25
// 	loggerService         *loggermock.LoggerMock
// 	cacheNotifyStore      *cachenotifymock.StoreMock
// 	cacheNotifyEntityMock *cachenotifymock.CacheNotifyEntityMock
// 	// allValues          = []interface{}{"fake-id-1", nil, nil, time.Now(), nil, cachenotify.CacheNotifyKeyDomain, "4"}
// )

// func initMocks(t *testing.T) {
// 	loggerService = &loggermock.LoggerMock{}
// 	cacheNotifyEntityMock = &cachenotifymock.CacheNotifyEntityMock{
// 		KeyFunc: func() string {
// 			return defaultNotifyValue
// 		},
// 	}
// 	cacheNotifyStore = &cachenotifymock.StoreMock{
// 		GetLatestFunc: func(interval time.Duration) ([]*cachenotify.CacheNotify, bool, error) {
// 			// TODO :: 77777 This needs a good global solution; returning a literal struct as a mock
// 			stub := []*cachenotify.CacheNotify{
// 				{},
// 			}
// 			return stub, true, nil
// 		},
// 		SaveFunc: func(target uint, key string) error {
// 			return nil
// 		},
// 	}
// }

// func TestSubscribe(t *testing.T) {
// 	initMocks(t)
// 	cacheSynchronizer := cachesynchronizer.New(loggerService, cacheNotifyStore, checkInterval)
// 	subscription, err := cacheSynchronizer.Subscribe(cachenotify.CacheNotifyKeyDomain)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if subscription.ID() == 0 {
// 		t.Fatal("ID should not be 0")
// 	}
// 	if cachenotify.CacheNotifyKeyDomain != subscription.Target() {
// 		t.Fatalf("Target `%d` isn't equal to the one in the subscription `%d`",
//  	subscription.Target(), cachenotify.CacheNotifyKeyDomain)
// 	}

// 	time.Sleep(time.Millisecond * 5)

// 	var receivedValue bool
// 	var value string

// 	select {
// 	case x, ok := <-subscription.Items():
// 		if ok {
// 			receivedValue = true
// 			value = x
// 		}
// 	default:
// 	}

// 	if !receivedValue {
// 		t.Fatal("There should be at least an item in the subscription")
// 	}
// 	if value != defaultNotifyValue {
// 		t.Fatalf("Expected returned value to be `%s` but got `%s`", defaultNotifyValue, value)
// 	}
// }

// func TestSubscribeNoItems(t *testing.T) {
// 	initMocks(t)
// 	cacheNotifyStore.GetLatestFunc = func(interval time.Duration) ([]cachenotify.CacheNotify, bool, error) {
// 		return nil, false, nil
// 	}
// 	cacheSynchronizer := cachesynchronizer.New(loggerService, cacheNotifyStore, checkInterval)
// 	subscription, err := cacheSynchronizer.Subscribe(cachenotify.CacheNotifyKeyDomain)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	time.Sleep(time.Millisecond * 5)

// 	var receivedValue bool

// 	select {
// 	case _, ok := <-subscription.Items():
// 		if ok {
// 			receivedValue = true
// 		}
// 	default:
// 	}

// 	if receivedValue {
// 		t.Fatal("There should be no items in the subscription yet")
// 	}
// }

// func TestUnsubscribe(t *testing.T) {
// 	initMocks(t)
// 	cacheSynchronizer := cachesynchronizer.New(loggerService, cacheNotifyStore, checkInterval)
// 	subscription, err := cacheSynchronizer.Subscribe(cachenotify.CacheNotifyKeyDomain)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	time.Sleep(time.Millisecond * 5)

// 	cacheSynchronizer.Unsubscribe(subscription)
// }

// func TestNotify(t *testing.T) {
// 	initMocks(t)
// 	returnValues := true
// 	newNotifyValue := "5"
// 	cacheNotifyStore.GetLatestFunc = func(interval time.Duration) ([]cachenotify.CacheNotify, bool, error) {
// 		if !returnValues {
// 			return nil, false, nil
// 		}
// 		return []cachenotify.CacheNotify{newCacheNotify(defaultNotifyValue), newCacheNotify(newNotifyValue)}, true, nil
// 	}
// 	cacheSynchronizer := cachesynchronizer.New(loggerService, cacheNotifyStore, checkInterval)
// 	subscription, err := cacheSynchronizer.Subscribe(cachenotify.CacheNotifyKeyDomain)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if err := cacheSynchronizer.Notify(cachenotify.CacheNotifyKeyDomain, newNotifyValue); err != nil {
// 		t.Fatal(err)
// 	}

// 	time.Sleep(time.Millisecond * 5)

// 	returnValues = false

// 	items := make([]string, 0)

// 	go func() {
// 		for {
// 			select {
// 			case item, ok := <-subscription.Items():
// 				if ok {
// 					items = append(items, item)
// 				}
// 			}
// 		}
// 	}()

// 	time.Sleep(time.Millisecond * 5)

// 	if len(items) != 2 {
// 		t.Fatalf("2 values should've been returned, only got `%d`", len(items))
// 	}
// }

// func TestGetLatestError(t *testing.T) {
// 	initMocks(t)
// 	var loggedMessage string
// 	latestError := errors.New("fake error")
// 	cacheNotifyStore.GetLatestFunc = func(interval time.Duration) ([]cachenotify.CacheNotify, bool, error) {
// 		return nil, false, latestError
// 	}
// 	loggerService.ErrorFunc = func(message string) {
// 		loggedMessage = message
// 	}
// 	cacheSynchronizer := cachesynchronizer.New(loggerService, cacheNotifyStore, checkInterval)
// 	subscription, err := cacheSynchronizer.Subscribe(cachenotify.CacheNotifyKeyDomain)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	time.Sleep(time.Millisecond * 5)

// 	var receivedValue bool

// 	select {
// 	case _, ok := <-subscription.Items():
// 		if ok {
// 			receivedValue = true
// 		}
// 	default:
// 	}

// 	if receivedValue {
// 		t.Fatal("There should be no items in the subscription because an error should've prevented propegation")
// 	}

// 	if loggedMessage != latestError.Error() {
// 		t.Fatalf("expected error message `%s` to be thrown, but got `%s`", latestError.Error(), loggedMessage)
// 	}
// }

// func TestSomeCyclesAndStop(t *testing.T) {
// 	initMocks(t)
// 	cacheSynchronizer := cachesynchronizer.New(loggerService, cacheNotifyStore, checkInterval)
// 	subscription, err := cacheSynchronizer.Subscribe(cachenotify.CacheNotifyKeyDomain)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	items := make([]string, 0)

// 	go func() {
// 		for {
// 			select {
// 			case item, ok := <-subscription.Items():
// 				if ok {
// 					items = append(items, item)
// 				}
// 			}
// 		}
// 	}()

// 	time.Sleep(time.Millisecond * 80)

// 	cacheSynchronizer.Unsubscribe(subscription)

// 	time.Sleep(time.Millisecond * 60)
// }
