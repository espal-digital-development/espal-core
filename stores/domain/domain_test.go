package domain_test

import (
	"strings"
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/database/databasemock"
	"github.com/espal-digital-development/espal-core/database/filters/filtersmock"
	"github.com/espal-digital-development/espal-core/routing/router/contexts/contextsmock"
	"github.com/espal-digital-development/espal-core/stores/domain"
	"github.com/espal-digital-development/espal-core/testtools"
	"github.com/juju/errors"
)

var (
	fakeDomainID = "fake-domain-id-1"
	fakeHost     = "example.com"
	allValues    = []interface{}{"fake-domain-id-1", "fake-user-id-2", nil, "fake-site-id-1", time.Now(), nil, true,
		"example.com", nil, "33"}

	transaction      *databasemock.TransactionMock
	dbRows           *databasemock.RowsMock
	selecterDatabase *databasemock.DatabaseMock
	updaterDatabase  *databasemock.DatabaseMock
	deletorDatabase  *databasemock.DatabaseMock
	filtersFactory   *filtersmock.FactoryMock
	filter           *filtersmock.FilterMock
	context          *contextsmock.ContextMock
)

func initMocks() {
	dbRows = databasemock.DefaultRowsMock()
	selecterDatabase = &databasemock.DatabaseMock{
		QueryFunc: func(query string, args ...interface{}) (database.Rows, error) {
			return dbRows, nil
		},
	}
	updaterDatabase = &databasemock.DatabaseMock{
		BeginFunc: func() (database.Transaction, error) {
			return transaction, nil
		},
	}
	transaction = &databasemock.TransactionMock{
		ExecFunc: func(query string, args ...interface{}) (database.Result, error) {
			return nil, nil
		},
		CommitFunc: func() error {
			return nil
		},
		RollbackFunc: func() error {
			return nil
		},
		QueryFunc: func(query string, args ...interface{}) (database.Rows, error) {
			return nil, nil
		},
	}
	deletorDatabase = &databasemock.DatabaseMock{
		BeginFunc: func() (database.Transaction, error) {
			return transaction, nil
		},
	}
	filtersFactory, filter, _, _, _, _, _ = filtersmock.DefaultMocks() // nolint:dogsled
}

func TestNew(t *testing.T) {
	initMocks()
	if _, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory); err != nil {
		t.Fatal(err)
	}
}

func TestCurrenciesCount(t *testing.T) {
	domain := domain.NewDomainEntity()
	domain.SetCurrencies("1,2,3")
	count := domain.CurrenciesCount()
	if count != 3 {
		t.Fatalf("3 currency IDs were set, but only %d were returned", count)
	}
}

func TestCurrenciesCountNone(t *testing.T) {
	domain := domain.NewDomainEntity()
	domain.SetCurrencies("")
	count := domain.CurrenciesCount()
	if count != 0 {
		t.Fatalf("0 currency IDs were set, but %d were returned", count)
	}
}

func TestHostWithProtocol(t *testing.T) {
	domain := domain.NewDomainEntity()
	expectedHostWithProtocol := "https://example.com"
	domain.SetHost("example.com")
	testtools.EqString(t, "HostWithProtocol", domain.HostWithProtocol(), expectedHostWithProtocol)
}

func TestHostWithProtocolAndWWW(t *testing.T) {
	domain := domain.NewDomainEntity()
	expectedHostWithProtocolAndWWW := "https://www.example.com"
	domain.SetHost("example.com")
	testtools.EqString(t, "HostWithProtocolAndWWW", domain.HostWithProtocolAndWWW(), expectedHostWithProtocolAndWWW)
}

func TestGetAll(t *testing.T) {
	initMocks()
	domainRows := [][]interface{}{
		allValues,
		allValues,
		allValues,
	}
	var nextCounter int
	rowsLength := len(domainRows)
	dbRows.NextFunc = func() bool {
		return nextCounter < rowsLength
	}
	dbRows.ScanFunc = func(dest ...interface{}) error {
		if err := testtools.ScanInterfaceValues(domainRows[nextCounter], dest); err != nil {
			return err
		}
		nextCounter++
		return nil
	}
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}

	domains, found, err := store.All()
	if err != nil {
		t.Fatal(err)
	}
	if !found {
		t.Fatal("Expected results but the function indicated none were found")
	}
	if len(domains) != 3 {
		t.Fatalf("Expected 3 results but got %d", len(domains))
	}
}

