package migrations_test

// var (
// 	configService   *configmock.ConfigMock
// 	loggerService   *loggermock.LoggerMock
// 	migratorDatabase *databasemock.DatabaseMock
// )

// func initMocks(t *testing.T) {
// 	configService = &configmock.ConfigMock{}
// 	loggerService = &loggermock.LoggerMock{}
// 	migratorDatabase = &databasemock.DatabaseMock{}
// }

// TODO :: 7 This is a bit harder to mock because it needs to mock the migrator calls
//         inside the Migrations, but also the INSERTs inside the migratorDatabase.
// func TestRun(t *testing.T) {
// 	migrations := getDefault(t)
// 	migratorMock := migrations.MigratorDBMock
// 	migratorMock.ExpectExec("*")
// 	// mock := migrations.migratorDatabase.Mocks().InserterDBMock
// 	// // mock.ExpectQuery("SELECT").WithArgs(1).WillReturnRows(rows)
// 	// mock.ExpectExec("INSERT INTO *")

// 	if err := migrations.Run(); err != nil {
// 		// t.Fatal(err)
// }
