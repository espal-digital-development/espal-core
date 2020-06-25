// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package databasemock

import (
	"github.com/espal-digital-development/espal-core/database"
	"sync"
)

var (
	lockDatabaseMockBegin    sync.RWMutex
	lockDatabaseMockClose    sync.RWMutex
	lockDatabaseMockExec     sync.RWMutex
	lockDatabaseMockOpen     sync.RWMutex
	lockDatabaseMockQuery    sync.RWMutex
	lockDatabaseMockQueryRow sync.RWMutex
)

// Ensure, that DatabaseMock does implement database.Database.
// If this is not the case, regenerate this file with moq.
var _ database.Database = &DatabaseMock{}

// DatabaseMock is a mock implementation of database.Database.
//
//     func TestSomethingThatUsesDatabase(t *testing.T) {
//
//         // make and configure a mocked database.Database
//         mockedDatabase := &DatabaseMock{
//             BeginFunc: func() (database.Transaction, error) {
// 	               panic("mock out the Begin method")
//             },
//             CloseFunc: func() error {
// 	               panic("mock out the Close method")
//             },
//             ExecFunc: func(query string, args ...interface{}) (database.Result, error) {
// 	               panic("mock out the Exec method")
//             },
//             OpenFunc: func(driver string, dsn string) error {
// 	               panic("mock out the Open method")
//             },
//             QueryFunc: func(query string, args ...interface{}) (database.Rows, error) {
// 	               panic("mock out the Query method")
//             },
//             QueryRowFunc: func(query string, args ...interface{}) database.Row {
// 	               panic("mock out the QueryRow method")
//             },
//         }
//
//         // use mockedDatabase in code that requires database.Database
//         // and then make assertions.
//
//     }
type DatabaseMock struct {
	// BeginFunc mocks the Begin method.
	BeginFunc func() (database.Transaction, error)

	// CloseFunc mocks the Close method.
	CloseFunc func() error

	// ExecFunc mocks the Exec method.
	ExecFunc func(query string, args ...interface{}) (database.Result, error)

	// OpenFunc mocks the Open method.
	OpenFunc func(driver string, dsn string) error

	// QueryFunc mocks the Query method.
	QueryFunc func(query string, args ...interface{}) (database.Rows, error)

	// QueryRowFunc mocks the QueryRow method.
	QueryRowFunc func(query string, args ...interface{}) database.Row

	// calls tracks calls to the methods.
	calls struct {
		// Begin holds details about calls to the Begin method.
		Begin []struct {
		}
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// Exec holds details about calls to the Exec method.
		Exec []struct {
			// Query is the query argument value.
			Query string
			// Args is the args argument value.
			Args []interface{}
		}
		// Open holds details about calls to the Open method.
		Open []struct {
			// Driver is the driver argument value.
			Driver string
			// Dsn is the dsn argument value.
			Dsn string
		}
		// Query holds details about calls to the Query method.
		Query []struct {
			// Query is the query argument value.
			Query string
			// Args is the args argument value.
			Args []interface{}
		}
		// QueryRow holds details about calls to the QueryRow method.
		QueryRow []struct {
			// Query is the query argument value.
			Query string
			// Args is the args argument value.
			Args []interface{}
		}
	}
}

// Begin calls BeginFunc.
func (mock *DatabaseMock) Begin() (database.Transaction, error) {
	if mock.BeginFunc == nil {
		panic("DatabaseMock.BeginFunc: method is nil but Database.Begin was just called")
	}
	callInfo := struct {
	}{}
	lockDatabaseMockBegin.Lock()
	mock.calls.Begin = append(mock.calls.Begin, callInfo)
	lockDatabaseMockBegin.Unlock()
	return mock.BeginFunc()
}

// BeginCalls gets all the calls that were made to Begin.
// Check the length with:
//     len(mockedDatabase.BeginCalls())
func (mock *DatabaseMock) BeginCalls() []struct {
} {
	var calls []struct {
	}
	lockDatabaseMockBegin.RLock()
	calls = mock.calls.Begin
	lockDatabaseMockBegin.RUnlock()
	return calls
}

// Close calls CloseFunc.
func (mock *DatabaseMock) Close() error {
	if mock.CloseFunc == nil {
		panic("DatabaseMock.CloseFunc: method is nil but Database.Close was just called")
	}
	callInfo := struct {
	}{}
	lockDatabaseMockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	lockDatabaseMockClose.Unlock()
	return mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedDatabase.CloseCalls())
