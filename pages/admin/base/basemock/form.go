// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package basemock

import (
	"sync"
)

var (
	lockFormMockContainsSelectSearch sync.RWMutex
	lockFormMockCreateUpdateActions  sync.RWMutex
	lockFormMockErrors               sync.RWMutex
	lockFormMockField                sync.RWMutex
	lockFormMockOpen                 sync.RWMutex
)

// FormMock is a mock implementation of Form.
//
//     func TestSomethingThatUsesForm(t *testing.T) {
//
//         // make and configure a mocked Form
//         mockedForm := &FormMock{
//             ContainsSelectSearchFunc: func() bool {
// 	               panic("mock out the ContainsSelectSearch method")
//             },
//             CreateUpdateActionsFunc: func(in1 string, in2 string) string {
// 	               panic("mock out the CreateUpdateActions method")
//             },
//             ErrorsFunc: func() string {
// 	               panic("mock out the Errors method")
//             },
//             FieldFunc: func(in1 string) string {
// 	               panic("mock out the Field method")
//             },
//             OpenFunc: func() string {
// 	               panic("mock out the Open method")
//             },
//         }
//
//         // use mockedForm in code that requires Form
//         // and then make assertions.
//
//     }
type FormMock struct {
	// ContainsSelectSearchFunc mocks the ContainsSelectSearch method.
	ContainsSelectSearchFunc func() bool

	// CreateUpdateActionsFunc mocks the CreateUpdateActions method.
	CreateUpdateActionsFunc func(in1 string, in2 string) string

	// ErrorsFunc mocks the Errors method.
	ErrorsFunc func() string

	// FieldFunc mocks the Field method.
	FieldFunc func(in1 string) string

	// OpenFunc mocks the Open method.
	OpenFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// ContainsSelectSearch holds details about calls to the ContainsSelectSearch method.
		ContainsSelectSearch []struct {
		}
		// CreateUpdateActions holds details about calls to the CreateUpdateActions method.
		CreateUpdateActions []struct {
			// In1 is the in1 argument value.
			In1 string
			// In2 is the in2 argument value.
			In2 string
		}
		// Errors holds details about calls to the Errors method.
		Errors []struct {
		}
		// Field holds details about calls to the Field method.
		Field []struct {
			// In1 is the in1 argument value.
			In1 string
		}
		// Open holds details about calls to the Open method.
		Open []struct {
		}
	}
}

// ContainsSelectSearch calls ContainsSelectSearchFunc.
func (mock *FormMock) ContainsSelectSearch() bool {
	if mock.ContainsSelectSearchFunc == nil {
		panic("FormMock.ContainsSelectSearchFunc: method is nil but Form.ContainsSelectSearch was just called")
	}
	callInfo := struct {
	}{}
	lockFormMockContainsSelectSearch.Lock()
	mock.calls.ContainsSelectSearch = append(mock.calls.ContainsSelectSearch, callInfo)
	lockFormMockContainsSelectSearch.Unlock()
	return mock.ContainsSelectSearchFunc()
}

// ContainsSelectSearchCalls gets all the calls that were made to ContainsSelectSearch.
// Check the length with:
//     len(mockedForm.ContainsSelectSearchCalls())
func (mock *FormMock) ContainsSelectSearchCalls() []struct {
} {
	var calls []struct {
	}
	lockFormMockContainsSelectSearch.RLock()
	calls = mock.calls.ContainsSelectSearch
	lockFormMockContainsSelectSearch.RUnlock()
	return calls
}

// CreateUpdateActions calls CreateUpdateActionsFunc.
func (mock *FormMock) CreateUpdateActions(in1 string, in2 string) string {
	if mock.CreateUpdateActionsFunc == nil {
		panic("FormMock.CreateUpdateActionsFunc: method is nil but Form.CreateUpdateActions was just called")
	}
	callInfo := struct {
		In1 string
		In2 string
	}{
		In1: in1,
		In2: in2,
	}
	lockFormMockCreateUpdateActions.Lock()
	mock.calls.CreateUpdateActions = append(mock.calls.CreateUpdateActions, callInfo)
	lockFormMockCreateUpdateActions.Unlock()
	return mock.CreateUpdateActionsFunc(in1, in2)
}

// CreateUpdateActionsCalls gets all the calls that were made to CreateUpdateActions.
// Check the length with:
//     len(mockedForm.CreateUpdateActionsCalls())
func (mock *FormMock) CreateUpdateActionsCalls() []struct {
	In1 string
	In2 string
} {
	var calls []struct {
		In1 string
		In2 string
	}
	lockFormMockCreateUpdateActions.RLock()
	calls = mock.calls.CreateUpdateActions
	lockFormMockCreateUpdateActions.RUnlock()
	return calls
}

// Errors calls ErrorsFunc.
func (mock *FormMock) Errors() string {
	if mock.ErrorsFunc == nil {
		panic("FormMock.ErrorsFunc: method is nil but Form.Errors was just called")
	}
	callInfo := struct {
	}{}
	lockFormMockErrors.Lock()
	mock.calls.Errors = append(mock.calls.Errors, callInfo)
	lockFormMockErrors.Unlock()
	return mock.ErrorsFunc()
}

// ErrorsCalls gets all the calls that were made to Errors.
// Check the length with:
//     len(mockedForm.ErrorsCalls())
func (mock *FormMock) ErrorsCalls() []struct {
} {
	var calls []struct {
	}
	lockFormMockErrors.RLock()
	calls = mock.calls.Errors
	lockFormMockErrors.RUnlock()
	return calls
}

// Field calls FieldFunc.
func (mock *FormMock) Field(in1 string) string {
	if mock.FieldFunc == nil {
		panic("FormMock.FieldFunc: method is nil but Form.Field was just called")
	}
	callInfo := struct {
		In1 string
	}{
		In1: in1,
	}
	lockFormMockField.Lock()
	mock.calls.Field = append(mock.calls.Field, callInfo)
	lockFormMockField.Unlock()
	return mock.FieldFunc(in1)
}

// FieldCalls gets all the calls that were made to Field.
// Check the length with:
//     len(mockedForm.FieldCalls())
func (mock *FormMock) FieldCalls() []struct {
	In1 string
} {
	var calls []struct {
		In1 string
	}
	lockFormMockField.RLock()
	calls = mock.calls.Field
	lockFormMockField.RUnlock()
	return calls
}

// Open calls OpenFunc.
func (mock *FormMock) Open() string {
	if mock.OpenFunc == nil {
		panic("FormMock.OpenFunc: method is nil but Form.Open was just called")
	}
	callInfo := struct {
	}{}
	lockFormMockOpen.Lock()
	mock.calls.Open = append(mock.calls.Open, callInfo)
	lockFormMockOpen.Unlock()
	return mock.OpenFunc()
}

// OpenCalls gets all the calls that were made to Open.
// Check the length with:
//     len(mockedForm.OpenCalls())
func (mock *FormMock) OpenCalls() []struct {
} {
	var calls []struct {
	}
	lockFormMockOpen.RLock()
	calls = mock.calls.Open
	lockFormMockOpen.RUnlock()
	return calls
}
