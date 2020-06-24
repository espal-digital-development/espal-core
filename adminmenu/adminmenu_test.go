package adminmenu_test

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"

	"github.com/espal-digital-development/espal-core/adminmenu"
	"github.com/espal-digital-development/espal-core/config/configmock"
	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/database/databasemock"
	"github.com/espal-digital-development/espal-core/repositories/translations/translationsmock"
	"github.com/espal-digital-development/espal-core/repositories/userrights/userrightsmock"
	"github.com/juju/errors"
)

var (
	configService          *configmock.ConfigMock
	dbRows                 *databasemock.RowsMock
	selecterDatabase       *databasemock.DatabaseMock
	translationsRepository *translationsmock.TranslationsMock
	userRightsRepository   *userrightsmock.RepositoryMock

	userRights    = []uint16{11, 12, 13, 51, 52, 53, 101, 102, 103}
	userRightRows = []string{
		"11,12,13",
		"51,52,53",
		"101,102,103",
	}
)

func initMocks() {
	configService = &configmock.ConfigMock{
		AdminURLFunc: func() string {
			return "fake-prefix-url"
		},
	}
	dbRows = &databasemock.RowsMock{
		ScanFunc: func(dest ...interface{}) error {
			return nil
		},
		NextFunc: func() bool {
			return false
		},
		CloseFunc: func() error {
			return nil
		},
		ErrFunc: func() error {
			return nil
		},
	}
	selecterDatabase = &databasemock.DatabaseMock{
		QueryFunc: func(query string, args ...interface{}) (database.Rows, error) {
			return dbRows, nil
		},
	}
	translationsRepository = &translationsmock.TranslationsMock{
		SingularFunc: func(languageID uint16, key string) string {
			return "fake-fast-trans"
		},
		PluralFunc: func(languageID uint16, key string) string {
			return "fake-fast-plural-trans"
		},
	}
	userRightsRepository = &userrightsmock.RepositoryMock{
		GetCodeFunc: func(name string) (uint16, error) {
			return userRights[rand.Intn(len(userRights)-1)], nil
		},
	}
}

func TestGenerateAdminMenuStructure(t *testing.T) {
	initMocks()
	adminMenu := adminmenu.New(configService, selecterDatabase, translationsRepository, userRightsRepository)
	var nextCounter int
	rowsLength := len(userRightRows)
	dbRows.NextFunc = func() bool {
		return nextCounter < rowsLength
	}
	dbRows.ScanFunc = func(dest ...interface{}) error {
		*dest[0].(*string) = userRightRows[nextCounter]
		nextCounter++
		return nil
	}
	menu, err := adminMenu.GenerateAdminMenuStructure("fake-id", 1)
	if err != nil {
		t.Fatal(err)
	}

	if nextCounter != rowsLength {
		t.Fatalf("Only consumed %d of the %d rows", nextCounter, rowsLength)
	}

	for k := range menu {
		if menu[k].Title() == "" {
			t.Errorf("menu entry `%d` Title shouldn't be empty", k)
			continue
		}
		items := menu[k].Items()
		for k2 := range items {
			if items[k2].Title() == "" {
				t.Errorf("menu entry `%d` Title shouldn't be empty", k2)
				continue
			}
			if items[k2].URL() == "" {
				t.Errorf("menu entry `%d` URL shouldn't be empty", k2)
				continue
			}
			if items[k2].AccessRight() == 0 {
				t.Errorf("menu entry `%d` AccessRight should not be 0", k2)
			}
		}
	}
}

func TestQueryError(t *testing.T) {
	initMocks()
	adminMenu := adminmenu.New(configService, selecterDatabase, translationsRepository, userRightsRepository)
	queryError := errors.New("fake error")
	selecterDatabase.QueryFunc = func(query string, args ...interface{}) (database.Rows, error) {
		return nil, queryError
	}
	menu, err := adminMenu.GenerateAdminMenuStructure("fake-id", 1)
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	if queryError != errors.Cause(err) {
		t.Fatalf("Expected error `%s` but got `%s`", queryError.Error(), err.Error())
	}
	if menu != nil {
		t.Fatal("menu should be nil when an error is thrown")
	}
}

