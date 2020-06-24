package user_test

// TODO :: 777 Implement again

// var (
// 	abcPassword = "$2a$12$RONtFJDTnoxlTHlUNQt0TO7xssRilt3PJN08ZKWfghfR05T3bJchy"
// 	fakeUserID  = "fake-user-id-1"
// 	allValues   = []interface{}{
// 		"fake-user-id-1", "fake-user-id-1", nil, nil,
// 		nil, time.Now(), nil, 1, nil, 388,
// 		"John", "Doe", time.Now(), "johndoe@example.com", abcPassword, nil, 0,
// 		nil, nil, nil, nil,
// 		nil, nil, nil, nil, "33",
// 	}

// 	dbRows                 *databasemock.RowsMock
// 	selecterDatabase       *databasemock.DatabaseMock
// 	filtersFactory         *filtersmock.FactoryMock
// 	filter                 *filtersmock.FilterMock
// 	translationsRepository *translationsmock.TranslationsMock
// 	userRightsRepository   *userrightsmock.RepositoryMock
// )

// func initMocks(t *testing.T) {
// 	dbRows = databasemock.DefaultRowsMock()
// 	selecterDatabase = &databasemock.DatabaseMock{
// 		QueryFunc: func(query string, args ...interface{}) (database.Rows, error) {
// 			return dbRows, nil
// 		},
// 	}
// 	filter = &filtersmock.FilterMock{}
// 	filtersFactory = &filtersmock.FactoryMock{
// 		NewFilterFunc: func(queryReader filters.QueryReader, m filters.Model) filters.Filter {
// 			return filter
// 		},
// 	}
// 	translationsRepository = &translationsmock.TranslationsMock{}
// 	userRightsRepository = &userrightsmock.RepositoryMock{}
// }

// func TestNew(t *testing.T) {
// 	getDefault(t)
// }

// func TestGetOne(t *testing.T) {
// 	store := getDefault(t)
// 	rows := sqlmock.NewRows(allFields).AddRow(allValues...)
// 	mock := store.databaseService.Mocks().SelecterDBMock()
// 	mock.ExpectQuery("SELECT").WithArgs(fakeUserID).WillReturnRows(rows)

