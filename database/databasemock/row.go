// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package databasemock

import (
	"github.com/espal-digital-development/espal-core/database"
	"sync"
)

var (
	lockRowMockScan sync.RWMutex
)

// Ensure, that RowMock does implement database.Row.
// If this is not the case, regenerate this file with moq.
var _ database.Row = &RowMock{}

// RowMock is a mock implementation of database.Row.
//
//     func TestSomethingThatUsesRow(t *testing.T) {
//
//         // make and configure a mocked database.Row
//         mockedRow := &RowMock{
//             ScanFunc: func(dest ...interface{}) error {
// 	               panic("mock out the Scan method")
//             },
//         }
//
//         // use mockedRow in code that requires database.Row
//         // and then make assertions.
//
//     }
type RowMock struct {
	// ScanFunc mocks the Scan method.
	ScanFunc func(dest ...interface{}) error

	// calls tracks calls to the methods.
	calls struct {
		// Scan holds details about calls to the Scan method.
		Scan []struct {
			// Dest is the dest argument value.
			Dest []interface{}
		}
	}
}

// Scan calls ScanFunc.
func (mock *RowMock) Scan(dest ...interface{}) error {
	if mock.ScanFunc == nil {
		panic("RowMock.ScanFunc: method is nil but Row.Scan was just called")
	}
	callInfo := struct {
		Dest []interface{}
	}{
		Dest: dest,
	}
	lockRowMockScan.Lock()
	mock.calls.Scan = append(mock.calls.Scan, callInfo)
	lockRowMockScan.Unlock()
	return mock.ScanFunc(dest...)
}

// ScanCalls gets all the calls that were made to Scan.
// Check the length with:
//     len(mockedRow.ScanCalls())
func (mock *RowMock) ScanCalls() []struct {
	Dest []interface{}
} {
	var calls []struct {
		Dest []interface{}
	}
	lockRowMockScan.RLock()
	calls = mock.calls.Scan
	lockRowMockScan.RUnlock()
	return calls
}
