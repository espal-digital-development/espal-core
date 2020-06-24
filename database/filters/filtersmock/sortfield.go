// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package filtersmock

import (
	"github.com/espal-digital-development/espal-core/database/filters"
	"sync"
)

var (
	lockSortFieldMockDescending    sync.RWMutex
	lockSortFieldMockName          sync.RWMutex
	lockSortFieldMockSetTableAlias sync.RWMutex
	lockSortFieldMockTableAlias    sync.RWMutex
)

// SortFieldMock is a mock implementation of SortField.
//
//     func TestSomethingThatUsesSortField(t *testing.T) {
//
//         // make and configure a mocked SortField
//         mockedSortField := &SortFieldMock{
//             DescendingFunc: func() bool {
// 	               panic("mock out the Descending method")
//             },
//             NameFunc: func() string {
// 	               panic("mock out the Name method")
//             },
//             SetTableAliasFunc: func(tableAlias string) filters.SortField {
// 	               panic("mock out the SetTableAlias method")
//             },
//             TableAliasFunc: func() string {
// 	               panic("mock out the TableAlias method")
//             },
//         }
//
//         // use mockedSortField in code that requires SortField
//         // and then make assertions.
//
//     }
type SortFieldMock struct {
	// DescendingFunc mocks the Descending method.
	DescendingFunc func() bool

	// NameFunc mocks the Name method.
	NameFunc func() string

	// SetTableAliasFunc mocks the SetTableAlias method.
	SetTableAliasFunc func(tableAlias string) filters.SortField

	// TableAliasFunc mocks the TableAlias method.
	TableAliasFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// Descending holds details about calls to the Descending method.
		Descending []struct {
		}
		// Name holds details about calls to the Name method.
		Name []struct {
		}
		// SetTableAlias holds details about calls to the SetTableAlias method.
		SetTableAlias []struct {
			// TableAlias is the tableAlias argument value.
			TableAlias string
		}
		// TableAlias holds details about calls to the TableAlias method.
		TableAlias []struct {
		}
	}
}

// Descending calls DescendingFunc.
func (mock *SortFieldMock) Descending() bool {
	if mock.DescendingFunc == nil {
		panic("SortFieldMock.DescendingFunc: method is nil but SortField.Descending was just called")
	}
	callInfo := struct {
	}{}
	lockSortFieldMockDescending.Lock()
	mock.calls.Descending = append(mock.calls.Descending, callInfo)
	lockSortFieldMockDescending.Unlock()
	return mock.DescendingFunc()
}

// DescendingCalls gets all the calls that were made to Descending.
// Check the length with:
//     len(mockedSortField.DescendingCalls())
func (mock *SortFieldMock) DescendingCalls() []struct {
} {
	var calls []struct {
	}
	lockSortFieldMockDescending.RLock()
	calls = mock.calls.Descending
	lockSortFieldMockDescending.RUnlock()
	return calls
}

// Name calls NameFunc.
func (mock *SortFieldMock) Name() string {
	if mock.NameFunc == nil {
		panic("SortFieldMock.NameFunc: method is nil but SortField.Name was just called")
	}
	callInfo := struct {
	}{}
	lockSortFieldMockName.Lock()
	mock.calls.Name = append(mock.calls.Name, callInfo)
	lockSortFieldMockName.Unlock()
	return mock.NameFunc()
}

// NameCalls gets all the calls that were made to Name.
// Check the length with:
//     len(mockedSortField.NameCalls())
func (mock *SortFieldMock) NameCalls() []struct {
} {
	var calls []struct {
	}
	lockSortFieldMockName.RLock()
	calls = mock.calls.Name
	lockSortFieldMockName.RUnlock()
	return calls
}

// SetTableAlias calls SetTableAliasFunc.
func (mock *SortFieldMock) SetTableAlias(tableAlias string) filters.SortField {
	if mock.SetTableAliasFunc == nil {
		panic("SortFieldMock.SetTableAliasFunc: method is nil but SortField.SetTableAlias was just called")
	}
	callInfo := struct {
		TableAlias string
	}{
		TableAlias: tableAlias,
	}
	lockSortFieldMockSetTableAlias.Lock()
	mock.calls.SetTableAlias = append(mock.calls.SetTableAlias, callInfo)
	lockSortFieldMockSetTableAlias.Unlock()
	return mock.SetTableAliasFunc(tableAlias)
}

// SetTableAliasCalls gets all the calls that were made to SetTableAlias.
// Check the length with:
//     len(mockedSortField.SetTableAliasCalls())
func (mock *SortFieldMock) SetTableAliasCalls() []struct {
	TableAlias string
} {
	var calls []struct {
		TableAlias string
	}
	lockSortFieldMockSetTableAlias.RLock()
	calls = mock.calls.SetTableAlias
	lockSortFieldMockSetTableAlias.RUnlock()
	return calls
}

// TableAlias calls TableAliasFunc.
func (mock *SortFieldMock) TableAlias() string {
	if mock.TableAliasFunc == nil {
		panic("SortFieldMock.TableAliasFunc: method is nil but SortField.TableAlias was just called")
	}
	callInfo := struct {
	}{}
	lockSortFieldMockTableAlias.Lock()
	mock.calls.TableAlias = append(mock.calls.TableAlias, callInfo)
	lockSortFieldMockTableAlias.Unlock()
	return mock.TableAliasFunc()
}

// TableAliasCalls gets all the calls that were made to TableAlias.
// Check the length with:
//     len(mockedSortField.TableAliasCalls())
func (mock *SortFieldMock) TableAliasCalls() []struct {
} {
	var calls []struct {
	}
	lockSortFieldMockTableAlias.RLock()
	calls = mock.calls.TableAlias
	lockSortFieldMockTableAlias.RUnlock()
	return calls
}