func TestGetAllWithNoResult(t *testing.T) {
	initMocks()
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}

	domains, found, err := store.All()
	if err != nil {
		t.Fatal(err)
	}
	if found {
		t.Fatal("Expected no results")
	}
	if len(domains) > 0 {
		t.Fatal("There shouldn't be any results returned")
	}
}

func TestGetAllWithError(t *testing.T) {
	initMocks()
	getAllError := errors.New("getAllError")
	selecterDatabase.QueryFunc = func(query string, args ...interface{}) (database.Rows, error) {
		return nil, getAllError
	}
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}

	domains, found, err := store.All()
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	testtools.EqError(t, "getAllError", err, getAllError)
	if found {
		t.Fatal("Expected no results")
	}
	if len(domains) > 0 {
		t.Fatal("There shouldn't be any results returned")
	}
}

func TestGetOne(t *testing.T) {
	initMocks()
	domainRows := [][]interface{}{
		allValues,
	}
	var nextCounter int
	rowsLength := len(domainRows)
	dbRows.NextFunc = func() bool {
		return nextCounter < rowsLength
	}
	dbRows.ScanFunc = func(dest ...interface{}) error {
		if err := testtools.ScanInterfaceValues(domainRows[nextCounter], dest); err != nil {
			return err
		}
		nextCounter++
		return nil
	}
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}

	domain, found, err := store.GetOne(fakeDomainID)
	if err != nil {
		t.Fatal(err)
	}
	if !found {
		t.Error("result should've been found")
	}
	if domain == nil {
		t.Error("result should not be nil")
	}
}

func TestGetOneWithNoResult(t *testing.T) {
	initMocks()
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}

	domain, found, err := store.GetOne(fakeDomainID)
	if err != nil {
		t.Fatal(err)
	}
	if found {
		t.Error("result should've not been found")
	}
	if domain != nil {
		t.Error("result should be nil")
	}
}

func TestGetOneWithError(t *testing.T) {
	initMocks()
	getOneError := errors.New("getOneError")
	selecterDatabase.QueryFunc = func(query string, args ...interface{}) (database.Rows, error) {
		return nil, getOneError
	}
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}

	domain, found, err := store.GetOne(fakeDomainID)
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	testtools.EqError(t, "getOneError", err, getOneError)
	if found {
		t.Fatal("Expected no results")
	}
	if domain != nil {
		t.Error("result should be nil")
	}
}

func TestGetOneByIDWithCreator(t *testing.T) {
	initMocks()
	domainRows := [][]interface{}{
		allValues,
	}
	var nextCounter int
	rowsLength := len(domainRows)
	dbRows.NextFunc = func() bool {
		return nextCounter < rowsLength
	}
	dbRows.ScanFunc = func(dest ...interface{}) error {
		if err := testtools.ScanInterfaceValues(domainRows[nextCounter], dest); err != nil {
			return err
		}
		nextCounter++
		return nil
	}
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}

	domain, found, err := store.GetOneByIDWithCreator(fakeDomainID)
	if err != nil {
		t.Fatal(err)
	}
	if !found {
		t.Error("result should've been found")
	}
	if domain == nil {
		t.Fatal("Result should not be nil")
	}
	testtools.EqString(t, "domain.Host", domain.Host(), fakeHost)
}

func TestGetOneByIDWithCreatorWithNoResult(t *testing.T) {
	initMocks()
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}

	domain, found, err := store.GetOneByIDWithCreator(fakeDomainID)
	if err != nil {
		t.Fatal(err)
	}
	if found {
		t.Error("result should've not been found")
	}
	if domain != nil {
		t.Error("result should be nil")
	}
}

func TestGetOneByIDWithCreatorWithError(t *testing.T) {
	initMocks()
	getOneByIDWithCreatorError := errors.New("getOneByIDWithCreatorError")
	selecterDatabase.QueryFunc = func(query string, args ...interface{}) (database.Rows, error) {
		return nil, getOneByIDWithCreatorError
	}
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}

	domain, found, err := store.GetOneByIDWithCreator(fakeDomainID)
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	testtools.EqError(t, "getOneByIDWithCreatorError", err, getOneByIDWithCreatorError)
	if found {
		t.Fatal("Expected no results")
	}
	if domain != nil {
		t.Error("result should be nil")
	}
}