func TestRowsError(t *testing.T) {
	initMocks()
	adminMenu := adminmenu.New(configService, selecterDatabase, translationsRepository, userRightsRepository)
	var nextCounter int
	rowsLength := len(userRightRows)
	dbRows.NextFunc = func() bool {
		return nextCounter < rowsLength
	}
	dbRows.ScanFunc = func(dest ...interface{}) error {
		*dest[0].(*string) = userRightRows[nextCounter]
		nextCounter++
		return nil
	}
	rowsError := errors.New("fake error")
	dbRows.ErrFunc = func() error {
		return rowsError
	}
	menu, err := adminMenu.GenerateAdminMenuStructure("fake-id", 1)
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	if rowsError != errors.Cause(err) {
		t.Fatalf("Expected error `%s` but got `%s`", rowsError.Error(), err.Error())
	}
	if menu != nil {
		t.Fatal("menu should be nil when an error is thrown")
	}
}

func TestRowsScanError(t *testing.T) {
	initMocks()
	adminMenu := adminmenu.New(configService, selecterDatabase, translationsRepository, userRightsRepository)
	var stopCounter int
	dbRows.NextFunc = func() bool {
		if stopCounter > 0 {
			return false
		}
		stopCounter++
		return true
	}
	scanError := errors.New("fake error")
	dbRows.ScanFunc = func(dest ...interface{}) error {
		return scanError
	}
	menu, err := adminMenu.GenerateAdminMenuStructure("fake-id", 1)
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	if scanError != errors.Cause(err) {
		t.Fatalf("Expected error `%s` but got `%s`", scanError.Error(), err.Error())
	}
	if menu != nil {
		t.Fatal("menu should be nil when an error is thrown")
	}
}

func TestRowsCloseError(t *testing.T) {
	initMocks()
	adminMenu := adminmenu.New(configService, selecterDatabase, translationsRepository, userRightsRepository)
	closeError := errors.New("fake error")
	dbRows.CloseFunc = func() error {
		return closeError
	}
	menu, err := adminMenu.GenerateAdminMenuStructure("fake-id", 1)
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	if closeError != errors.Cause(err) {
		t.Fatalf("Expected error `%s` but got `%s`", closeError.Error(), err.Error())
	}
	if menu != nil {
		t.Fatal("menu should be nil when an error is thrown")
	}
}

func TestParseUintError(t *testing.T) {
	initMocks()
	adminMenu := adminmenu.New(configService, selecterDatabase, translationsRepository, userRightsRepository)
	userRightRows = append(userRightRows, "")
	var nextCounter int
	rowsLength := len(userRightRows)
	dbRows.NextFunc = func() bool {
		return nextCounter < rowsLength
	}
	dbRows.ScanFunc = func(dest ...interface{}) error {
		*dest[0].(*string) = userRightRows[nextCounter]
		nextCounter++
		return nil
	}
	menu, err := adminMenu.GenerateAdminMenuStructure("fake-id", 1)
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	if !strings.Contains(err.Error(), strconv.ErrSyntax.Error()) {
		t.Fatalf("Expected error `%s` but got `%s`", strconv.ErrSyntax.Error(), err.Error())
	}
	if menu != nil {
		t.Fatal("menu should be nil when an error is thrown")
	}
}

func TestUserRightsCodeError(t *testing.T) {
	initMocks()
	adminMenu := adminmenu.New(configService, selecterDatabase, translationsRepository, userRightsRepository)
	userRightsError := errors.New("fake error")
	userRightsRepository.GetCodeFunc = func(name string) (uint16, error) {
		return 0, userRightsError
	}
	menu, err := adminMenu.GenerateAdminMenuStructure("fake-id", 1)
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	if userRightsError != errors.Cause(err) {
		t.Fatalf("Expected error `%s` but got `%s`", userRightsError.Error(), err.Error())
	}
	if menu != nil {
		t.Fatal("menu should be nil when an error is thrown")
	}
}
