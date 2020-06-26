package slug_test

// var (
// 	fakeUserID   = "fake-user-id-1"
// 	fakeDomainID = "fake-domain-id-1"
// 	fakeSlugID   = "fake-slug-id-1"
// 	fakeSlug     = "fakeSlugPath"
// 	allFields    = []string{
// 		"id", "createdByID", "updatedByID", "domainID",
// 		"createdAt", "updatedAt", "language", "path", "rerouteTo", "invalidWithStatus",
// 		"invalidMessage", "redirectToRawPath", "redirectStatusCode",
// 		// "createdByFirstName", "createdBySurname", "updatedByFirstName", "updatedBySurname",
// 	}
// 	allValues = []interface{}{
// 		"fake-domain-id-1", "fake-user-id-1", nil, "fake-domain-id-1",
// 		time.Now(), nil, 1, "PATH", "REROUTETO", nil, nil, nil, nil,
// 		// nil, nil, nil, nil,
// 	}
// 	errMock = errors.Errorf("mock SQL Error")
// )

// func getDefault(t *testing.T) *store {
// 	configService, err := config.NewMock(testtools.RequestNewTempDir(t))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	loggerService := logger.New()
// 	databaseService, err := database.NewMock(configService, loggerService)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	return new(databaseService)
// }

// func TestNew(t *testing.T) {
// 	getDefault(t)
// }

// TODO :: 777 Why does this fail? `arguments do not match: expected 1,
// but got 2 arguments` but there are $1 and $2 (2)..
// func TestGetOneByDomainIDAndPath(t *testing.T) {
// 	store := getDefault(t)
// 	rows := sqlmock.NewRows(allFields).AddRow(allValues...)
// 	mock := store.databaseService.Mocks().SelecterDBMock()
// 	mock.ExpectQuery("SELECT").WithArgs(fakeSlugID).WillReturnRows(rows)

// 	slug, found, err := store.GetOneByDomainIDAndPath(fakeDomainID, fakeSlug)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !found {
// 		t.Error("result should've been found")
// 	}
// 	if slug == nil {
// 		t.Error("result should not be nil")
// 	}
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }
