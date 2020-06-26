package sessions_test

// TODO :: 77 Test Cookie info correctness and the FileSystemStorage logic

// func getDefault(t *testing.T) *sessions {
// 	configService, err := config.NewMock(testtools.RequestNewTempDir(t))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	storageService := NewFileSystemStorage(configService.Paths().Store() + "sessions")
// 	sessions, err := new(storageService, configService.SessionCookieName(), configService.SessionExpiration(),
//  	configService.SessionRememberMeExpiration())
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	return sessions
// }

// func TestIDExists(t *testing.T) {
// 	sessions := getDefault(t)
// 	session, _, err := sessions.NewSession()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if !sessions.IDExists(session.ID()) {
// 		t.Fatal("ID should exist")
// 	}
// }

// func TestIDNotExists(t *testing.T) {
// 	sessions := getDefault(t)

// 	if sessions.IDExists("0") {
// 		t.Fatal("ID shouldn't exist")
// 	}

// 	if sessions.IDExists("") {
// 		t.Fatal("ID shouldn't exist")
// 	}
// }

// func TestCount(t *testing.T) {
// 	sessions := getDefault(t)

// 	for i := 1; i <= 10; i++ {
// 		_, _, err := sessions.NewSession()
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		count := sessions.Count()
// 		if int(count) != i {
// 			t.Errorf("expected count %d, but got %d", i, count)
// 		}
// 	}
// }

// func TestCountNul(t *testing.T) {
// 	sessions := getDefault(t)

// 	if sessions.Count() > 0 {
// 		t.Fatal("Count should be nul")
// 	}
// }

// func TestUintValue(t *testing.T) {
// 	sessions := getDefault(t)
// 	session, _, err := sessions.NewSession()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	testValue := uint(12345)
// 	session.SetUint(SessionKeyUserID, testValue)
// 	retrievedValue, ok := session.GetUint(SessionKeyUserID)

// 	if !ok {
// 		t.Fatal("Should return a value")
// 	}
// 	if testValue != retrievedValue {
// 		t.Fatalf("Expected `%d` but got `%d`", testValue, retrievedValue)
// 	}
// }

// func TestUnexistingUintValue(t *testing.T) {
// 	sessions := getDefault(t)
// 	session, _, err := sessions.NewSession()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	retrievedValue, ok := session.GetUint(SessionKeyUserID)
// 	if ok {
// 		t.Fatal("Value shouldn't exist")
// 	}
// 	if retrievedValue > 0 {
// 		t.Fatal("Value shouldn't be set")
// 	}
// }

// func TestBytesValue(t *testing.T) {
// 	sessions := getDefault(t)
// 	session, _, err := sessions.NewSession()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	testValue := []byte("12345")
// 	session.SetBytes(SessionKeyUserID, testValue)
// 	retrievedValue, ok := session.GetBytes(SessionKeyUserID)

// 	if !ok {
// 		t.Fatal("Should return a value")
// 	}
// 	if !bytes.Equal(testValue, retrievedValue) {
// 		t.Fatalf("Expected `%s` but got `%s`", string(testValue), string(retrievedValue))
// 	}
// }

// func TestUnexistingBytesValue(t *testing.T) {
// 	sessions := getDefault(t)
// 	session, _, err := sessions.NewSession()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	retrievedValue, ok := session.GetBytes(SessionKeyUserID)
// 	if ok {
// 		t.Fatal("Value shouldn't exist")
// 	}
// 	if len(retrievedValue) > 0 {
// 		t.Fatal("Value shouldn't be set")
// 	}
// }

// func TestUnexistingSessID(t *testing.T) {
// 	sessions := getDefault(t)
// 	fakeSessID := "12345"

// 	session := sessions.GetSession(fakeSessID)
// 	if session != nil {
// 		t.Fatal("Session shouldn't exist")
// 	}
// }

// func TestUintValueIDUnexisting(t *testing.T) {
// 	sessions := getDefault(t)
// 	session, _, err := sessions.NewSession()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	_, ok := session.GetUint(SessionKeyUserID)
// 	if ok {
// 		t.Fatal("Should not exist")
// 	}
// }

// func TestBytesValueIDUnexisting(t *testing.T) {
// 	sessions := getDefault(t)
// 	session, _, err := sessions.NewSession()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	_, ok := session.GetBytes(SessionKeyUserID)
// 	if ok {
// 		t.Fatal("Should not exist")
// 	}
// }