func TestGetOneActiveByHost(t *testing.T) {
	initMocks()
	domainRows := [][]interface{}{
		allValues,
	}
	var nextCounter int
	rowsLength := len(domainRows)
	dbRows.NextFunc = func() bool {
		return nextCounter < rowsLength
	}
	dbRows.ScanFunc = func(dest ...interface{}) error {
		if err := testtools.ScanInterfaceValues(domainRows[nextCounter], dest); err != nil {
			return err
		}
		nextCounter++
		return nil
	}
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}

	domain, found, err := store.GetOneActiveByHost(fakeDomainID)
	if err != nil {
		t.Fatal(err)
	}
	if !found {
		t.Error("result should've been found")
	}
	if domain == nil {
		t.Error("result should not be nil")
	}
}

func TestGetOneActiveByHostNoResult(t *testing.T) {
	initMocks()
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}

	domain, found, err := store.GetOneActiveByHost(fakeDomainID)
	if err != nil {
		t.Fatal(err)
	}
	if found {
		t.Error("result should've not been found")
	}
	if domain != nil {
		t.Error("result should be nil")
	}
}

func TestGetOneActiveByHostError(t *testing.T) {
	initMocks()
	getOneActiveByHostError := errors.New("getOneActiveByHostError")
	selecterDatabase.QueryFunc = func(query string, args ...interface{}) (database.Rows, error) {
		return nil, getOneActiveByHostError
	}
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}

	domain, found, err := store.GetOneActiveByHost(fakeDomainID)
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	testtools.EqError(t, "getOneActiveByHostError", err, getOneActiveByHostError)
	if found {
		t.Fatal("Expected no results")
	}
	if domain != nil {
		t.Error("result should be nil")
	}
}

func TestDelete(t *testing.T) {
	initMocks()
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}
	err = store.Delete([]string{fakeDomainID})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteBeginError(t *testing.T) {
	initMocks()
	beginError := errors.New("beginError")
	deletorDatabase.BeginFunc = func() (database.Transaction, error) {
		return nil, beginError
	}
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}
	err = store.Delete([]string{fakeDomainID})
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	testtools.EqError(t, "beginError", err, beginError)
}

func TestDeleteTransactionExecError(t *testing.T) {
	initMocks()
	transactionExecError := errors.New("transactionExecError")
	transaction.ExecFunc = func(query string, args ...interface{}) (database.Result, error) {
		return nil, transactionExecError
	}
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}
	err = store.Delete([]string{fakeDomainID})
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	testtools.EqError(t, "transactionExecError", err, transactionExecError)
	if len(transaction.RollbackCalls()) != 1 {
		t.Fatal("Rollback wasn't called after Exec failure")
	}
}

func TestDeleteTransactionExecRollbackError(t *testing.T) {
	initMocks()
	transactionExecError := errors.New("transactionExecError")
	transactionRollbackError := errors.New("transactionRollbackError")
	transaction.ExecFunc = func(query string, args ...interface{}) (database.Result, error) {
		return nil, transactionExecError
	}
	transaction.RollbackFunc = func() error {
		return transactionRollbackError
	}
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}
	err = store.Delete([]string{fakeDomainID})
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	testtools.EqError(t, "transactionRollbackError", err, transactionRollbackError)
}

func TestToggleActive(t *testing.T) {
	initMocks()
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}
	err = store.ToggleActive([]string{fakeDomainID})
	if err != nil {
		t.Fatal(err)
	}
}

func TestToggleActiveBeginError(t *testing.T) {
	initMocks()
	beginError := errors.New("beginError")
	updaterDatabase.BeginFunc = func() (database.Transaction, error) {
		return nil, beginError
	}
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}
	err = store.ToggleActive([]string{fakeDomainID})
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	testtools.EqError(t, "beginError", err, beginError)
}

func TestToggleActiveTransactionQueryError(t *testing.T) {
	initMocks()
	transactionQueryError := errors.New("transactionQueryError")
	transaction.QueryFunc = func(query string, args ...interface{}) (database.Rows, error) {
		return nil, transactionQueryError
	}
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}
	err = store.ToggleActive([]string{fakeDomainID})
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	testtools.EqError(t, "transactionQueryError", err, transactionQueryError)
	if len(transaction.RollbackCalls()) != 1 {
		t.Fatal("Rollback wasn't called after Exec failure")
	}
}

func TestToggleActiveTransactionQueryRollbackError(t *testing.T) {
	initMocks()
	transactionQueryError := errors.New("transactionQueryError")
	transactionRollbackError := errors.New("transactionRollbackError")
	transaction.QueryFunc = func(query string, args ...interface{}) (database.Rows, error) {
		return nil, transactionQueryError
	}
	transaction.RollbackFunc = func() error {
		return transactionRollbackError
	}
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}
	err = store.ToggleActive([]string{fakeDomainID})
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	testtools.EqError(t, "transactionRollbackError", err, transactionRollbackError)
	if len(transaction.RollbackCalls()) != 1 {
		t.Fatal("Rollback wasn't called after Exec failure")
	}
}

