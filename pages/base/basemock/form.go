// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package basemock

import (
	"github.com/espal-digital-development/espal-core/pages/base"
	"sync"
)

var (
	lockFormMockContainsSelectSearch sync.RWMutex
	lockFormMockErrors               sync.RWMutex
	lockFormMockField                sync.RWMutex
	lockFormMockOpen                 sync.RWMutex
)

// Ensure, that FormMock does implement base.Form.
// If this is not the case, regenerate this file with moq.
var _ base.Form = &FormMock{}

// FormMock is a mock implementation of base.Form.
//
//     func TestSomethingThatUsesForm(t *testing.T) {
//
//         // make and configure a mocked base.Form
//         mockedForm := &FormMock{
//             ContainsSelectSearchFunc: func() bool {
// 	               panic("mock out the ContainsSelectSearch method")
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
//         // use mockedForm in code that requires base.Form
//         // and then make assertions.
//
//     }
type FormMock struct {
	// ContainsSelectSearchFunc mocks the ContainsSelectSearch method.
	ContainsSelectSearchFunc func() bool

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
