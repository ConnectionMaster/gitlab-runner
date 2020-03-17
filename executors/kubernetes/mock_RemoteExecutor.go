// Code generated by mockery v1.0.0. DO NOT EDIT.

// This comment works around https://github.com/vektra/mockery/issues/155

package kubernetes

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	url "net/url"
)

// MockRemoteExecutor is an autogenerated mock type for the RemoteExecutor type
type MockRemoteExecutor struct {
	mock.Mock
}

// Execute provides a mock function with given fields: method, _a1, config, stdin, stdout, stderr, tty
func (_m *MockRemoteExecutor) Execute(method string, _a1 *url.URL, config *rest.Config, stdin io.Reader, stdout io.Writer, stderr io.Writer, tty bool) error {
	ret := _m.Called(method, _a1, config, stdin, stdout, stderr, tty)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *url.URL, *rest.Config, io.Reader, io.Writer, io.Writer, bool) error); ok {
		r0 = rf(method, _a1, config, stdin, stdout, stderr, tty)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