func setupFilter() {
	values := []interface{}{
		fakeDomainID, "fake-domain-id-2", nil, "fake-site-id", nil, nil, nil, nil,
		true, time.Now(), nil, "example.com", nil, "33",
	}
	domainRows := [][]interface{}{
		values,
		values,
		values,
	}
	var nextCounter int
	rowsLength := len(domainRows)
	dbRows.NextFunc = func() bool {
		return nextCounter < rowsLength
	}
	dbRows.ScanFunc = func(dest ...interface{}) error {
		if err := testtools.ScanInterfaceValues(domainRows[nextCounter], dest); err != nil {
			return err
		}
		nextCounter++
		return nil
	}
	filter.HasResultsFunc = func() bool {
		return true
	}
	filter.RowsFunc = func() database.Rows {
		return dbRows
	}
}

func TestFilter(t *testing.T) {
	initMocks()
	setupFilter()
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}

	domains, filter, err := store.Filter(context)
	if err != nil {
		t.Fatal(err)
	}
	if !filter.HasResults() {
		t.Error("result should have results")
	}
	if len(domains) == 0 {
		t.Error("result should be not be empty")
	}
}

func TestFilterProcessError(t *testing.T) {
	initMocks()
	setupFilter()
	processError := errors.New("processError")
	filter.ProcessFunc = func() error {
		return processError
	}
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}

	domains, filter, err := store.Filter(context)
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	testtools.EqError(t, "processError", err, processError)
	if !filter.HasResults() {
		t.Fatal("Expected the filter to say it had found results")
	}
	if len(domains) != 0 {
		t.Error("result should be empty because an error was thrown")
	}
}

func TestFilterRowsErrError(t *testing.T) {
	initMocks()
	setupFilter()
	rowsErrError := errors.New("rowsErrError")
	dbRows.ErrFunc = func() error {
		return rowsErrError
	}
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}

	domains, filter, err := store.Filter(context)
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	testtools.EqError(t, "rowsErrError", err, rowsErrError)
	if !filter.HasResults() {
		t.Fatal("Expected the filter to say it had found results")
	}
	if len(domains) != 0 {
		t.Error("result should be empty because an error was thrown")
	}
}

func TestFilterScanError(t *testing.T) {
	initMocks()
	setupFilter()
	scanError := errors.New("scanError")
	dbRows.ScanFunc = func(dest ...interface{}) error {
		return scanError
	}
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}

	domains, filter, err := store.Filter(context)
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	testtools.EqError(t, "scanError", err, scanError)
	if !filter.HasResults() {
		t.Fatal("Expected the filter to say it had found results")
	}
	if len(domains) != 0 {
		t.Error("result should be empty because an error was thrown")
	}
}

func TestFilterCloseError(t *testing.T) {
	initMocks()
	setupFilter()
	closeError := errors.New("closeError")
	dbRows.CloseFunc = func() error {
		return closeError
	}
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}

	domains, filter, err := store.Filter(context)
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	testtools.EqError(t, "closeError", err, closeError)
	if !filter.HasResults() {
		t.Fatal("Expected the filter to say it had found results")
	}
	if len(domains) != 3 {
		t.Error("result should not be empty because an the data was processed still")
	}
}

func TestFilterComboCloseError(t *testing.T) {
	initMocks()
	setupFilter()
	scanError := errors.New("scanError")
	dbRows.ScanFunc = func(dest ...interface{}) error {
		return scanError
	}
	closeError := errors.New("closeError")
	dbRows.CloseFunc = func() error {
		return closeError
	}
	store, err := domain.New(selecterDatabase, updaterDatabase, deletorDatabase, filtersFactory)
	if err != nil {
		t.Fatal(err)
	}

	domains, filter, err := store.Filter(context)
	if err == nil {
		t.Fatal("An error should've been thrown")
	}
	testtools.EqError(t, "closeError", err, closeError)
	if !strings.Contains(errors.Details(err), "scanError") {
		t.Fatal("Expected scanError to be wrapped on the closeError")
	}
	if !filter.HasResults() {
		t.Fatal("Expected the filter to say it had found results")
	}
	if len(domains) != 0 {
		t.Error("result should be empty because an error was thrown")
	}
}
