package setting_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/database/databasemock"
	"github.com/espal-digital-development/espal-core/stores/setting"
	"github.com/espal-digital-development/espal-core/testtools"
	"github.com/juju/errors"
)

var (
	stubUUID         = "ffecc727-9a0f-4905-a05d-20e2fec25f3f"
	dbRow            *databasemock.RowMock
	dbRows           *databasemock.RowsMock
	selecterDatabase *databasemock.DatabaseMock
)

func initMocks() {
	dbRow = databasemock.DefaultRowMock()
	dbRows = databasemock.DefaultRowsMock()
	selecterDatabase = &databasemock.DatabaseMock{
		QueryFunc: func(query string, args ...interface{}) (database.Rows, error) {
			return dbRows, nil
		},
		QueryRowFunc: func(query string, args ...interface{}) database.Row {
			return dbRow
		},
	}
}

func TestNew(t *testing.T) {
	initMocks()
	if _, err := setting.New(selecterDatabase); err != nil {
		t.Fatal(err)
	}
}

func TestGetOneForSite(t *testing.T) {
	initMocks()
	testSettingValue := "testSettingValue"
	dbRow.ScanFunc = func(dest ...interface{}) error {
		*dest[0].(*string) = testSettingValue
		return nil
	}
	store, err := setting.New(selecterDatabase)
	if err != nil {
		t.Fatal(err)
	}

	value, err := store.GetOneForSite(1, stubUUID, stubUUID, stubUUID)
	if err != nil {
		t.Fatal(err)
	}
	if value == "" {
		t.Error("result should not be empty")
	}
	testtools.EqString(t, "testSettingValue", value, testSettingValue)
}

func TestGetOneForSiteNoResult(t *testing.T) {
	initMocks()
	dbRow.ScanFunc = func(dest ...interface{}) error {
		return nil
	}
	store, err := setting.New(selecterDatabase)
	if err != nil {
		t.Fatal(err)
	}

	value, err := store.GetOneForSite(1, stubUUID, stubUUID, stubUUID)
	if err != nil {
		t.Fatal(err)
	}
	if value != "" {
		t.Error("result should be empty")
	}
}

func TestGetOneForSiteWithError(t *testing.T) {
	initMocks()
	rowError := errors.New("rowError")
	dbRow.ScanFunc = func(dest ...interface{}) error {
		return rowError
	}
	store, err := setting.New(selecterDatabase)
	if err != nil {
		t.Fatal(err)
	}

	value, err := store.GetOneForSite(1, stubUUID, stubUUID, stubUUID)
	if err == nil {
		t.Fatal("Expected an error to be returned")
	}
	if err.Error() != rowError.Error() {
		t.Fatal("Error was not the same as expected")
	}
	if value != "" {
		t.Fatal("There should be no result")
	}
}