func (mock *DatabaseMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	lockDatabaseMockClose.RLock()
	calls = mock.calls.Close
	lockDatabaseMockClose.RUnlock()
	return calls
}

// Exec calls ExecFunc.
func (mock *DatabaseMock) Exec(query string, args ...interface{}) (database.Result, error) {
	if mock.ExecFunc == nil {
		panic("DatabaseMock.ExecFunc: method is nil but Database.Exec was just called")
	}
	callInfo := struct {
		Query string
		Args  []interface{}
	}{
		Query: query,
		Args:  args,
	}
	lockDatabaseMockExec.Lock()
	mock.calls.Exec = append(mock.calls.Exec, callInfo)
	lockDatabaseMockExec.Unlock()
	return mock.ExecFunc(query, args...)
}

// ExecCalls gets all the calls that were made to Exec.
// Check the length with:
//     len(mockedDatabase.ExecCalls())
func (mock *DatabaseMock) ExecCalls() []struct {
	Query string
	Args  []interface{}
} {
	var calls []struct {
		Query string
		Args  []interface{}
	}
	lockDatabaseMockExec.RLock()
	calls = mock.calls.Exec
	lockDatabaseMockExec.RUnlock()
	return calls
}

// Open calls OpenFunc.
func (mock *DatabaseMock) Open(driver string, dsn string) error {
	if mock.OpenFunc == nil {
		panic("DatabaseMock.OpenFunc: method is nil but Database.Open was just called")
	}
	callInfo := struct {
		Driver string
		Dsn    string
	}{
		Driver: driver,
		Dsn:    dsn,
	}
	lockDatabaseMockOpen.Lock()
	mock.calls.Open = append(mock.calls.Open, callInfo)
	lockDatabaseMockOpen.Unlock()
	return mock.OpenFunc(driver, dsn)
}

// OpenCalls gets all the calls that were made to Open.
// Check the length with:
//     len(mockedDatabase.OpenCalls())
func (mock *DatabaseMock) OpenCalls() []struct {
	Driver string
	Dsn    string
} {
	var calls []struct {
		Driver string
		Dsn    string
	}
	lockDatabaseMockOpen.RLock()
	calls = mock.calls.Open
	lockDatabaseMockOpen.RUnlock()
	return calls
}

// Query calls QueryFunc.
func (mock *DatabaseMock) Query(query string, args ...interface{}) (database.Rows, error) {
	if mock.QueryFunc == nil {
		panic("DatabaseMock.QueryFunc: method is nil but Database.Query was just called")
	}
	callInfo := struct {
		Query string
		Args  []interface{}
	}{
		Query: query,
		Args:  args,
	}
	lockDatabaseMockQuery.Lock()
	mock.calls.Query = append(mock.calls.Query, callInfo)
	lockDatabaseMockQuery.Unlock()
	return mock.QueryFunc(query, args...)
}

// QueryCalls gets all the calls that were made to Query.
// Check the length with:
//     len(mockedDatabase.QueryCalls())
func (mock *DatabaseMock) QueryCalls() []struct {
	Query string
	Args  []interface{}
} {
	var calls []struct {
		Query string
		Args  []interface{}
	}
	lockDatabaseMockQuery.RLock()
	calls = mock.calls.Query
	lockDatabaseMockQuery.RUnlock()
	return calls
}

// QueryRow calls QueryRowFunc.
func (mock *DatabaseMock) QueryRow(query string, args ...interface{}) database.Row {
	if mock.QueryRowFunc == nil {
		panic("DatabaseMock.QueryRowFunc: method is nil but Database.QueryRow was just called")
	}
	callInfo := struct {
		Query string
		Args  []interface{}
	}{
		Query: query,
		Args:  args,
	}
	lockDatabaseMockQueryRow.Lock()
	mock.calls.QueryRow = append(mock.calls.QueryRow, callInfo)
	lockDatabaseMockQueryRow.Unlock()
	return mock.QueryRowFunc(query, args...)
}

// QueryRowCalls gets all the calls that were made to QueryRow.
// Check the length with:
//     len(mockedDatabase.QueryRowCalls())
func (mock *DatabaseMock) QueryRowCalls() []struct {
	Query string
	Args  []interface{}
} {
	var calls []struct {
		Query string
		Args  []interface{}
	}
	lockDatabaseMockQueryRow.RLock()
	calls = mock.calls.QueryRow
	lockDatabaseMockQueryRow.RUnlock()
	return calls
}