// 	user, found, err := store.GetOne(fakeUserID)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !found {
// 		t.Error("result should've been found")
// 	}
// 	if user == nil {
// 		t.Error("result should not be nil")
// 	}
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestGetOneWithNoResult(t *testing.T) {
// 	store := getDefault(t)
// 	mock := store.databaseService.Mocks().SelecterDBMock()
// 	mock.ExpectQuery("SELECT").WillReturnError(sql.ErrNoRows)

// 	user, found, err := store.GetOne("fake-user-id-2")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if found {
// 		t.Error("result should've not been found")
// 	}
// 	if user != nil {
// 		t.Error("result should be nil")
// 	}
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestGetOneWithError(t *testing.T) {
// 	store := getDefault(t)
// 	mock := store.databaseService.Mocks().SelecterDBMock()
// 	mock.ExpectQuery("SELECT").WithArgs(fakeUserID).WillReturnError(errMock)

// 	_, _, err := store.GetOne(fakeUserID)
// 	if err == nil{
// 		t.Fatal("Expected an error to be returned")
// 	}
// 	if err.Error() != errMock.Error() {
// 		t.Fatal("Error was not the same as expected")
// 	}
// }

// func TestGetOneActive(t *testing.T) {
// 	store := getDefault(t)
// 	rows := sqlmock.NewRows(allFields).AddRow(allValues...)
// 	mock := store.databaseService.Mocks().SelecterDBMock()
// 	mock.ExpectQuery("SELECT").WithArgs(fakeUserID).WillReturnRows(rows)

// 	user, found, err := store.GetOneActive(fakeUserID)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !found {
// 		t.Error("result should've been found")
// 	}
// 	if user == nil {
// 		t.Error("result should not be nil")
// 	}
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestGetOneActiveWithNoResult(t *testing.T) {
// 	store := getDefault(t)
// 	mock := store.databaseService.Mocks().SelecterDBMock()
// 	mock.ExpectQuery("SELECT").WillReturnError(sql.ErrNoRows)

// 	user, found, err := store.GetOneActive("fake-user-id-2")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if found {
// 		t.Error("result should've not been found")
// 	}
// 	if user != nil {
// 		t.Error("result should be nil")
// 	}
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestGetOneActiveWithError(t *testing.T) {
// 	store := getDefault(t)
// 	mock := store.databaseService.Mocks().SelecterDBMock()
// 	mock.ExpectQuery("SELECT").WithArgs(fakeUserID).WillReturnError(errMock)

// 	_, _, err := store.GetOneActive(fakeUserID)
// 	if err == nil{
// 		t.Fatal("Expected an error to be returned")
// 	}
// 	if err.Error() != errMock.Error() {
// 		t.Fatal("Error was not the same as expected")
// 	}
// }

// func TestGetOneByIDWithCreator(t *testing.T) {
// 	store := getDefault(t)

// 	fields := allFields
// 	extraFields := []string{"firstName", "surname", "firstName", "surname"}
// 	for k := range extraFields {
// 		fields = append(fields, extraFields[k])
// 	}
// 	values := allValues
// 	extraValues := []interface{}{"John", "Doe", nil, nil}
// 	for k := range extraValues {
// 		values = append(values, extraValues[k])
// 	}
// 	rows := sqlmock.NewRows(fields).AddRow(values...)
// 	mock := store.databaseService.Mocks().SelecterDBMock()
// 	mock.ExpectQuery("SELECT").WithArgs(fakeUserID).WillReturnRows(rows)

// 	user, found, err := store.GetOneByIDWithCreator(fakeUserID)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !found {
// 		t.Error("result should've been found")
// 	}
// 	if user == nil {
// 		t.Fatal("Result should not be nil")
// 	}
// 	if user.CreatedByFirstName() == nil {
// 		t.Fatal("Expected CreatedByFirstName to not be nil")
// 	}
// 	if "John" != *user.CreatedByFirstName() {
// 		t.Fatalf("Expected CreatedBy FirstName to be `%s`, but got `%s`", "John", *user.CreatedByFirstName())
// 	}
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestGetOneByIDWithCreatorWithNoResult(t *testing.T) {
// 	store := getDefault(t)
// 	mock := store.databaseService.Mocks().SelecterDBMock()
// 	mock.ExpectQuery("SELECT").WillReturnError(sql.ErrNoRows)

// 	user, found, err := store.GetOneByIDWithCreator("fake-user-id-2")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if found {
// 		t.Error("result should've not been found")
// 	}
// 	if user != nil {
// 		t.Error("result should be nil")
// 	}
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestGetOneByIDWithCreatorWithError(t *testing.T) {
// 	store := getDefault(t)
// 	mock := store.databaseService.Mocks().SelecterDBMock()
// 	mock.ExpectQuery("SELECT").WithArgs(fakeUserID).WillReturnError(errMock)

// 	_, _, err := store.GetOneByIDWithCreator(fakeUserID)
// 	if err == nil{
// 		t.Fatal("Expected an error to be returned")
// 	}
// 	if err.Error() != errMock.Error() {
// 		t.Fatal("Error was not the same as expected")
// 	}
// }

// func TestGetOneActiveByEmail(t *testing.T) {
// 	store := getDefault(t)
// 	rows := sqlmock.NewRows(allFields).AddRow(allValues...)
// 	mock := store.databaseService.Mocks().SelecterDBMock()
// 	mock.ExpectQuery("SELECT").WithArgs("johndoe@example.com").WillReturnRows(rows)

// 	user, found, err := store.GetOneActiveByEmail("johndoe@example.com")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !found {
// 		t.Error("result should've been found")
// 	}
// 	if user == nil {
// 		t.Error("result should not be nil")
// 	}
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestGetOneActiveByEmailWithNoResult(t *testing.T) {
// 	store := getDefault(t)
// 	mock := store.databaseService.Mocks().SelecterDBMock()
// 	mock.ExpectQuery("SELECT").WillReturnError(sql.ErrNoRows)

// 	user, found, err := store.GetOneActiveByEmail("johndoe@example.com")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if found {
// 		t.Error("result should've not been found")
// 	}
// 	if user != nil {
// 		t.Error("result should be nil")
// 	}
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestGetOneActiveByEmailWithError(t *testing.T) {
// 	store := getDefault(t)
// 	mock := store.databaseService.Mocks().SelecterDBMock()
// 	mock.ExpectQuery("SELECT").WithArgs("johndoe@example.com").WillReturnError(errMock)

// 	_, _, err := store.GetOneActiveByEmail("johndoe@example.com")
// 	if err == nil{
// 		t.Fatal("Expected an error to be returned")
// 	}
// 	if err.Error() != errMock.Error() {
// 		t.Fatal("Error was not the same as expected")
// 	}
// }

// func TestGetOneIDAndPasswordForActiveByEmail(t *testing.T) {
// 	store := getDefault(t)
// 	fields := []string{"id", "password"}
// 	values := []interface{1, abcPassword}
// 	rows := sqlmock.NewRows(fields).AddRow(values...)
// 	mock := store.databaseService.Mocks().SelecterDBMock()
// 	mock.ExpectQuery("SELECT").WithArgs("johndoe@example.com").WillReturnRows(rows)

// 	user, found, err := store.GetOneIDAndPasswordForActiveByEmail("johndoe@example.com")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !found {
// 		t.Error("result should've been found")
// 	}
// 	if user == nil {
// 		t.Error("result should not be nil")
// 	}
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestGetOneIDAndPasswordForActiveByEmailWithNoResult(t *testing.T) {
// 	store := getDefault(t)
// 	mock := store.databaseService.Mocks().SelecterDBMock()
// 	mock.ExpectQuery("SELECT").WillReturnError(sql.ErrNoRows)

// 	user, found, err := store.GetOneIDAndPasswordForActiveByEmail("johndoe@example.com")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if found {
// 		t.Error("result should've not been found")
// 	}
// 	if user != nil {
// 		t.Error("result should be nil")
// 	}
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestGetOneIDAndPasswordForActiveByEmailWithError(t *testing.T) {
// 	store := getDefault(t)
// 	mock := store.databaseService.Mocks().SelecterDBMock()
// 	mock.ExpectQuery("SELECT").WithArgs("johndoe@example.com").WillReturnError(errMock)

// 	_, _, err := store.GetOneIDAndPasswordForActiveByEmail("johndoe@example.com")
// 	if err == nil{
// 		t.Fatal("Expected an error to be returned")
// 	}
// 	if err.Error() != errMock.Error() {
// 		t.Fatal("Error was not the same as expected")
// 	}
// }

// func TestHasUserRight(t *testing.T) {
// 	store := getDefault(t)
// 	rows := sqlmock.NewRows([]string{"1"}).AddRow("1")
// 	mock := store.databaseService.Mocks().SelecterDBMock()
// 	mock.ExpectQuery("SELECT").WithArgs(30000, fakeUserID).WillReturnRows(rows)

// 	user := newUser()
// 	user.id = fakeUserID
// 	user.createdByID = fakeUserID
// 	user.createdAt = time.Now()

// 	hasUserRight, err := store.HasUserRight(user, "AccessAuth")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !hasUserRight {
// 		t.Fatal("Expecting user to have the tested UserRight")
// 	}
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestHasUserRightError(t *testing.T) {
// 	store := getDefault(t)
// 	mock := store.databaseService.Mocks().SelecterDBMock()
// 	mock.ExpectQuery("SELECT").WillReturnError(errMock)

// 	user := newUser()
// 	user.id = gofakeit.UUID()
// 	user.createdByID = gofakeit.UUID()
// 	user.createdAt = time.Now()

// 	ok, err := store.HasUserRight(user, "AccessAuth")
// 	if err == nil{
// 		t.Fatal("Should give an error")
// 	}
// 	if ok {
// 		t.Fatal("Should not be a true return")
// 	}
// }

// func TestUserCalls(t *testing.T) {
// 	store := getDefault(t)
// 	rows := sqlmock.NewRows(allFields).AddRow(allValues...)
// 	mock := store.databaseService.Mocks().SelecterDBMock()
// 	mock.ExpectQuery("SELECT").WithArgs(fakeUserID).WillReturnRows(rows)

// 	user, ok, err := store.GetOne(fakeUserID)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !ok {
// 		t.Fatal("Entry should exist")
// 	}

// 	if user.TableName() != "User" {
// 		t.Errorf("expected table name to be `%s` but got `%s`", "User", user.TableName())
// 	}
// 	if user.TableAlias() != "u" {
// 		t.Errorf("expected table alias to be `%s` but got `%s`", "u", user.TableAlias())
// 	}
// 	for count, currencies := range []string{"", "1", "1,2", "1,2,3"} {
// 		user.SetCurrencies(currencies)
// 		if uint(count) != user.CurrenciesCount() {
// 			t.Errorf("expected currencies count to be `%d` but got `%d`", count, user.CurrenciesCount())
// 		}
// 	}
// 	name := store.Name(user, 141)
// 	if name != "John Doe" {
// 		t.Errorf("expected Name to be `%s` but got `%s`", "John Doe", name)
// 	}
// 	nameWithEmail := store.NameWithEmail(user, 141)
// 	if nameWithEmail != "John Doe (johndoe@example.com)" {
// 		t.Errorf("expected Name to be `%s` but got `%s`", "John Doe (johndoe@example.com)", nameWithEmail)
// 	}
// 	user.SetFirstName(nil)
// 	user.SetSurname(nil)
// 	nameEmptyCheck := store.Name(user, 141)
// 	if "User "+fakeUserID != nameEmptyCheck {
// 		t.Errorf("expected Name to be `%s` but got `%s`", "User "+fakeUserID, nameEmptyCheck)
// 	}
// }
