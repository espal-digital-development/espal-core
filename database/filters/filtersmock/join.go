// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package filtersmock

import (
	"sync"
)

var (
	lockJoinMockAlias     sync.RWMutex
	lockJoinMockStatement sync.RWMutex
)

// JoinMock is a mock implementation of Join.
//
//     func TestSomethingThatUsesJoin(t *testing.T) {
//
//         // make and configure a mocked Join
//         mockedJoin := &JoinMock{
//             AliasFunc: func() string {
// 	               panic("mock out the Alias method")
//             },
//             StatementFunc: func() string {
// 	               panic("mock out the Statement method")
//             },
//         }
//
//         // use mockedJoin in code that requires Join
//         // and then make assertions.
//
//     }
type JoinMock struct {
	// AliasFunc mocks the Alias method.
	AliasFunc func() string

	// StatementFunc mocks the Statement method.
	StatementFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// Alias holds details about calls to the Alias method.
		Alias []struct {
		}
		// Statement holds details about calls to the Statement method.
		Statement []struct {
		}
	}
}

// Alias calls AliasFunc.
func (mock *JoinMock) Alias() string {
	if mock.AliasFunc == nil {
		panic("JoinMock.AliasFunc: method is nil but Join.Alias was just called")
	}
	callInfo := struct {
	}{}
	lockJoinMockAlias.Lock()
	mock.calls.Alias = append(mock.calls.Alias, callInfo)
	lockJoinMockAlias.Unlock()
	return mock.AliasFunc()
}

// AliasCalls gets all the calls that were made to Alias.
// Check the length with:
//     len(mockedJoin.AliasCalls())
func (mock *JoinMock) AliasCalls() []struct {
} {
	var calls []struct {
	}
	lockJoinMockAlias.RLock()
	calls = mock.calls.Alias
	lockJoinMockAlias.RUnlock()
	return calls
}

// Statement calls StatementFunc.
func (mock *JoinMock) Statement() string {
	if mock.StatementFunc == nil {
		panic("JoinMock.StatementFunc: method is nil but Join.Statement was just called")
	}
	callInfo := struct {
	}{}
	lockJoinMockStatement.Lock()
	mock.calls.Statement = append(mock.calls.Statement, callInfo)
	lockJoinMockStatement.Unlock()
	return mock.StatementFunc()
}

// StatementCalls gets all the calls that were made to Statement.
// Check the length with:
//     len(mockedJoin.StatementCalls())
func (mock *JoinMock) StatementCalls() []struct {
} {
	var calls []struct {
	}
	lockJoinMockStatement.RLock()
	calls = mock.calls.Statement
	lockJoinMockStatement.RUnlock()
	return calls
}
